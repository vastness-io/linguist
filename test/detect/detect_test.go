package detect

import (
	"github.com/vastness-io/linguist/pkg/action"
	"github.com/vastness-io/linguist/pkg/detect"
	"testing"
)

var (
	repoFiles = action.RepoFiles{
		{
			Name: "main.go",
			Size: 1,
		},
		{
			Name: "pom.xml",
			Size: 1,
		},
		{
			Name: "build.xml",
			Size: 1,
		},

		{
			Name: "build.gradle",
			Size: 1,
		},
	}
)

func TestDetermineLanguages(t *testing.T) {
	determineLanguagesTests := []struct {
		repoFiles         action.RepoFiles
		expectedLanguages []detect.Language
	}{
		{
			repoFiles: repoFiles,
			expectedLanguages: []detect.Language{
				{
					Language:   "Go",
					Percentage: float64(25),
				},
				{
					Language:   "Maven POM",
					Percentage: float64(25),
				},
				{
					Language:   "Gradle",
					Percentage: float64(25),
				},
				{
					Language:   "Ant Build System",
					Percentage: float64(25),
				},
			},
		},
	}

	for _, test := range determineLanguagesTests {

		langs := detect.DetermineLanguages(test.repoFiles)

		for _, iLang := range test.expectedLanguages {

			for _, jLang := range langs {
				if jLang.Language != iLang.Language {
					continue
				}
				if jLang.Percentage != iLang.Percentage {
					t.Errorf("expected: %f, got: %f", iLang.Percentage, jLang.Percentage)
				}
			}
		}
	}
}
