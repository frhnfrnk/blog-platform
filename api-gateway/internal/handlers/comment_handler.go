package handlers

import (
	"github.com/frhnfrnk/blog-platform-microservices/api-gateway/internal/graphql/comment"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

func NewCommentHandler(schema graphql.Schema) http.Handler {
	return handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})
}

func CommentHandler(resolver *comment.Resolver) http.Handler {
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query:    comment.NewRootQuery(resolver),
		Mutation: comment.NewRootMutation(resolver),
	})
	if err != nil {
		panic(err)
	}

	h := handler.New(&handler.Config{
		Schema:     &schema,
		Pretty:     true,
		GraphiQL:   true,
		Playground: true,
	})

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.ContextHandler(r.Context(), w, r)
	})
}
