package main

type RaceSnippet struct {
	Snippet string `bson:"snippet", json:"snippet"`
}

func main() {
	// // set client options -- get URI
	// clientOptions := options.Client().ApplyURI(mongoURI)

	// // Connet to MongoDB
	// client, err := mongo.Connect(context.TODO(), clientOptions)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// defer client.Disconnect(context.TODO())

	// // Check the connection
	// err = client.Ping(context.TODO(), nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println("Conneted to MongoDB!")

	// collection := client.Database("typers").Collection("racesnippets")

	// randomSnippet1 := RaceSnippet{Snippet: "randomSnippet1"}
	// randomSnippet2 := RaceSnippet{Snippet: "randomSnippet2"}
	// randomSnippet3 := RaceSnippet{Snippet: "randomSnippet3"}

	// insertResult, err := collection.InsertOne(context.TODO(), randomSnippet1)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println("Inserted a single document: ", insertResult.InsertedID)

	// manySnippets := []interface{}{randomSnippet2, randomSnippet3}
	// insertManyResult, err := collection.InsertMany(context.TODO(), manySnippets)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println("Inserted multiple documents: ", insertManyResult.InsertedIDs)

	// filter := bson.D{{"snippet", "randomSnippet1"}}

	// update := bson.D{
	// 	{"$set", bson.D{{"snippet", "updatedSnippet"}}},
	// }
	// updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)

}
