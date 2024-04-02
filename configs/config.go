package configs

import (
	"log"

	"github.com/spf13/viper"
)

type cfg struct {
	WeatherAPIKey string
	WeatherAPIUrl string
	ViaCepAPIUrl  string
	ServerPort    string
}

func GetConfig() *cfg {
	viper.SetDefault("WEATHER_API_URL", "https://api.weatherapi.com/v1/current.json")
	viper.SetDefault("VIA_CEP_API_URL", "https://viacep.com.br/ws/")
	viper.SetDefault("SERVER_PORT", "8080")

	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		log.Println("Error reading config file", err)
	}

	return &cfg{
		WeatherAPIKey: viper.GetString("WEATHER_API_KEY"),
		WeatherAPIUrl: viper.GetString("WEATHER_API_URL"),
		ViaCepAPIUrl:  viper.GetString("VIA_CEP_API_URL"),
		ServerPort:    viper.GetString("SERVER_PORT"),
	}
}
