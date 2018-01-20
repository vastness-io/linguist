package route

import (
	"net/http"

	"encoding/json"
	"github.com/vastness-io/linguist/pkg/action"
	"github.com/vastness-io/linguist/pkg/detect"
	"github.com/vastness-io/linguist/pkg/model"
	"io/ioutil"
	"os"
)

const (
	FromVCSDetectPath = "/vcs/detect"
)

func RepositoryHandler(w http.ResponseWriter, r *http.Request) {

	b, err := ioutil.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var repositoryInfo model.RepositoryInfo

	if err := json.Unmarshal(b, &repositoryInfo); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	repo, err := action.CloneRepository(repositoryInfo.String())

	if err != nil {

		if repo != nil {
			os.RemoveAll(repo.LocalPath())
		}

		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer os.RemoveAll(repo.LocalPath())

	lister, err := action.NewRepoFileLister(repo)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	files, err := lister.GetFilesFromCurrent()

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	langs := detect.DetermineLanguages(files)

	w.Header().Add("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(langs); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
