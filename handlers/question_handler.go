package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"question-answer/models"
	"question-answer/repositories"

	"gorm.io/gorm"
)

type QuestionRepository interface {
	Create(question *models.Question) error
	GetByID(id uint) (*models.Question, error)
	GetWithAnswers(id uint) (*models.Question, error)
	GetAll() ([]models.Question, error)
	Delete(id uint) error
}

type QuestionHandler struct {
	Repo QuestionRepository
}

func NewQuestionHandler(repo QuestionRepository) *QuestionHandler {
	if repo != nil {
		return &QuestionHandler{
			Repo: repo,
		}
	} else {
		return &QuestionHandler{
			Repo: repositories.NewQuestionRepository(),
		}
	}
}

func (h *QuestionHandler) GetQuestions(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		errorResponse(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	questions, err := h.Repo.GetAll()
	if err != nil {
		errorResponse(w, "Failed to get questions", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(questions)
}

func (h *QuestionHandler) CreateQuestion(w http.ResponseWriter, r *http.Request) {
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

	question := &models.Question{
		Text: strings.TrimSpace(req.Text),
	}

	if err := h.Repo.Create(question); err != nil {
		errorResponse(w, "Failed to create question", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(question)
}

func (h *QuestionHandler) GetQuestion(w http.ResponseWriter, r *http.Request) {
	idStr := ExtractID(r.URL.Path, "/questions/")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		errorResponse(w, "Invalid question ID", http.StatusBadRequest)
		return
	}

	question, err := h.Repo.GetWithAnswers(uint(id))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			errorResponse(w, "Question not found", http.StatusNotFound)
			return
		}
		errorResponse(w, "Failed to get question", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(question)
}

func (h *QuestionHandler) DeleteQuestion(w http.ResponseWriter, r *http.Request) {
	idStr := ExtractID(r.URL.Path, "/questions/")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		errorResponse(w, "Invalid question ID", http.StatusBadRequest)
		return
	}

	if err := h.Repo.Delete(uint(id)); err != nil {
		errorResponse(w, "Failed to delete question", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
