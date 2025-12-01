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
)

func TestQuestionHandler_CreateQuestion_Success(t *testing.T) {
	questionRepo := new(MockQuestionRepository)
	handler := handlers.NewQuestionHandler(questionRepo)

	questionRepo.On("Create", mock.AnythingOfType("*models.Question")).
		Run(func(args mock.Arguments) {
			question := args.Get(0).(*models.Question)
			question.ID = 1
		}).
		Return(nil)

	questionData := models.CreateQuestionRequest{
		Text: "Как работает Go?",
	}
	jsonData, _ := json.Marshal(questionData)

	req := httptest.NewRequest("POST", "/questions/", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	handler.CreateQuestion(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var response models.Question
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, uint(1), response.ID)
	assert.Equal(t, "Как работает Go?", response.Text)

	questionRepo.AssertExpectations(t)
}

func TestQuestionHandler_GetQuestions_Success(t *testing.T) {
	questionRepo := new(MockQuestionRepository)
	handler := handlers.NewQuestionHandler(questionRepo)

	expectedQuestions := []models.Question{
		{ID: 1, Text: "Первый вопрос"},
		{ID: 2, Text: "Второй вопрос"},
		{ID: 3, Text: "Третий вопрос"},
	}

	questionRepo.On("GetAll").Return(expectedQuestions, nil)

	req := httptest.NewRequest("GET", "/questions/", nil)
	w := httptest.NewRecorder()

	handler.GetQuestions(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var questions []models.Question
	json.Unmarshal(w.Body.Bytes(), &questions)

	assert.Len(t, questions, 3)
	assert.Equal(t, "Первый вопрос", questions[0].Text)
	assert.Equal(t, "Третий вопрос", questions[2].Text)

	questionRepo.AssertExpectations(t)
}

func TestQuestionHandler_GetQuestion_Success(t *testing.T) {
	questionRepo := new(MockQuestionRepository)
	handler := handlers.NewQuestionHandler(questionRepo)

	expectedQuestion := &models.Question{
		ID:   1,
		Text: "Тестовый вопрос",
		Answers: []models.Answer{
			{ID: 1, Text: "Первый ответ", UserID: 100},
			{ID: 2, Text: "Второй ответ", UserID: 200},
		},
	}

	questionRepo.On("GetWithAnswers", uint(1)).Return(expectedQuestion, nil)

	req := httptest.NewRequest("GET", "/questions/1", nil)
	w := httptest.NewRecorder()

	handler.GetQuestion(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var question models.Question
	json.Unmarshal(w.Body.Bytes(), &question)

	assert.Equal(t, uint(1), question.ID)
	assert.Equal(t, "Тестовый вопрос", question.Text)
	assert.Len(t, question.Answers, 2)
	assert.Equal(t, "Первый ответ", question.Answers[0].Text)

	questionRepo.AssertExpectations(t)
}

func TestQuestionHandler_DeleteQuestion_Success(t *testing.T) {
	questionRepo := new(MockQuestionRepository)
	handler := handlers.NewQuestionHandler(questionRepo)

	questionRepo.On("Delete", uint(1)).Return(nil)

	req := httptest.NewRequest("DELETE", "/questions/1", nil)
	w := httptest.NewRecorder()

	handler.DeleteQuestion(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
	assert.Empty(t, w.Body.String())

	questionRepo.AssertExpectations(t)
}
