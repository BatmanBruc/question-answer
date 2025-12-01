package handlers

import (
	"encoding/json"
	"net/http"
	"question-answer/models"
	"question-answer/repositories"
	"strconv"
	"strings"

	"gorm.io/gorm"
)

type AnswerRepository interface {
	Create(answer *models.Answer) error
	GetByID(id uint) (*models.Answer, error)
	Delete(id uint) error
	GetByQuestionID(questionID uint) ([]models.Answer, error)
}

type UserRepository interface {
	GetByToken(token string) (*models.User, error)
}

type AnswerHandler struct {
	repo         AnswerRepository
	questionRepo QuestionRepository
	userRepo     UserRepository
}

func NewAnswerHandler(repo AnswerRepository, questionRepo QuestionRepository, userRepo UserRepository) *AnswerHandler {
	var r AnswerRepository
	if repo != nil {
		r = repo
	} else {
		r = repositories.NewAnswerRepository()
	}

	var qR QuestionRepository
	if repo != nil {
		qR = questionRepo
	} else {
		qR = repositories.NewQuestionRepository()
	}

	var uR UserRepository
	if repo != nil {
		uR = userRepo
	} else {
		uR = repositories.NewUserRepository()
	}

	return &AnswerHandler{
		repo:         r,
		questionRepo: qR,
		userRepo:     uR,
	}
}

func (h *AnswerHandler) CreateAnswer(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	path = strings.TrimSuffix(path, "/answers/")
	questionIDStr := ExtractID(path, "/questions/")
	questionID, err := strconv.ParseUint(questionIDStr, 10, 32)
	if err != nil {
		errorResponse(w, "Invalid question ID", http.StatusBadRequest)
		return
	}

	_, err = h.questionRepo.GetByID(uint(questionID))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			errorResponse(w, "Question not found", http.StatusNotFound)
			return
		}
		errorResponse(w, "Failed to get question", http.StatusInternalServerError)
		return
	}

	var req struct {
		Text string
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		errorResponse(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if strings.TrimSpace(req.Text) == "" {
		errorResponse(w, "Text is required", http.StatusBadRequest)
		return
	}

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

	user, err := h.userRepo.GetByToken(token)
	if err != nil {
		errorResponse(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	answer := &models.Answer{
		QuestionID: uint(questionID),
		UserID:     user.ID,
		Text:       strings.TrimSpace(req.Text),
	}

	if err := h.repo.Create(answer); err != nil {
		errorResponse(w, "Failed to create answer", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(answer)
}

func (h *AnswerHandler) GetAnswer(w http.ResponseWriter, r *http.Request) {
	answerIDStr := ExtractID(r.URL.Path, "/answers/")
	answerID, err := strconv.ParseUint(answerIDStr, 10, 32)
	if err != nil {
		errorResponse(w, "Invalid answer ID", http.StatusBadRequest)
		return
	}

	answer, err := h.repo.GetByID(uint(answerID))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			errorResponse(w, "Answer not found", http.StatusNotFound)
			return
		}
		errorResponse(w, "Failed to get answer", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(answer)
}

func (h *AnswerHandler) DeleteAnswer(w http.ResponseWriter, r *http.Request) {
	answerIDStr := ExtractID(r.URL.Path, "/answers/")
	answerID, err := strconv.ParseUint(answerIDStr, 10, 32)
	if err != nil {
		errorResponse(w, "Invalid answer ID", http.StatusBadRequest)
		return
	}

	if err := h.repo.Delete(uint(answerID)); err != nil {
		errorResponse(w, "Failed to delete answer", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
