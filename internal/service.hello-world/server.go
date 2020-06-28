package main

import (
	"net/http"

	"github.com/samstarling/twirp-test/internal/service.hello-world/handler"
	helloworldproto "github.com/samstarling/twirp-test/rpc/service.hello-world"
)

func main() {
	server := &handler.Server{}
	handler := helloworldproto.NewHelloWorldServer(server, nil)
	http.ListenAndServe(":8080", handler)
}
