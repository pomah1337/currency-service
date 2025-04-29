package handler

import (
	"currencyService/gateway/internal/dto"
	"currencyService/gateway/internal/service"
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"net/http"
)

var validate = validator.New()

type AuthorizationHandler struct {
	svc *service.AuthService
}

func NewAuthorizationHandler(svc *service.AuthService) *AuthorizationHandler {
	return &AuthorizationHandler{svc: svc}
}

func (h *AuthorizationHandler) RegisterRoutes() {
	r := mux.NewRouter()
	api := r.PathPrefix("/api/v1").Subrouter()

	api.HandleFunc("/login", h.login).Methods("POST")
	api.HandleFunc("/register", h.register).Methods("POST")

}

func (h *AuthorizationHandler) login(w http.ResponseWriter, r *http.Request) {
	var creds dto.Credentials
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "Invalid payload", http.StatusBadRequest)
		return
	}

	if err := validate.Struct(creds); err != nil {
		http.Error(w, "Validation failed", http.StatusBadRequest)
		return
	}

	token, err := h.svc.Login(r.Context(), creds.Login, creds.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(token))
}

func (h *AuthorizationHandler) register(w http.ResponseWriter, r *http.Request) {
	var creds dto.Credentials
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "Invalid payload", http.StatusBadRequest)
		return
	}

	if err := validate.Struct(creds); err != nil {
		http.Error(w, "Validation failed", http.StatusBadRequest)
		return
	}

	err := h.svc.Registration(r.Context(), creds.Login, creds.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User registered successfully"))
}
