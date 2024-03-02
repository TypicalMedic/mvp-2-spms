package github

import (
	"context"
	"time"

	"github.com/google/go-github/v56/github"
)

// думаю есть возможность сделать как у гугла генерацию токена
const AUTH_TOKEN = "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"

type githubAPI struct {
	Context context.Context
	Client  *github.Client
}

func InitGithubAPI() githubAPI {
	context := context.Background()
	client := github.NewClient(nil).WithAuthToken(AUTH_TOKEN)

	return githubAPI{Context: context, Client: client}
}

func (g *githubAPI) GetRepoBranchCommitsFromTime(owner, repoName string, fromTime time.Time, branch string) ([]*github.RepositoryCommit, error) {
	opt := &github.CommitsListOptions{SHA: branch, Since: fromTime}
	commits, _, err := g.Client.Repositories.ListCommits(g.Context, owner, repoName, opt)
	if err == nil {
		return commits, nil
	}
	return nil, err
}

func (g *githubAPI) GetRepoBranches(owner, repoName string) ([]*github.Branch, error) {
	opt := &github.BranchListOptions{}
	branches, _, err := g.Client.Repositories.ListBranches(g.Context, owner, repoName, opt)
	if err == nil {
		return branches, nil
	}
	return nil, err
}
