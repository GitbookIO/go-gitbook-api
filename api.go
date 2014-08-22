package gitbook

import (
	"github.com/GitbookIO/go-gitbook-api/api"
	"github.com/GitbookIO/go-gitbook-api/client"
)

type API struct {
	// Authentication API client
	Account *api.Account
	// Individual book API client
	Book *api.Book
	// Book listing API client
	Books *api.Books
	// User API client
	User *api.User

	// Internal client
	Client *client.Client
}

type APIOpts ClientOpts

func NewAPI(opts APIOpts) *API {
	c := client.NewClient(opts)
	return &API{
		Account: &api.Account{c},
		Book:    &api.Book{c},
		Books:   &api.Books{c},
		User:    &api.User{c},
	}
}
