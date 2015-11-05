package api

import (
	"testing"

	"github.com/GitbookIO/go-gitbook-api/client"
)

func TestBooksList(t *testing.T) {
	c := client.NewClient(client.ClientOptions{
		Host:     "http://localhost:5000/api/",
		Username: "james",
		Password: "730e0de8-ca53-42f9-9ad3-49c547b0cdc0",
	})
	b := Books{c}

	books, err := b.List()
	if err != nil {
		t.Error(err)
	}

	if len(books) < 1 {
		t.Errorf("Should have at least one book, found %d instead", len(books))
	}
}
