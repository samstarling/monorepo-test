package addition

//go:generate protoc --proto_path=$GOPATH/src --twirp_out=$GOPATH/src --go_out=$GOPATH/src github.com/samstarling/twirp-test/rpc/service.addition/service.proto
