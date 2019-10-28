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

	gitRepo := "../.git"
	notGitRepo := "./hello"
	notEvenDir := "lol"

	repoListWork := Repolist{
		&gitRepo,
	}

	repoListNotRepo := Repolist{
		&notGitRepo,
	}

	repoListNotURL := Repolist{
		&notEvenDir,
	}

	fr, er := OpenReposAndFilter(repoListWork)
	assert.Equal(t, 0, len(er))
	assert.EqualValues(t, repoListWork, fr)

	errorString := errors.New("repository does not exist")
	fr, er = OpenReposAndFilter(repoListNotRepo)
	assert.Equal(t, 1, len(er))
	assert.EqualValues(t, errorString, er[0].e)

	fr, er = OpenReposAndFilter(repoListNotURL)
	assert.Equal(t, 1, len(er))
	assert.EqualValues(t, errorString, er[0].e)

}

func TestRun(t *testing.T) {
	i := Run()
	assert.Equal(t, 0, i)
}
