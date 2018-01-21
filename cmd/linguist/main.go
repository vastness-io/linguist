package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"github.com/vastness-io/linguist-svc"
	"github.com/vastness-io/linguist/pkg/svc"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

const (
	name        = "linguist"
	description = "Detects languages and reports assumptions."
)

var (
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
			Value:       8082,
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

	address := fmt.Sprintf("%s:%s", addr, strconv.Itoa(port))

	lis, err := net.Listen("tcp", address)

	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()

	go func() {
		log.Infof("Listening on %s", address)
		if err := s.Serve(lis); err != nil {
			log.Fatal(err)
		}
	}()

	linguist.RegisterLinguistServer(s, &svc.LinguistService{})

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	for {
		select {
		case <-signalChan:
			log.Info("Exiting linguist")
			s.GracefulStop()
			os.Exit(0)
		}
	}
}
