package handlers

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"question-answer/models"
	"question-answer/repositories"
	"strings"
)

type UserHandler struct {
	repo *repositories.UserRepository
}

func NewUserHandler() *UserHandler {
	return &UserHandler{
		repo: repositories.NewUserRepository(),
	}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	token := GenerateToken()

	user := &models.User{
		Token: token,
	}

	if err := h.repo.Create(user); err != nil {
		errorResponse(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	response := models.TokenResponse{
		Token: token,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func (h *UserHandler) GetUserID(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		errorResponse(w, "Authorization header required", http.StatusUnauthorized)
		return
	}

	var token string
	if strings.HasPrefix(authHeader, "Bearer ") {
		token = strings.TrimPrefix(authHeader, "Bearer ")
	} else {
		token = authHeader
	}

	user, err := h.repo.GetByToken(token)
	if err != nil {
		errorResponse(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	response := map[string]interface{}{
		"user_id": user.ID,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func GenerateToken() string {
	bytes := make([]byte, 16)
	rand.Read(bytes)
	return hex.EncodeToString(bytes)
}
