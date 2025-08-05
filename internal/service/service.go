package service

import (
	"fmt"
	"github.com/smolyaninov/go-task-tracker-cli/internal/domain"
)

type Service struct {
	tasks  []domain.Task
	nextID int
}

func NewServiceWithData(tasks []domain.Task) *Service {
	maxID := 0
	for _, t := range tasks {
		if t.ID > maxID {
			maxID = t.ID
		}
	}

	return &Service{
		tasks:  tasks,
		nextID: maxID + 1,
	}
}

func (s *Service) Add(description string) (domain.Task, error) {
	task, err := domain.NewTask(s.nextID, description)
	if err != nil {
		return domain.Task{}, err
	}

	s.tasks = append(s.tasks, *task)
	s.nextID++

	return *task, nil
}

func (s *Service) List(statusFilter *domain.Status) []domain.Task {
	if statusFilter == nil {
		return s.tasks
	}

	filtered := make([]domain.Task, 0)
	for _, t := range s.tasks {
		if t.Status == *statusFilter {
			filtered = append(filtered, t)
		}
	}

	return filtered
}

func (s *Service) GetAll() []domain.Task {
	return s.tasks
}

func (s *Service) FindByID(id int) (*domain.Task, error) {
	for i := range s.tasks {
		if s.tasks[i].ID == id {
			return &s.tasks[i], nil
		}
	}
	return nil, fmt.Errorf("task with ID %d not found", id)
}

func (s *Service) Update(id int, newDesc string) error {
	task, err := s.FindByID(id)
	if err != nil {
		return err
	}

	return task.UpdateDescription(newDesc)
}

func (s *Service) ChangeStatus(id int, newStatus domain.Status) error {
	task, err := s.FindByID(id)
	if err != nil {
		return err
	}
	return task.ChangeStatus(newStatus)
}

func (s *Service) Delete(id int) error {
	for i, t := range s.tasks {
		if t.ID == id {
			s.tasks = append(s.tasks[:i], s.tasks[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("task with ID %d not found", id)
}
