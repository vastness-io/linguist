package action

import (
	"errors"
	"github.com/Masterminds/vcs"
	"github.com/sirupsen/logrus"
	"strconv"
	"strings"
)

const (
	GitFileSizeInBytesIndex = 3
	GitFileNameIndex        = 4
)

type GitFileLister struct {
	*logrus.Entry
	*vcs.GitRepo
}

func NewGitFileLister(repo *vcs.GitRepo) RepoFileLister {
	return &GitFileLister{
		logrus.WithFields(logrus.Fields{
			"pkg":  "action",
			"type": repo.Vcs(),
		}),
		repo,
	}
}

func (l *GitFileLister) GetFilesFromCurrent() (FileNames, error) {

	ver, err := l.Version()

	if err != nil {
		return nil, err
	}

	l.WithField("version", ver).
		Debug("Detected version.")

	b, err := l.RunFromDir("git", "ls-tree", "-r", "-l", ver)

	if err != nil {
		return nil, err
	}

	gitLogOutput := string(b)

	gitLogLines := strings.Split(gitLogOutput, "\n")

	if gitLogLines[0] == gitLogOutput {
		l.WithField("output", gitLogOutput).Debug("Unexpected log output")
		return nil, errors.New("expected git ls-tree format")
	}

	fileNames := make(FileNames, len(gitLogLines))

	for _, line := range gitLogLines[:len(gitLogLines)-1] {

		gitCommitFileHistory := strings.Fields(line)

		if len(gitCommitFileHistory) == 0 {
			return nil, errors.New("shouldn't be empty")
		}

		name := gitCommitFileHistory[GitFileNameIndex]
		sizeInBytes, err := strconv.Atoi(gitCommitFileHistory[GitFileSizeInBytesIndex])

		if err != nil {
			return nil, err
		}

		fileNames = append(fileNames, RepoFile{
			Name: name,
			Size: sizeInBytes,
		})
	}

	return fileNames, nil

}
