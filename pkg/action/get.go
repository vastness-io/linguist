package action

import (
	"github.com/Masterminds/vcs"
	"io/ioutil"
)

func NewRepoFileLister(repo vcs.Repo) (RepoFileLister, error) {

	vcsType := repo.Vcs()

	switch vcsType {
	case vcs.Git:
		return NewGitFileLister(repo.(*vcs.GitRepo)), nil
	default:
		return nil, vcs.ErrCannotDetectVCS
	}
}

func NewRepository(remoteURL string) (vcs.Repo, error) {
	local, _ := ioutil.TempDir("", "linguist")

	repo, err := vcs.NewRepo(remoteURL, local)

	if err != nil {
		return nil, err
	}

	return repo, nil
}

func CloneRepository(remoteURL string) (vcs.Repo, error) {
	repo, err := NewRepository(remoteURL)

	if err != nil {
		return nil, err
	}

	return repo, repo.Get()
}
