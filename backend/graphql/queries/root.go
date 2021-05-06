package queries

import (
	"fmt"

	"github.com/graphql-go/graphql"
	"github.com/hujoseph99/typing/backend/db"
	"github.com/hujoseph99/typing/backend/graphql/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// RootQuery is the root query object for our graphql schema
var RootQuery = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{
			"practiceRace": &graphql.Field{
				Type: newPracticeRaceType,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					snippetID, ok := p.Args["snippetId"]
					if !ok {
						// could not extract snippet id, return random snippet
						snippet, err := db.GetRandomSnippet(p.Context)
						if err != nil {
							return nil, fmt.Errorf("cannot get random snippet")
						}
						return models.NewNewPracticeRace(snippet, models.GetTimeLimit(snippet)), nil
					}

					// get snippet by id
					snippetIDString, ok := snippetID.(string)
					if !ok {
						return nil, fmt.Errorf("given invalid id, id is: %v", snippetID)
					}
					objectID, err := primitive.ObjectIDFromHex(snippetIDString)
					if err != nil {
						return nil, fmt.Errorf("given invalid id, id is: %v", snippetID)
					}
					snippet, err := db.GetSnippetByID(p.Context, objectID)
					if err != nil {
						return nil, fmt.Errorf("cannot get the snippet by id")
					}

					return models.NewNewPracticeRace(snippet, models.GetTimeLimit(snippet)), nil
				},
			},
		},
	})
