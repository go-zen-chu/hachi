// go:generate mockgen -source=$GOFILE -destination=mock_$GOFILE -package=$GOPACKAGE
package gh

import (
	"context"
	"errors"
	"fmt"
	"log"
	"path"
	"time"

	"github.com/google/go-github/v26/github"
	"golang.org/x/oauth2"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/http"
)

// GitHubClient is interface for working with GitHub
type GitHubClient interface {
	Clone(repoURI string, dir string) (*git.Repository, error)
	Commit(r *git.Repository, msg string) error
	Push(r *git.Repository) error
	PullRequest(owner, repo, title, head, body, baseBranch string) (string, error)
}

type ghclient struct {
	client *github.Client
	user   string
	mail   string
	token  string
}

// NewClient create github client
func NewClient(baseURL string, token string, user string, mail string) (GitHubClient, error) {
	// validation
	if baseURL == "" {
		return nil, errors.New("need to set baseURL")
	}
	uploadURL := path.Join(baseURL, "upload")
	if token == "" {
		return nil, errors.New("need to set token")
	}
	if user == "" || mail == "" {
		return nil, errors.New("need to set user, mail for git operation")
	}
	// create client
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	cli, err := github.NewEnterpriseClient(baseURL, uploadURL, tc)
	if err != nil {
		return nil, fmt.Errorf("creating github enterprise client: %s", err)
	}
	ghc := &ghclient{
		client: cli,
		user:   user,
		mail:   mail,
		token:  token,
	}
	return ghc, nil
}

// Clone is function of 'git clone'
func (c *ghclient) Clone(repoURI string, dir string) (*git.Repository, error) {
	o := &git.CloneOptions{
		Auth: &http.BasicAuth{
			Username: c.user,
			Password: c.token,
		},
		URL: repoURI,
	}
	return git.PlainClone(dir, false, o)
}

// Commit is function of 'git commit'
func (c *ghclient) Commit(r *git.Repository, msg string) error {
	w, err := r.Worktree()
	if err != nil {
		return err
	}
	o := &git.CommitOptions{
		Author: &object.Signature{
			Name:  c.user,
			Email: c.mail,
			When:  time.Now(),
		},
	}
	_, err = w.Commit(msg, o)
	return err
}

// Push is function of 'git push'
func (c *ghclient) Push(r *git.Repository) error {
	o := &git.PushOptions{
		Auth: &http.BasicAuth{
			Username: c.user,
			Password: c.token,
		},
		// TODO: add refspec
	}
	return r.Push(o)
}

// PullRequest is function which create new pull request
func (c *ghclient) PullRequest(owner, repo, title, head, body, baseBranch string) (string, error) {
	pr := github.NewPullRequest{
		Title: &title,
		Head:  &head,
		Base:  &baseBranch,
		Body:  &body,
	}
	ctx := context.Background()
	r, _, err := c.client.PullRequests.Create(ctx, owner, repo, &pr)
	if err != nil {
		return "", fmt.Errorf("creating PullRequest : %s", err)
	}
	prURL := r.GetHTMLURL()
	log.Printf("PullRequest created: %s", prURL)
	return prURL, nil
}
