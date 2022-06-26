package models

type Repo struct {
	Owner       string `json:"owner"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func NewRepo(owner string, name string, description string) Repo {
	if description == "" {
		description = "This is a auto-generated repository"
	}

	return Repo{
		Owner:       owner,
		Name:        name,
		Description: description,
	}
}
