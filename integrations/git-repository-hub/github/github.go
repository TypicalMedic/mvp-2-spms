package github

import (
	"log"
	entities "mvp-2-spms/domain-aggregate"
	"time"

	"github.com/google/go-github/v56/github"
)

type Github struct {
	api githubAPI
}

func InitGithub(api githubAPI) Github {
	return Github{api: api}
}

func (g *Github) GetRepositoryCommitsFromTime(repo entities.ProjectInRepository, fromTime time.Time) []entities.Commit {
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
	commits := []entities.Commit{}
	for _, ghcommit := range ghCommitsUnique {
		cm := mapCommitToEntity(*ghcommit)
		commits = append(commits, cm)
	}
	return commits
}

func mapCommitToEntity(commit github.RepositoryCommit) entities.Commit {
	return entities.Commit{
		SHA:         *commit.SHA,
		Description: *commit.Commit.Message,
		Date:        commit.Commit.Committer.Date.Time,
		Author:      *commit.Commit.Committer.Name,
	}
}
