package review_test

import (
	"context"
	"fmt"
	"strings"
	"testing"

	pb "svetozar12/e-com/v2/api/v1/review/dist/proto"
	"svetozar12/e-com/v2/apps/services/review/internal/app/entities"
	reviewrepository "svetozar12/e-com/v2/apps/services/review/internal/app/repositories/reviewRepository"
	reviewConstants "svetozar12/e-com/v2/apps/services/review/internal/pkg/constants"
	"svetozar12/e-com/v2/libs/api/constants"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestGetReview(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	client := pb.NewReviewServiceClient(conn)
	review, err := reviewrepository.CreateReview(&entities.ReviewEntity{ProductId: 1, UserId: 1, Comment: "dummy comment", Rating: 2})
	t.Run("rpc GetReview(expected behavior)", func(t *testing.T) {
		resp, err := client.GetReview(ctx, &pb.GetReviewRequest{ReviewId: int32(review.ID)})
		if err != nil {
			t.Fatalf("GetReview failed: %v", err)
		}

		if resp.ReviewId != int32(review.ID) {
			t.Fatalf(constants.InvalidFieldValueMessage("id"))
		}
		if resp.ProductId != int32(review.ProductId) {
			fmt.Println(resp.ProductId, review.ProductId, "error")
			t.Fatalf(constants.InvalidFieldValueMessage("ProductId"))
		}
		if resp.UserId != int32(review.UserId) {
			t.Fatalf(constants.InvalidFieldValueMessage("UserId"))
		}
		if resp.Comment != review.Comment {
			t.Fatalf(constants.InvalidFieldValueMessage("comment"))
		}
		if resp.Rating != review.Rating {
			t.Fatalf(constants.InvalidFieldValueMessage("rating"))
		}
	})

	t.Run("rpc GetReview(invalid input)", func(t *testing.T) {
		_, err := client.GetReview(ctx, &pb.GetReviewRequest{ReviewId: int32(0)})
		if !strings.Contains(err.Error(), constants.GTEValueMessage("1")) {
			t.Errorf(constants.InvalidFieldMessage("id"))
		}
	})

	t.Run("rpc GetReview(not found)", func(t *testing.T) {
		_, err := client.GetReview(ctx, &pb.GetReviewRequest{ReviewId: int32(99999)})

		if err.Error() != status.Error(codes.NotFound, reviewConstants.ReviewNotFound).Error() {
			t.Errorf(err.Error())
		}
	})

	t.Cleanup(func() {
		reviewrepository.HardDeleteReview(&entities.ReviewEntity{}, "id = ?", review.ID)
	})
}
