package models

// Books data structure as returned by the API
type Book struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Urls        struct {
		Access    string `json:"access"`
		Homepage  string `json:"homepage"`
		Read      string `json:"read"`
		Reviews   string `json:"reviews"`
		Subscribe string `json:"subscribe"`

		Download struct {
			Epub string `json:"epub"`
			Mobi string `json:"mobi"`
			Pdf  string `json:"pdf"`
		} `json:"download"`
	} `json:"urls"`

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
