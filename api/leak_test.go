package api

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"sync"

	"testing"

	"github.com/GitbookIO/go-gitbook-api/client"
)

func TestLeaks(t *testing.T) {
	// Get current count
	c1 := openDescriptors()

	wg := &sync.WaitGroup{}

	// Do some work
	for i := 0; i < 100; i++ {
		go func() {
			wg.Add(1)
			c := client.NewClient(client.ClientOptions{
				Host:     "http://localhost:5000",
				Username: "aaronomullan",
				Password: "0c72ca47-5145-481d-bed8-d8a076d1b3ad",
			})
			b := Book{c}

			_, err := b.Get("aaronomullan/abc")
			if err != nil {
				t.Error(err)
			}
		}()

		t.Log(c1, "...", openDescriptors())
	}

	wg.Wait()

	// See how many files are open now
	c2 := openDescriptors()

	t.Log(c1, "=>", c2)

	// Check for leak
	if c2 > c1 {
		t.Errorf("Leak: %d to %d descriptors", c1, c2)
	}
}

func openDescriptors() int {
	out, err := lsof(os.Getpid())
	if err != nil {
		return 0
	}
	return bytes.Count(out, []byte("\n"))
}

func lsof(pid int) ([]byte, error) {
	return exec.Command(
		"bash", "-c",
		fmt.Sprintf("lsof -p '%d'", pid),
	).Output()
}
