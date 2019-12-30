package main

import (
	"fmt"

	"github.com/cndjp/Sugi275/src/handler"
	"github.com/cndjp/Sugi275/src/httprouter"
)

func main() {
	r := httprouter.MakeRouter(handler.CallbackHandler)

	fmt.Println("test")
}
