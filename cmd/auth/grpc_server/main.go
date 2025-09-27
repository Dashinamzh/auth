package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/timestamppb"

	desc "github.com/Dashinamzh/auth/pkg/auth_v1"
	"github.com/brianvoe/gofakeit"
)

const grpc_port = 50051

type server struct {
	desc.UnimplementedAuthV1Server
}

// Get\
func (s *server) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	log.Printf("Auth id: %d")

	return &desc.GetResponse{
		User: &desc.User{
			Id: req.GetId(),
			Info: &desc.UserInfo{
				Name:  gofakeit.Name(),
				Email: gofakeit.Email(),
			},
			CreatedAt: timestamppb.New(gofakeit.Date()),
			UpdatedAt: timestamppb.New(gofakeit.Date()),
		},
	}, nil
}

/*func (s *server) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error){
	log.Printf("Created New User")
}*/

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpc_port))
	if err != nil {
		log.Fatalf("listen error")
	}
	s := grpc.NewServer() // создаем сервер
	reflection.Register(s)
	desc.RegisterAuthV1Server(s, &server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("is bad")
	}
}
