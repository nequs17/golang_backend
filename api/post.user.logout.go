package api

import (
	"backend/cookie"
	"backend/internal/net"
	"net/http"
)

func UserLogout(w http.ResponseWriter, r *http.Request) {
	sessions, _ := cookie.Store.Get(r, "session-name")
	sessions.Values["authenticated"] = false
	sessions.Save(r, w)
	net.Respond(w, http.StatusOK, net.Msg{
		"response": "logout",
	})
}
