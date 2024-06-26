package usecase

import (
	"cloudrun/internal/gateway"
	"context"
	"errors"
	"fmt"
)

var (
	ErrNotFound = fmt.Errorf("not found")
)

type GetTemperatureFromCepResponse struct {
	Celsius    string `json:"temp_C"`
	Fahrenheit string `json:"temp_F"`
	Kelvin     string `json:"temp_K"`
}

type GetTemperatureFromCepUseCase interface {
	Execute(ctx context.Context, cep string) (*GetTemperatureFromCepResponse, error)
}

type getTemperatureFromCepUseCase struct {
	locationProvider    gateway.LocationGateway
	temperatureProvider gateway.TemperatureGateway
}

func NewGetTemperatureFromCepUseCase(locationProvider gateway.LocationGateway, temperatureGateway gateway.TemperatureGateway) *getTemperatureFromCepUseCase {
	return &getTemperatureFromCepUseCase{
		locationProvider:    locationProvider,
		temperatureProvider: temperatureGateway,
	}
}

func (g *getTemperatureFromCepUseCase) Execute(ctx context.Context, cep string) (*GetTemperatureFromCepResponse, error) {
	location, err := g.locationProvider.FetchLocationByCep(ctx, cep)
	if errors.Is(err, gateway.ErrCepNotFound) {
		err = fmt.Errorf("%w; original: %w", ErrNotFound, err)
	}
	if err != nil {
		return nil, fmt.Errorf("get_temperature_from_cep_use_case.Execute: %w", err)
	}
	temperature, err := g.temperatureProvider.FetchTemperatureByCity(ctx, location.City)
	if err != nil {
		return nil, fmt.Errorf("get_temperature_from_cep_use_case.Execute: %w", err)
	}
	return &GetTemperatureFromCepResponse{
		Celsius:    fmt.Sprintf("%.2f", temperature.Celsius),
		Fahrenheit: fmt.Sprintf("%.2f", celsiusToFahrenheit(temperature.Celsius)),
		Kelvin:     fmt.Sprintf("%.2f", celsiusToKelvin(temperature.Celsius)),
	}, nil
}

func celsiusToFahrenheit(celsius float64) float64 {
	return celsius*1.8 + 32
}

func celsiusToKelvin(celsius float64) float64 {
	return celsius + 273
}
