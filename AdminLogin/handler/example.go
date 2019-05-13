package handler

import (
	"context"

	"github.com/micro/go-log"

	example "MyShop/AdminLogin/proto/example"
	"github.com/astaxie/beego"
)

type Example struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Example) AdminLogin(ctx context.Context, req *example.Request, rsp *example.Response) error {
	//log.Log("Received Example.Call request")
	beego.Info(req.UserName)
	beego.Info(req.Password)
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
