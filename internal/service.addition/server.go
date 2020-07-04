package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/samstarling/twirp-test/internal/service.addition/handler"
	additionproto "github.com/samstarling/twirp-test/rpc/service.addition"
)

func main() {
	server := &handler.Server{}
	handler := additionproto.NewAdditionServer(server, nil)
	port := "8080"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}
	http.ListenAndServe(fmt.Sprintf(":%s", port), handler)
	fmt.Println(fmt.Sprintf("Listening on port :%s", port))
}
