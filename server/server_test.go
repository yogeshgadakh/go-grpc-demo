package main

import (
	"context"
	"testing"

	"fmt"
	"net"

	pb "yogeshgadakh/go-grpc-demo/protos"

	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func init() {
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	pb.RegisterTrainTicketServiceServer(s, NewServer())
	go func() {
		if err := s.Serve(lis); err != nil {
			fmt.Println("Failed to serve:", err)
		}
	}()
}

func dialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func TestPurchaseTicket(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "", grpc.WithContextDialer(dialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewTrainTicketServiceClient(conn)

	resp, err := client.PurchaseTicket(ctx, &pb.PurchaseRequest{
		From:      "London",
		To:        "France",
		User:      &pb.User{FirstName: "John", LastName: "Doe", Email: "john.doe@example.com"},
		PricePaid: 20.0,
		Section:   "A",
	})
	if err != nil {
		t.Fatalf("could not purchase ticket: %v", err)
	}

	if resp.ReceiptId == "" {
		t.Fatalf("expected non-empty receipt ID")
	}
}

func TestGetReceipt(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "", grpc.WithContextDialer(dialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewTrainTicketServiceClient(conn)

	_, err = client.PurchaseTicket(ctx, &pb.PurchaseRequest{
		From:      "London",
		To:        "France",
		User:      &pb.User{FirstName: "John", LastName: "Doe", Email: "john.doe@example.com"},
		PricePaid: 20.0,
		Section:   "A",
	})
	if err != nil {
		t.Fatalf("could not purchase ticket: %v", err)
	}

	resp, err := client.GetReceipt(ctx, &pb.PurchaseRequest{User: &pb.User{Email: "john.doe@example.com"}})
	if err != nil {
		t.Fatalf("could not get receipt: %v", err)
	}

	if resp.User.Email != "john.doe@example.com" {
		t.Fatalf("expected email john.doe@example.com, got %v", resp.User.Email)
	}
}

func TestGetUsersInSection(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "", grpc.WithContextDialer(dialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewTrainTicketServiceClient(conn)

	_, err = client.PurchaseTicket(ctx, &pb.PurchaseRequest{
		From:      "London",
		To:        "France",
		User:      &pb.User{FirstName: "John", LastName: "Doe", Email: "john.doe@example.com"},
		PricePaid: 20.0,
		Section:   "A",
	})
	if err != nil {
		t.Fatalf("could not purchase ticket: %v", err)
	}

	resp, err := client.GetUsersInSection(ctx, &pb.UsersInSectionRequest{Section: "A"})
	if err != nil {
		t.Fatalf("could not get users in section: %v", err)
	}

	if len(resp.Users) == 0 {
		t.Fatalf("expected to find users in section A")
	}
}

func TestRemoveUser(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "", grpc.WithContextDialer(dialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewTrainTicketServiceClient(conn)

	_, err = client.PurchaseTicket(ctx, &pb.PurchaseRequest{
		From:      "London",
		To:        "France",
		User:      &pb.User{FirstName: "John", LastName: "Doe", Email: "john.doe@example.com"},
		PricePaid: 20.0,
		Section:   "A",
	})
	if err != nil {
		t.Fatalf("could not purchase ticket: %v", err)
	}

	resp, err := client.RemoveUser(ctx, &pb.RemoveUserRequest{UserId: "1"})
	if err != nil {
		t.Fatalf("could not remove user: %v", err)
	}

	if !resp.Success {
		t.Fatalf("expected removal to be successful")
	}
}

func TestModifySeat(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "", grpc.WithContextDialer(dialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewTrainTicketServiceClient(conn)

	_, err = client.PurchaseTicket(ctx, &pb.PurchaseRequest{
		From:      "London",
		To:        "France",
		User:      &pb.User{FirstName: "John", LastName: "Doe", Email: "john.doe@example.com"},
		PricePaid: 20.0,
		Section:   "A",
	})
	if err != nil {
		t.Fatalf("could not purchase ticket: %v", err)
	}

	resp, err := client.ModifySeat(ctx, &pb.ModifySeatRequest{ReceiptId: "1", NewSeatNumber: "B2"})
	if err != nil {
		t.Fatalf("could not modify seat: %v", err)
	}

	if !resp.Success {
		t.Fatalf("expected seat modification to be successful")
	}
}
