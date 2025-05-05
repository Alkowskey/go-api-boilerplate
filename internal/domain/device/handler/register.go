package handler

import (
	"encoding/json"
	"net/http"

	"github.com/aleksander/Go_API/internal/domain/device/usecase"
)

type DeviceHandler struct {
	registerUseCase *usecase.RegisterDeviceUseCase
}

func NewDeviceHandler(registerUseCase *usecase.RegisterDeviceUseCase) *DeviceHandler {
	return &DeviceHandler{
		registerUseCase: registerUseCase,
	}
}

type registerDeviceRequest struct {
	Name string `json:"name"`
}

func (h *DeviceHandler) RegisterDevice(w http.ResponseWriter, r *http.Request) {
	var req registerDeviceRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error": "Invalid request body"}`, http.StatusBadRequest)
		return
	}

	output := h.registerUseCase.Execute(usecase.RegisterDeviceInput{
		Name: req.Name,
	})

	if output.Err != nil {
		http.Error(w, `{"error": "`+output.Err.Error()+`"}`, http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(output.Device)
}
