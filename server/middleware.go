package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
)

func doNothingIntercepter() grpc.UnaryServerInterceptor {
	return func(tx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		log.Println("do nothing")
		return handler(tx, req)
	}
}

func doNothingStreamIntercepter() grpc.StreamServerInterceptor {
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		log.Println("do nothing stream")
		return handler(srv, ss)
	}
}
