package main

import (
	"github.com/frhnfrnk/blog-platform-microservices/api-gateway/internal/graphql/user"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"

	"github.com/frhnfrnk/blog-platform-microservices/api-gateway/internal/handlers"
	"github.com/frhnfrnk/blog-platform-microservices/user-service/pb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to user service: %v", err)
	}
	defer conn.Close()

	userServiceClient := pb.NewUserServiceClient(conn)
	resolver := user.NewResolver(userServiceClient)

	http.Handle("/graphql/user", handlers.UserHandler(resolver))

	log.Println("Starting API Gateway on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
