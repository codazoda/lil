package lil

import (
	"testing"

	"golang.org/x/exp/slices"
)

func TestFileToSlice(t *testing.T) {
	name := "FileToSlice()"

	// Read the file into an array
	repos, err := FileToSlice("testfile/file.txt")
	if err != nil {
		t.Fatalf(`%v: unable to read file.txt: %v`, name, err)
	}

	// Verify that we got 6 elements
	wantLen := 6
	gotLen := len(repos)
	if gotLen != wantLen {
		t.Fatalf(`%v: want length of %v got %v`, name, wantLen, gotLen)
	}

	// Verify that "one" is there
	want := "one"
	got := slices.Contains(repos, want)
	if got != true {
		t.Fatalf(`%v: want true for %v got %v`, name, want, got)
	}

	// Verify that "six" is there
	want = "six"
	got = slices.Contains(repos, want)
	if got != true {
		t.Fatalf(`%v: want true for %v got %v`, name, want, got)
	}

}

func TestFile(t *testing.T) {
	name := "File()"

	// Read the file into an array
	repos, err := File("testfile/file.txt")
	if err != nil {
		t.Fatalf(`%v: unable to read file.txt: %v`, name, err)
	}

	// Verify that we got 6 elements
	wantLen := 6
	gotLen := len(repos)
	if gotLen != wantLen {
		t.Fatalf(`%v: want length of %v got %v`, name, wantLen, gotLen)
	}

	// Verify that "one" is there
	want := "one"
	got := slices.Contains(repos, want)
	if got != true {
		t.Fatalf(`%v: want true for %v got %v`, name, want, got)
	}

	// Verify that "six" is there
	want = "six"
	got = slices.Contains(repos, want)
	if got != true {
		t.Fatalf(`%v: want true for %v got %v`, name, want, got)
	}

}
