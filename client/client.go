package client

import (
	"github.com/jmcvetta/napping"
	"net/http"
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
	// Default hostname
	if opts.Host == "" {
		opts.Host = "https://www.gitbook.io"
	}

	// Setup session
	// for authentication and custom headers
	session := &napping.Session{
		Userinfo: url.UserPassword(opts.Username, opts.Password),
		Header:   &http.Header{},
		Client:   &http.Client{},

		// Authorize use of client to HTTP endpoints
		UnsafeBasicAuth: true,
	}

	// We want JSON responses (for errors especially)
	session.Header.Set("Accept", "application/json")

	return &Client{
		Session:       session,
		ClientOptions: &opts,
	}
}

func (c *Client) Delete(url string, result interface{}) (*napping.Response, error) {
	return errorPatch(func(errMsg *Error) (*napping.Response, error) {
		return c.Session.Delete(c.Url(url), result, errMsg)
	})
}

func (c *Client) Get(url string, p *napping.Params, result interface{}) (*napping.Response, error) {
	return errorPatch(func(errMsg *Error) (*napping.Response, error) {
		return c.Session.Get(c.Url(url), p, result, errMsg)
	})
}

func (c *Client) Head(url string, result interface{}) (*napping.Response, error) {
	return errorPatch(func(errMsg *Error) (*napping.Response, error) {
		return c.Session.Head(c.Url(url), result, errMsg)
	})
}

func (c *Client) Options(url string, result interface{}) (*napping.Response, error) {
	return errorPatch(func(errMsg *Error) (*napping.Response, error) {
		return c.Session.Options(c.Url(url), result, errMsg)
	})
}

func (c *Client) Patch(url string, payload, result interface{}) (*napping.Response, error) {
	return errorPatch(func(errMsg *Error) (*napping.Response, error) {
		return c.Session.Patch(c.Url(url), payload, result, errMsg)
	})
}

func (c *Client) Post(url string, payload, result interface{}) (*napping.Response, error) {
	return errorPatch(func(errMsg *Error) (*napping.Response, error) {
		return c.Session.Post(c.Url(url), payload, result, errMsg)
	})
}

func (c *Client) Put(url string, payload, result interface{}) (*napping.Response, error) {
	return errorPatch(func(errMsg *Error) (*napping.Response, error) {
		return c.Session.Put(c.Url(url), payload, result, errMsg)
	})
}

// Url returns the full http url including host
func (c *Client) Url(urlpath string) string {
	// Ignore errors for now
	parsed, _ := url.Parse(c.Host)

	// Rewrite path
	parsed.Path = path.Join(parsed.Path, urlpath)

	// Return string URL
	return parsed.String()
}

// This is so we include API errors as well as protocol errors here
func errorPatch(f func(err *Error) (*napping.Response, error)) (*napping.Response, error) {
	errMsg := &Error{}
	resp, err := f(errMsg)
	// API error
	if err == nil && errMsg.Code != 0 {
		return resp, errMsg
	}
	// Normal or protcol error
	return resp, err
}
