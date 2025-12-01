package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"question-answer/handlers"
	"question-answer/models"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func TestAnswerHandler_CreateAnswer_Success(t *testing.T) {
	answerRepo := new(MockAnswerRepository)
	questionRepo := new(MockQuestionRepository)
	userRepo := new(MockUserRepository)
	handler := handlers.NewAnswerHandler(answerRepo, questionRepo, userRepo)

	question := &models.Question{ID: 1, Text: "Вопрос?"}
	user := &models.User{ID: 123, Token: "valid_token"}

	questionRepo.On("GetByID", uint(1)).Return(question, nil)
	userRepo.On("GetByToken", "valid_token").Return(user, nil)
	answerRepo.On("Create", mock.AnythingOfType("*models.Answer")).
		Run(func(args mock.Arguments) {
			answer := args.Get(0).(*models.Answer)
			answer.ID = 1
			answer.QuestionID = 1
			answer.UserID = 123
		}).
		Return(nil)

	answerData := models.CreateAnswerRequest{
		Text: "Это ответ на вопрос",
	}
	jsonData, _ := json.Marshal(answerData)

	req := httptest.NewRequest("POST", "/questions/1/answers/", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer valid_token")
	w := httptest.NewRecorder()

	handler.CreateAnswer(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var answer models.Answer
	json.Unmarshal(w.Body.Bytes(), &answer)

	assert.Equal(t, uint(1), answer.ID)
	assert.Equal(t, uint(1), answer.QuestionID)
	assert.Equal(t, uint(123), answer.UserID)
	assert.Equal(t, "Это ответ на вопрос", answer.Text)

	questionRepo.AssertExpectations(t)
	userRepo.AssertExpectations(t)
	answerRepo.AssertExpectations(t)
}

func TestAnswerHandler_GetAnswer_Success(t *testing.T) {
	answerRepo := new(MockAnswerRepository)
	questionRepo := new(MockQuestionRepository)
	userRepo := new(MockUserRepository)
	handler := handlers.NewAnswerHandler(answerRepo, questionRepo, userRepo)

	expectedAnswer := &models.Answer{
		ID:         1,
		QuestionID: 1,
		UserID:     123,
		Text:       "Тестовый ответ",
		Question: models.Question{
			ID:   1,
			Text: "Вопрос",
		},
	}

	answerRepo.On("GetByID", uint(1)).Return(expectedAnswer, nil)

	req := httptest.NewRequest("GET", "/answers/1", nil)
	w := httptest.NewRecorder()

	handler.GetAnswer(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var answer models.Answer
	json.Unmarshal(w.Body.Bytes(), &answer)

	assert.Equal(t, uint(1), answer.ID)
	assert.Equal(t, "Тестовый ответ", answer.Text)
	assert.Equal(t, uint(1), answer.QuestionID)
	assert.Equal(t, "Вопрос", answer.Question.Text)

	answerRepo.AssertExpectations(t)
}

func TestAnswerHandler_GetAnswer_NotFound(t *testing.T) {
	answerRepo := new(MockAnswerRepository)
	questionRepo := new(MockQuestionRepository)
	userRepo := new(MockUserRepository)
	handler := handlers.NewAnswerHandler(answerRepo, questionRepo, userRepo)

	answerRepo.On("GetByID", uint(999)).Return(nil, gorm.ErrRecordNotFound)

	req := httptest.NewRequest("GET", "/answers/999", nil)
	w := httptest.NewRecorder()

	handler.GetAnswer(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)

	var errorResponse models.ErrorResponse
	json.Unmarshal(w.Body.Bytes(), &errorResponse)

	assert.Equal(t, "Not Found", errorResponse.Error)
	assert.Contains(t, errorResponse.Message, "Answer not found")

	answerRepo.AssertExpectations(t)
}
