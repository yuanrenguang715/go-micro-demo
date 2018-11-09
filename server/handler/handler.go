package handler

import (
	"context"
	"go-micro-demo/server/proto/message"
	"time"
	"utils/errors"
)

type MessSrv struct{}

var users []*message.UserInfo

func (*MessSrv) Send(ctx context.Context, req *message.SendRequest, rsp *message.SendReply) error {
	time.Sleep(time.Millisecond * 200) //mock timeout
	//log.Println("Received MessSrv.Send request", userInfo.Userid, userInfo.Username, userInfo.Intro)

	if req.Userinfo.Userid <= 0 {
		err := errors.New(int(message.MessSrvErrors_InvalidUserId), "userid必须大于0") //mock error
		return err
	}

	users = append(users, req.Userinfo)
	rsp.Result = true
	return nil
}

func (*MessSrv) Receive(ctx context.Context, req *message.Empty, rsp *message.ReceiveReply) error {
	//time.Sleep(time.Millisecond * 800) //mock timeout
	//log.Println("Received MessSrv.Receive request")

	rsp.Userinfo = users
	return nil
}
