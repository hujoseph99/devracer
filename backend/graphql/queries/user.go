package queries

import (
	"github.com/graphql-go/graphql"
)

var userType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"profile": &graphql.Field{
				Type: profileType,
			},
			"preferences": &graphql.Field{
				Type: preferencesType,
			},
		},
	},
)
