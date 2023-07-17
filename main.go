package main

import (
	"fmt"
	"log"
	"mygoproject/api"
	"mygoproject/repository"
	rpc_servers "mygoproject/rpc/servers"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	mongoConn := repository.OpenConnection()
	fmt.Println("Successfully connected to MongoDB")
	NewGrpcServer()
	fmt.Println("Successfully started gRPC server")
	err := NewServer(mongoConn)
	if err != nil {
		log.Fatal("Could not start server")
	}
	fmt.Println("Hello, World!")
}

func NewGrpcServer() {
	mongoConn := repository.OpenConnection()
	rpc_servers.InitializeGRPCServer(mongoConn)
}

func NewServer(mongoConn *mongo.Client) error {

	router := mux.NewRouter()

	fmt.Println("Initializing server...")
	server := api.InitializeServer(mongoConn)
	router.HandleFunc("/create-user", server.CreateUserHandler).Methods("GET")

	fmt.Println("Server initialized successfully")
	os.Getenv("PORT")
	err := http.ListenAndServe("PORT", router)
	if err != nil {
		return err
	}

	return nil
}
