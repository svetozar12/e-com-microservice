package user_test

import (
	"context"
	"strings"
	"testing"

	pb "svetozar12/e-com/v2/api/v1/user/dist/proto"
	"svetozar12/e-com/v2/apps/services/user/internal/app/entities"
	"svetozar12/e-com/v2/apps/services/user/internal/app/repositories/userRepository"
	userConstants "svetozar12/e-com/v2/apps/services/user/internal/pkg/constants"
	"svetozar12/e-com/v2/apps/services/user/internal/pkg/jwtUtils"
	"svetozar12/e-com/v2/libs/api/constants"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestGetUser(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	userClient := pb.NewUserServiceClient(conn)
	user, _ := userRepository.CreateUser(&entities.UserEntity{Email: testEmail, Password: jwtUtils.HashAndSalt([]byte(testPassword)), Fname: testFname, Lname: testLname})
	usersToDelete[0] = user.Email
	t.Run("rpc GetUser(expected behavior)", func(t *testing.T) {
		resp, err := userClient.GetUser(ctx, &pb.GetUserRequest{Id: int32(user.ID)})
		if err != nil {
			t.Fatalf("GetUser failed: %v", err)
		}
		if resp.Id != int32(user.ID) {
			t.Fatalf(constants.InvalidFieldValueMessage("id"))
		}
		if resp.Email != testEmail {
			t.Fatalf(constants.InvalidFieldValueMessage("email"))
		}
		if resp.Fname != user.Fname {
			t.Fatalf(constants.InvalidFieldValueMessage("fname"))
		}
		if resp.Lname != user.Lname {
			t.Fatalf(constants.InvalidFieldValueMessage("lname"))
		}
	})

	t.Run("rpc GetUser(invalid input)", func(t *testing.T) {
		_, err := userClient.GetUser(ctx, &pb.GetUserRequest{Id: int32(0)})
		if !strings.Contains(err.Error(), constants.GTEValueMessage("1")) {
			t.Errorf(constants.InvalidFieldMessage("id"))
		}
	})

	t.Run("rpc GetUser(not found)", func(t *testing.T) {
		_, err := userClient.GetUser(ctx, &pb.GetUserRequest{Id: int32(999999)})

		if err.Error() != status.Error(codes.NotFound, userConstants.UserNotFoundMessage).Error() {
			t.Errorf(err.Error())
		}
	})

	t.Cleanup(func() {
		userRepository.HardDeleteUser(user, nil)
	})
}
