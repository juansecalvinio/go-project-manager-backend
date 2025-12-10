package services

import (
	"time"

	"go-project-manager-backend/internal/domain/models"
)

type TaskRepository interface {
	Create(task *models.Task) error
	GetByID(id string) (*models.Task, error)
	Update(task *models.Task) error
	Delete(id string) error
	ListByProject(projectID string) ([]*models.Task, error)
	ListByAssignee(assigneeID string) ([]*models.Task, error)
	ListBySprint(sprintID string) ([]*models.Task, error)
	ListBacklog(projectID string) ([]*models.Task, error)
}

type TaskService struct {
	repository TaskRepository
}

func NewTaskService(repository TaskRepository) *TaskService {
	return &TaskService{repository: repository}
}

func (s *TaskService) CreateTask(title, description, projectID, assigneeID string) (*models.Task, error) {
	task := &models.Task{
		ID:          generateID(),
		Title:       title,
		Description: description,
		ProjectID:   projectID,
		AssigneeID:  assigneeID,
		SprintID:    nil,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	err := s.repository.Create(task)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (s *TaskService) GetTask(id string) (*models.Task, error) {
	return s.repository.GetByID(id)
}

func (s *TaskService) UpdateTask(task *models.Task) error {
	task.UpdatedAt = time.Now()
	return s.repository.Update(task)
}

func (s *TaskService) DeleteTask(id string) error {
	return s.repository.Delete(id)
}

func (s *TaskService) AssignToSprint(taskID, sprintID string) error {
	task, err := s.repository.GetByID(taskID)
	if err != nil {
		return err
	}

	task.SprintID = &sprintID
	task.UpdatedAt = time.Now()
	return s.repository.Update(task)
}

func (s *TaskService) MoveToBacklog(taskID string) error {
	task, err := s.repository.GetByID(taskID)
	if err != nil {
		return err
	}

	task.SprintID = nil
	task.UpdatedAt = time.Now()
	return s.repository.Update(task)
}

func (s *TaskService) ListTasksByProject(projectID string) ([]*models.Task, error) {
	return s.repository.ListByProject(projectID)
}

func (s *TaskService) ListTasksByAssignee(assigneeID string) ([]*models.Task, error) {
	return s.repository.ListByAssignee(assigneeID)
}

func (s *TaskService) ListTasksBySprint(sprintID string) ([]*models.Task, error) {
	return s.repository.ListBySprint(sprintID)
}

func (s *TaskService) ListBacklog(projectID string) ([]*models.Task, error) {
	return s.repository.ListBacklog(projectID)
}
