package action

type RepoFile struct {
	Name string
	Size int
}

type FileNames []RepoFile

type RepoFileLister interface {
	GetFilesFromCurrent() (FileNames, error)
}
