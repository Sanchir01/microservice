package main

import (
	"fmt"
	"github.com/Sanchir01/microservice/internal/config"
	"github.com/Sanchir01/microservice/pkg/lib/logger/handlers/slogpretty"
	"log/slog"
	"os"
)

var (
	development = "development"
	production  = "production"
)

func main() {
	cfg := config.MustLoad()
	fmt.Println(cfg)
	lg := setupLogger(cfg.Env)
	lg.Warn("test logger", lg)

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
