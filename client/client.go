package client

import (
	"github.com/jmcvetta/napping"
	"net/url"
	"path"
)

type Client struct {
	*napping.Session
	*ClientOptions
}

type ClientOptions struct {
	// Hostname of gitbookio endpoint
	Host string

	// Auth info
	Username string
	Password string
}

func NewClient(opts ClientOptions) *Client {
	return &Client{
		Session: napping.Session{
			Userinfo: url.UserPassword(opts.Username, opts.Password),
		},
		ClientOptions: &opts,
	}
}

func (c *Client) Books()

// Wrap napping.Session.Send
// So that we insert the client's host in the URL
func (c *Client) Send(r *napping.Request) (*napping.Response, error) {
	r.Url = path.Join(c.Host, r.Url)
	return c.Session.Send(r)
}
