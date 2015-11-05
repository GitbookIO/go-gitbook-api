package models

type Build struct {
	Branch  string      `json:"branch"`
	Message string      `json:"message"`
	Author  BuildAuthor `json:"author"`
}

type BuildAuthor struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
