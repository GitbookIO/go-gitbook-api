package api

import (
	"fmt"

	"github.com/GitbookIO/go-gitbook-api/client"
	"github.com/GitbookIO/go-gitbook-api/models"
)

type Book struct {
	Client *client.Client
}

// Get returns a books details for a given "bookId"
// (for example "gitbookio/javascript")
func (b *Book) Get(bookId string) (models.Book, error) {
	book := models.Book{}

	_, err := b.Client.Get(
		fmt.Sprintf("/api/book/%s", bookId),
		nil,
		&book,
	)

	return book, err
}
