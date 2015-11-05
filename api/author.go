package api

import (
	"fmt"

	"github.com/GitbookIO/go-gitbook-api/client"
	"github.com/GitbookIO/go-gitbook-api/models"
)

type Author struct {
	Client *client.Client
}

func (a *Author) Get(username string) (models.Author, error) {
	author := models.Author{}

	_, err := a.Client.Get(
		fmt.Sprintf("/author/%s", username),
		nil,
		&author,
	)

	return author, err
}
