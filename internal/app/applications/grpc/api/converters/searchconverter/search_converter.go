package searchconverter

import (
	"github.com/alexeysamorodov/creator-applications/internal/app/applications/domain/valueobjects"
	pb "github.com/alexeysamorodov/creator-applications/internal/pb"
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

func FromApplicationAttributesDomain(attributes []valueobjects.ApplicationAttribute) []*pb.SearchResponseApplicationAttribute {
	result := make([]*pb.SearchResponseApplicationAttribute, len(attributes))

	for i := 0; i < len(attributes); i++ {
		result[i] = fromDomain(attributes[i])
	}

	return result
}

func fromDomain(attribute valueobjects.ApplicationAttribute) *pb.SearchResponseApplicationAttribute {
	result := &pb.SearchResponseApplicationAttribute{
		Name: attribute.Name,
	}

	return result
}
