package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/gerladeno/otus_homeworks/hw12_13_14_15_calendar/internal/app"
	"github.com/gerladeno/otus_homeworks/hw12_13_14_15_calendar/internal/logger"
	internalhttp "github.com/gerladeno/otus_homeworks/hw12_13_14_15_calendar/internal/server/http"
	memorystorage "github.com/gerladeno/otus_homeworks/hw12_13_14_15_calendar/internal/storage/memory"
	sqlstorage "github.com/gerladeno/otus_homeworks/hw12_13_14_15_calendar/internal/storage/sql"
	_ "github.com/jackc/pgx/v4/stdlib"
	"os"
	"os/signal"
	"syscall"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "config", "/etc/calendar/config.toml", "Path to configuration file")
}

func main() {
	flag.Parse()

	if flag.Arg(0) == "version" {
		printVersion()
		return
	}

	config := NewConfig(configFile)
	log := logger.New(config.Logger.Level, config.Logger.Path)

	var (
		storage app.Storage
		err     error
	)
	if config.Storage.Remote {
		dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
			config.Storage.Host,
			config.Storage.Port,
			"calendar",
			"calendar",
			config.Storage.Database,
			config.Storage.Ssl)
		storage, err = sqlstorage.New(log, dsn)
		if err != nil {
			log.Fatalf("failed to connect to database: %s", err)
		}
	} else {
		storage = memorystorage.New(log)
	}

	calendar := app.New(log, storage)
	handler := internalhttp.NewEventHandler(calendar, log)
	router := internalhttp.NewRouter(handler, log, version)
	httpServer := internalhttp.NewServer(router, config.HTTP.Port)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		signals := make(chan os.Signal, 1)
		signal.Notify(signals, syscall.SIGINT, syscall.SIGHUP)

		select {
		case <-ctx.Done():
			return
		case <-signals:
		}

		signal.Stop(signals)
		cancel()

		if err := httpServer.Stop(); err != nil {
			log.Error("failed to stop http server: " + err.Error())
		}
	}()

	log.Infof("starting server on %d", config.HTTP.Port)
	if err := httpServer.Start(ctx); err != nil {
		log.Error("failed to start http server: " + err.Error())
		cancel()
		os.Exit(1) //nolint:gocritic
	}
}
