package client

import (
	"fmt"
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
		Session: &napping.Session{
			Userinfo: url.UserPassword(opts.Username, opts.Password),
		},
		ClientOptions: &opts,
	}
}

func (c *Client) Delete(url string, result, errMsg interface{}) (*napping.Response, error) {
	return c.Session.Delete(c.Url(url), result, errMsg)
}

func (c *Client) Get(url string, p *napping.Params, result, errMsg interface{}) (*napping.Response, error) {
	return c.Session.Get(c.Url(url), p, result, errMsg)
}

func (c *Client) Head(url string, result, errMsg interface{}) (*napping.Response, error) {
	return c.Session.Head(c.Url(url), result, errMsg)
}

func (c *Client) Options(url string, result, errMsg interface{}) (*napping.Response, error) {
	return c.Session.Options(c.Url(url), result, errMsg)
}

func (c *Client) Patch(url string, payload, result, errMsg interface{}) (*napping.Response, error) {
	return c.Session.Patch(c.Url(url), payload, result, errMsg)
}

func (c *Client) Post(url string, payload, result, errMsg interface{}) (*napping.Response, error) {
	return c.Session.Post(c.Url(url), payload, result, errMsg)
}

func (c *Client) Put(url string, payload, result, errMsg interface{}) (*napping.Response, error) {
	return c.Session.Put(c.Url(url), payload, result, errMsg)
}

// Url returns the full http url including host
func (c *Client) Url(urlpath string) string {
	parsed, _ := url.Parse(c.Host)
	parsed.Path = path.Join(parsed.Path, urlpath)
	str := parsed.String()
	fmt.Printf("Going to '%s'\n", str)
	return str
}
