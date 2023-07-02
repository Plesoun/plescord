package config

import "github.com/spf13/viper"

type Config struct {
	WebsocketBindAddress string `mapstructure:"websocket.bind_address"`
	WebsocketPort        string `mapstructure:"websocket.port"`
	LoggingType          string `mapstructure:"logging.type"`
	LoggingPath          string `mapstructure:"logging.path"`
}

func LoadConfig(configPath *string) (Config, error) {
	var config Config

	viper.SetConfigFile(*configPath)

	if err := viper.ReadInConfig(); err != nil {
		return config, err
	}

	config.WebsocketBindAddress = viper.GetString("websocket.bind_address")
	config.WebsocketPort = viper.GetString("websocket.port")
	config.LoggingType = viper.GetString("logging.type")
	config.LoggingPath = viper.GetString("logging.path")

	return config, nil
}
