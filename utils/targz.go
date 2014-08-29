package utils

import (
	"io"
	"os/exec"
)

var (
	tarArgs = []string{"tar", "-cz", "-", "."}
)

// TarGz returns a stream of tar.gz data of the directory
func TarGz(dir string) (io.ReadCloser, error) {
	return tarCommand(tarArgs, dir)
}

// TarGzExclude returns a stream of tar.gz data of the directory
// excluding the specified files
func TarGzExclude(dir string, exclude ...string) (io.ReadCloser, error) {
	return tarCommand(tarExcludeArgs(exclude...), dir)
}

// Run tar command for a folder given the provided args
func tarCommand(args []string, dir string) (io.ReadCloser, error) {
	cmd := exec.Command(args[0], args[1:]...)

	// Set target directory
	cmd.Dir = dir

	// Get stream
	return CmdStream(cmd, nil)
}

// Generate args for excluding files
func tarExcludeArgs(files ...string) []string {
	excluding := []string{"tar"}

	for _, f := range files {
		excluding = append(excluding, "--exclude", f)
	}

	return append(
		excluding,
		"-cz", "-", ".",
	)
}
