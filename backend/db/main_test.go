package db

import (
	"context"
	"os"
	"testing"
)

var client *Client

func TestMain(m *testing.M) {
	var err error
	client, err = ConnectToDB(context.TODO())

	if err != nil {
		os.Exit(0)
	}

	os.Exit(m.Run())
}
