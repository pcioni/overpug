package session

import (
	"net/http"

	"github.com/gorilla/sessions"
)

var (
	Store       *sessions.CookieStore
	SessionName string
)

type SessionInfo struct {
	SecretKey string `json:"SecretKey"`
	Name      string `json:"SessionName"`
}

func ConfigureSessions(s SessionInfo) {
	sessions.NewCookieStore([]byte(s.SecretKey))
	SessionName = s.Name
}

func GetSession(r *http.Request) *sessions.Session {
	session, _ := Store.Get(r, SessionName)
	return session
}
