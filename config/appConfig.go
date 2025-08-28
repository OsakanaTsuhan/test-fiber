package config

import (
	"github.com/spf13/viper"
)

type AppConfig struct {
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
}

func SetupEnv(path string) (cfg AppConfig, err error) {

	// // APP_ENV=devならpathをに
	// if os.Getenv("APP_ENV") == "dev" {
	// 	path = "."
	// }

	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env") // json, xml

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&cfg)
	return
}
