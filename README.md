# Jump App Go gRPC Repository

This repository includes a microservice based on Golang and gRPC that is a component develop for Jump App application. The idea of this microservice is implement an API based on gRPC that emulates the current features implemented in the original Jump App Golang microservice.

# Test Golang Code

- Run the gRPC server

```$bash
go run ./cmd/main.go
```

- Create a test/main.go test function

```$bash
package main

import (
	"log"

	grpcclient "github.com/acidonper/jump-app-golang-grpc/internal/client"
	pb "github.com/acidonper/jump-app-protos/jump"
)

func main() {
	log.Printf("Starting Server Process...")

	p := &pb.JumpReq{
		Count: 0,
		Message: "hola",
		Jumps: []string{"localhost:50051","localhost:50051",},
	}

	r, _ := grpcclient.Jump(p)
	log.Println(r)
}
```

- Execute the test

```$bash
go run ./test/main.go
...
2021/10/23 01:47:03 code:200  message:"/ - Greetings from Golang gRPC! | Jumps: 3"
```

# Test Docker image

- Build a container image

```$bash
podman build . -t jump-app-golang-grpc
```

- Run the new container image

```$bash
podman run -it -d -p 50051:50051 jump-app-golang-grpc
```

- Use a **grpcurl** command to test the app running (_* or create a main.go function as shown before_)

```$bash
grpcurl -plaintext -d '{"count": 0, "message": "hola", "jumps": ["localhost:50051","localhost:50051"]}' localhost:50051 jump.JumpService/Jump
{
  "code": 200,
  "message": "/ - Greetings from Golang gRPC! | Jumps: 3"
}
```

# Author

Asier Cidon

asier.cidon@gmail.com
