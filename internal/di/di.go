package di

import (
	"cloudrun/internal/infra/webserver/handler"
	"cloudrun/internal/provider"
	"cloudrun/internal/usecase"
)

func NewTemperatureHandler(viacepUrl, weatherApiUrl, weatherApiKey string) *handler.TemperatureHandler {
	locationGateway := provider.NewViaCepLocationProvider(viacepUrl)
	temperatureGateway := provider.NewWeatherAPIProvider(weatherApiKey, weatherApiUrl)
	usecase := usecase.NewGetTemperatureFromCepUseCase(locationGateway, temperatureGateway)
	return handler.NewTemperatureHandler(usecase)
}
