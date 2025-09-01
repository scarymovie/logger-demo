package main

import (
	"context"
	"github.com/scarymovie/logger-demo/grpc-server/proto"

	"github.com/scarymovie/logger/slogx"
)

type greeterServer struct {
	proto.UnimplementedGreeterServer
}

func (s *greeterServer) SayHello(ctx context.Context, req *proto.HelloRequest) (*proto.HelloReply, error) {
	slogx.Info(ctx, "handling SayHello", slogx.String("name", req.Name))
	return &proto.HelloReply{Message: "Hello, " + req.Name + "!"}, nil
}
