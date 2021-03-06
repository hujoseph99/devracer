package queries

import (
	"github.com/graphql-go/graphql"
)

type world struct {
	message string
}

var RootQuery = graphql.NewObject(
	graphql.ObjectConfig{Name: "RootQuery", Fields: graphql.Fields{
		Type: worldType,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			newWorld := &world{
				message: "Hello World",
			}
			return newWorld, nil
		},
	}},
)
