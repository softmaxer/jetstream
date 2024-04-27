package main

import (
	"context"
	"errors"
	"flag"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	// tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
	"github.com/charmbracelet/ssh"
	"github.com/charmbracelet/wish"
	//"github.com/charmbracelet/wish/bubbletea"
	"github.com/charmbracelet/wish/logging"
	//"github.com/muesli/termenv"
)

const (
	host = "localhost"
	port = "5531"
)

func parseCommandLineArgs() map[string]*string {
	host := flag.String("host", "localhost", "The host address to run on. Defaults to localhost")
	port := flag.String("port", "5531", "The port number to run the application on. Default: 5531")
	flag.Parse()
	args := make(map[string]*string)
	args["host"] = host
	args["port"] = port
	return args
}

func main() {
	args := parseCommandLineArgs()

	s, err := wish.NewServer(
		wish.WithAddress(net.JoinHostPort(*args["host"], *args["port"])),
		wish.WithHostKeyPath(".ssh/id_ed25519"),
		wish.WithMiddleware(
			logging.Middleware(),
		),
	)
	if err != nil {
		log.Error("Could not start server", "error", err)
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	log.Info("Starting SSH server", "host", host, "port", port)
	go func() {
		if err = s.ListenAndServe(); err != nil && !errors.Is(err, ssh.ErrServerClosed) {
			log.Error("Could not start server", "error", err)
			done <- nil
		}
	}()

	<-done
	log.Info("Stopping SSH server")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer func() { cancel() }()
	if err := s.Shutdown(ctx); err != nil && !errors.Is(err, ssh.ErrServerClosed) {
		log.Error("Could not stop server", "error", err)
	}
}
