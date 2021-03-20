package graphql

import (
	"context"
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/hujoseph99/typing/backend/api"
	"github.com/hujoseph99/typing/backend/graphql/queries"
)

// RegisterEndpoints registers the endpoints for graphql
func RegisterEndpoints(api *api.API) {
	schemaConfig := graphql.SchemaConfig{Query: queries.RootQuery}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatal("Schema was unable to be started up")
	}

	graphqlHandler := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})

	// passes the database conntext through the context in each graphql request
	api.Router.
		HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), queries.ContextKey(queries.DatabaseContextKey), api.Database)
			graphqlHandler.ContextHandler(ctx, w, r)
		})
}
