package main

import (
	"fmt"

	"github.com/cndjp/Sugi275/serless_auth_backend/src/handler"
	"github.com/cndjp/Sugi275/serless_auth_backend/src/httprouter"
)

func main() {
	r := httprouter.MakeRouter(handler.CallbackHandler)

	fmt.Println("test")
}
