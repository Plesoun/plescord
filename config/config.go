package config

import "github.com/spf13/viper"

type Config struct {
	WebsocketBindAddress string `mapstructure:"websocket.bind_address"`
	WebsocketPort        string `mapstructure:"websocket.port"`
}

func LoadConfig(configPath *string) (Config, error) {
	var config Config

	viper.SetConfigFile(*configPath)

	if err := viper.ReadInConfig(); err != nil {
		return config, err
	}

	config.WebsocketBindAddress = viper.GetString("websocket.bind_address")
	config.WebsocketPort = viper.GetString("websocket.port")

	return config, nil
}
