package v4

import "fmt"

type Group struct {
	ID   int64
	Name string
}

func (g Group) String() string {
	return fmt.Sprintf("project %s id is %d\n", g.Name, g.ID)
}
