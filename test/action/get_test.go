package action

import (
	"github.com/Masterminds/vcs"
	"github.com/vastness-io/linguist/pkg/action"
	"io/ioutil"
	"strings"
	"testing"
)

const repoURL = "https://github.com/vastness-io/language-detection-data.git"

func TestNewRepoFileLister(t *testing.T) {

	repo := getTestRepo(t, repoURL)
	defer deleteTestRepo(t, repo)

	var fileListerTests = []struct {
		expectedType  vcs.Type
		expectedError error
	}{
		{
			expectedType:  vcs.Git,
			expectedError: nil,
		},
	}

	for _, test := range fileListerTests {

		repoFileLister, err := action.NewRepoFileLister(repo)

		if err != test.expectedError {
			t.Errorf("expected: %s, got: %s", test.expectedError, err)
		}

		if repoFileLister.VCS() != test.expectedType {
			t.Errorf("expected: %s, got: %s", test.expectedType, repoFileLister.VCS())
		}

	}
}

func TestNewRepository(t *testing.T) {

	repo, err := action.NewRepository(repoURL)
	defer deleteTestRepo(t, repo)

	if err != nil {
		t.Fail()
	}

	if !strings.Contains(repo.LocalPath(), "linguist") {
		t.Error("Local path should be set to a temp direction with prefix")
	}

}

func TestCloneRepository(t *testing.T) {

	vcsTypeTests := []struct {
		expectedType    vcs.Type
		remoteUrl       string
		repoExistsCheck func(t *testing.T, repo vcs.Repo)
	}{
		{
			expectedType: vcs.Git,
			remoteUrl:    repoURL,
			repoExistsCheck: func(t *testing.T, repo vcs.Repo) {
				b, err := repo.RunFromDir("git", "rev-parse", "--is-bare-repository")

				if err != nil {
					t.Error(err)
				}

				isBare := string(b)

				if strings.TrimSpace(isBare) != "true" {
					t.Error("Repo is not bare")
				}
				deleteTestRepo(t, repo)
			},
		},
	}

	for _, test := range vcsTypeTests {

		repo, err := action.CloneRepository(test.remoteUrl)

		if err != nil {
			t.Error(err)
		}

		_, err = ioutil.ReadDir(repo.LocalPath())

		if err != nil {
			t.Error("Dir should exist")
		}

		test.repoExistsCheck(t, repo)

	}

}

func TestGetFilesFromCurrent(t *testing.T) {

	vcsTypeTests := []struct {
		expectedType vcs.Type
		remoteUrl    string
	}{
		{
			expectedType: vcs.Git,
			remoteUrl:    repoURL,
		},
	}

	for _, test := range vcsTypeTests {

		repo, err := action.CloneRepository(test.remoteUrl)

		if err != nil {
			t.Fail()
		}

		fileLister, err := action.NewRepoFileLister(repo)

		if err != nil {
			t.Error(err)
		}

		if fileLister.VCS() != test.expectedType {
			t.Errorf("vcs type mismatch")
		}

		files, err := fileLister.GetFilesFromCurrent()

		if err != nil {
			t.Error(err)
		}

		for _, file := range files {

			if file.Name == "" {
				t.Error("Name can't be empty")
			}

			if file.Size <= 0 {
				t.Error("Size can't be below 1")
			}
		}

		deleteTestRepo(t, repo)
	}
}
