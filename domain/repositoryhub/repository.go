package repositoryhub

type Repository struct {
	Id       uint
	Name     string
	IsPublic bool
	Commits  []Commit
	Branches []Branch
}
