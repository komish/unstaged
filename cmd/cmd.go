// Package cmd exposes the CLI for Unstaged.
//
// Copyright 2019 Jose R. Gonzalez. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.
package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/fatih/color"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/yaml.v2"
)

var (
	home       = homeDir()
	configPath = path.Join(home, ".unstaged.yaml")
)

// Repo is an indirect reference to a string.
type Repo *string

// Repolist is a slice containing multiple Repo types.
type Repolist []Repo

// RepoFileInput is used to unmarshal the yaml structure.
type RepoFileInput struct {
	R Repolist `yaml:"repos"`
}

// RepoWithErrors takes our Repo path and stores
// the error object encountered.
type RepoWithErrors struct {
	r Repo
	e error
}

// ReposWithErrors is a slice containing multiple RepoWithErrors types.
type ReposWithErrors []RepoWithErrors

// Run will execute the command line interface and
// then return an exit code.
func Run() int {
	// Bug(komish): Does not properly check against an upstream
	// to compare if changes have been pushed.
	var configExists = true
	yamlContents, err := ioutil.ReadFile(configPath)
	if err != nil {
		// TODO(komish): Change output streams
		configExists = false
	}
	if len(os.Args[1:]) < 1 && !configExists {
		fmt.Printf("Not enough arguments provided and the config file ")
		fmt.Printf("at path %s was not parsed successfully\n(Error: %s).\n\n", configPath, err)
		PrintHelp()
		return 15
	}

	i := os.Args[1:]
	var d RepoFileInput
	err = yaml.Unmarshal(yamlContents, &d)
	if err != nil {
		fmt.Println(err)
		os.Exit(9)
	}
	cliargs := stringToRepo(i)
	d.R = append(d.R, cliargs...)
	var problemRepos ReposWithErrors
	d.R, problemRepos = OpenReposAndFilter(d.R)
	for _, path := range d.R {
		repo, _ := git.PlainOpen(*path)
		wt, err := repo.Worktree()
		if err != nil {
			problemRepos = append(problemRepos, RepoWithErrors{path, err})
		}
		if stat, err := wt.Status(); err == nil {
			if stat.IsClean() {
				c := color.New(color.Bold, color.FgCyan)
				c.Printf("  CLEAN ")
				fmt.Println(*path)
			} else {
				c := color.New(color.Bold, color.FgYellow)
				c.Printf("UNCLEAN ")
				fmt.Println(*path)
			}
		}
	}

	for _, prepo := range problemRepos {
		c := color.New(color.Bold, color.FgRed)
		c.Printf("  ERROR ")
		fmt.Printf("%s ==> %s\n", *prepo.r, prepo.e)
	}
	return 0
}

// PrintHelp returns help output.
func PrintHelp() {
	fmt.Printf("Usage:\n %s [/path/to/repo] ...\n\n", os.Args[0])
	fmt.Println("Check through a list of provided git repositories and")
	fmt.Println("report back if they have unstaged changes Useful for.")
	fmt.Println("cases where knowledge bases are stored in source control")
	fmt.Println("and need to be pushed upstream frequently.")
}

// stringToRepo performs the indrecting of a String slice
// to a Repo slice and returns the resulting Repo slice.
func stringToRepo(s []string) Repolist {
	var r Repolist
	for _, e := range s {
		r = append(r, Repo(&e))
	}
	return r
}

// OpenReposAndFilter will check that we can open a path
// and that it represents a git repository.
func OpenReposAndFilter(u Repolist) (Repolist, ReposWithErrors) {
	// TODO(komish): We perform a git.PlainOpen() but don't do
	// anything with the resulting Repository. Redesign structs
	// to avoid having to do this again later to get a Worktree
	var (
		fr Repolist
		er ReposWithErrors
	)
	for _, path := range u {
		_, err := git.PlainOpen(*path)
		if err != nil {
			er = append(er, RepoWithErrors{path, err})
		} else {
			fr = append(fr, path)
		}
	}
	return fr, er
}

func homeDir() string {
	homedir, err := os.UserHomeDir()
	if err != nil {
		return "./"
	}
	return homedir
}
