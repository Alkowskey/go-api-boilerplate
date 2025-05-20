package handler

import (
	"encoding/json"
	"net/http"

	"github.com/aleksander/Go_API/internal/domain/acceleration/usecase"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type AccelerationHandler struct {
	importAccelerationDataUseCase *usecase.ImportAccelerationDataUseCase
	getAccelerationDataUsecase    *usecase.GetAccelerationDataUseCase
}

func NewAccelerationHandler(
	importAccelerationDataUseCase *usecase.ImportAccelerationDataUseCase,
	getAccelerationDataUsecase *usecase.GetAccelerationDataUseCase,
) *AccelerationHandler {
	return &AccelerationHandler{
		importAccelerationDataUseCase: importAccelerationDataUseCase,
		getAccelerationDataUsecase:    getAccelerationDataUsecase,
	}
}

type importAccelerationDataRequest struct {
	Data []struct {
		X float64 `json:"x"`
		Y float64 `json:"y"`
		Z float64 `json:"z"`
	} `json:"data"`
	DeviceID uuid.UUID `json:"device_id"`
}

func (h *AccelerationHandler) ImportAccelerationData(w http.ResponseWriter, r *http.Request) {
	var req importAccelerationDataRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error": "Invalid request body"}`, http.StatusBadRequest)
		return
	}

	output := h.importAccelerationDataUseCase.Execute(usecase.ImportAccelerationDataInput{
		X: req.Data[0].X,
		Y: req.Data[0].Y,
		Z: req.Data[0].Z,
	}, req.DeviceID)

	if output.Err != nil {
		http.Error(w, `{"error": "`+output.Err.Error()+`"}`, http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(output.Acceleration)
}

func (h *AccelerationHandler) GetAccelerationData(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := uuid.Parse(vars["id"])
	if err != nil {
		http.Error(w, `{"error": "Invalid ID format"}`, http.StatusBadRequest)
		return
	}

	output := h.getAccelerationDataUsecase.Execute(usecase.GetAccelerationDataInput{
		ID: id,
	})

	if output.Err != nil {
		http.Error(w, `{"error": "`+output.Err.Error()+`"}`, http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output.Acceleration)
}
