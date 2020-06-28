package main

import (
	"net/http"

	"github.com/samstarling/twirp-test/internal/service.addition/handler"
	additionproto "github.com/samstarling/twirp-test/rpc/service.addition"
)

func main() {
	server := &handler.Server{}
	handler := additionproto.NewAdditionServer(server, nil)
	http.ListenAndServe(":8080", handler)
}
