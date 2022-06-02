package main

import (
	"log"

	"github.com/go-micro/starter/handler"
	pb "github.com/go-micro/starter/proto"
	"go-micro.dev/v4"
)

func main() {
	service := micro.NewService(
		micro.Name("starter"),
	)

	service.Init()

	pb.RegisterExampleHandler(service.Server(), handler.New())

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
