package queries

import (
	"github.com/graphql-go/graphql"
)

var profileType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Profile",
		Fields: graphql.Fields{
			"totalWordsTyped": &graphql.Field{
				Type: graphql.Int,
			},
			"racesCompleted": &graphql.Field{
				Type: graphql.Int,
			},
			"racesWon": &graphql.Field{
				Type: graphql.Int,
			},
			"maxTPM": &graphql.Field{
				Type: graphql.Float,
			},
			"averageTPMAllTime": &graphql.Field{
				Type: graphql.Float,
			},
			"averageTPMLast10": &graphql.Field{
				Type: graphql.Float,
			},
		},
	},
)
