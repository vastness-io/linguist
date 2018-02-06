package main

import (
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"github.com/vastness-io/linguist-svc"
	toolkit "github.com/vastness-io/toolkit/pkg/grpc"
	"net"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"github.com/vastness-io/linguist/pkg/service"
	"github.com/vastness-io/linguist/pkg/server"
)

const (
	name        = "linguist"
	description = "Detects languages and reports assumptions."
)

var (
	log       = logrus.WithField("component", "linguist")
	commit    string
	version   string
	addr      string
	port      int
	debugMode bool
)

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
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

	var (
		address  = net.JoinHostPort(addr, strconv.Itoa(port))
		tracer   = opentracing.GlobalTracer()
		lis, err = net.Listen("tcp", address)
		srv      = toolkit.NewGRPCServer(tracer, log)
		svc = service.NewLinguistService()

	)

	if err != nil {
		log.Fatal(err)
	}

	go func() {
		log.Infof("Listening on %s", address)
		if err := srv.Serve(lis); err != nil {
			log.Fatal(err)
		}
	}()

	linguist.RegisterLinguistServer(srv, server.NewLinguistServer(svc))

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	for {
		select {
		case <-signalChan:
			log.Info("Exiting linguist")
			srv.GracefulStop()
			os.Exit(0)
		}
	}
}
