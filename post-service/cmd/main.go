package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/frhnfrnk/blog-platform-microservices/post-service/internal/handlers"
	"github.com/frhnfrnk/blog-platform-microservices/post-service/internal/models"
	"github.com/frhnfrnk/blog-platform-microservices/post-service/internal/repositories"
	"github.com/frhnfrnk/blog-platform-microservices/post-service/internal/services"
	"github.com/frhnfrnk/blog-platform-microservices/post-service/pb"
	userPb "github.com/frhnfrnk/blog-platform-microservices/user-service/pb"
	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Load environment variables
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	grpcPort := os.Getenv("GRPC_PORT")
	redisAddr := os.Getenv("REDIS_ADDR")
	redisPassword := os.Getenv("REDIS_PASSWORD")
	userServiceAddr := os.Getenv("USER_SERVICE_ADDR")


	dsn := "host=" + dbHost + " user=" + dbUser + " password=" + dbPassword + " dbname=" + dbName + " port=" + dbPort + " sslmode=disable"
	fmt.Println(dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	if err := db.AutoMigrate(&models.Post{}); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	redisClient := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: redisPassword,
		DB:       0, // Use default DB
	})

	// Connect to user-service
	userConn, err := grpc.NewClient(userServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to user service: %v", err)
	}
	defer userConn.Close()
	userClient := userPb.NewUserServiceClient(userConn)

	postRepo := repositories.NewPostRepository(db)
	postService := services.NewPostService(postRepo, redisClient, userClient)
	postHandler := handlers.NewPostHandler(postService)

	lis, err := net.Listen("tcp", ":"+grpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterPostServiceServer(grpcServer, postHandler)

	log.Println("Starting Post Service on port :50052...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
