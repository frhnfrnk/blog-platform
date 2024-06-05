package graphql

import (
	"context"

	"github.com/frhnfrnk/blog-platform-microservices/user-service/pb"
	"github.com/graphql-go/graphql"
)

type Resolver struct {
	userServiceClient pb.UserServiceClient
}

func NewResolver(userServiceClient pb.UserServiceClient) *Resolver {
	return &Resolver{
		userServiceClient: userServiceClient,
	}
}

var userType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"email": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

func (r *Resolver) resolveUsers(p graphql.ResolveParams) (interface{}, error) {
	response, err := r.userServiceClient.GetAllUsers(context.Background(), &pb.GetAllUsersRequest{})
	if err != nil {
		return nil, err
	}
	return response.Users, nil
}

func (r *Resolver) createUser(p graphql.ResolveParams) (interface{}, error) {
	name, _ := p.Args["name"].(string)
	email, _ := p.Args["email"].(string)

	response, err := r.userServiceClient.CreateUser(context.Background(), &pb.CreateUserRequest{Name: name, Email: email})
	if err != nil {
		return nil, err
	}
	return response.User, nil
}

func (r *Resolver) deleteUser(p graphql.ResolveParams) (interface{}, error) {
	id, _ := p.Args["id"].(string)

	response, err := r.userServiceClient.DeleteUser(context.Background(), &pb.DeleteUserRequest{Id: id})
	if err != nil {
		return nil, err
	}
	return response, nil
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

var RootMutation = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"createUser": &graphql.Field{
				Type: userType,
				Args: graphql.FieldConfigArgument{
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"email": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: r.createUser,
			},
			"deleteUser": &graphql.Field{
				Type: graphql.String,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: r.deleteUser,
			},
			"updateUser": &graphql.Field{
				Type: userType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"email": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: r.updateUser,
			},
		},
	},
)

var RootQuery = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"users": &graphql.Field{
				Type:    graphql.NewList(userType),
				Resolve: r.resolveUsers,
			},
		},
	},
)

var Schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query:    RootQuery,
		Mutation: RootMutation,
	},
)
