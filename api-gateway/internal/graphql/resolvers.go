package graphql

import (
	"github.com/frhnfrnk/blog-platform-microservices/user-service/pb"
)

type Resolver struct {
	userServiceClient pb.UserServiceClient
}

func NewResolver(userServiceClient pb.UserServiceClient) *Resolver {
	return &Resolver{
		userServiceClient: userServiceClient,
	}
}
