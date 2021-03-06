package queries

import (
	"github.com/graphql-go/graphql"
)

// RootQuery is the root query object for our graphql schema
var RootQuery = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{
			"hello": &graphql.Field{
				Type: worldType,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					newWorld := &world{
						message: "Hello World",
					}
					return newWorld, nil
				},
			},
		},
	})
