package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/textproto"

	"github.com/GitbookIO/go-gitbook-api/client"
	"github.com/GitbookIO/go-gitbook-api/models"
	"github.com/GitbookIO/go-gitbook-api/streams"

	"mime/multipart"
)

type Builds struct {
	Client *client.Client
}

// BuildOpts are optional data passed along when doing a build (e.g: branch, message, author, ...)
type BuildOpts models.Build

type postStream func(bookId, version, branch string, r io.Reader) error

// Build should only be used by internal clients, Publish by others
// Build starts a build and will not update the backing git repository
func (b *Builds) Build(bookId, version, path string, opts BuildOpts) error {
	return b.doStreamPublish(bookId, version, path, opts, streams.PickStream)
}

// PublishGit packages a git repo as tar.gz and uploads it to gitbook.io
func (b *Builds) BuildGit(bookId, version, path, ref string, opts BuildOpts) error {
	return b.doStreamPublish(bookId, version, path, opts, streams.GitRef(ref))
}

// PublishFolder packages a folder as tar.gz and uploads it to gitbook.io
func (b *Builds) BuildFolder(bookId, version, path string, opts BuildOpts) error {
	return b.doStreamPublish(bookId, version, path, opts, streams.Folder)
}

// PublishTarGz publishes a book based on a tar.gz file
func (b *Builds) BuildTarGz(bookId, version, path string, opts BuildOpts) error {
	return b.doStreamPublish(bookId, version, path, opts, streams.File)
}

func (b *Builds) doStreamPublish(bookId, version, path string, opts BuildOpts, streamfn streams.StreamFunc) error {
	stream, err := streamfn(path)
	if err != nil {
		return err
	}
	defer stream.Close()

	return b.PublishBuildStream(bookId, version, stream, opts)
}

func (b *Builds) PublishBuildStream(bookId, version string, reader io.Reader, opts BuildOpts) error {
	return b.publishStream(
		fmt.Sprintf("/book/%s/build/%s", bookId, version),
		version,
		reader,
		opts,
	)
}

// PublishStream
func (b *Builds) publishStream(_url, version string, reader io.Reader, opts BuildOpts) error {
	// Build request
	req, err := newfileUploadRequest(
		b.Client.Url(_url),
		opts,
		reader,
	)
	if err != nil {
		return err
	}

	uinfo := b.Client.Userinfo

	// Auth
	pwd, _ := uinfo.Password()
	req.SetBasicAuth(uinfo.Username(), pwd)

	// Execute request
	response, err := b.Client.Client.Do(req)
	if err != nil {
		return err
	}
	// Close body immediately to avoid leaks
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		data, _ := ioutil.ReadAll(response.Body)
		return fmt.Errorf(string(data[:]))
	}

	// Some error to code
	if response.StatusCode >= 400 {
		errMsg, err := client.DecodeError(response.Body)
		if err != nil {
			return err
		}
		return errMsg
	}

	return nil
}

// Creates a new file upload http request with optional extra params
func newfileUploadRequest(uri string, opts BuildOpts, reader io.Reader) (*http.Request, error) {
	// Buffer for body
	body := &bytes.Buffer{}
	// Multipart data
	writer := multipart.NewWriter(body)

	// Write JSON metadata
	metadataPart, err := writer.CreatePart(textproto.MIMEHeader{
		"Content-Disposition": {"form-data; name=metadata"},
		"Content-Type":        {"application/json"},
	})
	if err != nil {
		return nil, err
	}
	metadataPart.Write([]byte(jsonString(opts)))

	// File part
	part, err := writer.CreateFormFile("book", "book.tar.gz")
	if err != nil {
		return nil, err
	}

	// Copy over data for file
	_, err = io.Copy(part, reader)
	if err != nil {
		return nil, err
	}

	// Close writer
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

func jsonString(v interface{}) string {
	if data, err := json.Marshal(v); err == nil {
		return string(data)
	}
	return ""
}
