package service

import (
	"context"
	"testing"
)

// Example repository interface for testing
type MockRepository interface {
	GetByID(ctx context.Context, id string) (string, error)
}

// Example service using BaseService
type ExampleService struct {
	BaseService[MockRepository]
}

func NewExampleService(repo MockRepository) *ExampleService {
	return &ExampleService{
		BaseService: NewBaseService(repo),
	}
}

func TestBaseService(t *testing.T) {
	// This test demonstrates that BaseService can be embedded and used
	repo := &mockRepo{data: "test"}
	service := NewExampleService(repo)

	if service.Repo == nil {
		t.Error("Repo should not be nil")
	}
}

type mockRepo struct {
	data string
}

func (m *mockRepo) GetByID(ctx context.Context, id string) (string, error) {
	return m.data, nil
}
