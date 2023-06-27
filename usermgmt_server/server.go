package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"

	pb "github.com/amankr1279/grpcEx/usermgmt"
	"google.golang.org/grpc"
)

const port = ":50051"

type UserManagementServer struct {
	pb.UnimplementedUserManagementServer
	usersList *pb.UserList // array of user pointers
}

func NewUserManagementServer() *UserManagementServer {
	return &UserManagementServer{
		usersList: &pb.UserList{},
	}
}
func (s *UserManagementServer) CreateNewUser(ctx context.Context, in *pb.NewUser) (*pb.User, error) {
	log.Printf("Received New user reeuest with age : %v\n", in.GetAge())
	log.Printf("Context : %+v\n", ctx)
	user_id := int32(rand.Intn(1000))

	created_user := &pb.User{
		Name: in.GetName(),
		Age:  in.GetAge(),
		Id:   user_id,
	}
	s.usersList.Users = append(s.usersList.Users, created_user)

	return created_user, nil
}

func (s *UserManagementServer) GetUsers(ctx context.Context, in *pb.GetUsersParams) (*pb.UserList, error) {
	return s.usersList, nil
}
func main() {
	fmt.Println("Hello server")

	user_mgmt_server := NewUserManagementServer()
	err := user_mgmt_server.RunServer()
	if err != nil {
		log.Fatalf("Failed to serve : %v\n", err)
	}

}

func (server *UserManagementServer) RunServer() error {
	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("Failed to listen : %v\n", err)
	}

	s := grpc.NewServer()
	pb.RegisterUserManagementServer(s, server)
	log.Printf("Server listening at %v:%v", lis, port)

	return s.Serve(lis)
}
