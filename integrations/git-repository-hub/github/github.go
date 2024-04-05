package github

import (
	"encoding/base64"
	"fmt"
	"log"
	"sort"
	"time"

	"mvp-2-spms/services/models"

	"github.com/google/go-github/v56/github"
	"golang.org/x/oauth2"
)

type Github struct {
	api githubAPI
}

func InitGithub(api githubAPI) *Github {
	return &Github{api: api}
}

func (g *Github) GetRepositoryCommitsFromTime(repo models.Repository, fromTime time.Time) []models.Commit {
	// to get all commits we need to check all the branches
	ghbranches, err := g.api.GetRepoBranches(repo.OwnerName, repo.RepoId)
	if err != nil {
		log.Fatal(err)
	}

	// finding all branches commits
	ghAllBranchesCommits := []*github.RepositoryCommit{}
	for _, branch := range ghbranches {
		ghbrcommits, err := g.api.GetRepoBranchCommitsFromTime(repo.OwnerName, repo.RepoId, fromTime, *branch.Name)
		if err != nil {
			log.Fatal(err)
		}
		ghAllBranchesCommits = append(ghAllBranchesCommits, ghbrcommits...)
	}

	// throwing away repeated commits (brnches might have the same history)
	ghCommitsUnique := map[string]*github.RepositoryCommit{}
	for _, c := range ghAllBranchesCommits {
		ghCommitsUnique[*c.SHA] = c
	}

	// transforming to entity
	commits := []models.Commit{}
	for _, ghcommit := range ghCommitsUnique {
		cm := mapCommitToEntity(*ghcommit)
		commits = append(commits, cm)
	}

	// sorting by publication date
	sort.Slice(commits, func(i, j int) bool {
		return commits[i].Date.Unix() > commits[j].Date.Unix()
	})

	return commits
}

func (g *Github) GetAuthLink(redirectURI string, accountId int, returnURL string) string {
	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	// encode as JSON!
	statestr := base64.URLEncoding.EncodeToString([]byte(fmt.Sprint(accountId, ",", returnURL)))
	url := g.api.GetAuthLink(redirectURI, statestr)
	return url
}

func (g *Github) Authentificate(token *oauth2.Token) {
	g.api.SetupClient(token)
}

func (g *Github) GetToken(code string) *oauth2.Token {
	token := g.api.GetToken(code)
	return token
}

func mapCommitToEntity(commit github.RepositoryCommit) models.Commit {
	return models.Commit{
		SHA:         *commit.SHA,
		Description: *commit.Commit.Message,
		Date:        commit.Commit.Committer.Date.Time,
		Author:      *commit.Commit.Author.Name,
	}
}
