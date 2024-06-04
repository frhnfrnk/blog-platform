package graphql

import (
	"context"

	"github.com/frhnfrnk/blog-platform-microservices/user-service/internal/services"
)

type Resolver struct {
	userService services.UserService
}

func NewResolver(userService services.UserService) *Resolver {
	return &Resolver{
		userService: userService,
	}
}

func (r *Resolver) GetUserByID(ctx context.Context, id string) (*User, error) {
	user, err := r.userService.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return &User{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

// Implementasi resolver untuk query getAllUsers, createUser, updateUser, dan deleteUser
// ...
