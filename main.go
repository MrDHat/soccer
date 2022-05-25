package main

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"soccer-manager/config"
	"soccer-manager/instance"
	"soccer-manager/runner"

	"github.com/urfave/cli"
)

func shutDown(shutDownChannel chan *bool) {
	shutDown := true
	shutDownChannel <- &shutDown
	close(shutDownChannel)
}

func main() {
	config.Load()
	instance.Init()

	var shutDownChannel chan *bool

	defer instance.Destroy()

	clientApp := cli.NewApp()
	clientApp.Name = "soccer-manager-app"
	clientApp.Version = "0.0.1"
	clientApp.Commands = []cli.Command{
		{
			Name:  "start",
			Usage: "Start the service",
			Action: func(c *cli.Context) error {
				ctx := context.Background()

				var wg sync.WaitGroup

				wg.Add(1)
				go runner.NewAPI().Go(ctx, &wg)
				wg.Wait()
				return nil
			},
		},
	}
	if err := clientApp.Run(os.Args); err != nil {
		panic(err)
	}
	signalChannel := make(chan os.Signal, 2)
	signal.Notify(
		signalChannel,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)
	go func() {
		<-signalChannel
		shutDown(shutDownChannel)
	}()
}
