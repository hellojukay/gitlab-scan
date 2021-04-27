package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/hellojukay/gitlab/v4/client"
	"github.com/xlab/treeprint"
)

var api string
var token string
var clt client.Client
var group int64

func init() {
	flag.StringVar(&api, "api", "", "gitlab api: https://gitlab.com/api/v4/")
	flag.StringVar(&token, "token", "", "gitlab person access token")
	flag.Int64Var(&group, "group", 1, "gitlab group id")
	flag.Parse()
	if len(api) == 0 {
		flag.PrintDefaults()
		os.Exit(1)
	}
	if !strings.HasSuffix(api, "/") {
		api = api + "/"
	}
	clt = client.New(api, token)
}
func main() {
	if !clt.Ping() {
		os.Exit(1)
	}
	tree := buildTree(group)
	fmt.Println(tree)
}

func buildTree(id int64) treeprint.Tree {
	g, err := clt.Group(id)
	if err != nil {
		return treeprint.New()
	}
	tree := treeprint.NewWithRoot(g.Name)
	projects, err := clt.Projects(*g)
	if err != nil {
		log.Printf("can not get projects from %s, %s", api, err)
		os.Exit(1)
	}
	for _, project := range projects {
		node := tree.AddBranch(project.Name)
		branches, err := clt.Branches(project)
		if err != nil {
			log.Printf("project %s can not get branches ", project.WebURL)
			continue
		}
		for _, branch := range branches {
			node.AddNode(fmt.Sprintf("%s      %s", branch.Name, branch.Commit.Author))
		}
	}
	groups, err := clt.Subgroups(*g)
	if err != nil {
		log.Printf("can not get subgroups, %s", err)
	}
	if len(groups) == 0 {
		return tree
	}
	for _, group := range groups {
		t := buildTree(group.ID)
		tree.AddMetaBranch(group.Name, t)
	}
	return tree
}
