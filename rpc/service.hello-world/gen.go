package helloworld

//go:generate protoc --proto_path=$GOPATH/src --twirp_out=$GOPATH/src --go_out=$GOPATH/src github.com/samstarling/twirp-test/rpc/service.hello-world/service.proto
