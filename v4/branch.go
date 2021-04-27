package v4

type Branch struct {
	Name   string
	Commit Commit
}

type Commit struct {
	Author string `json:"author_name"`
}
