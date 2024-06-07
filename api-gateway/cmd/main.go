package main

import (
	"github.com/frhnfrnk/blog-platform-microservices/api-gateway/internal/graphql/comment"
	"github.com/frhnfrnk/blog-platform-microservices/api-gateway/internal/graphql/post"
	"github.com/frhnfrnk/blog-platform-microservices/api-gateway/internal/graphql/user"
	"github.com/frhnfrnk/blog-platform-microservices/api-gateway/internal/handlers"
	"github.com/joho/godotenv"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
	"os"

	commentPb "github.com/frhnfrnk/blog-platform-microservices/comment-service/pb"
	postPb "github.com/frhnfrnk/blog-platform-microservices/post-service/pb"
	userPb "github.com/frhnfrnk/blog-platform-microservices/user-service/pb"
	"google.golang.org/grpc"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	userPort := os.Getenv("USER_GRPC_PORT")
	postPort := os.Getenv("POST_GRPC_PORT")
	commentPort := os.Getenv("COMMENT_GRPC_PORT")

	userConn, err := grpc.NewClient(userPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to user service: %v", err)
	}
	defer userConn.Close()

	postConn, err := grpc.NewClient(postPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to post service: %v", err)
	}
	defer postConn.Close()

	commentConn, err := grpc.NewClient(commentPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to comment service: %v", err)
	}
	defer commentConn.Close()

	userServiceClient := userPb.NewUserServiceClient(userConn)
	postServiceClient := postPb.NewPostServiceClient(postConn)
	commentServiceClient := commentPb.NewCommentServiceClient(commentConn)
	userResolver := user.NewResolver(userServiceClient)
	postResolver := post.NewResolver(postServiceClient)
	commentResolver := comment.NewResolver(commentServiceClient)

	http.Handle("/graphql/user", handlers.UserHandler(userResolver))
	http.Handle("/graphql/post", handlers.PostHandler(postResolver))
	http.Handle("/graphql/comment", handlers.CommentHandler(commentResolver))

	log.Println("Starting API Gateway on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
