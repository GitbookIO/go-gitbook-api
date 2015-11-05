package api

import (
	"testing"

	"github.com/GitbookIO/go-gitbook-api/client"
)

func TestBuildsCreate(t *testing.T) {
	c := client.NewClient(client.ClientOptions{
		Host:     "stupid_host",
		Username: "badboy",
		Password: "password",
	})
	b := Builds{c}

	err := b.BuildGit("james/test", "master", "/Users/aaron/git/documentation", "master", BuildOpts{
		Branch: "master",
	})
	if err != nil {
		t.Error(err)
	}
}
