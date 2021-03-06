package graphql

import (
	"github.com/graphql-go/graphql"
	"github.com/hujoseph99/typing/backend/graphql/queries"
)

var schemaConfig = graphql.SchemaConfig{Query: queries.RootQuery}
