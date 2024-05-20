package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8080"
)

func main() {
	conn, err := grpc.NewClient("localhost:"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to the server %v", err)
	}
	defer conn.Close()

	// client := pb.NewGreetServiceClient(conn)

	// names = &pb.NamesList{
	// 	Names: []string{"Harshal", "Rahul", "Rohit"},
	// }
}