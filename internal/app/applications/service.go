package applications

import pb "github.com/alexeysamorodov/creator-applications/internal/pb"

type Implementation struct {
	pb.UnimplementedApplicationsServiceServer
}

func NewApplicationsService() *Implementation {
	return &Implementation{}
}
