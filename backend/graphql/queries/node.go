package queries

import (
	"github.com/graphql-go/graphql"
)

var nodeType = graphql.NewInterface(
	graphql.InterfaceConfig{
		Name: "Node",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.ID,
				// TODO: Resolver - If following GUID thing, add __typename in front of given id
			},
		},
	},
)
