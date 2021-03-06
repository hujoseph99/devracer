package graphql

import (
	"log"

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

	graphiqlHandler := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})

	api.Router.
		Path("/graphql").
		Handler(graphiqlHandler)
}
