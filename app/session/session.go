package session

import (
	"log"
	"net/http"

	"github.com/gorilla/sessions"
)

var (
	store       *sessions.CookieStore
	sessionName string
)

type SessionInfo struct {
	SecretKey string `json:"SecretKey"`
	Name      string `json:"Name"`
}

func ConfigureSessions(s SessionInfo) {
	store = sessions.NewCookieStore([]byte(s.SecretKey))
	sessionName = s.Name
}

func GetSession(r *http.Request) *sessions.Session {
	sess, err := store.Get(r, sessionName)
	if err != nil {
		log.Println(err)
	}
	return sess
}
