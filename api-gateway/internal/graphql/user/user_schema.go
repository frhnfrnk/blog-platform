package user

import (
	"github.com/graphql-go/graphql"
)

// Define GraphQL user type
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

func NewRootQuery(resolver *Resolver) *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"users": &graphql.Field{
					Type:    graphql.NewList(userType),
					Resolve: resolver.resolveUsers,
				},
				"user": &graphql.Field{
					Type: userType,
					Args: graphql.FieldConfigArgument{
						"id": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.String),
						},
					},
					Resolve: resolver.resolveUserByID,
				},
			},
		},
	)
}

func NewRootMutation(resolver *Resolver) *graphql.Object {
	return graphql.NewObject(
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
					Resolve: resolver.createUser,
				},
				"deleteUser": &graphql.Field{
					Type: graphql.String,
					Args: graphql.FieldConfigArgument{
						"id": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.String),
						},
					},
					Resolve: resolver.deleteUser,
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
					Resolve: resolver.updateUser,
				},
			},
		},
	)
}

func NewSchema(resolver *Resolver) (graphql.Schema, error) {
	return graphql.NewSchema(
		graphql.SchemaConfig{
			Query:    NewRootQuery(resolver),
			Mutation: NewRootMutation(resolver),
		},
	)
}
