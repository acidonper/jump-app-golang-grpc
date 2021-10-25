package grpcserver

import (
	"context"
	"log"
	"net"

	grpcclient "github.com/acidonper/jump-app-golang-grpc/internal/client"
	pb "github.com/acidonper/jump-app-protos/jump"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedJumpServiceServer
}

func (s *server) Jump(ctx context.Context, jump *pb.JumpReq) (*pb.Response, error) {
	log.Printf("gRPC Server: Request received %v", jump)

	// Add a new count
	jump.Count = jump.Count + 1
	log.Printf("gRPC Server: Steps count %v", jump.Count)

	// Evaluate jumps to send response or perform a jump 
	if len(jump.Jumps) == 0 || jump.Jumps[0] == "" {
		log.Printf("gRPC Server: Send response 200")
		return &pb.Response{Code: 200, Message: "/jump - Greetings from Golang gRPC!"}, nil
	} else {
		r, err := grpcclient.Jump(jump)
		if err != nil {
			log.Fatalf("Error local calling grpcclient from grpcserver - %v", err)
			return &pb.Response{Code: 500, Message: "/jump - Farewell from Python! Error Jumping"}, nil
		}
		log.Printf("gRPC Server: Response received %v", r)
		return r, nil
	}
}

func Start() {
	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterJumpServiceServer(s, &server{})

	reflection.Register(s)
	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
