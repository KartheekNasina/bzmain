// package service/user_service.go
package service

import (
	"context"

	pb "github.com/vivekbnwork/bz-backend/bz-main/grpc/proto/grpc/proto"

	"github.com/jinzhu/copier"
	"github.com/sirupsen/logrus"
	models "github.com/vivekbnwork/bz-backend/bz-main/models/db"
	dto "github.com/vivekbnwork/bz-backend/bz-main/models/dto"

	"github.com/vivekbnwork/bz-backend/bz-main/repository"
)

type UserService struct {
	userRepo *repository.UserRepository
}

// UserServiceServer is the server API for UserService service.
type UserServiceServer struct {
	pb.UnimplementedUserServiceServer
	userRepo *repository.UserRepository
}

func NewUserServiceServer(r *repository.UserRepository) *UserServiceServer {
	return &UserServiceServer{userRepo: r}
}

func NewUserService(r *repository.UserRepository) *UserService {
	return &UserService{userRepo: r}
}

func (us *UserService) ListUsers() ([]dto.UserDTO, error) {
	logrus.WithFields(logrus.Fields{
		"service": "bz-main",
		"event":   "ListUsers",
	}).Debug("List Users - Start")

	defer logrus.WithFields(logrus.Fields{
		"service": "bz-main",
		"event":   "ListUsers",
	}).Debug("List Users - End")

	users, err := us.userRepo.ListUsers()
	if err != nil {
		return nil, err
	}

	var userDtos []dto.UserDTO
	copier.Copy(&userDtos, &users)

	return userDtos, nil
}

func (s *UserServiceServer) GetAllUsers(ctx context.Context, req *pb.GetAllUsersRequest) (*pb.GetAllUsersResponse, error) {

	logrus.WithFields(logrus.Fields{
		"service": "bz-main",
		"event":   "ListUsers",
	}).Debug("List Users - Start")

	defer logrus.WithFields(logrus.Fields{
		"service": "bz-main",
		"event":   "ListUsers",
	}).Debug("List Users - End")

	users, err := s.userRepo.ListUsers()

	if err != nil {
		return nil, err
	}

	var usersProto []*pb.User
	err = copier.Copy(&usersProto, &users)
	if err != nil {
		// handle the error
		return nil, err
	}

	return &pb.GetAllUsersResponse{Users: usersProto}, nil
}

func (us *UserService) UpdateUser(id string, userDTO *dto.UserDTO) error {
	logrus.WithFields(logrus.Fields{
		"service": "bz-main",
		"event":   "UpdateUser",
	}).Debug("Update User - Start")

	defer logrus.WithFields(logrus.Fields{
		"service": "bz-main",
		"event":   "UpdateUser",
	}).Debug("Update User - End")

	// Create a new instance of models.User
	var user models.User

	// Copy data from userDTO to the new user instance
	if err := copier.Copy(&user, userDTO); err != nil {
		logrus.WithFields(logrus.Fields{
			"service": "bz-main",
			"event":   "UpdateUser",
		}).Errorf("Failed to copy data from userDTO to user: %v", err)
		return err
	}

	// Call the repository's UpdateUser function with the created user model
	err := us.userRepo.UpdateUser(id, user)
	if err != nil {
		return err
	}

	return nil
}

func (us *UserService) DeleteUser(userID string) error {
	logrus.WithFields(logrus.Fields{
		"service": "bz-main",
		"event":   "DeleteUser",
	}).Debug("Delete User - Start")

	defer logrus.WithFields(logrus.Fields{
		"service": "bz-main",
		"event":   "DeleteUser",
	}).Debug("Delete User - End")

	err := us.userRepo.DeleteUser(userID)
	if err != nil {
		return err
	}

	return nil
}

func (us *UserService) CreateUser(userDTO *dto.UserDTO) error {
	logrus.WithFields(logrus.Fields{
		"service": "bz-main",
		"event":   "CreateUser",
	}).Debug("Create User - Start")

	defer logrus.WithFields(logrus.Fields{
		"service": "bz-main",
		"event":   "CreateUser",
	}).Debug("Create User - End")

	// Create a new instance of models.User
	var user models.User

	// Copy data from userDTO to the new user instance
	if err := copier.Copy(&user, userDTO); err != nil {
		logrus.WithFields(logrus.Fields{
			"service": "bz-main",
			"event":   "CreateUser",
		}).Errorf("Failed to copy data from userDTO to user: %v", err)
		return err
	}

	// Call the repository's CreateUser function with the created user model
	err := us.userRepo.CreateUser(user)
	if err != nil {
		return err
	}

	return nil
}

func (us *UserService) GetUsers(limit, offset int) ([]dto.UserDTO, error) {
	logrus.WithFields(logrus.Fields{
		"service": "bz-main",
		"event":   "GetUsers",
	}).Debug("Get Users - Start")

	defer logrus.WithFields(logrus.Fields{
		"service": "bz-main",
		"event":   "GetUsers",
	}).Debug("Get Users - End")

	users, err := us.userRepo.GetUsers(limit, offset)
	if err != nil {
		return nil, err
	}

	var userDtos []dto.UserDTO
	copier.Copy(&userDtos, &users)

	return userDtos, nil
}
