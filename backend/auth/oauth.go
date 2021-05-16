package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/hujoseph99/typing/backend/common/api"
	"github.com/hujoseph99/typing/backend/db"
	"github.com/hujoseph99/typing/backend/secret"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

const githubUserURI = "https://api.github.com/user"

var (
	githubOAuthConfig *oauth2.Config
)

type GitHubCallback struct {
	Id    json.Number `json:"id"`
	Email string      `json:"email"`
	Name  string      `json:"name"`
	Login string      `json:"login"`
}

func init() {
	githubOAuthConfig = &oauth2.Config{
		RedirectURL:  secret.FrontendCallback,
		ClientID:     getGithubClientID(),
		ClientSecret: getGithubClientSecret(),
		Scopes:       []string{"read:user"},
		Endpoint:     github.Endpoint,
	}
}

func HandleGithubLogin(w http.ResponseWriter, r *http.Request) {
	// TODO: Add error handling
	oauthStateString, err := generateAndSetStateOAuthCookie(w)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	url := githubOAuthConfig.AuthCodeURL(oauthStateString)

	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func HandleGithubCallback(w http.ResponseWriter, r *http.Request) {
	content, err := getUserInfo(r)
	if err != nil {
		fmt.Println(err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	var c GitHubCallback
	err = json.Unmarshal(content, &c)
	if err != nil {
		fmt.Println(err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	ctx := context.TODO()
	existingUser, err := db.GetUserByID(ctx, string(c.Id), db.GithubID)

	// if user doesn't exist, create one
	if err != nil {
		newUser := db.NewUser("", "", c.Email, "", string(c.Id), "", time.Now())
		var nickname string
		if len(c.Name) != 0 {
			nickname = c.Name
		} else {
			nickname = c.Login
		}
		newPreferences := db.NewPreferences(newUser.ID, nickname)
		newProfile := db.NewProfile(newUser.ID, 0, 0, 0, 0, 0, 0)

		err = db.RegisterUser(ctx, newUser, newProfile, newPreferences)
		if err != nil {
			api.DefaultError(w, r, http.StatusInternalServerError, api.DefaultErrorMessage)
			return
		}
		existingUser = newUser
	}

	// generate jwt
	jwtPayload := newJwtPayload(existingUser.ID)
	accessToken, err := jwtPayload.convertToJwt()
	if err != nil {
		api.DefaultError(w, r, http.StatusInternalServerError, api.DefaultErrorMessage)
		return
	}

	// create session
	refreshToken, err := createSession(ctx, existingUser.ID, false)
	if err != nil {
		api.DefaultError(w, r, http.StatusInternalServerError, api.DefaultErrorMessage)
		return
	}

	res := map[string]string{
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
	}
	json.NewEncoder(w).Encode(res)
}

func getUserInfo(r *http.Request) ([]byte, error) {
	oauthState, err := r.Cookie(oauthCookieName)
	if err != nil {
		return nil, err
	}

	state := r.FormValue("state")
	code := r.FormValue("code")

	if state != oauthState.Value {
		return nil, fmt.Errorf("invalid github oauth state")
	}

	token, err := githubOAuthConfig.Exchange(context.TODO(), code)
	if err != nil {
		return nil, fmt.Errorf("code exchange failed: %s", err.Error())
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", githubUserURI, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "token "+token.AccessToken)
	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}
	defer res.Body.Close()

	contents, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed reading response body: %s", err.Error())
	}
	return contents, nil
}
