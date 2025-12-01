package handlers_test

import (
	"fmt"
	"question-answer/models"

	"github.com/stretchr/testify/mock"
)

type MockQuestionRepository struct {
	mock.Mock
}

func (m *MockQuestionRepository) Create(question *models.Question) error {
	args := m.Called(question)
	return args.Error(0)
}

func (m *MockQuestionRepository) GetByID(id uint) (*models.Question, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Question), args.Error(1)
}

func (m *MockQuestionRepository) GetWithAnswers(id uint) (*models.Question, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Question), args.Error(1)
}

func (m *MockQuestionRepository) GetAll() ([]models.Question, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]models.Question), args.Error(1)
}

func (m *MockQuestionRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

type MockAnswerRepository struct {
	mock.Mock
}

func (m *MockAnswerRepository) Create(answer *models.Answer) error {
	args := m.Called(answer)
	return args.Error(0)
}

func (m *MockAnswerRepository) GetByID(id uint) (*models.Answer, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Answer), args.Error(1)
}

func (m *MockAnswerRepository) GetByQuestionID(questionID uint) ([]models.Answer, error) {
	args := m.Called(questionID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]models.Answer), args.Error(1)
}

func (m *MockAnswerRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) GetByToken(token string) (*models.User, error) {
	args := m.Called(token)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	fmt.Print(token)
	return args.Get(0).(*models.User), args.Error(1)
}
