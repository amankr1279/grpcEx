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
}

func (s *UserManagementServer) CreaeNewUser(ctx context.Context, in *pb.NewUser) (*pb.User, error) {
	log.Printf("Received : %v\n", in.GetAge())
	user_id := int32(rand.Intn(1000))

	return &pb.User{
		Name: in.GetName(),
		Age:  in.GetAge(),
		Id:   user_id,
	}, nil
}
func main() {
	fmt.Println("Hello server")

	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("Failed to listen : %v\n", err)
	}

	server := grpc.NewServer()
	pb.RegisterUserManagementServer(server, &UserManagementServer{})
	log.Printf("Server lsitening at %v:%v", lis, port)

	if err = server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve : %v\n", err)
	}
}
