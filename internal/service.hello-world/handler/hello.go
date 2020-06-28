package handler

import (
	"context"
	"fmt"

	helloworldproto "github.com/samstarling/twirp-test/rpc/service.hello-world"
)

// Server is a server
type Server struct{}

// Hello returns a hello
func (s *Server) Hello(ctx context.Context, req *helloworldproto.HelloReq) (*helloworldproto.HelloResp, error) {
	return &helloworldproto.HelloResp{
		Text: fmt.Sprintf("Hello %s", req.Subject),
	}, nil
}
