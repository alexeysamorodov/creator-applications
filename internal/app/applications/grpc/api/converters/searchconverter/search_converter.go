package searchconverter

import (
	"github.com/alexeysamorodov/creator-applications/internal/app/applications/domain"
	"github.com/alexeysamorodov/creator-applications/internal/app/applications/domain/valueobjects"
	"github.com/alexeysamorodov/creator-applications/internal/app/applications/grpc/api/converters/enumsconverter"
	pb "github.com/alexeysamorodov/creator-applications/internal/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToApplicationAttributesDomain(attributes []*pb.SearchResponseApplicationAttribute) []*valueobjects.ApplicationAttribute {
	result := make([]*valueobjects.ApplicationAttribute, len(attributes))

	for i := 0; i < len(attributes); i++ {
		result[i] = toDomain(attributes[i])
	}

	return result
}

func toDomain(attribute *pb.SearchResponseApplicationAttribute) *valueobjects.ApplicationAttribute {
	result := &valueobjects.ApplicationAttribute{
		Name: attribute.Name,
	}

	return result
}

func ToSearchResponseApplicationPb(application *domain.Application) *pb.SearchResponseApplication {
	state, err := enumsconverter.ToApplicationStatePb(application.State)
	if err != nil {
		panic(err)
	}

	result := &pb.SearchResponseApplication{
		Id:         application.ID.String(),
		Name:       application.Name,
		State:      state,
		RequestId:  application.RequestID,
		CreatedAt:  timestamppb.New(application.CreatedAt),
		UpdatedAt:  timestamppb.New(application.UpdatedAt),
		Attributes: ToSearchResponseApplicationAttributesPb(application.Attributes),
	}

	return result
}

func ToSearchResponseApplicationAttributesPb(attributes []valueobjects.ApplicationAttribute) []*pb.SearchResponseApplicationAttribute {
	result := make([]*pb.SearchResponseApplicationAttribute, len(attributes))

	for i := 0; i < len(attributes); i++ {
		result[i] = toSearchResponseApplicationAttributePb(attributes[i])
	}

	return result
}

func toSearchResponseApplicationAttributePb(attribute valueobjects.ApplicationAttribute) *pb.SearchResponseApplicationAttribute {
	result := &pb.SearchResponseApplicationAttribute{
		Name: attribute.Name,
	}

	return result
}
