package cmd

import (
	"context"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/zhtheyun/fibonacci/lib/config"
	"github.com/zhtheyun/fibonacci/web"
)

const (
	// RestCmdName represents the cobra rest sub-command.
	RestCmdName = "rest"

	// RestCmdDesp represents the cobra rest sub-command short description.
	RestCmdDesp = "Start to serve a RESTful API service"
)

func commandRest() *cobra.Command {
	return &cobra.Command{
		Use:   RestCmdName,
		Short: RestCmdDesp,
		RunE: func(*cobra.Command, []string) error {
			return executeRest()
		},
	}
}

func executeRest() error {
	cfg, err := config.Build()
	if err != nil {
		return errors.Wrap(err, "failed to build config")
	}
	setupLogLevel(cfg.LogLevel)

	server, err := web.NewServer(*cfg)
	if err != nil {
		return errors.Wrap(err, "failed to init the server")
	}

	if err := server.Run(context.Background()); err != nil {
		return errors.Wrap(err, "failed to run the server")
	}

	return nil
}

func setupLogLevel(level string) error {
	switch level {
	case "debug":
		logrus.SetLevel(logrus.DebugLevel)
	case "info":
		logrus.SetLevel(logrus.InfoLevel)
	case "warning":
		logrus.SetLevel(logrus.WarnLevel)
	case "error":
		logrus.SetLevel(logrus.ErrorLevel)
	case "fatal":
		logrus.SetLevel(logrus.FatalLevel)
	case "panic":
		logrus.SetLevel(logrus.PanicLevel)
	default:
		return errors.New("invalid log level")
	}
	return nil
}
