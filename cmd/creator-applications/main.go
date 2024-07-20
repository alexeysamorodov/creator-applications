package main

import (
	"fmt"
	"log"
	"net"

	"github.com/alexeysamorodov/creator-applications/internal/app/applications"
	"github.com/alexeysamorodov/creator-applications/internal/app/applications/database"
	"github.com/alexeysamorodov/creator-applications/internal/config"
	pb "github.com/alexeysamorodov/creator-applications/internal/pb"
	"github.com/alexeysamorodov/creator-applications/internal/pkg/db"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	if err := config.ReadConfigYML("config.yml"); err != nil {
		log.Fatal("Failed init configuration")
	}
	cfg := config.GetConfigInstance()

	dsn := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=%v",
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Name,
		cfg.Database.SslMode,
	)

	db, err := db.ConnectDB(dsn, cfg.Database.Driver)
	if err != nil {
		log.Fatal("Failed init postgres")
	}
	defer db.Close()

	// Создаем gRPC сервер с перехватчиком
	server := grpc.NewServer()

	// Регистрируем ваш сервис
	applicationRepository := database.NewApplicationRepository()

	applicationService := applications.NewApplicationsService(applicationRepository)

	pb.RegisterApplicationsServiceServer(server, applicationService)

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
