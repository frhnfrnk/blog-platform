package user

import (
	"context"

	"github.com/frhnfrnk/blog-platform-microservices/user-service/pb"
	"github.com/graphql-go/graphql"
)

// Resolver struct
type Resolver struct {
	userServiceClient pb.UserServiceClient
}

// NewResolver creates a new Resolver
func NewResolver(userServiceClient pb.UserServiceClient) *Resolver {
	return &Resolver{
		userServiceClient: userServiceClient,
	}
}

func (r *Resolver) resolveUsers(p graphql.ResolveParams) (interface{}, error) {
	response, err := r.userServiceClient.GetAllUser(context.Background(), &pb.GetAllUserRequest{})
	if err != nil {
		return nil, err
	}
	return response.Users, nil
}

func (r *Resolver) resolveUserByID(params graphql.ResolveParams) (interface{}, error) {
	id, _ := params.Args["id"].(string)

	user, err := r.userServiceClient.GetUser(context.Background(), &pb.GetUserRequest{Id: id})
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *Resolver) createUser(p graphql.ResolveParams) (interface{}, error) {
	name, _ := p.Args["name"].(string)
	email, _ := p.Args["email"].(string)

	response, err := r.userServiceClient.CreateUser(context.Background(), &pb.CreateUserRequest{Name: name, Email: email})
	if err != nil {
		return nil, err
	}
	return response.CreatedUser, nil
}

func (r *Resolver) deleteUser(p graphql.ResolveParams) (interface{}, error) {
	id, _ := p.Args["id"].(string)

	response, err := r.userServiceClient.DeleteUser(context.Background(), &pb.DeleteUserRequest{Id: id})
	if err != nil {
		return nil, err
	}
	return response.Message, nil
}

func (r *Resolver) updateUser(p graphql.ResolveParams) (interface{}, error) {
	id, _ := p.Args["id"].(string)
	name, _ := p.Args["name"].(string)
	email, _ := p.Args["email"].(string)

	response, err := r.userServiceClient.UpdateUser(context.Background(), &pb.UpdateUserRequest{Id: id, Name: name, Email: email})
	if err != nil {
		return nil, err
	}
	return response.UpdatedUser, nil
}
