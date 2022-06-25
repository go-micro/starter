package handler

import (
	"context"

	pb "github.com/go-micro/starter/proto"
)

type Example struct{}

func New() *Example {
	return &Example{}
}

func (e *Example) Call(ctx context.Context, req *pb.CallRequest, rsp *pb.CallResponse) error {
	rsp.Message = "Hello " + req.Name
	return nil
}
