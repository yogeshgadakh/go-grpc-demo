package main

import (
	"context"
	"fmt"
	"net"
	"sync"
	pb "yogeshgadakh/go-grpc-demo/protos"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedTrainTicketServiceServer
	receipts     map[string]*pb.PurchaseResponse
	userSeats    map[string]string
	sectionSeats map[string][]*pb.UserDetails
	mu           sync.Mutex
}

func NewServer() *server {
	return &server{
		receipts:     make(map[string]*pb.PurchaseResponse),
		userSeats:    make(map[string]string),
		sectionSeats: make(map[string][]*pb.UserDetails),
	}
}

func (s *server) PurchaseTicket(ctx context.Context, req *pb.PurchaseRequest) (*pb.PurchaseResponse, error) {
	fmt.Printf("Received purchase request for %s to %s from %s\n", req.From, req.To, req.User.Email)
	s.mu.Lock()
	defer s.mu.Unlock()

	receiptID := fmt.Sprintf("%d", len(s.receipts)+1)
	seatNumber := "A1" // For simplicity, use a fixed seat or implement seat allocation logic
	response := &pb.PurchaseResponse{
		ReceiptId:  receiptID,
		From:       req.From,
		To:         req.To,
		User:       req.User,
		PricePaid:  req.PricePaid,
		Section:    req.Section,
		SeatNumber: seatNumber,
	}

	s.receipts[receiptID] = response
	s.userSeats[req.User.Email] = seatNumber
	s.sectionSeats[req.Section] = append(s.sectionSeats[req.Section], &pb.UserDetails{
		UserId:     receiptID,
		SeatNumber: seatNumber,
		User:       req.User,
	})
	fmt.Printf("Purchase successful. Receipt ID: %s\n", receiptID)
	return response, nil
}

func (s *server) GetReceipt(ctx context.Context, req *pb.PurchaseRequest) (*pb.PurchaseResponse, error) {
	fmt.Printf("Received receipt request for %s\n", req.User.Email)
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, receipt := range s.receipts {
		if receipt.User.Email == req.User.Email {
			fmt.Printf("Receipt found for %s\n", req.User.Email)
			return receipt, nil
		}
	}
	fmt.Printf("Receipt not found for %s\n", req.User.Email)
	return nil, fmt.Errorf("receipt not found")
}

func (s *server) GetUsersInSection(ctx context.Context, req *pb.UsersInSectionRequest) (*pb.UsersInSectionResponse, error) {
	fmt.Printf("Received users in section request for %s\n", req.Section)
	s.mu.Lock()
	defer s.mu.Unlock()

	users, exists := s.sectionSeats[req.Section]
	if !exists {
		return &pb.UsersInSectionResponse{}, nil
	}
	fmt.Printf("Found %d users in section %s\n", len(users), req.Section)
	return &pb.UsersInSectionResponse{Users: users}, nil
}

func (s *server) RemoveUser(ctx context.Context, req *pb.RemoveUserRequest) (*pb.RemoveUserResponse, error) {
	fmt.Printf("Received remove user request for %s\n", req.UserId)
	s.mu.Lock()
	defer s.mu.Unlock()

	receipt, exists := s.receipts[req.UserId]
	if !exists {
		return &pb.RemoveUserResponse{Success: false}, nil
	}

	sectionUsers := s.sectionSeats[receipt.Section]
	for i, user := range sectionUsers {
		if user.UserId == req.UserId {
			s.sectionSeats[receipt.Section] = append(sectionUsers[:i], sectionUsers[i+1:]...)
			break
		}
	}

	delete(s.receipts, req.UserId)
	delete(s.userSeats, receipt.User.Email)
	fmt.Printf("User %s removed\n", req.UserId)
	return &pb.RemoveUserResponse{Success: true}, nil
}

func (s *server) ModifySeat(ctx context.Context, req *pb.ModifySeatRequest) (*pb.ModifySeatResponse, error) {
	fmt.Printf("Received modify seat request for %s\n", req.ReceiptId)
	s.mu.Lock()
	defer s.mu.Unlock()

	receipt, exists := s.receipts[req.ReceiptId]
	if !exists {
		return &pb.ModifySeatResponse{Success: false}, nil
	}

	// oldSeat := receipt.SeatNumber
	receipt.SeatNumber = req.NewSeatNumber
	s.userSeats[receipt.User.Email] = req.NewSeatNumber

	sectionUsers := s.sectionSeats[receipt.Section]
	for i, user := range sectionUsers {
		if user.UserId == req.ReceiptId {
			sectionUsers[i].SeatNumber = req.NewSeatNumber
			break
		}
	}
	fmt.Printf("Seat modified for %s\n", req.ReceiptId)
	return &pb.ModifySeatResponse{Success: true}, nil
}

func main() {
	fmt.Println("Starting server. . .")
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		fmt.Printf("failed to listen: %v\n", err)
		return
	}
	s := grpc.NewServer()
	pb.RegisterTrainTicketServiceServer(s, NewServer())
	fmt.Println("Server listening on port 50051")
	if err := s.Serve(lis); err != nil {
		fmt.Printf("failed to serve: %v\n", err)
	}
}
