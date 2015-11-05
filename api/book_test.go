package api

import (
	"testing"

	"github.com/GitbookIO/go-gitbook-api/client"
)

func TestBasic(t *testing.T) {
	c := client.NewClient(client.ClientOptions{
		Host:     "http://localhost:5000/api/",
		Username: "aaronomullan",
		Password: "0c72ca47-5145-481d-bed8-d8a076d1b3ad",
	})
	b := Book{c}

	_, err := b.Get("aaronomullan/abc")
	if err != nil {
		t.Error(err)
	}
}
