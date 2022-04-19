package handler

import (
	pb "github.com/go-micro/starter/proto"
)

type Example struct{}

func (e *Example) Call(ctx context.Context, req *pb.CallRequest, rsp *pb.CallResponse) error {
	rsp.Message = "Hello " + req.Name
	return nil
}
