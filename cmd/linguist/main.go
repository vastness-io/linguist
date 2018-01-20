package main

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"github.com/vastness-io/linguist/pkg/route"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

const (
	name        = "linguist"
	description = "Detects languages and reports assumptions."
)

var (
	router    = mux.NewRouter()
	log       = logrus.WithField("pkg", "main")
	commit    string
	version   string
	addr      string
	port      int
	debugMode bool
)

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)
	router.Methods(http.MethodPost).Path(route.FromVCSDetectPath).HandlerFunc(route.RepositoryHandler)
}

func main() {
	app := cli.NewApp()
	app.Name = name
	app.Usage = description
	app.Version = fmt.Sprintf("%s (%s)", version, commit)
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "addr,a",
			Usage:       "TCP address to listen on",
			Value:       "127.0.0.1",
			Destination: &addr,
		},
		cli.IntFlag{
			Name:        "port,p",
			Usage:       "Port to listen on",
			Value:       8080,
			Destination: &port,
		},
		cli.BoolFlag{
			Name:        "debug,d",
			Usage:       "Debug mode",
			Destination: &debugMode,
		},
	}
	app.Action = func(_ *cli.Context) { run() }
	app.Run(os.Args)
}

func run() {

	if debugMode {
		logrus.SetLevel(logrus.DebugLevel)
	}

	log.Info("Starting linguist")

	srv := http.Server{
		Addr:    fmt.Sprintf("%s:%s", addr, strconv.Itoa(port)),
		Handler: router,
	}

	go func() {
		log.Infof("Listening on %s", srv.Addr)
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	for {
		select {
		case <-signalChan:
			log.Info("Exiting linguist")
			ctx, _ := context.WithTimeout(context.Background(), time.Minute)
			if err := srv.Shutdown(ctx); err != nil {
				os.Exit(1)
			}
			os.Exit(0)
		}
	}
}
