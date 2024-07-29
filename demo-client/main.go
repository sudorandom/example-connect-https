package main

import (
	"context"
	"log"
	"net/http"

	"connectrpc.com/connect"

	"buf.build/gen/go/connectrpc/eliza/connectrpc/go/connectrpc/eliza/v1/elizav1connect"
	elizav1 "buf.build/gen/go/connectrpc/eliza/protocolbuffers/go/connectrpc/eliza/v1"
)

const url = "https://demo.connectrpc.com"

var reqBody = &elizav1.SayRequest{Sentence: "Hello World!"}

func main() {
	svcClient := elizav1connect.NewElizaServiceClient(http.DefaultClient, url, connect.WithGRPC())

	log.Println("connect: ", url)
	log.Println("send: ", reqBody)
	resp, err := svcClient.Say(context.Background(), connect.NewRequest(reqBody))
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	log.Println("recv: ", resp.Msg)
}
