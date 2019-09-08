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
