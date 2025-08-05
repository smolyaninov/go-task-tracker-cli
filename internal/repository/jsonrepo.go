package repository

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/smolyaninov/go-task-tracker-cli/internal/domain"
)

type TaskRepository interface {
	Load() ([]domain.Task, error)
	Save([]domain.Task) error
}

type JSONRepository struct {
	filename string
}

func NewJSONRepository(filename string) *JSONRepository {
	return &JSONRepository{filename: filename}
}

func (r *JSONRepository) Load() ([]domain.Task, error) {
	file, err := os.Open(r.filename)
	if errors.Is(err, os.ErrNotExist) {
		return []domain.Task{}, nil
	}
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var tasks []domain.Task
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *JSONRepository) Save(tasks []domain.Task) error {
	tempFile := r.filename + ".tmp"

	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}

	if err := os.WriteFile(tempFile, data, 0644); err != nil {
		return err
	}

	return os.Rename(tempFile, r.filename)
}
