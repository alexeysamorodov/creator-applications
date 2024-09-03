package grpcapi

import (
	"context"
	"fmt"
	"log"

	"github.com/alexeysamorodov/creator-applications/internal/app/applications/domain"
	createapplicationconverter "github.com/alexeysamorodov/creator-applications/internal/app/applications/grpc/api/converters/createapplicationconverter"
	pb "github.com/alexeysamorodov/creator-applications/internal/pb"
)

func (i *Implementation) CreateApplication(ctx context.Context, request *pb.CreateApplicationRequest) (*pb.CreateApplicationResponse, error) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
	}()

	applicationAttributes := createapplicationconverter.ToApplicationAttributesDomain(request.ApplicationAttributes)

	application := domain.NewApplication(
		request.RequestId,
		request.Name,
		applicationAttributes)

	err := i.ApplicationRepository.Save(ctx, application)

	if err != nil {
		log.Fatal(err)
	}

	response := pb.CreateApplicationResponse{
		ApplicationId: application.ID.String(),
	}

	return &response, nil
}
