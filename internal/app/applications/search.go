package applications

import (
	"context"
	"errors"

	pb "github.com/alexeysamorodov/creator-applications/internal/pb"
	"github.com/google/uuid"
)

func (i *Implementation) Search(ctx context.Context, in *pb.SearchRequest) (*pb.SearchResponse, error) {
	var applicationIds []uuid.UUID

	for _, applicationId := range in.ApplicationIds {
		parsedId, err := uuid.Parse(applicationId)

		if err != nil {
			// TODO: return InvalidArgument
			return nil, errors.New("InvalidArgument")
		}

		applicationIds = append(applicationIds, parsedId)
	}

	applications := i.ApplicationRepository.SearchApplications(ctx, applicationIds)

	var responseApplications []*pb.SearchResponseApplication

	for _, application := range applications {
		application := pb.SearchResponseApplication{
			Id:        application.ID.String(),
			TaxId:     application.TaxID,
			RequestId: application.RequestID,
		}

		responseApplications = append(responseApplications, &application)
	}

	response := pb.SearchResponse{
		Applications: responseApplications,
	}

	return &response, nil
}
