package language

import (
	"github.com/vastness-io/linguist-svc"
	"testing"
)

var (
	repoFiles = FileNames{

		"main.go",
		"pom.xml",
		"build.xml",
		"README.md",
		"invalid_file_format",
	}
)

func TestDetermineLanguages(t *testing.T) {
	calculateLanguagesTests := []struct {
		fileNames         FileNames
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
					Name:       "Ant Build System",
					Percentage: float64(25),
				},
				{
					Name:       "Other",
					Percentage: float64(25),
				},
			},
		},
	}

	for _, test := range calculateLanguagesTests {

		langs := CalculateLanguages(test.fileNames)

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
