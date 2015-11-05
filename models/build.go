package models

type Build struct {
	Version string      `json:"version"`
	Branch  string      `json:"branch"`
	Message string      `json:"message"`
	Author  BuildAuthor `json:"author"`
}

type BuildAuthor struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
