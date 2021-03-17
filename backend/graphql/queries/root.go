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
			"practiceRace": &graphql.Field{
				Type: newPracticeRaceType,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					snippetID, ok := p.Args["snippetId"]
					if !ok {
						// return random snippet
					}

					snippetIDString, ok := snippetID.(string)
					if !ok {
						// not a string
					}

					// get snippet by id
					return snippetIDString, nil
				},
			},
		},
	})
