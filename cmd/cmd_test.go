// Copyright 2019 Jose R. Gonzalez. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.
package cmd

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOpenReposAndFilter(t *testing.T) {

	git_repo := "../.git"
	not_git_repo := "./hello"
	not_even_dir := "lol"

	repoList_work := Repolist{
		&git_repo,
	}

	repoList_not_repo := Repolist{
		&not_git_repo,
	}

	repoList_not_url := Repolist{
		&not_even_dir,
	}

	fr, er := OpenReposAndFilter(repoList_work)
	assert.Equal(t, 0, len(er))
	assert.EqualValues(t, repoList_work, fr)

	errorString := errors.New("repository does not exist")
	fr, er = OpenReposAndFilter(repoList_not_repo)
	assert.Equal(t, 1, len(er))
	assert.EqualValues(t, errorString, er[0].e)

	fr, er = OpenReposAndFilter(repoList_not_url)
	assert.Equal(t, 1, len(er))
	assert.EqualValues(t, errorString, er[0].e)

}

func TestRun(t *testing.T) {
	i := Run()
	assert.Equal(t, 0, i)
}
