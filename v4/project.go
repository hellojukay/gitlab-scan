package v4

import "fmt"

// Project gitlab project
type Project struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	WebURL    string `json:"web_url"`
	Namespace Group  `json:"namespace"`
}

func (p Project) String() string {
	return fmt.Sprintf("project %s id is %d link url %s\n", p.Name, p.ID, p.WebURL)
}
