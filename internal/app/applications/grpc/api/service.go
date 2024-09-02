package grpcapi

import (
	db "github.com/alexeysamorodov/creator-applications/internal/app/applications/database"
	pb "github.com/alexeysamorodov/creator-applications/internal/pb"
)

type Implementation struct {
	pb.UnimplementedApplicationsServiceServer
	ApplicationRepository db.IApplicationRepository
}

func NewApplicationsService(applicationRepository db.IApplicationRepository) *Implementation {
	return &Implementation{
		ApplicationRepository: applicationRepository,
	}
}
