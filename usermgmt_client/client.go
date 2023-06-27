package main

import (
	"context"
	"log"
	"time"

	pb "github.com/amankr1279/grpcEx/usermgmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const address = "localhost:50051"

func main() {
	// Blockling call
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Failed to connect : %v\n", err)
	}
	defer conn.Close()

	client := pb.NewUserManagementClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	new_users := make(map[string]int32)

	new_users["Alice"] = 30
	new_users["Bob"] = 31

	for name, age := range new_users {
		resp, err := client.CreateNewUser(ctx, &pb.NewUser{
			Name: name,
			Age:  age,
		})
		if err != nil {
			log.Fatalf("Failed to get response of create user : %v\n", err)
		}
		log.Printf(` User Details : 
Name : %v
Age: %v
Id :  %v
`, resp.GetName(), resp.GetAge(), resp.GetId())
	}

	params := &pb.GetUsersParams{}
	resp, err := client.GetUsers(ctx, params)
	if err != nil {
		log.Fatalf("Failed to get response of get all users : %v\n", err)
	}
	log.Println(resp.GetUsers())
}
