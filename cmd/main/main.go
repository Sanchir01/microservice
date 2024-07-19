package main

import (
	"fmt"
	"github.com/Sanchir01/microservice/internal/app"
	"github.com/Sanchir01/microservice/internal/config"
	"github.com/Sanchir01/microservice/pkg/lib/logger/handlers/slogpretty"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

var (
	development = "development"
	production  = "production"
)

func main() {
	cfg := config.MustLoad()
	fmt.Println(cfg)
	lg := setupLogger(cfg.Env)

	application := app.NewAppSrv(lg, cfg.GRPC.Port, cfg.StoragePath)

	go application.GrpcSrv.MustRun()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	sign := <-stop

	lg.Info("stoppping application", slog.String("signal", sign.String()))

	application.GrpcSrv.Stop()
}
func setupLogger(env string) *slog.Logger {
	var lg *slog.Logger
	switch env {
	case development:
		lg = setupPrettySlog()
	case production:
		lg = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}
	return lg
}

func setupPrettySlog() *slog.Logger {
	opts := slogpretty.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}

	handler := opts.NewPrettyHandler(os.Stdout)

	return slog.New(handler)
}
