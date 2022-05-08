package main

import (
	"strings"
	"testing"
)

// CODE UNDER TEST:

// Interface for the GitHub "Contents API"
type ContentsAPI interface {
	GetReadme(owner, repo string) (string, error)
	GetContents(owner, repo, path string) (string, error)
	PutContents(owner, repo, path, contents string) error
	DeleteFile(owner, repo, path string) error
	GetArchiveLink(owner, repo, ref string) (string, error)
}

// Function under test: it calls the API's "GET readme" endpoint to fetch
// the data, and then returns the first non-empty line as the title. Note
// that it's written to take the ContentsAPI interface, but it only uses
// the GetReadme() method.
func GetReadmeTitle(api ContentsAPI, owner, repo string) (string, error) {
	readme, err := api.GetReadme(owner, repo)
	if err != nil {
		return "", err
	}
	for _, line := range strings.Split(readme, "\n") {
		line = strings.TrimSpace(line)
		if line != "" {
			return line, nil
		}
	}
	return "", nil
}

// TESTS:

// We create a fakeAPI type that embeds the entire ContentsAPI using an
// interface field, and overrides only the relevant method, GetReadme().
type fakeAPI struct {
	ContentsAPI
	readmes map[string]string // maps owner/repo name to readme text
}

func (a *fakeAPI) GetReadme(owner, repo string) (string, error) {
	return a.readmes[owner+"/"+repo], nil
}

func TestGetReadmeTitle(t *testing.T) {
	api := &fakeAPI{readmes: map[string]string{
		"my/project": "\nMy Project\n==========\n\nThis is a project.",
	}}
	title, err := GetReadmeTitle(api, "my", "project")
	if err != nil {
		t.Fatalf("error fetching readme: %v", err)
	}
	expected := "My Project"
	if title != expected {
		t.Errorf("got title %q, expected %q", title, expected)
	}
}
