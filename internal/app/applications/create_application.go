package applications

import (
	"context"

	"github.com/alexeysamorodov/creator-applications/internal/app/applications/domain"
	pb "github.com/alexeysamorodov/creator-applications/internal/pb"
)

func (i *Implementation) CreateApplication(ctx context.Context, request *pb.CreateApplicationRequest) (*pb.CreateApplicationResponse, error) {
	application := domain.NewApplication(request.RequestId, request.TaxId)

	i.ApplicationRepository.Save(ctx, *application)

	response := pb.CreateApplicationResponse{
		ApplicationId: application.ID.String(),
	}

	return &response, nil
}
