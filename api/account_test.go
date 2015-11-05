package api

import (
	"testing"

	"github.com/GitbookIO/go-gitbook-api/client"
)

func TestAccount(t *testing.T) {
	c := client.NewClient(client.ClientOptions{
		Host:     "http://localhost:5000/api/",
		Username: "james",
		Password: "730e0de8-ca53-42f9-9ad3-49c547b0cdc0",
	})
	a := Account{c}

	_, err := a.Get()
	if err != nil {
		t.Error(err)
	}
}
