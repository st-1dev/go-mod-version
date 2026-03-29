package main

import (
	"fmt"
	"os"

	"github.com/go-git/go-git/v6"
	"github.com/go-git/go-git/v6/plumbing/object"
	"golang.org/x/mod/modfile"
)

func main() {
	exitCode, err := Main(os.Args[0])
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
	}
	os.Exit(exitCode)
}

func Main(program string) (exitCode int, err error) {
	var goModule string
	goModule, err = getGoModule()
	if err != nil {
		return 1, fmt.Errorf("could not get Go module from git repository: %w", err)
	}

	var gitHash string
	gitHash, err = getLatestCommitHash(".")
	if err != nil {
		return 1, fmt.Errorf("could not get latest commit hash from git repository: %w", err)
	}

	fmt.Printf("%s@%s\n", goModule, gitHash)
	return 0, nil
}

// getLatestCommitHash returns the hash of the latest commit in the repository.
func getLatestCommitHash(path string) (hash string, err error) {
	var repository *git.Repository
	repository, err = git.PlainOpen(path)
	if err != nil {
		return
	}

	var logs object.CommitIter
	logs, err = repository.Log(&git.LogOptions{})
	if err != nil {
		return
	}
	defer logs.Close()

	var commit *object.Commit
	commit, err = logs.Next()
	if err != nil {
		return
	}

	return commit.Hash.String(), nil
}

// getGoModule returns the Go module name.
func getGoModule() (goModule string, err error) {
	var content []byte
	content, err = os.ReadFile("go.mod")
	if err != nil {
		return "", fmt.Errorf("could not read module go.mod: %w", err)
	}
	return modfile.ModulePath(content), nil
}
