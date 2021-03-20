package queries

import "github.com/graphql-go/graphql"

var snippetType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Snippet",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.ID,
			},
			"raceContent": &graphql.Field{
				Type: graphql.String,
			},
			"tokenCount": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)
