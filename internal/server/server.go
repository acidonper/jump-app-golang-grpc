package grpcserver

import (
	"context"
	"log"
	"net"

	grpcclient "github.com/acidonper/jump-app-golang-grpc/internal/client"
	pb "github.com/acidonper/jump-app-protos/jump"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedJumpServiceServer
}

func (s *server) Jump(ctx context.Context, in *pb.JumpReq) (*pb.JumpRes, error) {
	log.Printf("gRPC Server: Request received %v", in.GetJump())

	// Add a new count
	in.Jump.Count = in.Jump.Count + 1
	log.Printf("gRPC Server: Steps count %v", in.Jump.Count)

	// Evaluate jumps to send response or perform a jump 
	if len(in.Jump.Jumps) == 0 || in.Jump.Jumps[0] == "" {
		log.Printf("gRPC Server: Send response 200")
		return &pb.JumpRes{Response: &pb.Response{Code: 200, Message: "/ - Greetings from Golang gRPC!",}}, nil
	} else {
		r, err := grpcclient.Jump(in)
		if err != nil {
			log.Fatalf("Error local calling grpcclient from grpcserver - %v", err)
			return &pb.JumpRes{Response: &pb.Response{Code: 500, Message: "/jump - Farewell from Python! Error Jumping" }}, nil
		}
		log.Printf("gRPC Server: Response received %v", r.GetResponse())
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

	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
