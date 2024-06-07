package api

import (
	"backend/cookie"
	"backend/internal/net"
	"net/http"
)

// UserLogout logs out a user by invalidating their session.
//
// @Summary Logs out a user
// @Description This endpoint logs out a user by invalidating their session.
// @ID userLogout
// @Produce json
// @Success 200 {object} net.Msg "Successful logout"
// @Router /api/user/logout [get]
func UserLogout(w http.ResponseWriter, r *http.Request) {
	sessions, _ := cookie.Store.Get(r, "session-name")
	sessions.Options.MaxAge = -1
	sessions.Save(r, w)
	net.Respond(w, http.StatusOK, net.Msg{
		"response": "logout",
	})
}
