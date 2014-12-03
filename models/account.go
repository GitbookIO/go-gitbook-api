package models

// Account data structure as returned by the API
type Account struct {
	Name     string             `json:"name"`
	Email    string             `json:"email"`
	Token    string             `json:"token"`
	Username string             `json:"username"`
	GitHub   *GitHubAccountInfo `json:"github"`
}

type GitHubAccountInfo struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}
