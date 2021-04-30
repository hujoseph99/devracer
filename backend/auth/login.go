package auth

// var (
// 	githubOAuthConfig *oauth2.Config
// )

// func init() {
// 	githubOAuthConfig = &oauth2.Config{
// 		RedirectURL:  "http://localhost:8080/auth/githubCallback",
// 		ClientID:     getGithubClientID(),
// 		ClientSecret: getGithubClientSecret(),
// 		Scopes:       []string{"read:user"},
// 		Endpoint:     github.Endpoint,
// 	}
// }

// // RegisterAuthEndpoints adds the endpoints for the auth package to a given
// // 	api client
// func RegisterAuthEndpoints(api *api.API) {
// 	const routePrefix = "/auth"

// 	api.Router.
// 		Methods("GET").
// 		Path(routePrefix + "/githubLogin").
// 		HandlerFunc(handleGithubLogin)

// 	api.Router.
// 		Methods("GET").
// 		Path(routePrefix + "/githubCallback").
// 		HandlerFunc(handleGithubCallback)
// }

// func handleGithubLogin(w http.ResponseWriter, r *http.Request) {
// 	// TODO: Add error handling
// 	oauthStateString, err := generateStateOAuthCookie(w)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}
// 	url := githubOAuthConfig.AuthCodeURL(oauthStateString)
// 	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
// }

// func handleGithubCallback(w http.ResponseWriter, r *http.Request) {
// 	content, err := getUserInfo(r)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
// 		return
// 	}
// 	fmt.Fprintf(w, "Content: %s\n", content)
// }

// func getUserInfo(r *http.Request) ([]byte, error) {
// 	oauthState, err := r.Cookie(oauthCookieName)
// 	if err != nil {
// 		return nil, err
// 	}

// 	state := r.FormValue("state")
// 	code := r.FormValue("code")

// 	if state != oauthState.Value {
// 		return nil, fmt.Errorf("invalid github oauth state")
// 	}

// 	token, err := githubOAuthConfig.Exchange(oauth2.NoContext, code)
// 	if err != nil {
// 		return nil, fmt.Errorf("code exchange failed: %s", err.Error())
// 	}

// 	client := &http.Client{}
// 	req, err := http.NewRequest("GET", githubUserURI, nil)
// 	if err != nil {
// 		return nil, err
// 	}
// 	req.Header.Add("Authorization", "token "+token.AccessToken)
// 	res, err := client.Do(req)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed getting user info: %s", err.Error())
// 	}
// 	defer res.Body.Close()

// 	contents, err := ioutil.ReadAll(res.Body)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed reading response body: %s", err.Error())
// 	}
// 	return contents, nil
// }
