package grpcapi

import (
	"context"
	"errors"
	"fmt"
	"log"

	searchconverter "github.com/alexeysamorodov/creator-applications/internal/app/applications/grpc/api/converters/searchconverter"
	pb "github.com/alexeysamorodov/creator-applications/internal/pb"
	"github.com/google/uuid"
)

func (i *Implementation) Search(ctx context.Context, in *pb.SearchRequest) (*pb.SearchResponse, error) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
	}()

	var applicationIds []uuid.UUID

	for _, applicationId := range in.ApplicationIds {
		parsedId, err := uuid.Parse(applicationId)

		if err != nil {
			// TODO: return InvalidArgument
			return nil, errors.New("InvalidArgument")
		}

		applicationIds = append(applicationIds, parsedId)
	}

	applications, err := i.ApplicationRepository.SearchApplications(ctx, applicationIds)

	if err != nil {
		log.Fatal(err)
	}

	var responseApplications []*pb.SearchResponseApplication

	for _, application := range applications {
		applicationPb := searchconverter.ToSearchResponseApplicationPb(&application)

		responseApplications = append(responseApplications, applicationPb)
	}

	response := pb.SearchResponse{
		Applications: responseApplications,
	}

	return &response, nil
}
