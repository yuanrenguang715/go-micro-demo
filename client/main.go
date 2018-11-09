package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go-micro-demo/client/cli"
	"go-micro-demo/server/proto/message"
)

const timeout = 5 //grpc client超时(秒)

func main() {
	cli.Init()

	// Set arbitrary headers in context
	/* 	ctx := metadata.NewContext(context.Background(), map[string]string{
		"X-User-Id": "john",
		"X-From-Id": "script",
	}) */

	// Set timeout(call new instance)
	ctx := context.Background()
	c, cancel := context.WithTimeout(ctx, time.Duration(timeout)*time.Second)
	defer cancel()

	// call Send
	rsp, err := cli.MessClient.Send(c, &message.SendRequest{
		&message.UserInfo{
			Userid:   1,
			Username: "nikolas",
			Intro:    "i am ok",
		},
	})
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(rsp.Result)

	// call Receive
	rs, err := cli.MessClient.Receive(c, &message.Empty{})
	if err != nil {
		log.Fatalln(err)
	}

	for _, item := range rs.Userinfo {
		fmt.Println(item.Userid, item.Username, item.Intro)
	}
}
