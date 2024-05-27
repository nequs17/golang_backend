package cookie

import (
	"encoding/gob"

	"github.com/gorilla/sessions"
)

var Store = sessions.NewCookieStore([]byte("session-key"))

func init() {
	gob.Register(map[string]interface{}{})
}
