package action

import (
	"github.com/Masterminds/vcs"
	"github.com/vastness-io/linguist/pkg/action"
	"os"
	"testing"
)

func getTestRepo(t *testing.T, remoteUrl string) vcs.Repo {
	repo, err := action.NewRepository(remoteUrl)

	if err != nil {
		t.Error(err)
	}

	return repo
}

func deleteTestRepo(t *testing.T, repo vcs.Repo) {
	if err := os.RemoveAll(repo.LocalPath()); err != nil {
		t.Errorf("Failed to remove: %s", repo.LocalPath())
	}
}
