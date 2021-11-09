package grpcclient

import (
	"context"
	"log"
	"strconv"
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
	addr := jump.Jumps[len(jump.Jumps)-1]
	log.Printf("gRPC Client: Jump to %v", addr)

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
		return &pb.Response{Code: 400, Message: "/jump - Farewell from Golang gRPC | Jumps: " + strconv.FormatInt(int64(jump.Count), 10)}, nil
	}
	defer conn.Close()

	// Perform Jump
	c := pb.NewJumpServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()
	r, err := c.Jump(ctx, jump)
	if err != nil {
		log.Fatalf("Error gRPC server negotiation: %v", err)
		return &pb.Response{Code: 400, Message: "/jump - Farewell from Golang gRPC  | Jumps: " + strconv.FormatInt(int64(jump.Count), 10)}, nil
	}

	// Response
	log.Printf("gRPC Client: Send received response %v", r)
	return r, nil
}
