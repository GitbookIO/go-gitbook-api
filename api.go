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
	// Builds API client
	Builds *api.Builds
	// User API client
	User *api.User

	// Internal client
	Client *client.Client
}

type APIOptions client.ClientOptions

func NewAPI(opts APIOptions) *API {
	c := client.NewClient(client.ClientOptions(opts))
	return NewAPIFromClient(c)
}

func NewAPIFromClient(c *client.Client) *API {
	return &API{
		Account: &api.Account{c},
		Book:    &api.Book{c},
		Books:   &api.Books{c},
		Books:   &api.Builds{c},
		User:    &api.User{c},
		Client:  c,
	}
}

func (a *API) Fork(opts APIOptions) *API {
	forkedClient := a.Client.Fork(client.ClientOptions(opts))
	return NewAPIFromClient(forkedClient)
}

func (a *API) AuthFork(username, password string) *API {
	forkedClient := a.Client.AuthFork(username, password)
	return NewAPIFromClient(forkedClient)
}
