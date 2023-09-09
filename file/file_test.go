package file

import (
	"os"
	"path/filepath"
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

func TestGetConfigDir(t *testing.T) {
	name := "GetConfigDir()"

	// Read the config directory
	dir, err := GetConfigDir("test")
	if err != nil {
		t.Fatalf(`%v: unable to set config dir (%v): %v`, name, dir, err)
	}

	// Verify that we got the expected home directory
	// TODO: Unset the XDG_CONFIG_PATH env
	userDir, _ := os.UserHomeDir()
	wantDir := filepath.Join(userDir, ".config", "test")
	gotDir := dir
	if gotDir != wantDir {
		t.Fatalf(`%v: want string of %v got %v`, name, wantDir, gotDir)
	}

}
