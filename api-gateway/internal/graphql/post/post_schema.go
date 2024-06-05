package post

import "github.com/graphql-go/graphql"

var postType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Post",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"title": &graphql.Field{
				Type: graphql.String,
			},
			"content": &graphql.Field{
				Type: graphql.String,
			},
			"authorID": &graphql.Field{
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
				"posts": &graphql.Field{
					Type:    graphql.NewList(postType),
					Resolve: resolver.resolvePosts,
				},
				"post": &graphql.Field{
					Type: postType,
					Args: graphql.FieldConfigArgument{
						"id": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.String),
						},
					},
					Resolve: resolver.resolvePostByID,
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
				"createPost": &graphql.Field{
					Type: postType,
					Args: graphql.FieldConfigArgument{
						"title": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.String),
						},
						"content": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.String),
						},
						"authorID": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.String),
						},
					},
					Resolve: resolver.createPost,
				},
				"updatePost": &graphql.Field{
					Type: postType,
					Args: graphql.FieldConfigArgument{
						"id": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.String),
						},
						"title": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.String),
						},
						"content": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.String),
						},
					},
					Resolve: resolver.updatePost,
				},
				// Mutation for deleting a post
				"deletePost": &graphql.Field{
					Type: graphql.String,
					Args: graphql.FieldConfigArgument{
						"id": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.String),
						},
					},
					Resolve: resolver.deletePost,
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
