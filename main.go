package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/vivekbnwork/bz-backend/bz-main/driver"
	grpcserver "github.com/vivekbnwork/bz-backend/bz-main/grpc/server"
	"github.com/vivekbnwork/bz-backend/bz-main/routes"
	"google.golang.org/grpc"
)

func init() {
	loadEnv()
}

var ctx = context.Background()

func loadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func connectDatabase() *driver.DB {
	dsn := os.Getenv("DATABASE_URL")
	db, err := driver.ConnectSQL(dsn)
	if err != nil {
		logrus.Fatal(err)
	}
	return db
}

// Your gRPC server function
func startGRPCServer(wg *sync.WaitGroup, db *driver.DB) {
	defer wg.Done() // Decrement the counter when the goroutine completes.

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	grpcserver.RegisterUserService(s, db)
	log.Println("gRPC Server is running on port 50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// Your HTTP server function
func startHTTPServer(wg *sync.WaitGroup, db *driver.DB) {
	defer wg.Done() // Decrement the counter when the goroutine completes.

	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	// Get Redis host and port from environment variables
	redisHost := os.Getenv("REDIS_HOST")
	redisPortStr := os.Getenv("REDIS_PORT")
	redisPassword := os.Getenv("REDIS_PASSWORD")

	// Convert the port string to an integer
	redisPort, err := strconv.Atoi(redisPortStr)
	if err != nil {
		fmt.Printf("Error converting Redis port to integer: %v\n", err)
		return
	}

	// Create a Redis client with the specified address
	rdb := redis.NewClient(&redis.Options{
		//Addr:     "localhost:6379",
		Addr:     fmt.Sprintf("%s:%d", redisHost, redisPort),
		Password: redisPassword, // No password
		DB:       0,             // Default DB
	})

	// Check the connection by trying to ping the server
	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		fmt.Printf("Error pinging Redis server: %v\n", err)
		return
	}
	fmt.Printf("Redis Ping Response: %s\n", pong)

	env := os.Getenv("ENV")

	// Create an AWS session
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(os.Getenv("AWS_REGION")),
		Credentials: credentials.NewStaticCredentials(os.Getenv("AWS_ACCESS_KEY_FOR_CLIENT"), os.Getenv("AWS_ACCESS_KEY_FOR_CLIENT_SECRET"), ""),
	})
	if err != nil {
		log.Fatalf("Failed to create AWS session: %v", err)
	}

	// Create S3 service client
	s3Client := s3.New(sess)

	switch env {
	case "PROD":
		logrus.SetLevel(logrus.InfoLevel) // Only log INFO and above in production
	case "DEV":
		logrus.SetLevel(logrus.DebugLevel) // Log everything in development
	default:
		logrus.SetLevel(logrus.WarnLevel) // Default to WARN level
	}

	logrus.SetFormatter(&logrus.JSONFormatter{})

	// Initialize routes
	routes.InitializeRoutes(router, db, rdb, s3Client) // Pass the db to InitializeRoutes

	port := os.Getenv("PORT")
	if port == "" {
		logrus.Fatal("PORT is not set in .env file")
	}

	log.Println("HTTP Server is running on port", port)
	router.Run(":" + port)
}

func main() {
	var wg sync.WaitGroup

	wg.Add(2) // Add 2 to the wait group counter.

	db := connectDatabase() // Use the local function directly without any prefix
	defer db.Pool.Close()   // Defer the closing of the database connection pool to when the app exits

	go startGRPCServer(&wg, db) // Start the gRPC server in a new goroutine.
	go startHTTPServer(&wg, db) // Start the HTTP server in a new goroutine.

	wg.Wait() // Wait for all goroutines to complete.
}
