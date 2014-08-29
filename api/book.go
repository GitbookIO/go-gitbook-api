package api

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/GitbookIO/go-gitbook-api/client"
	"github.com/GitbookIO/go-gitbook-api/models"

	"mime/multipart"
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

// PublishStream
func (b *Book) PublishStream(bookId, version string, r io.Reader) error {
	// Build request
	req, err := newfileUploadRequest(
		b.Client.Url(fmt.Sprintf("/api/book/%s/builds", bookId)),
		// No params
		nil,
		"book",
		r,
	)
	if err != nil {
		return err
	}

	uinfo := b.Client.Userinfo

	// Auth
	pwd, _ := uinfo.Password()
	req.SetBasicAuth(uinfo.Username(), pwd)

	// Set version
	values := url.Values{}
	values.Set("version", version)
	req.URL.RawQuery = values.Encode()

	// Execute request
	_, err = b.Client.Client.Do(req)
	return err
}

// Creates a new file upload http request with optional extra params
func newfileUploadRequest(uri string, params map[string]string, paramName string, reader io.Reader) (*http.Request, error) {
	// Buffer for body
	body := &bytes.Buffer{}
	// Multipart data
	writer := multipart.NewWriter(body)

	// File part
	part, err := writer.CreateFormFile(paramName, "book.tar.gz")
	if err != nil {
		return nil, err
	}

	// Copy over data for file
	_, err = io.Copy(part, reader)
	if err != nil {
		return nil, err
	}

	// Write extra fields
	for key, val := range params {
		_ = writer.WriteField(key, val)
	}
	err = writer.Close()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", uri, body)
	if err != nil {
		return nil, err
	}

	// Set header
	req.Header.Set("Content-Type", writer.FormDataContentType())

	return req, nil
}
