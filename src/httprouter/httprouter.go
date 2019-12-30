package httprouter

import (
	"net/http"

	"github.com/gorilla/mux"
)

// MakeRouter muxのroute設定用関数
func MakeRouter(callbackFunc, corsPrelightFunc func(w http.ResponseWriter, r *http.Request)) *mux.Router {
	r := mux.NewRouter()

	// route callback
	r.Path("/callback").
		Methods("GET").
		HandlerFunc(callbackFunc)

	return r
}
