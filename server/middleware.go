package main

import (
	"context"
	"log"

	"google.golang.org/grpc/codes"

	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"

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

func authenticateFunc(ctx context.Context) (context.Context, error) {
	token, err := grpc_auth.AuthFromMD(ctx, "bearer")
	if err != nil {
		return nil, err
	}
	log.Println("recived token")
	// 固定文字列
	if token != "p@ssword" {
		return nil, grpc.Errorf(codes.Unauthenticated, "invalid authorization token: %v", token)
	}
	log.Println("success to authenticate")
	// grpc_ctxtags.Extract(ctx).Set("auth.sub", userClaimFromToken(tokenInfo))
	return ctx, nil
}
