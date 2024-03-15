package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/priyanshu360/remindnator/src/config"
	"github.com/priyanshu360/remindnator/src/sink/slackmessage"
	gcal "github.com/priyanshu360/remindnator/src/source/googlecalendar"
	"github.com/priyanshu360/remindnator/src/watcher"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func main() {
	go oauth()
	for config.CLIENT == nil {
		// fmt.Println(config.TOKEN)
	}
	fmt.Println(config.TOKEN)
	if err := gcal.Init(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("gcal init")
	cal, err := gcal.New("priyanshurajput360@gmail.com")
	if err != nil {
		log.Fatal(err)
	}

	cal.Fetch()

	slackmessage.Init()
	notifier := slackmessage.New("C04KQEF85D5", "* * * * *")
	cal.Subscribe(notifier)

	w := watcher.NewWatcher()
	w.Subscribe(cal)
	w.Run()
}

var (
	googleOauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:8080/callback",
		ClientID:     config.CLIENT_ID,     // from https://console.developers.google.com/project/<your-project-id>/apiui/credential
		ClientSecret: config.CLIENT_SECRET, // from https://console.developers.google.com/project/<your-project-id>/apiui/credential
		Scopes:       []string{"https://www.googleapis.com/auth/calendar"},
		Endpoint:     google.Endpoint,
	}
	// Some random string, random for each request
	oauthStateString = "random"
)

const htmlIndex = `<html><body>
<a href="/GoogleLogin">Log in with Google</a>
</body></html>
`

func oauth() {
	http.HandleFunc("/", handleMain)
	http.HandleFunc("/login", handleGoogleLogin)
	http.HandleFunc("/callback", handleGoogleCallback)
	fmt.Println(http.ListenAndServe(":8080", nil))
}

func handleMain(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, htmlIndex)
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
}
