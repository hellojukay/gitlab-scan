package v4

import (
	"time"
)

type Branch struct {
	Name   string
	Commit Commit
}

type Commit struct {
	Author        string    `json:"author_name"`
	CommittedDate time.Time `json:"committed_date"`
}
