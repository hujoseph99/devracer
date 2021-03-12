package queries

import "github.com/graphql-go/graphql"

var snippet = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Snippet",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.ID,
			},
			"raceContent": &graphql.Field{
				Type: graphql.String,
			},
			"wordCount": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)
