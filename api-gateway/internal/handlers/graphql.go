package graphql

import (
	"net/http"

	"github.com/frhnfrnk/blog-platform-microservices/api-gateway/graphql"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

func NewGraphQLHandler(schema Schema) http.Handler {
	return handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})
}

func GraphQLHandler(resolver *Resolver) http.Handler {
	schemaConfig := graphql.SchemaConfig{
		Query:    RootQuery,
		Mutation: RootMutation,
	}
	schema, err := graphql.NewSchema(schemaConfig)
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
