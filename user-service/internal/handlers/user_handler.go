package handlers

import (
	"context"
	"github.com/frhnfrnk/blog-platform-microservices/user-service/internal/models"
	"strconv"

	"github.com/frhnfrnk/blog-platform-microservices/user-service/internal/services"
	"github.com/frhnfrnk/blog-platform-microservices/user-service/pb"
)

type UserHandler struct {
	userService *services.UserService
	pb.UnimplementedUserServiceServer
}

func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	user := &models.User{
		Name:  req.GetName(),
		Email: req.GetEmail(),
	}

	if err := h.userService.CreateUser(user); err != nil {
		return nil, err
	}

	return &pb.CreateUserResponse{
		CreatedUser: &pb.User{
			Id:    strconv.Itoa(int(user.ID)),
			Name:  user.Name,
			Email: user.Email,
		},
	}, nil
}

func (h *UserHandler) GetAllUser(context.Context, *pb.GetAllUserRequest) (*pb.GetAllUsersResponse, error) {
	users, err := h.userService.GetAllUsers()
	if err != nil {
		return nil, err
	}

	var pbUsers []*pb.User
	for _, user := range users {
		pbUser := &pb.User{
			Id:    strconv.Itoa(int(user.ID)),
			Name:  user.Name,
			Email: user.Email,
		}
		pbUsers = append(pbUsers, pbUser)
	}

	return &pb.GetAllUsersResponse{Users: pbUsers}, nil
}

func (h *UserHandler) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	user, err := h.userService.GetUserByID(req.Id)
	if err != nil {
		return nil, err
	}

		userID := strconv.Itoa(int(user.ID))
	return &pb.GetUserResponse{
		User: &pb.User{
			Id:    userID,
			Name:  user.Name,
			Email: user.Email,
		},
	}, nil
}

func (h *UserHandler) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	// Find the user by ID
	user, err := h.userService.GetUserByID(req.GetId())
	if err != nil {
		return nil, err
	}

	// Update user details
	user.Name = req.GetName()
	user.Email = req.GetEmail()

	// Save the updated user to the database
	if err := h.userService.UpdateUser(user); err != nil {
		return nil, err
	}

	// Return the updated user
	return &pb.UpdateUserResponse{
		UpdatedUser: &pb.User{
			Id:    strconv.Itoa(int(user.ID)),
			Name:  user.Name,
			Email: user.Email,
		},
	}, nil
}

func (h *UserHandler) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	// Delete the user by ID
	if err := h.userService.DeleteUser(req.GetId()); err != nil {
		return nil, err
	}

	// Return success message
	return &pb.DeleteUserResponse{
		Message: "User deleted successfully",
	}, nil
}
