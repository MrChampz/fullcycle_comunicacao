package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/MrChampz/fullcycle-comunicacao/grpc/pb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to gRPC server: %v", err)
	}
	defer conn.Close()

	client := pb.NewUserServiceClient(conn)
	//AddUser(client)
	//AddUserVerbose(client)
	//AddUsers(client)
	AddUserStreamBoth(client)
}

func AddUser(client pb.UserServiceClient) {
	req := &pb.User{
		Id:    "0",
		Name:  "Robbye",
		Email: "robbye@wow.com",
	}

	res, err := client.AddUser(context.Background(), req)
	if err != nil {
		log.Fatalf("Could not make gRPC request: %v", err)
	}

	fmt.Println(res)
}

func AddUserVerbose(client pb.UserServiceClient) {
	req := &pb.User{
		Id:    "0",
		Name:  "Robbye",
		Email: "robbye@wow.com",
	}

	res, err := client.AddUserVerbose(context.Background(), req)
	if err != nil {
		log.Fatalf("Could not make gRPC request: %v", err)
	}

	for {
		stream, err := res.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Could not receive the message: %v", err)
		}
		fmt.Println("Status: ", stream.Status, " - ", stream.GetUser())
	}
}

func AddUsers(client pb.UserServiceClient) {
	reqs := []*pb.User{
		&pb.User{
			Id:    "0",
			Name:  "Roberto",
			Email: "rob@gmail.com",
		},
		&pb.User{
			Id:    "1",
			Name:  "Giuliano",
			Email: "giu@gmail.com",
		},
		&pb.User{
			Id:    "2",
			Name:  "Marcela",
			Email: "mac@gmail.com",
		},
		&pb.User{
			Id:    "3",
			Name:  "Felipe",
			Email: "fel@gmail.com",
		},
		&pb.User{
			Id:    "4",
			Name:  "Gabriel",
			Email: "gab@gmail.com",
		},
		&pb.User{
			Id:    "5",
			Name:  "Steve",
			Email: "steve@gmail.com",
		},
	}

	stream, err := client.AddUsers(context.Background())
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	for _, req := range reqs {
		stream.Send(req)
		time.Sleep(time.Second * 3)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error receiving response: %v", err)
	}

	fmt.Println(res)
}

func AddUserStreamBoth(client pb.UserServiceClient) {
	reqs := []*pb.User{
		&pb.User{
			Id:    "0",
			Name:  "Roberto",
			Email: "rob@gmail.com",
		},
		&pb.User{
			Id:    "1",
			Name:  "Giuliano",
			Email: "giu@gmail.com",
		},
		&pb.User{
			Id:    "2",
			Name:  "Marcela",
			Email: "mac@gmail.com",
		},
		&pb.User{
			Id:    "3",
			Name:  "Felipe",
			Email: "fel@gmail.com",
		},
		&pb.User{
			Id:    "4",
			Name:  "Gabriel",
			Email: "gab@gmail.com",
		},
		&pb.User{
			Id:    "5",
			Name:  "Steve",
			Email: "steve@gmail.com",
		},
	}

	stream, err := client.AddUserStreamBoth(context.Background())
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	wait := make(chan int)

	go func() {
		for _, req := range reqs {
			fmt.Println("Sending user: ", req.GetName())
			stream.Send(req)
			time.Sleep(time.Second * 2)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error receiving data: %v", err)
				break
			}
			fmt.Printf("Receiving user: %v, with status: %v\n", res.GetUser().GetName(), res.GetStatus())
		}
		close(wait)
	}()

	<-wait
}
