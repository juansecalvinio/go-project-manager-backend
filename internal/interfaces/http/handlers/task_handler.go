package handlers

import (
	"encoding/json"
	"go-project-manager-backend/internal/domain/models"
	"go-project-manager-backend/internal/domain/services"
	"net/http"
)

type TaskHandler struct {
	taskService *services.TaskService
}

func NewTaskHandler(taskService *services.TaskService) *TaskHandler {
	return &TaskHandler{
		taskService: taskService,
	}
}

type CreateTaskRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	ProjectID   string `json:"project_id"`
	AssigneeID  string `json:"assignee_id"`
}

type UpdateTaskRequest struct {
	Title       string            `json:"title"`
	Description string            `json:"description"`
	Status      models.TaskStatus `json:"status,omitempty"`
	AssigneeID  string            `json:"assignee_id"`
}

func (h *TaskHandler) CreateTask(w http.ResponseWriter, req *http.Request) {
	var taskRequest CreateTaskRequest
	if err := json.NewDecoder(req.Body).Decode(&taskRequest); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	task, err := h.taskService.CreateTask(taskRequest.Title, taskRequest.Description, taskRequest.ProjectID, taskRequest.AssigneeID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

func (h *TaskHandler) GetTask(w http.ResponseWriter, req *http.Request) {
	id := req.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Task ID required", http.StatusBadRequest)
		return
	}

	task, err := h.taskService.GetTask(id)
	if err != nil {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func (h *TaskHandler) UpdateTask(w http.ResponseWriter, req *http.Request) {
	id := req.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Task ID required", http.StatusBadRequest)
		return
	}

	task, err := h.taskService.GetTask(id)
	if err != nil {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	var updateRequest UpdateTaskRequest
	if err := json.NewDecoder(req.Body).Decode(&updateRequest); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if updateRequest.Title != "" {
		task.Title = updateRequest.Title
	}
	if updateRequest.Description != "" {
		task.Description = updateRequest.Description
	}
	if updateRequest.Status != "" {
		task.Status = updateRequest.Status
	}
	if updateRequest.AssigneeID != "" {
		task.AssigneeID = updateRequest.AssigneeID
	}

	err = h.taskService.UpdateTask(task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func (h *TaskHandler) DeleteTask(w http.ResponseWriter, req *http.Request) {
	id := req.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Task ID required", http.StatusBadRequest)
		return
	}

	err := h.taskService.DeleteTask(id)
	if err != nil {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *TaskHandler) AssignToSprint(w http.ResponseWriter, req *http.Request) {
	taskID := req.URL.Query().Get("task_id")
	sprintID := req.URL.Query().Get("sprint_id")
	if taskID == "" || sprintID == "" {
		http.Error(w, "Task ID and Sprint ID required", http.StatusBadRequest)
		return
	}

	err := h.taskService.AssignToSprint(taskID, sprintID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *TaskHandler) MoveToBacklog(w http.ResponseWriter, req *http.Request) {
	taskID := req.URL.Query().Get("task_id")
	if taskID == "" {
		http.Error(w, "Task ID required", http.StatusBadRequest)
		return
	}

	err := h.taskService.MoveToBacklog(taskID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *TaskHandler) ListTasks(w http.ResponseWriter, req *http.Request) {
	projectID := req.URL.Query().Get("project_id")
	if projectID == "" {
		http.Error(w, "Project ID required", http.StatusBadRequest)
		return
	}

	tasks, err := h.taskService.ListTasksByProject(projectID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func (h *TaskHandler) ListBacklog(w http.ResponseWriter, req *http.Request) {
	projectID := req.URL.Query().Get("project_id")
	if projectID == "" {
		http.Error(w, "Project ID required", http.StatusBadRequest)
		return
	}

	tasks, err := h.taskService.ListBacklog(projectID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}
