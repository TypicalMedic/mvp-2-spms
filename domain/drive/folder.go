package drive

type Folder struct {
	ParentFolder *Folder
	Name         string
}
