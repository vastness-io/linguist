package route

import (
	"net/http"

	"encoding/json"
	"github.com/vastness-io/linguist/pkg/action"
	"github.com/vastness-io/linguist/pkg/detect"
	"github.com/vastness-io/linguist/pkg/model"
	"io/ioutil"
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
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	repo, err := action.CloneRepository(repositoryInfo.String())

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	langs, err := detect.DetermineLanguages(repo.LocalPath())

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(langs); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
