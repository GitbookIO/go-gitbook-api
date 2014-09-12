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

type APIOptions client.ClientOptions

func NewAPI(opts APIOptions) *API {
	c := client.NewClient(client.ClientOptions(opts))

	return &API{
		Account: &api.Account{c},
		Book:    &api.Book{c},
		Books:   &api.Books{c},
		User:    &api.User{c},
		Client: c,
	}
}
