go-gitbook-api
==============

GitBook API client in GO (golang)

## Documentation

See [![GoDoc](https://godoc.org/github.com/GitbookIO/go-gitbook-api?status.png)](https://godoc.org/github.com/GitbookIO/go-gitbook-api)
for automatically generated API documentation.

Check out the **examples** below for quick and simple ways to start.

### Simple Example

```go
package main

import (
    "fmt"
    "github.com/GitbookIO/go-gitbook-api"
)

func main() {
    // Make API client
    api := gitbook.NewAPI(gitbook.APIOptions{})

    // Get book
    book, err := api.Book.Get("gitbookio/javascript")

    // Print results
    fmt.Printf("book = %q\n", book)
    fmt.Printf("book = %q\n", err)
}
```

### Advanced Example

```go
package main

import (
    "fmt"
    "github.com/GitbookIO/go-gitbook-api"
)

func main() {
    // Make API client
    api := gitbook.NewAPI(gitbook.APIOptions{
        // Custom host instead of "https://www.gitbook.io"
        Host: "http://localhost:5000",

        // Hit API with a specific user
        Username: "username",
        Password: "token or password",
    })

    // Get book
    book, err := api.Book.Get("gitbookio/javascript")

    // Print results
    fmt.Printf("book = %q\n", book)
    fmt.Printf("book = %q\n", err)
}
```
