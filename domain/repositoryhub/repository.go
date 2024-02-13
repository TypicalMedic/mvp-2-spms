package repositoryhub

type Repository struct {
	Name     string
	IsPublic bool
	Commits  []Commit
	Branches []Branch
}
