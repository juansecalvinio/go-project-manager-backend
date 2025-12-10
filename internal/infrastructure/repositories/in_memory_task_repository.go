package repositories

import (
	"errors"
	"sync"

	"go-project-manager-backend/internal/domain/models"
)

type InMemoryTaskRepository struct {
	tasks map[string]*models.Task
	mu    sync.RWMutex
}

func NewInMemoryTaskRepository() *InMemoryTaskRepository {
	return &InMemoryTaskRepository{
		tasks: make(map[string]*models.Task),
	}
}

func (r *InMemoryTaskRepository) Create(task *models.Task) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.tasks[task.ID]; exists {
		return errors.New("task already exists")
	}

	r.tasks[task.ID] = task
	return nil
}

func (r *InMemoryTaskRepository) GetByID(id string) (*models.Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	task, exists := r.tasks[id]
	if !exists {
		return nil, errors.New("task not found")
	}
	return task, nil
}

func (r *InMemoryTaskRepository) Update(task *models.Task) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.tasks[task.ID]; !exists {
		return errors.New("task not found")
	}

	r.tasks[task.ID] = task
	return nil
}

func (r *InMemoryTaskRepository) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.tasks[id]; !exists {
		return errors.New("task not found")
	}

	delete(r.tasks, id)
	return nil
}

func (r *InMemoryTaskRepository) ListByProject(projectID string) ([]*models.Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	tasks := make([]*models.Task, 0)
	for _, task := range r.tasks {
		if task.ProjectID == projectID {
			tasks = append(tasks, task)
		}
	}

	return tasks, nil
}

func (r *InMemoryTaskRepository) ListByAssignee(assigneeID string) ([]*models.Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	tasks := make([]*models.Task, 0)
	for _, task := range r.tasks {
		if task.AssigneeID == assigneeID {
			tasks = append(tasks, task)
		}
	}

	return tasks, nil
}

func (r *InMemoryTaskRepository) ListBySprint(sprintID string) ([]*models.Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	tasks := make([]*models.Task, 0)
	for _, task := range r.tasks {
		if task.SprintID != nil && *task.SprintID == sprintID {
			tasks = append(tasks, task)
		}
	}

	return tasks, nil
}

func (r *InMemoryTaskRepository) ListBacklog(projectID string) ([]*models.Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	tasks := make([]*models.Task, 0)
	for _, task := range r.tasks {
		if task.ProjectID == projectID && task.SprintID == nil {
			tasks = append(tasks, task)
		}
	}

	return tasks, nil
}
