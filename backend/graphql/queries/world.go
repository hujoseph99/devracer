package queries

import (
	"github.com/graphql-go/graphql"
)

type world struct {
	message string
}

var worldType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "World",
		Fields: graphql.Fields{
			"message": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					world, ok := p.Source.(*world)
					if ok {
						return world.message, nil
					}
					return nil, nil
				},
			},
		},
	})
