package main

import (
	"context"
	"fmt"
	"os"
	"time"

	//pb "path/to/your/generated/protobuf"
	//pb "protos/train-ticket-service/"
	pb "yogeshgadakh/go-grpc-demo/protos"

	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	fmt.Println("Starting client. . .")
	serverAddress := os.Getenv("SERVER_ADDRESS")
	if serverAddress == "" {
		serverAddress = "localhost:50051"
	}
	fmt.Printf("Connecting to Server address: %s\n", serverAddress)
	conn, err := grpc.NewClient(serverAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Printf("did not connect: %v\n", err)
		return
	}
	defer conn.Close()

	// Wait for the connection to be ready
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Use the connection to check if it is ready
	state := conn.GetState()
	if state == connectivity.Ready {
		fmt.Println("Connection is already ready")
	} else {
		fmt.Println("Waiting for connection to be ready...")
		if isConnected := conn.WaitForStateChange(ctx, connectivity.Ready); !isConnected {
			fmt.Printf("Failed to connect: %v\n", err)
			return
		}
		fmt.Println("Connection established successfully")
	}

	c := pb.NewTrainTicketServiceClient(conn)

	// Example of calling PurchaseTicket
	fmt.Println("Example of calling PurchaseTicket. . .")

	purchaseResp, err := c.PurchaseTicket(context.Background(), &pb.PurchaseRequest{
		From:      "London",
		To:        "France",
		User:      &pb.User{FirstName: "Yogesh", LastName: "Gadakh", Email: "yogee@gmail.com"},
		PricePaid: 20.0,
		Section:   "A",
	})
	if err != nil {
		fmt.Printf("could not purchase ticket: %v\n", err)
		return
	}
	fmt.Printf("Purchase Response: %v\n", purchaseResp)

	time.Sleep(5 * time.Second) // Adding delay to see the output in simulated time

	// Example of calling GetReceipt
	fmt.Println("Example of calling GetReceipt . . .")

	receiptResp, err := c.GetReceipt(context.Background(), &pb.PurchaseRequest{User: &pb.User{Email: "yogee@gmail.com"}})
	if err != nil {
		fmt.Printf("could not get receipt: %v\n", err)
		return
	}
	fmt.Printf("Receipt Response: %v\n", receiptResp)
	fmt.Printf("Receipt ID: %s\n", receiptResp.ReceiptId)

	time.Sleep(5 * time.Second)

	// Example of calling GetUsersInSection
	fmt.Println("Example of calling GetUsersInSection . . .")
	sectionResp, err := c.GetUsersInSection(context.Background(), &pb.UsersInSectionRequest{Section: "A"})
	if err != nil {
		fmt.Printf("could not get User In Section: %v\n", err)
		return
	}

	time.Sleep(5 * time.Second)
	fmt.Printf("Section Response: %v\n", sectionResp)
	for _, user := range sectionResp.Users {
		fmt.Printf("User email: %s, Name: %s, last Name: %s, User ID: %s\n",
			user.User.Email, user.User.FirstName, user.User.LastName, user.UserId)
	}

	time.Sleep(5 * time.Second)

	// Example of calling ModifySeat
	fmt.Println("Example of calling ModifySeat. . .")
	modifyRes, err := c.ModifySeat(context.Background(), &pb.ModifySeatRequest{ReceiptId: receiptResp.ReceiptId, NewSeatNumber: "100"})
	if err != nil {
		fmt.Printf("could not Modify the seat number: %v\n", err)
		return
	}
	fmt.Printf("Modify Seat request status: %v\n", modifyRes.Success)

	time.Sleep(5 * time.Second)

	// Example of calling RemoveUser
	fmt.Println("Example of calling RemoveUser. . .")
	removeResp, err := c.RemoveUser(context.Background(), &pb.RemoveUserRequest{UserId: receiptResp.ReceiptId})
	if err != nil {
		fmt.Printf("could not Remove the user: %v\n", err)
		return
	}
	fmt.Printf("Remove User request status: %v\n", removeResp.Success)
}
