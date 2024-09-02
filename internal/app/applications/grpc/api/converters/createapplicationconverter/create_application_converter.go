package createapplicationconverter

import (
	"github.com/alexeysamorodov/creator-applications/internal/app/applications/domain/valueobjects"
	pb "github.com/alexeysamorodov/creator-applications/internal/pb"
)

func ToApplicationAttributesDomain(attributes []*pb.CreateApplicationRequestAttribute) []*valueobjects.ApplicationAttribute {
	result := make([]*valueobjects.ApplicationAttribute, len(attributes))

	for i := 0; i < len(attributes); i++ {
		result[i] = toDomain(attributes[i])
	}

	return result
}

func toDomain(attribute *pb.CreateApplicationRequestAttribute) *valueobjects.ApplicationAttribute {
	result := &valueobjects.ApplicationAttribute{
		Name: attribute.Name,
	}

	return result
}

func FromApplicationAttributesDomain(attributes []*valueobjects.ApplicationAttribute) []*pb.CreateApplicationRequestAttribute {
	result := make([]*pb.CreateApplicationRequestAttribute, len(attributes))

	for i := 0; i < len(attributes); i++ {
		result[i] = fromDomain(attributes[i])
	}

	return result
}

func fromDomain(attribute *valueobjects.ApplicationAttribute) *pb.CreateApplicationRequestAttribute {
	result := &pb.CreateApplicationRequestAttribute{
		Name: attribute.Name,
	}

	return result
}
