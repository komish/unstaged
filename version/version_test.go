package version

import (
	"testing"
)

// TestVersionVar checks that the defined placeholder for Version doesn't change.
func TestVersionVar(t *testing.T) {
	got := Version
	want := "0.0.0"
	if got != want {
		t.Errorf("Got %s Wanted %s", got, want)
	}
}

// TestCommitHashVar checks that the defined placeholder for CommitHash doesn't change.
func TestCommitHashVar(t *testing.T) {
	got := CommitHash; 	want := "-"
	if got != want {
		t.Errorf("Got %s Wanted %s", got, want)
	}
}
