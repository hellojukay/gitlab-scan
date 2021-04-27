package main

import (
	"fmt"
	"log"
	"os"

	"github.com/hellojukay/gitlab/v4/client"
	"github.com/xlab/treeprint"
)

var api = os.Getenv("API")
var clt = client.New(api, os.Getenv("TOKEN"))

func main() {

	if !clt.Ping() {
		log.Fatalln("can not connect to gitlab")
	} else {
		log.Printf("auth %s success", api)
	}
	tree := buildTree(5802)
	fmt.Println(tree)
}

func buildTree(id int64) treeprint.Tree {
	g, err := clt.Group(id)
	if err != nil {
		return treeprint.New()
	}
	projects, err := clt.Projects(*g)
	if err != nil {
		log.Printf("can not get projects from %s, %s", api, err)
		os.Exit(1)
	}
	tree := treeprint.New()
	for _, project := range projects {
		node := tree.AddBranch(project.Name)
		branches, err := clt.Branches(project)
		if err != nil {
			log.Printf("project %s can not get branches ", project.WebURL)
			continue
		}
		for _, branch := range branches {
			node.AddNode(fmt.Sprintf("%s      (%s)", branch.Name, branch.Commit.Author))
		}
	}
	groups, err := clt.Subgroups(*g)
	if err != nil {
		log.Println("can not get subgroups")
	}
	if len(groups) == 0 {
		return tree
	}
	for _, group := range groups {
		t := buildTree(group.ID)
		tree.AddMetaBranch(t, group.Name)
	}
	return tree
}
