package graphql

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/hujoseph99/typing/backend/graphql/queries"
	"github.com/hujoseph99/typing/backend/secret"
)

// CORS Middleware, have to change this in the future to be more secure
func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		if r.Method == "OPTIONS" {
			w.WriteHeader((http.StatusOK))
			return
		}
		next.ServeHTTP(w, r)
	})
}

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

	router.Handle(
		"/graphql",
		CorsMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			graphqlHandler.ServeHTTP(w, r)
		})))
}
