package action

import (
	"github.com/Masterminds/vcs"
	"github.com/sirupsen/logrus"
)

var (
	log = logrus.WithFields(logrus.Fields{
		"pkg": "action",
	})
)

type RepoFile struct {
	Name string
	Size int
}

type RepoFiles []RepoFile

type RepoFileLister interface {
	VCS() vcs.Type
	GetFilesFromCurrent() (RepoFiles, error)
}
