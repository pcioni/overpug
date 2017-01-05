package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/albshin/overpug/app/session"

	"github.com/julienschmidt/httprouter"
	"golang.org/x/oauth2"
)

var (
	oauthConf = &oauth2.Config{
		ClientID:     "",
		ClientSecret: "",
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://us.battle.net/oauth/authorize",
			TokenURL: "https://us.battle.net/oauth/token",
		},
		RedirectURL: "https://localhost:3000/auth",
	}
	oauthStateString = "randomstringxd"
)

type OAuth2Info struct {
	ClientID     string
	ClientSecret string
}

// Handle login using Battle.net account
func LoginGET(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	url := oauthConf.AuthCodeURL(oauthStateString, oauth2.AccessTypeOnline)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func ConfigureOAuth2(o OAuth2Info) {
	oauthConf.ClientID = o.ClientID
	oauthConf.ClientSecret = o.ClientSecret
}

func AuthGET(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	type User struct {
		ID        int    `json:"id"`
		Battletag string `json:"battletag"`
	}

	state := r.FormValue("state")
	if state != oauthStateString {
		fmt.Printf("invalid oauth state, expected '%s', got '%s'\n", oauthStateString, state)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	code := r.FormValue("code")
	// See: https://github.com/golang/oauth2/pull/74
	token, err := oauthConf.Exchange(oauth2.NoContext, code)
	if err != nil {
		fmt.Printf("oauthConf.Exchange() failed with '%s'\n", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	oauthClient := oauthConf.Client(oauth2.NoContext, token)
	req, _ := http.NewRequest("GET", "https://us.api.battle.net/account/user", nil)
	req.Header.Add("User-Agent", "overmatch/0.1")
	resp, err := oauthClient.Do(req)

	if err != nil {
		log.Fatal(err)
	} else {
		defer resp.Body.Close()
		var user User
		decoder := json.NewDecoder(resp.Body)
		err = decoder.Decode(&user)
		if err != nil {
			log.Fatal(err)
		}

		sess := session.GetSession(r)
		sess.Values["id"] = user.ID
		sess.Values["battletag"] = user.Battletag
		sess.Save(r, w)

		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
}
