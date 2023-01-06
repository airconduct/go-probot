package probot

import (
	"github.com/go-logr/logr"
	"github.com/go-logr/zapr"
	"github.com/spf13/pflag"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ServerOptions struct {
	Address string
	Port    int
	Path    string
}

func (opts *ServerOptions) AddFlags(flags *pflag.FlagSet) {
	flags.StringVar(&opts.Address, "address", "127.0.0.1", "Address listen on")
	flags.IntVar(&opts.Port, "port", 7771, "Port listen on")
	flags.StringVar(&opts.Path, "path", "/hook", "Path listen on")
}

type loggerOptions struct {
	logger logr.Logger
}

func (opts *loggerOptions) AddFlags(flags *pflag.FlagSet) {
	logconfig := zap.NewDevelopmentConfig()
	logconfig.Encoding = "console"
	logconfig.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	zlogger, err := logconfig.Build()
	if err != nil {
		panic(err)
	}
	opts.logger = zapr.NewLogger(zlogger)
}

func (opts *loggerOptions) GetLogger() logr.Logger {
	return opts.logger
}
