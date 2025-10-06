package main

import (
	"context"
	"flag"
	"log"
	"net"

	"github.com/Dashinamzh/auth/intenal/config"
	"github.com/Dashinamzh/auth/intenal/config/env"
	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	flag.Parse()
	ctx := context.Background()

	err := config.Load("configPath")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	grpcConfig, err := env.NewGRPCConfig()
	if err != nil {
		log.Fatalf("failed to get grpc config: %v", err)
	}

	pgConfig, err := env.NewPGConfig()
	if err != nil {
		log.Fatalf("failed to get postgres config: %v", err)
	}

	lis, err := net.Listen("tcp", grpcConfig.Address())
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	pool, err := pgxpool.Connect(ctx, pgConfig.DSN())
	if err != nil {
		log.Fatalf("failed to connect database %f", err)
	}
	defer pool.Close()

	s := grpc.NewServer()
	reflection.Register(s)

	log.Printf("server listening at: %d", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve %v", err)
	}
}
