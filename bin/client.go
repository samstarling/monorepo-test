package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	additionproto "github.com/samstarling/twirp-test/rpc/service.addition"
	helloworldproto "github.com/samstarling/twirp-test/rpc/service.hello-world"
)

func main() {
	addClient := additionproto.NewAdditionProtobufClient("http://127.0.0.1:8081", &http.Client{})
	client := helloworldproto.NewHelloWorldProtobufClient("http://127.0.0.1:8082", &http.Client{})

	ctx := context.Background()

	rsp, err := client.Hello(ctx, &helloworldproto.HelloReq{Subject: "World"})
	if err != nil {
		fmt.Printf("An error occurred: %v", err)
		os.Exit(1)
	}
	fmt.Println(fmt.Sprintf("The reply was: %s", rsp.Text))

	rsp, err = client.Hello(ctx, &helloworldproto.HelloReq{Subject: "Potato"})
	if err != nil {
		fmt.Printf("An error occurred: %v", err)
		os.Exit(1)
	}
	fmt.Println(fmt.Sprintf("The reply was: %s", rsp.Text))

	rspTwo, err := addClient.Add(ctx, &additionproto.AddRequest{
		First:  1,
		Second: 45,
	})
	if err != nil {
		fmt.Printf("An error occurred: %v", err)
		os.Exit(1)
	}
	fmt.Println(fmt.Sprintf("The result was: %s", rspTwo))
}
