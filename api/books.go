package api

import (
	"github.com/GitbookIO/go-gitbook-api/client"
	"github.com/GitbookIO/go-gitbook-api/models"
)

type Books struct {
	Client *client.Client
}

type booksListResponse struct {
	List []models.Book `json:"list"`
}

func (b *Books) List() ([]models.Book, error) {
	resp := booksListResponse{}

	if _, err := b.Client.Get(
		"/books",
		nil,
		&resp,
	); err != nil {
		return nil, err
	}

	return resp.List, nil
}
