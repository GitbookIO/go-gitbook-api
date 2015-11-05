package api

import (
	"github.com/GitbookIO/go-gitbook-api/client"
	"github.com/GitbookIO/go-gitbook-api/models"
)

type Account struct {
	Client *client.Client
}

// Get returns a books details for a given "bookId"
// (for example "gitbookio/javascript")
func (a *Account) Get() (models.Account, error) {
	account := models.Account{}

	_, err := a.Client.Get(
		"/account",
		nil,
		&account,
	)

	return account, err
}
