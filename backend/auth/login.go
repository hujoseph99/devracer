
func (myAPI *API) loginUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := context.TODO()

	user, err := decodeUserBody(w, r)
	if err != nil {
		defaultError(w, r, http.StatusBadRequest, "Invalid request.")
		return
	}

	// check if username already exists
	existingUser, err := myAPI.Database.FindUserByUsername(ctx, user.Username)
	if err != nil || existingUser == nil {
		defaultError(w, r, http.StatusUnauthorized, "The username was not found. Please try again.")
		return
	}

	byteHash := []byte(existingUser.Password)
	bytePassword := []byte(user.Password)
	err = bcrypt.CompareHashAndPassword(byteHash, bytePassword)
	if err != nil {
		defaultError(w, r, http.StatusUnauthorized, "The password is incorrect. Please try again.")
		return
	}

	returnUserToClient(w, r, existingUser)
}
