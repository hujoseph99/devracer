package queries

import (
	"github.com/graphql-go/graphql"
)

var newPracticeRaceType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "NewPracticeRace",
		Fields: graphql.Fields{
			"snippet": &graphql.Field{
				Type: snippetType,
			},
			"timeLimit": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)
