package models

// Account data structure as returned by the API
// Account extends Author with private fields
type Account struct {
	Author

	Email  string             `json:"email"`
	Token  string             `json:"token"`
	GitHub *GitHubAccountInfo `json:"github"`
}

type GitHubAccountInfo struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}
