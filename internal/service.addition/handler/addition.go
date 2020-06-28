package handler

import (
	"context"

	additionproto "github.com/samstarling/twirp-test/rpc/service.addition"
)

// Server is a server
type Server struct{}

// Add will add some numbers
func (s *Server) Add(ctx context.Context, req *additionproto.AddRequest) (*additionproto.AddResponse, error) {
	return &additionproto.AddResponse{
		Result: req.First + req.Second,
	}, nil
}
