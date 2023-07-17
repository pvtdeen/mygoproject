package rpc_servers

import (
	"context"
	"fmt"
	"mygoproject/repository"
	"mygoproject/rpc/servers/users"
	"net"

	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

type Server struct {
	db *repository.MongoDatabase
	users.UnimplementedGetUsersServiceServer
}

func (s *Server) GetUsers(ctx context.Context, req *users.GetUsersRequest) (*users.GetUsersResponse, error) {
	dbUsers, err := s.db.GetUsersByIds(ctx, req.Ids)
	if err != nil {
		return nil, err
	}
	var usersDto users.GetUsersResponse
	for _, user := range dbUsers {
		userDto := users.GetUsersDto{
			Id:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		}
		usersDto.Users = append(usersDto.Users, &userDto)
	}
	return &usersDto, nil
}

func InitializeGRPCServer(conn *mongo.Client) {
	s := grpc.NewServer()
	server := Server{
		db: repository.NewDB("MyDatabase", conn),
	}

	users.RegisterGetUsersServiceServer(s, &server)

	lis, err := net.Listen("tcp", ":6868")
	if err != nil {
		fmt.Println("Error while listening: ", err)
	}

	go func() {
		s.Serve(lis)
	}()
}
