package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"strconv"

	"github.com/dty1er/kubecept/microservices/proto/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service struct {
	pb.UnimplementedItemServiceServer
	data map[string]*pb.Item
}

func (s *Service) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	i, ok := s.data[req.GetId()]
	if !ok {
		return nil, status.New(codes.NotFound, "item is not found").Err()
	}

	return &pb.GetResponse{Item: i}, nil
}

func main() {
	service := &Service{data: loadData()}

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("[ERROR] listen: %v", err)
	}

	server := grpc.NewServer()
	pb.RegisterItemServiceServer(server, service)

	log.Printf("[INFO] item server is now up and running on :50051")
	if err := server.Serve(lis); err != nil {
		log.Fatalf("[ERROR] serve: %v", err)
	}
}

func loadData() map[string]*pb.Item {
	d := make(map[string]*pb.Item)
	for i := 1; i <= 100; i++ {
		d[strconv.Itoa(i)] = &pb.Item{Id: strconv.Itoa(i), Name: fmt.Sprintf("Item%d", i), Price: uint64(i * 100)}
	}

	return d
}
