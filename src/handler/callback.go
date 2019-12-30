package handler

import (
	"fmt"
	"net/http"
)

// CallbackHandler CallbackHandler
func CallbackHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("i am callback")
}
