package main

import (
	"context"
	"log"
	"net/http"

	"github.com/hujoseph99/typingBackend/auth"

	"github.com/hujoseph99/typingBackend/api"
)

func main() {
	ctx := context.TODO()
	myAPI, err := api.NewAPI(ctx)
	if err != nil {
		log.Fatal(err)
	}
	myAPI.SetupRouter()
	auth.RegisterAuthEndpoints(myAPI)
	http.ListenAndServe(":8080", myAPI.Router)

	// ctx := context.Background()

	// client, _ := db.ConnectToDB(ctx)

	// newUser := db.NewUser("email", "foo", "bar", "pw", 50, time.Now())
	// err := client.AddUser(ctx, newUser)

	// delID, _ := newUser.ID.(string)
	// err = client.DeleteUserByID(ctx, delID)

	// if err != nil {
	// 	log.Fatal(err)
	// }
	// newSnippet := db.NewRaceSnippet("hello test snippet")
	// res, err := client.AddRaceSnippet(ctx, newSnippet)

	// err = client.DeleteRaceSnippetByID(ctx, res)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// findUser, err := client.FindUserByEmail(ctx, "random@gmail.com")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(findUser)

	// id := "5da995caa4b5ab52e8458670"
	// snippet, err := client.GetRaceSnippetByID(ctx, id)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(snippet)

	// id = "5db820a77a1b9513ec733b3f"
	// user, err := client.FindUserByID(ctx, id)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(user)

	// snippet, err = client.GetRandomRaceSnippet(ctx)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(snippet)
}
