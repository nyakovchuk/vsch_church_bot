package config

import (
	"errors"
	"fmt"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

var (
	ErrUnknownError = errors.New("unknown error type")
)

type Config struct {
	TelegramBotToken string `env:"TELEGRAM_BOT_TOKEN"`
	LogType          string `env:"LOG_TYPE" envDefault="Screen"`
	LogFilePath      string `env:"LOG_FILE_PATH"`
	DbDriver         string `env:"DB_DRIVER"`
	Dsn              string `env:"DSN"`
	Platform         string `env:"PLATFORM_BOT"`
}

func LoadConfig(configFile string) (*Config, error) {
	err := godotenv.Load(configFile)
	if err != nil {
		return nil, err
	}

	var config Config

	if err := env.Parse(&config); err != nil {
		if errors.Is(err, env.EmptyVarError{}) {
			return nil, err
		}

		aggErr := env.AggregateError{}
		if ok := errors.As(err, &aggErr); ok {
			for _, er := range aggErr.Errors {
				switch v := er.(type) {
				// Handle the error types you need:
				// ParseError
				// NotStructPtrError
				// NoParserError
				// NoSupportedTagOptionError
				// EnvVarIsNotSetError
				// EmptyEnvVarError
				// LoadFileContentError
				case env.ParseValueError:
					return nil, err
				case env.EmptyVarError:
					return nil, err
				default:
					return nil, fmt.Errorf("%v %w %v", ErrUnknownError, err, v)
				}
			}
		}
	}

	return &config, nil
}
