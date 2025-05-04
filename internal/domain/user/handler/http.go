package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"Go_API/internal/auth"
	"Go_API/internal/domain/user/usecase"

	"github.com/gorilla/mux"
)

type UserHandler struct {
	registerUseCase     *usecase.RegisterUseCase
	authenticateUseCase *usecase.AuthenticateUseCase
	getUserUseCase      *usecase.GetUserUseCase
	listUsersUseCase    *usecase.ListUsersUseCase
}

func NewUserHandler(
	registerUseCase *usecase.RegisterUseCase,
	authenticateUseCase *usecase.AuthenticateUseCase,
	getUserUseCase *usecase.GetUserUseCase,
	listUsersUseCase *usecase.ListUsersUseCase,
) *UserHandler {
	return &UserHandler{
		registerUseCase:     registerUseCase,
		authenticateUseCase: authenticateUseCase,
		getUserUseCase:      getUserUseCase,
		listUsersUseCase:    listUsersUseCase,
	}
}

type registerRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req registerRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error": "Invalid request body"}`, http.StatusBadRequest)
		return
	}

	output := h.registerUseCase.Execute(usecase.RegisterInput{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	})

	if output.Err != nil {
		http.Error(w, `{"error": "`+output.Err.Error()+`"}`, http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(output.User)
}

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type loginResponse struct {
	Token string `json:"token"`
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req loginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error": "Invalid request body"}`, http.StatusBadRequest)
		return
	}

	output := h.authenticateUseCase.Execute(usecase.AuthenticateInput{
		Email:    req.Email,
		Password: req.Password,
	})

	if output.Err != nil {
		http.Error(w, `{"error": "`+output.Err.Error()+`"}`, http.StatusUnauthorized)
		return
	}

	// Generate JWT token
	token, err := auth.GenerateToken(output.User.ID)
	if err != nil {
		http.Error(w, `{"error": "Failed to generate token"}`, http.StatusInternalServerError)
		return
	}

	response := loginResponse{
		Token: token,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		http.Error(w, `{"error": "Invalid user ID"}`, http.StatusBadRequest)
		return
	}

	output := h.getUserUseCase.Execute(usecase.GetUserInput{
		ID: uint(id),
	})

	if output.Err != nil {
		http.Error(w, `{"error": "`+output.Err.Error()+`"}`, http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output.User)
}

func (h *UserHandler) ListUsers(w http.ResponseWriter, r *http.Request) {
	output := h.listUsersUseCase.Execute()
	if output.Err != nil {
		http.Error(w, `{"error": "Failed to fetch users"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output.Users)
}
