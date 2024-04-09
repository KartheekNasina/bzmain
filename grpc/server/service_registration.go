package server

import (
	// other imports

	repository "github.com/vivekbnwork/bz-backend/bz-main/repository"
	service "github.com/vivekbnwork/bz-backend/bz-main/service"

	pb "github.com/vivekbnwork/bz-backend/bz-main/grpc/proto/grpc/proto"

	"github.com/vivekbnwork/bz-backend/bz-main/driver"
	"google.golang.org/grpc"
)

// RegisterUserService registers the user service with the given gRPC server.
func RegisterUserService(s *grpc.Server, db *driver.DB) {
	userRepo := repository.NewUserRepository(db)
	userServiceServer := service.NewUserServiceServer(userRepo)
	pb.RegisterUserServiceServer(s, userServiceServer)
}
