package animalservice

import (
	"context"
	"errors"
	"sync"

	"github.com/google/uuid"
	"github.com/sonix66/animalito-bot/internal/entity"
)

type AnimalServiceMock struct {
	mu             sync.Mutex
	animals        map[string]*entity.Animal
	CreateFunc     func(ctx context.Context, animal *entity.Animal) error
	GetByIDFunc    func(ctx context.Context, id string) (*entity.Animal, error)
	GetListFunc    func(ctx context.Context, count, offset int) ([]*entity.Animal, error)
	UpdateByIDFunc func(ctx context.Context, animal *entity.Animal) error
	DeleteByIDFunc func(ctx context.Context, id string) error
}

func NewAnimalServiceMock() *AnimalServiceMock {
	return &AnimalServiceMock{
		animals: make(map[string]*entity.Animal),
	}
}

func (m *AnimalServiceMock) CreateAnimal(ctx context.Context, animal *entity.Animal) error {
	if m.CreateFunc != nil {
		return m.CreateFunc(ctx, animal)
	}

	animal.ID = uuid.New().String()
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, exists := m.animals[animal.ID]; exists {
		return errors.New("animal already exists")
	}

	m.animals[animal.ID] = animal
	return nil
}

func (m *AnimalServiceMock) GetAnimalByID(ctx context.Context, id string) (*entity.Animal, error) {
	if m.GetByIDFunc != nil {
		return m.GetByIDFunc(ctx, id)
	}

	m.mu.Lock()
	defer m.mu.Unlock()

	animal, exists := m.animals[id]
	if !exists {
		return nil, errors.New("animal not found")
	}

	return animal, nil
}

func (m *AnimalServiceMock) GetAnimalList(ctx context.Context, count, offset int) ([]*entity.Animal, error) {
	if m.GetListFunc != nil {
		return m.GetListFunc(ctx, count, offset)
	}

	m.mu.Lock()
	defer m.mu.Unlock()

	animals := make([]*entity.Animal, 0, len(m.animals))
	for _, animal := range m.animals {
		animals = append(animals, animal)
	}

	start := offset
	end := offset + count
	if start > len(animals) {
		start = len(animals)
	}
	if end > len(animals) {
		end = len(animals)
	}

	return animals[start:end], nil
}

func (m *AnimalServiceMock) UpdateAnimalByID(ctx context.Context, animal *entity.Animal) error {
	if m.UpdateByIDFunc != nil {
		return m.UpdateByIDFunc(ctx, animal)
	}

	m.mu.Lock()
	defer m.mu.Unlock()

	if _, exists := m.animals[animal.ID]; !exists {
		return errors.New("animal not found")
	}

	m.animals[animal.ID] = animal
	return nil
}

func (m *AnimalServiceMock) DeleteAnimalByID(ctx context.Context, id string) error {
	if m.DeleteByIDFunc != nil {
		return m.DeleteByIDFunc(ctx, id)
	}

	m.mu.Lock()
	defer m.mu.Unlock()

	if _, exists := m.animals[id]; !exists {
		return errors.New("animal not found")
	}

	delete(m.animals, id)
	return nil
}
