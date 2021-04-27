package tree

import (
	v4 "github.com/hellojukay/gitlab/v4"
)

type Node struct {
	Name     string
	Branches []v4.Branch
	Sub      []Node
}

func (node Node) String() string {
	return ""
}
