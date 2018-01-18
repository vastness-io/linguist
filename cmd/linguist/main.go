package main

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/vastness-io/linguist/pkg/route"
	"net/http"
	"os"
)

var (
	router = mux.NewRouter()
	log    = logrus.WithField("pkg", "main")
)

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)
	router.Methods(http.MethodPost).Path(route.FromVCSDetectPath).HandlerFunc(route.RepositoryHandler)
	http.Handle("/", router)
}

func main() {
	log.Info("Starting Linguist")
	defer log.Info("Exiting Linguist")
	panic(http.ListenAndServe(":8080", router))
}
