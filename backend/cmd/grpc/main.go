package main

import (
	"log"
	"net"

	pb "github.com/oyal2/LIFXMaster/internal/proto"
	"github.com/oyal2/LIFXMaster/internal/svc"
	"github.com/oyal2/LIFXMaster/pkg/connection"
	"google.golang.org/grpc"
)

func main() {
	// Keep the program running
	c, err := connection.NewLXClient("192.168.1.255")
	if err != nil {
		log.Fatalln(err)
	}

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterDeviceServiceServer(s, svc.NewDeviceSvc(c))
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
