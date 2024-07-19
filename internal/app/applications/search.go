package applications

import (
	"context"

	pb "github.com/alexeysamorodov/creator-applications/internal/pb"
)

func (i *Implementation) Search(ctx context.Context, in *pb.SearchRequest) (*pb.SearchResponse, error) {
	application := &pb.SearchResponseApplication{
		Id:        "42",
		TaxId:     "777",
		RequestId: 7,
	}

	response := pb.SearchResponse{
		Applications: []*pb.SearchResponseApplication{application},
	}

	return &response, nil
}
