package auth

import (
	"fmt"
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

var (
	githubOAuthConfig *oauth2.Config
)

func init() {
	githubOAuthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:8080/auth/githubCalback",
		ClientID:     getGithubClientID(),
		ClientSecret: getGithubClientSecret(),
		Scopes:       []string{"read:user"},
		Endpoint:     github.Endpoint,
	}
}

func handleGithubLogin(w http.ResponseWriter, r *http.Request) {
	// TODO: Add error handling
	oauthStateString, _ := generateEncryptedOAuthStateString()
	url := githubOAuthConfig.AuthCodeURL(oauthStateString)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func handleGithubCallback(w http.ResponseWriter, r *http.Request) {
	content, err := getUserInfo(r.FormValue("state"), r.FormValue("code"))
	if err != nil {
		fmt.Println(err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	fmt.Fprintf(w, "Content: %s\n", content)
}

func getUserInfo(state string, code string) ([]byte, error) {
	validState, err := verifyOAuthStateString(state)
	if err != nil {
		return nil, err
	}
	if !validState {
		return nil, fmt.Errorf("invalid oauth state")
	}
	// token, err := githubOAuthConfig.Exchange(oauth2.NoContext, code)
	// if err != nil {
	// 	return nil, fmt.Errorf("code exchange failed: %s", err.Error())
	// }
	// response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	// if err != nil {
	// 	return nil, fmt.Errorf("failed getting user info: %s", err.Error())
	// }
	// defer response.Body.Close()
	// contents, err := ioutil.ReadAll(response.Body)
	// if err != nil {
	// 	return nil, fmt.Errorf("failed reading response body: %s", err.Error())
	// }
	return nil, nil
}
