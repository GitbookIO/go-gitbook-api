package api

import (
	"testing"

	"github.com/GitbookIO/go-gitbook-api/client"
)

func TestBasic(t *testing.T) {
	c := client.NewClient(client.ClientOptions{
		Host:     "http://localhost:5000",
		Username: "aaronomullan",
		Password: "0c72ca47-5145-481d-bed8-d8a076d1b3ad",
	})
	b := Book{c}

	if err := b.Publish("aaronomullan/some-paid-book", "6.6.6", "/Users/aaron/git/jsbook"); err != nil {
		t.Error(err)
	}
}
