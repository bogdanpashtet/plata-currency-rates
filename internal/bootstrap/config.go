package bootstrap

import (
	"github.com/BurntSushi/toml"
	"github.com/bogdanpashtet/plata-currency-rates/internal/models/config"
	"github.com/rs/zerolog"
)

const configPath = "./config/config.toml"

func InitConfig(logger zerolog.Logger) config.Config {
	var conf config.Config
	if _, err := toml.DecodeFile(configPath, &conf); err != nil {
		logger.Fatal().Msg("can`t unmarshall toml config")
	}

	return conf
}
