package main

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/priyanshu360/remindnator/config"
	"github.com/priyanshu360/remindnator/util"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/tasks/v1"
)

var (
	googleOauthConfig = &oauth2.Config{
		RedirectURL: "http://localhost:8080/callback",
		Scopes:      []string{calendar.CalendarScope, tasks.TasksScope},
		Endpoint:    google.Endpoint,
	}
	// Some random string, random for each request
	oauthStateString, _ = generateRandomString(8)
)

const htmlIndex = `<html><body>
<a href="/login">Log in with Google</a>
</body></html>
`

var server *http.Server

func generateRandomString(length int) (string, error) {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes)[:length], nil
}

func oauth() {
	if err := util.LoadTokenFromFile(); err == nil {
		return
	}

	http.HandleFunc("/", handleMain)
	http.HandleFunc("/login", handleGoogleLogin)
	http.HandleFunc("/callback", handleGoogleCallback)
	server = &http.Server{Addr: ":8080"}
	fmt.Println(server.ListenAndServe())
}

func saveTokenToFile(token *oauth2.Token) {
	file, err := json.MarshalIndent(token, "", " ")
	if err != nil {
		fmt.Println("Error marshalling token:", err)
		return
	}

	err = os.WriteFile("token.json", file, 0644)
	if err != nil {
		fmt.Println("Error writing token to file:", err)
		return
	}

	fmt.Println("Token saved to file.")
}

func handleMain(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, htmlIndex)
}

func handleGoogleLogin(w http.ResponseWriter, r *http.Request) {
	googleOauthConfig.ClientID = config.CLIENT_ID
	googleOauthConfig.ClientSecret = config.CLIENT_SECRET
	url := googleOauthConfig.AuthCodeURL(oauthStateString)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func handleGoogleCallback(w http.ResponseWriter, r *http.Request) {
	fmt.Println("callback")
	state := r.FormValue("state")
	if state != oauthStateString {
		fmt.Printf("invalid oauth state, expected '%s', got '%s'\n", oauthStateString, state)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	code := r.FormValue("code")
	token, err := googleOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		fmt.Printf("oauthConf.Exchange() failed with '%s'\n", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	saveTokenToFile(token)

	config.CLIENT = oauth2.NewClient(context.Background(), oauth2.StaticTokenSource(token))
	fmt.Println("Token exchanged successfully.")
	server.Shutdown(context.Background())
}
