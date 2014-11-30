package client

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"

	"github.com/jmcvetta/napping"
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
	}

	// We want JSON responses (for errors especially)
	session.Header.Set("Accept", "application/json")

	return &Client{
		Session:       session,
		ClientOptions: &opts,
	}
}

// Fork creates a new client off of the base client
// however it shares the same http.Client for efficiency reasons
// this prevents socket leaks from happening etc ...
func (c *Client) Fork(opts ClientOptions) *Client {
	if opts.Host == "" {
		opts.Host = c.Host
	}
	if opts.Username == "" {
		opts.Username = c.Username
	}
	if opts.Password == "" {
		opts.Password = c.Password
	}

	session := &napping.Session{
		Userinfo: url.UserPassword(opts.Username, opts.Password),
		Header:   c.Session.Header,
		Client:   c.Session.Client,
	}

	return &Client{
		Session:       session,
		ClientOptions: &opts,
	}
}

// AuthFork is a shorthand of Fork, when you simply want to change the auth
func (c *Client) AuthFork(username, password string) *Client {
	return c.Fork(ClientOptions{
		Username: username,
		Password: password,
	})
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

func DecodeError(reader io.Reader) (*Error, error) {
	errMsg := &Error{}
	decoder := json.NewDecoder(reader)
	err := decoder.Decode(errMsg)
	if err != nil {
		// Failed to decode, error must be string not JSON
		data, err := ioutil.ReadAll(decoder.Buffered())
		if err != nil {
			return nil, err
		}
		return &Error{
			Msg:  string(data[:]),
			Code: 500,
		}, nil
	}
	return errMsg, nil
}
