package main

import (
	"github.com/frhnfrnk/blog-platform-microservices/comment-service/internal/models"
	"github.com/go-redis/redis/v8"
	"log"
	"net"
	"os"

	"github.com/frhnfrnk/blog-platform-microservices/comment-service/internal/handlers"
	"github.com/frhnfrnk/blog-platform-microservices/comment-service/internal/repositories"
	"github.com/frhnfrnk/blog-platform-microservices/comment-service/internal/services"
	"github.com/frhnfrnk/blog-platform-microservices/comment-service/pb"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	grpcPort := os.Getenv("GRPC_PORT")
	redisAddr := os.Getenv("REDIS_ADDR")
	redisPassword := os.Getenv("REDIS_PASSWORD")

	// Set up database connection
	dsn := "host=" + dbHost + " user=" + dbUser + " password=" + dbPassword + " dbname=" + dbName + " port=" + dbPort + " sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// Migrate the schema
	if err := db.AutoMigrate(&models.Comment{}); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	redisClient := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: redisPassword,
		DB:       0, // Use default DB
	})

	// Set up repositories and services
	commentRepo := repositories.NewCommentRepository(db)
	commentService := services.NewCommentService(commentRepo, redisClient)
	commentHandler := handlers.NewCommentHandler(commentService)

	// Set up gRPC server
	grpcServer := grpc.NewServer()
	pb.RegisterCommentServiceServer(grpcServer, commentHandler)

	// Start listening for incoming gRPC requests
	listener, err := net.Listen("tcp", ":"+grpcPort)
	if err != nil {
		log.Fatalf("failed to listen on port 50053: %v", err)
	}

	log.Println("Starting Comment Service on port 50053...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve gRPC server: %v", err)
	}
}
