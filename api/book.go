package api

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/GitbookIO/go-gitbook-api/client"
	"github.com/GitbookIO/go-gitbook-api/models"
	"github.com/GitbookIO/go-gitbook-api/utils"

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

// Publish packages the desired book as a tar.gz and pushes it to gitbookio
// bookpath can be a path to a tar.gz file, git repo or folder
func (b *Book) Publish(bookId, version, bookpath string) error {
	basepath := filepath.Base(bookpath)

	if !exists(bookpath) {
		return fmt.Errorf("Book path '%s' does not exist", bookpath)
	}

	// Tar.gz
	if strings.HasSuffix(basepath, ".tar.gz") || strings.HasSuffix(basepath, ".tgz") {
		return b.PublishTarGz(bookId, version, bookpath)
	}

	// Git repo
	if isGitDir(bookpath) {
		return b.PublishGit(bookId, version, bookpath, "HEAD")
	} else if dir := path.Join(bookpath, ".git"); isGitDir(dir) {
		return b.PublishGit(bookId, version, dir, "HEAD")
	}

	// Standard folder
	return b.PublishFolder(bookId, version, bookpath)
}

// PublishGit packages a git repo as tar.gz and uploads it to gitbook.io
func (b *Book) PublishGit(bookId, version, bookpath, ref string) error {
	tar, err := utils.GitTarGz(bookpath, ref)
	if err != nil {
		return err
	}
	defer tar.Close()

	return b.PublishStream(bookId, version, tar)
}

// PublishFolder packages a folder as tar.gz and uploads it to gitbook.io
func (b *Book) PublishFolder(bookId, version, bookpath string) error {
	// Build tar from folder, exclude unwanted directories
	tar, err := utils.TarGzExclude(
		bookpath,

		// Excluded files & folders
		".git",
		"node_modules",
		"bower",
		"_book",
		"book.pdf",
		"book.mobi",
		"book.epub",
	)
	if err != nil {
		return err
	}
	defer tar.Close()

	return b.PublishStream(bookId, version, tar)
}

// PublishTarGz publishes a book based on a tar.gz file
func (b *Book) PublishTarGz(bookId, version, bookpath string) error {
	file, err := os.Open(bookpath)
	if err != nil {
		return err
	}
	defer file.Close()

	return b.PublishStream(bookId, version, file)
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

func isGitDir(dirpath string) bool {
	return (exists(path.Join(dirpath, "HEAD")) &&
		exists(path.Join(dirpath, "objects")) &&
		exists(path.Join(dirpath, "refs")))
}

// Does a file exist on disk ?
func exists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}
