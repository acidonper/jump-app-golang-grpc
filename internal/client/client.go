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

func Jump(jump *pb.JumpReq) (*pb.JumpRes, error) {
	log.Printf("gRPC Client: Request received %v", jump.GetJump())

	// Obtaining Jump Step
	a := jump.Jump.Jumps[:len(jump.Jump.Jumps)]
	addr := a[0]

	// Control the number of jumps
	if len(jump.Jump.Jumps) != 1 {
		jump.Jump.Jumps = jump.Jump.Jumps[:len(jump.Jump.Jumps)-1]
	} else {
		jump.Jump.Jumps[0] = ""
	}

	// Connect with gRPC server
	conn, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Error gRPC server connection: %v", err)
		return &pb.JumpRes{Response: &pb.Response{Code: 400, Message: "/jump - Farewell from Golang gRPC" }}, nil
	}
	defer conn.Close()

	// Perform Jump
	c := pb.NewJumpServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()
	r, err := c.Jump(ctx, &pb.JumpReq{Jump: jump.Jump})
	if err != nil {
		log.Fatalf("Error gRPC server negotiation: %v", err)
		return &pb.JumpRes{Response: &pb.Response{Code: 400, Message: "/jump - Farewell from Golang gRPC" }}, nil
	}

	// Response
	log.Printf("gRPC Client: Send received response %v", r.GetResponse())
	return &pb.JumpRes{Response: &pb.Response{Code: 200, Message: "/ - Greetings from Golang gRPC!" }}, nil
}
