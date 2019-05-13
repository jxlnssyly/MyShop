package handler

import (
	"context"

	"github.com/micro/go-log"

	example "MyShop/AdminLogin/proto/example"
	"github.com/astaxie/beego/orm"
	"MyShop/Router/models"
)

type Example struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Example) AdminLogin(ctx context.Context, req *example.Request, rsp *example.Response) error {

	userName := req.UserName
	password := req.Password

	o := orm.NewOrm()
	admin := models.AdminUser{Name: userName}
	o.Read(&admin,"Name")
	if password != admin.Password {
		rsp.Msg = "密码错误"
		rsp.Code = "3"
	} else {
		rsp.Msg = "登录成功"
		rsp.Code = "1"
	}

	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Example) Stream(ctx context.Context, req *example.StreamingRequest, stream example.Example_StreamStream) error {
	log.Logf("Received Example.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Logf("Responding: %d", i)
		if err := stream.Send(&example.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Example) PingPong(ctx context.Context, stream example.Example_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Logf("Got ping %v", req.Stroke)
		if err := stream.Send(&example.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
