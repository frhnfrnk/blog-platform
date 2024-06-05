package main

import (
	"github.com/frhnfrnk/blog-platform-microservices/post-service/internal/handlers"
	"github.com/frhnfrnk/blog-platform-microservices/post-service/internal/models"
	"github.com/frhnfrnk/blog-platform-microservices/post-service/internal/repositories"
	"github.com/frhnfrnk/blog-platform-microservices/post-service/internal/services"
	"github.com/frhnfrnk/blog-platform-microservices/post-service/pb"
	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net"
	"os"
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

	dsn := "host=" + dbHost + " user=" + dbUser + " password=" + dbPassword + " dbname=" + dbName + " port=" + dbPort + " sslmode=disable"

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

	postRepo := repositories.NewPostRepository(db)
	postService := services.NewPostService(postRepo, redisClient)
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
