package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/priyanshu360/remindnator/src/config"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/tasks/v1"
)

var (
	googleOauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:8080/callback",
		ClientID:     config.CLIENT_ID,
		ClientSecret: config.CLIENT_SECRET,
		Scopes:       []string{calendar.CalendarScope, tasks.TasksScope},
		Endpoint:     google.Endpoint,
	}
	// Some random string, random for each request
	oauthStateString = "random"
)

const htmlIndex = `<html><body>
<a href="/login">Log in with Google</a>
</body></html>
`

var server *http.Server

func oauth() {
	// TODO #4 : Write oauth token to file for reuse
	http.HandleFunc("/", handleMain)
	http.HandleFunc("/login", handleGoogleLogin)
	http.HandleFunc("/callback", handleGoogleCallback)
	server = &http.Server{Addr: ":8080"}
	fmt.Println(server.ListenAndServe())
}

func handleMain(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, htmlIndex)
}

func handleGoogleLogin(w http.ResponseWriter, r *http.Request) {
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

	fmt.Println("token -> ", token) // Print the token for testing purposes (

	config.CLIENT = oauth2.NewClient(context.Background(), oauth2.StaticTokenSource(token))
	fmt.Println("client -> ", config.CLIENT)
	server.Shutdown(context.Background())
}
