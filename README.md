# Jump App Go gRPC Repository

This repository includes a microservice based on Golang and gRPC that is a component develop for Jump App application. The idea of this microservice is implement an API based on gRPC that emulates the current features implemented in the original Jump App Golang microservice.

# Test

- Create a main.go function

```$bash
package main

import (
	"log"
	grpcclient "github.com/acidonper/jump-app-golang-grpc/internal/client"
)

func main() {
	log.Printf("Starting Server Process...")
	grpcclient.PerformJump()
}
```

- Run the gRPC server

```$bash
go run ./cmd/main.go
```

- Execute the test

```$bash
go run ./test/main.go
```

# Author

Asier Cidon

asier.cidon@gmail.com
