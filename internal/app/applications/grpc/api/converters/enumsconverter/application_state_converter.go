package enumsconverter

import (
	"github.com/alexeysamorodov/creator-applications/internal/app/applications/domain"
	"github.com/alexeysamorodov/creator-applications/internal/app/applications/domain/valueobjects/errors"
	pb "github.com/alexeysamorodov/creator-applications/internal/pb"
)

func ToApplicationStatePb(applicationState domain.ApplicationState) (pb.ApplicationState, error) {
	var result pb.ApplicationState

	switch applicationState {
	case domain.ApplicationState_None:
		result = *pb.ApplicationState_ApplicationState_None.Enum()
	case domain.ApplicationState_Created:
		result = *pb.ApplicationState_ApplicationState_Created.Enum()
	default:
		return pb.ApplicationState_ApplicationState_None, errors.NewEnumOutOfRangeError(applicationState)
	}

	return result, nil
}
