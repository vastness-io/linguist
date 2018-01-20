package action

import (
	"errors"
	"github.com/Masterminds/vcs"
	"strconv"
	"strings"
)

const (
	GitFileSizeInBytesIndex = 3
	GitFileNameIndex        = 4
)

type GitFileLister struct {
	*vcs.GitRepo
}

func NewGitFileLister(repo *vcs.GitRepo) RepoFileLister {
	return &GitFileLister{
		repo,
	}
}

func (l *GitFileLister) VCS() vcs.Type {
	return l.GitRepo.Vcs()
}

func (l *GitFileLister) GetFilesFromCurrent() (RepoFiles, error) {

	ver, err := l.Version()

	if err != nil {
		return nil, err
	}

	log.WithField("version", ver).
		Debug("Detected version")

	b, err := l.RunFromDir("git", "ls-tree", "-r", "-l", ver)

	if err != nil {
		return nil, err
	}

	gitLogOutput := string(b)

	gitLogLines := strings.Split(gitLogOutput, "\n")

	if gitLogLines[0] == gitLogOutput {
		log.WithField("output", gitLogOutput).Debug("Unexpected log output")
		return nil, errors.New("expected git ls-tree format")
	}

	fileNames := make(RepoFiles, 0)

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
			Size: sanitizeFileSize(sizeInBytes),
		})
	}

	return fileNames, nil

}
