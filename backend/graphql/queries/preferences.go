package queries

import (
	"github.com/graphql-go/graphql"
)

var preferencesType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Preferences",
		Fields: graphql.Fields{
			"displayName": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)
