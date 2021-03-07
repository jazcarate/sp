// Playground with go-git
package main

import (
	"fmt"
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
)

func checkIfError(err error) {
	if err == nil {
		return
	}

	fmt.Printf("\x1b[31;1m%s\x1b[0m\n", fmt.Sprintf("error: %s", err))
	os.Exit(1)
}

func foo() {
	r, err := git.PlainOpen("./workspace")
	checkIfError(err)

	ref, err := r.Head()
	checkIfError(err)

	cIter, err := r.Log(&git.LogOptions{From: ref.Hash()})
	checkIfError(err)

	err = cIter.ForEach(func(c *object.Commit) error {
		fmt.Println(c.Author.Email)
		fmt.Println(c.Message)

		return nil
	})
	checkIfError(err)
}
