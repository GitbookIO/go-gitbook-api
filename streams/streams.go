package streams

import (
	"io"
	"os"
	"path"
	"path/filepath"

	"github.com/GitbookIO/go-gitbook-api/utils"
)

type StreamFunc func(path string) (io.ReadCloser, error)

func PickStream(p string) (io.ReadCloser, error) {
	basepath := filepath.Base(p)

	if !exists(p) {
		return nil, fmt.Errorf("PickStream: Path '%s' does not exist", p)
	}

	// Tar.gz
	if strings.HasSuffix(basepath, ".tar.gz") || strings.HasSuffix(basepath, ".tgz") {
		return File(p)
	}

	// Git repo
	if isGitDir(p) {
		return GitHead(p)
	} else if dir := path.Join(p, ".git"); isGitDir(dir) {
		return GitHead(dir)
	}

	// Standard folder
	return Folder(p)
}

func GitHead(p string) (io.ReadCloser, error) {
	return Git(p, "HEAD")
}

// Git returns an io.ReadCloser of a repo as a tar.gz
func Git(p, ref string) (io.ReadCloser, error) {
	return utils.GitTarGz(p, ref)
}

func Folder(p string) (io.ReadCloser, error) {
	return utils.TarGzExclude(
		p,

		// Excluded files & folders
		".git",
		"node_modules",
		"bower",
		"_book",
		"book.pdf",
		"book.mobi",
		"book.epub",
	)
}

func File(p string) (io.ReadCloser, error) {
	return os.Open(p)
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
