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
	defaultName = "world"
)

func PerformJump() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewJumpServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.FinalJump(ctx, &pb.FinalJumpReq{})
	if err != nil {
		log.Fatalf("GET - could not greet: %v", err)
	}
	log.Printf("GET - Greeting: %s", r.GetResponse())

	p := &pb.Jump{
		Message: "hola",
		LastPath: "/jump",
		JumpPath: "/jump",
		Jumps: []*pb.Jump_JumpStep{{
			Jump: "localhost",
		}},
	}
	r, err = c.PerformJump(ctx, &pb.PerformJumpReq{Jump: p})
	if err != nil {
		log.Fatalf("POST - could not greet: %v", err)
	}
	log.Printf("POST - Greeting: %s", r.GetResponse())
}
