package detect

import (
	"github.com/vastness-io/linguist-svc"
	"github.com/vastness-io/linguist/pkg/detect"
	"testing"
)

var (
	repoFiles = detect.FileNames{

		"main.go",
		"pom.xml",
		"build.xml",
		"build.gradle",
	}
)

func TestDetermineLanguages(t *testing.T) {
	determineLanguagesTests := []struct {
		fileNames         detect.FileNames
		expectedLanguages []*linguist.Language
	}{
		{
			fileNames: repoFiles,
			expectedLanguages: []*linguist.Language{
				{
					Name:       "Go",
					Percentage: float64(25),
				},
				{
					Name:       "Maven POM",
					Percentage: float64(25),
				},
				{
					Name:       "Gradle",
					Percentage: float64(25),
				},
				{
					Name:       "Ant Build System",
					Percentage: float64(25),
				},
			},
		},
	}

	for _, test := range determineLanguagesTests {

		langs := detect.DetermineLanguages(test.fileNames)

		for _, iLang := range test.expectedLanguages {

			for _, jLang := range langs {
				if jLang.Name != iLang.Name {
					continue
				}
				if jLang.Percentage != iLang.Percentage {
					t.Errorf("expected: %f, got: %f", iLang.Percentage, jLang.Percentage)
				}
			}
		}
	}
}
