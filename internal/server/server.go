package grpcserver

import (
	"context"
	"log"
	"net"

	pb "github.com/acidonper/jump-app-protos/jump"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedJumpServiceServer
}

func (s *server) PerformJump(ctx context.Context, in *pb.PerformJumpReq) (*pb.JumpRes, error) {
	log.Printf("POST: %v", in.GetJump())
	return &pb.JumpRes{Response: &pb.JumpRes_JumpResponse{Code: 123, Message: "pepe" }}, nil
}

func (s *server) FinalJump(ctx context.Context, in *pb.FinalJumpReq) (*pb.JumpRes, error) {
	log.Println("GET")
	return &pb.JumpRes{Response: &pb.JumpRes_JumpResponse{Code: 123, Message: "pepe" }}, nil
}

func Start() {
	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterJumpServiceServer(s, &server{})

	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
