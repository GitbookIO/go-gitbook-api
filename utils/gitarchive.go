package utils

import (
	"compress/gzip"
	"io"
	"os/exec"
)

// GitTar return a stream of tar data of the repo
// at a specific ref
func GitTar(dir, ref string) (io.ReadCloser, error) {
	// Build tar using git-archive
	args := []string{"git-archive", "--format=tar", ref}

	cmd := exec.Command(args[0], args[1:]...)
	// Set directory to repo's
	cmd.Dir = dir

	// Get stream
	return CmdStream(cmd, nil)
}

// GitTarGz returns a stream tar.gz data of the repo
func GitTarGz(dir, ref string) (io.ReadCloser, error) {
	reader, err := GitTar(dir, ref)
	if err != nil {
		return nil, err
	}

	// Create pipe for compression
	pipeReader, pipeWriter := io.Pipe()

	// Compress stuff
	gzipWriter := gzip.NewWriter(pipeWriter)

	// Flush data in async
	go func() {
		// Copy over data
		io.Copy(gzipWriter, reader)
		// Close writers
		gzipWriter.Close()
		pipeWriter.Close()
	}()

	return pipeReader, nil
}
