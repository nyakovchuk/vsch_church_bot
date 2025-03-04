package app

import (
	"fmt"
	"log"
	"log/slog"
	"sync"

	"github.com/nyakovchuk/vsch_church_bot/config"
	"github.com/nyakovchuk/vsch_church_bot/internal/logger"
)

type App struct {
	config *config.Config
	logger *slog.Logger
}

// Singleton
var (
	instance *App
	once     sync.Once
)

func (a *App) Config() *config.Config {
	return a.config
}

func (a *App) Logger() *slog.Logger {
	return a.logger
}

func NewApp(config *config.Config, logger *slog.Logger) *App {
	return &App{
		config: config,
		logger: logger,
	}
}

// GetApp returns singleton
func GetApp() *App {
	once.Do(func() {
		fmt.Print("Loading configuration...")
		config, err := config.LoadConfig()
		if err != nil {
			fmt.Println("Error loading configuration:", err)
			return
		}
		fmt.Println("DONE")

		fmt.Print("Setting up logging...")
		logger, err := logger.SetupLogger(config, nil)
		if err != nil {
			fmt.Println("Error setting up logging:", err)
			return
		}
		fmt.Println("DONE")

		instance = NewApp(config, logger)
	})

	if instance == nil {
		log.Fatal("Failed to set up the configuration")
	}

	return instance
}
