package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	helloworldproto "github.com/samstarling/twirp-test/rpc/service.hello-world"
)

func main() {
	client := helloworldproto.NewHelloWorldProtobufClient("http://localhost:8080", &http.Client{})

	ctx := context.Background()

	rsp, err := client.Hello(ctx, &helloworldproto.HelloReq{
		Subject: "World",
	})
	if err != nil {
		fmt.Printf("An error occurred: %v", err)
		os.Exit(1)
	}
	fmt.Println(fmt.Sprintf("The reply was: %s", rsp.Text))

	rsp, err = client.Hello(ctx, &helloworldproto.HelloReq{
		Subject: "Potato",
	})
	if err != nil {
		fmt.Printf("An error occurred: %v", err)
		os.Exit(1)
	}
	fmt.Println(fmt.Sprintf("The reply was: %s", rsp.Text))
}
