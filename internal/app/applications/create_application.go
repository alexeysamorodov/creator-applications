package applications

import (
	"context"

	pb "github.com/alexeysamorodov/creator-applications/internal/pb"
)

func (i *Implementation) CreateApplication(ctx context.Context, request *pb.CreateApplicationRequest) (*pb.CreateApplicationResponse, error) {
	response := pb.CreateApplicationResponse{
		ApplicationId: "42",
	}

	return &response, nil
}
