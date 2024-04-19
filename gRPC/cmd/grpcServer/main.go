package main

import (
	"database/sql"
	"net"

	"github.com/FelpsCorrea/GoExpertPostgraduation/gRPC/internal/database"
	"github.com/FelpsCorrea/GoExpertPostgraduation/gRPC/internal/pb"
	"github.com/FelpsCorrea/GoExpertPostgraduation/gRPC/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	_ "github.com/mattn/go-sqlite3"
)

// This code sets up the main function for the gRPC server.
// It initializes a SQLite database connection, creates a category service,
// and registers the service with the gRPC server.

func main() {
	// Open a connection to the SQLite database
	db, err := sql.Open("sqlite3", "./db.sqlite")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Create a category database instance
	categoryDb := database.NewCategoryDb(db)

	// Create a category service instance
	categoryService := service.NewCategoryService(*categoryDb)

	// Create a new gRPC server
	grpcServer := grpc.NewServer()

	reflection.Register(grpcServer)

	// Register the category service with the gRPC server
	pb.RegisterCategoryServiceServer(grpcServer, categoryService)

	lis, err := net.Listen("tcp", ":50051")

	if err != nil {
		panic(err)
	}

	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
