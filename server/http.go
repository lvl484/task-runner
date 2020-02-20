package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lvl484/task-runner/model"
	"github.com/lvl484/task-runner/service"
)

type Operations string

const (
	CreateTask   Operations = "create task"
	CreateAction Operations = "create action"
	GetStatus    Operations = "get status"
	UpdateTask   Operations = "update task"
	DeleteTask   Operations = "delete task"
	GetOutput    Operations = "get output"
)

const TaskID string = "ID"

type HTTP struct {
	service *service.Service
	address string
}

func NewHTTP(s *service.Service, addr string) *HTTP {
	return &HTTP{
		service: s,
		address: addr,
	}
}

// CreateTask Decodes request and adds new task into memory
func (h *HTTP) CreateTask(w http.ResponseWriter, r *http.Request) {
	var task model.Task

	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		log.Println(CreateTask, err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id, err := h.service.CreateTask(r.Context(), &task)
	if err != nil {
		log.Println(CreateAction, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, err = w.Write([]byte(id))
	if err != nil {
		log.Println(CreateTask, err)
	}
}

// CreateTask Decodes request and adds new action into memory
func (h *HTTP) CreateAction(w http.ResponseWriter, r *http.Request) {
	var task model.Task

	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		log.Println(CreateAction, err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id, err := h.service.CreateAction(r.Context(), &task)
	if err != nil {
		log.Println(CreateAction, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, err = w.Write([]byte(id))
	if err != nil {
		log.Println(CreateAction, err)
	}
}

// GetTaskStatus reads id from request and gets task from memory
// Write into response task Status
func (h *HTTP) GetTaskStatus(w http.ResponseWriter, r *http.Request) {
	id := GetIdFromRequest(r)
	task, err := h.service.GetTask(r.Context(), id)
	if err != nil {
		log.Println(GetStatus, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, err = w.Write([]byte(task.Status))
	if err != nil {
		log.Println(GetStatus, err)
	}
}

// UpdateTask reads id from request, decodes request to get task, and updates memory
func (h *HTTP) UpdateTask(w http.ResponseWriter, r *http.Request) {
	id := GetIdFromRequest(r)
	var task model.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		log.Println(UpdateTask, err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = h.service.UpdateTask(r.Context(), id, &task)
	if err != nil {
		log.Println(UpdateTask, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// DeleteTask reads id from request and deletes row from memory
func (h *HTTP) DeleteTask(w http.ResponseWriter, r *http.Request) {
	id := GetIdFromRequest(r)
	err := h.service.DeleteTask(r.Context(), id)
	if err != nil {
		log.Println(DeleteTask, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// GetTaskOutput reads id from request and gets task from memory
// Write into response task Output
func (h *HTTP) GetTaskOutput(w http.ResponseWriter, r *http.Request) {
	id := GetIdFromRequest(r)
	task, err := h.service.GetTask(r.Context(), id)
	if err != nil {
		log.Println(GetOutput, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, err = w.Write([]byte(task.Output))
	if err != nil {
		log.Println(GetOutput, err)
	}
}

// Start create all routes and starting server
func (h *HTTP) Start() error {
	mainRoute := mux.NewRouter()
	mainRoute.HandleFunc("/tasks", h.CreateTask).Methods(http.MethodPost)
	mainRoute.HandleFunc("/tasks/{ID}", h.UpdateTask).Methods(http.MethodPut)
	mainRoute.HandleFunc("/tasks/{ID}", h.DeleteTask).Methods(http.MethodDelete)
	mainRoute.HandleFunc("/tasks/{ID}/output", h.GetTaskOutput).Methods(http.MethodGet)
	mainRoute.HandleFunc("/tasks/{ID}/status", h.GetTaskStatus).Methods(http.MethodGet)
	mainRoute.HandleFunc("/tasks/action", h.CreateAction).Methods(http.MethodPost)

	fmt.Printf("Server Listening...")
	return http.ListenAndServe(h.address, mainRoute)
}
