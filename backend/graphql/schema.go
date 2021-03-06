package graphql

import (
	"log"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/hujoseph99/typing/backend/api"
	"github.com/hujoseph99/typing/backend/graphql/queries"
)

// func executeQuery(query string, schema graphql.Schema) *graphql.Result {
// 	result := graphql.Do(graphql.Params{
// 		Schema:        schema,
// 		RequestString: query,
// 	})
// 	if len(result.Errors) > 0 {
// 		fmt.Printf("wrong result, unexpected errors: %v", result.Errors)
// 	}
// 	return result
// }

// func handleGraphql(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	result := executeQuery(r.URL.Query().Get("query"), schema)
// 	json.NewEncoder(w).Encode(result)
// }

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
