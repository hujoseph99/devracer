package db

import (
	"context"
	"testing"
	"time"
)

func getTestSnippet(t *testing.T) *Snippet {
	snippet := "func main() {\n\tfmt.Println(\"Hello World\")\n}"
	res := NewSnippet(snippet, LanguageGo, time.Now())

	return res
}

func TestSnippetAddAndDelete(t *testing.T) {
	testSnippet := getTestSnippet(t)

	collection := db.Database(DatabaseTypers).Collection(CollectionsSnippets)

	startingNum := getNumDocuments(t, collection)
	err := AddSnippet(context.Background(), testSnippet)
	changedNum := getNumDocuments(t, collection)

	if err != nil || changedNum != startingNum+1 {
		t.Fatal("Document was not added")
	}

	err = DeleteSnippetByID(context.Background(), testSnippet.ID)
	changedNum = getNumDocuments(t, collection)

	if err != nil || changedNum != startingNum {
		t.Fatal("Document was not deleted")
	}
}

func TestGetSnippetByID(t *testing.T) {
	testSnippet := getTestSnippet(t)

	err := AddSnippet(context.Background(), testSnippet)
	if err != nil {
		t.Fatal("Document was not added")
	}

	foundSnippet, err := GetSnippetByID(context.Background(), testSnippet.ID)

	// checking username and password is good enough for me
	if err != nil || foundSnippet.ID.Hex() != testSnippet.ID.Hex() ||
		foundSnippet.Language != testSnippet.Language ||
		foundSnippet.RaceContent != testSnippet.RaceContent {

		t.Fatal("Document was not found correctly")
	}

	err = DeleteSnippetByID(context.Background(), testSnippet.ID)
	if err != nil {
		t.Fatal("Document was not deleted")
	}
}

func TestGetRandomSnippet(t *testing.T) {
	_, err := GetRandomSnippet(context.Background())
	if err != nil {
		t.Fatal("Random document was not found")
	}
}
