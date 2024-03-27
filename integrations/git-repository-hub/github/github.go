package github

import (
	"log"
	"sort"
	"time"

	"mvp-2-spms/services/models"

	"github.com/google/go-github/v56/github"
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

func (g *Github) GetAuthLink(redirectURI string) string

func (g *Github) Authentificate(token string)

func mapCommitToEntity(commit github.RepositoryCommit) models.Commit {
	return models.Commit{
		SHA:         *commit.SHA,
		Description: *commit.Commit.Message,
		Date:        commit.Commit.Committer.Date.Time,
		Author:      *commit.Commit.Author.Name,
	}
}
