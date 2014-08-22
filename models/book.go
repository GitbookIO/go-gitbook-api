package models

// Books data structure as returned by the API
type Book struct {
	Id          string            `json:"id"`
	Name        string            `json:"name"`
	Title       string            `json:"title"`
	Description string            `json:"description"`
	Urls        map[string]string `json:"urls"`

	Author struct {
		Name     string `json:"name"`
		Username string `json:"username"`
	} `json:"author"`

	Permissions struct {
		Read   bool `json:"read"`
		Write  bool `json:"write"`
		Manage bool `json:"manage"`
	} `json:"permissions"`

	LatestBuild struct {
		Version  string `json:"version"`
		Finished string `json:"finished"`
		Started  string `json:"started"`
	} `json:"latestBuild"`
}
