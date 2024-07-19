package main

import (
	"log"
	"net"

	"github.com/alexeysamorodov/creator-applications/internal/app/applications"
	pb "github.com/alexeysamorodov/creator-applications/internal/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// Создаем gRPC сервер с перехватчиком
	server := grpc.NewServer()

	// Регистрируем ваш сервис
	pb.RegisterApplicationsServiceServer(server, &applications.Implementation{})

	// Включаем рефлексию (для gRPC CLI и других инструментов)
	reflection.Register(server)

	// Настраиваем прослушивание порта
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Запускаем сервер
	log.Printf("Server is listening on port %s", lis.Addr())
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
