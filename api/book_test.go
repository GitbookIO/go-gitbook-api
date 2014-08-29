package api

import (
	"bytes"

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

	buf := &bytes.Buffer{}
	buf.Write([]byte("heelo"))

	if err := b.PublishStream("aaronomullan/some-paid-book", "99.55.5", buf); err != nil {
		t.Error(err)
	}
}
