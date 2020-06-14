package gh

import (
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
)

// Add is function of 'git add'
func Add(r *git.Repository, file string) error {
	w, err := r.Worktree()
	if err != nil {
		return err
	}
	_, err = w.Add(file)
	return err
}

// AddBranch is function of 'git branch ${branch}'
func AddBranch(r *git.Repository, branch string) error {
	headRef, err := r.Head()
	if err != nil {
		return err
	}
	refName := plumbing.ReferenceName("refs/heads/" + branch)
	ref := plumbing.NewHashReference(refName, headRef.Hash())
	return r.Storer.SetReference(ref)
}

// Checkout is function of 'git checkout ${branch}'
func Checkout(r *git.Repository, branch string) error {
	w, err := r.Worktree()
	if err != nil {
		return err
	}
	refName := plumbing.ReferenceName("refs/heads/" + branch)
	return w.Checkout(&git.CheckoutOptions{
		Branch: refName,
	})
}
