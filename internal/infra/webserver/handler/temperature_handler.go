package handler

import (
	"cloudrun/internal/usecase"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"regexp"

	"github.com/go-chi/chi/v5"
)

type TemperatureHandler struct {
	getTemperatureFromCepUseCase usecase.GetTemperatureFromCepUseCase
}

func NewTemperatureHandler(getTemperatureFromCepUseCase usecase.GetTemperatureFromCepUseCase) *TemperatureHandler {
	return &TemperatureHandler{
		getTemperatureFromCepUseCase: getTemperatureFromCepUseCase,
	}
}

func (t *TemperatureHandler) GetTemperatureFromCep(w http.ResponseWriter, r *http.Request) {
	cep := chi.URLParam(r, "cep")
	if cep == "" || !isValidCep(cep) {
		err := errors.New("invalid zipcode")
		log.Printf("error: %s", err.Error())
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	response, err := t.getTemperatureFromCepUseCase.Execute(r.Context(), cep)
	if errors.Is(err, usecase.ErrNotFound) {
		log.Printf("error: %s", err.Error())
		http.Error(w, "can not find zipcode", http.StatusNotFound)
		return
	}
	if err != nil {
		log.Printf("error: %s", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("error: %s", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

const validCepRegex string = `^\d{5}-?\d{3}$`

var (
	cepRegex = regexp.MustCompile(validCepRegex)
)

func isValidCep(cep string) bool {
	return cepRegex.MatchString(cep)
}
