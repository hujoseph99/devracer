package graphql

import (
	"log"

	"github.com/gorilla/mux"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/hujoseph99/typing/backend/graphql/queries"
	"github.com/hujoseph99/typing/backend/secret"
)

// RegisterEndpoints registers the endpoints for graphql
func RegisterEndpoints(router *mux.Router) {
	schemaConfig := graphql.SchemaConfig{Query: queries.RootQuery}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatal("Schema was unable to be started up")
	}

	graphqlHandler := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: !secret.Production,
	})

	router.Handle("/graphql", graphqlHandler)
}
