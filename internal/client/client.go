package grpcclient

import (
	"context"
	"log"
	"time"

	pb "github.com/acidonper/jump-app-protos/jump"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "jump"
)

func Jump(jump *pb.JumpReq) (*pb.Response, error) {
	log.Printf("gRPC Client: Request received %v", jump)

	// Obtaining Jump Step
	a := jump.Jumps[:len(jump.Jumps)]
	addr := a[0]

	// Control the number of jumps
	if len(jump.Jumps) != 1 {
		jump.Jumps = jump.Jumps[:len(jump.Jumps)-1]
	} else {
		jump.Jumps[0] = ""
	}

	// Connect with gRPC server
	conn, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Error gRPC server connection: %v", err)
		return &pb.Response{Code: 400, Message: "/jump - Farewell from Golang gRPC"}, nil
	}
	defer conn.Close()

	// Perform Jump
	c := pb.NewJumpServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()
	r, err := c.Jump(ctx, jump)
	if err != nil {
		log.Fatalf("Error gRPC server negotiation: %v", err)
		return &pb.Response{Code: 400, Message: "/jump - Farewell from Golang gRPC"}, nil
	}

	// Response
	log.Printf("gRPC Client: Send received response %v", r)
	return &pb.Response{Code: 200, Message: "/jump - Greetings from Golang gRPC!"}, nil
}
