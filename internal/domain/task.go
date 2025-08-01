package domain

import (
	"errors"
	"time"
)

type Status string

const (
	StatusTodo       Status = "todo"
	StatusInProgress Status = "in-progress"
	StatusDone       Status = "done"
)

type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Status      Status    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func NewTask(id int, desc string) (*Task, error) {
	if desc == "" {
		return nil, errors.New("description cannot be empty")
	}

	now := time.Now().UTC()

	return &Task{
		ID:          id,
		Description: desc,
		Status:      StatusTodo,
		CreatedAt:   now,
		UpdatedAt:   now,
	}, nil
}

func (t *Task) UpdateDescription(desc string) error {
	if desc == "" {
		return errors.New("description cannot be empty")
	}

	t.Description = desc
	t.UpdatedAt = time.Now().UTC()

	return nil
}

func (t *Task) ChangeStatus(newStatus Status) error {
	switch newStatus {
	case StatusTodo, StatusInProgress, StatusDone:
		t.Status = newStatus
		t.UpdatedAt = time.Now().UTC()
		return nil
	default:
		return errors.New("invalid status")
	}
}
