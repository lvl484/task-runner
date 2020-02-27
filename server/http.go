package server

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lvl484/task-runner/model"
)

const TaskID string = "ID"

type HTTP struct {
	service Service
	address string
}

type Service interface {
	CreateTask(ctx context.Context, input *model.TaskInput) (string, error)
	CreateAction(ctx context.Context, input *model.TaskInput) (string, error)
	DeleteTask(ctx context.Context, id string) error
	UpdateTask(ctx context.Context, id string, input *model.TaskInput) error
	UpdateAction(ctx context.Context, id string, input *model.TaskInput) error
	GetTask(ctx context.Context, id string) (*model.Task, error)
}

func NewHTTP(s Service, addr string) *HTTP {
	return &HTTP{
		service: s,
		address: addr,
	}
}

func (h *HTTP) writingResponse(w http.ResponseWriter, value []byte, str string) {
	_, err := w.Write(value)
	if err != nil {
		log.Println(str, err)
	}
}

// CreateTask Decodes request and adds new task into memory
func (h *HTTP) CreateTask(w http.ResponseWriter, r *http.Request) {
	var task model.TaskInput

	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		log.Println("create task [json decoding]:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id, err := h.service.CreateTask(r.Context(), &task)
	if err != nil {
		log.Println("create task:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	h.writingResponse(w, []byte(id), "create task [writing into response]:")
}

// CreateTask Decodes request and adds new action into memory
func (h *HTTP) CreateAction(w http.ResponseWriter, r *http.Request) {
	var task model.TaskInput

	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		log.Println("create task [json decoding]:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id, err := h.service.CreateAction(r.Context(), &task)
	if err != nil {
		log.Println("create action:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	h.writingResponse(w, []byte(id), "create action [writing into response]:")
}

// GetTaskStatus reads id from request and gets task from memory
// Write into response task Status
func (h *HTTP) GetTaskStatus(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars[TaskID]

	task, err := h.service.GetTask(r.Context(), id)
	if err != nil {
		log.Println("get status:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	h.writingResponse(w, []byte(task.Executions[len(task.Executions)-1].Status), "get status [writing into response]:")
}

// UpdateTask reads id from request, decodes request to get task, and updates memory
func (h *HTTP) UpdateTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars[TaskID]

	var task model.TaskInput

	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		log.Println("create task [json decoding]:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.service.UpdateTask(r.Context(), id, &task)
	if err != nil {
		log.Println("update task:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h *HTTP) UpdateAction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars[TaskID]

	var task model.TaskInput

	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		log.Println("create task [json decoding]:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.service.UpdateAction(r.Context(), id, &task)
	if err != nil {
		log.Println("update task:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// DeleteTask reads id from request and deletes row from memory
func (h *HTTP) DeleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars[TaskID]

	err := h.service.DeleteTask(r.Context(), id)
	if err != nil {
		log.Println("delete task:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// GetTaskOutput reads id from request and gets task from memory
// Write into response task Output
func (h *HTTP) GetTaskOutput(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars[TaskID]

	task, err := h.service.GetTask(r.Context(), id)
	if err != nil {
		log.Println("get output:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	h.writingResponse(w, []byte(task.Executions[len(task.Executions)-1].Output), "get output [writing into response]:")
}

// GetTaskOutput reads id from request and gets history of task executions
// Write into response task outputs History
func (h *HTTP) GetTaskHistory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars[TaskID]

	task, err := h.service.GetTask(r.Context(), id)
	if err != nil {
		log.Println("get output", err)
		w.WriteHeader(http.StatusInternalServerError)
		return

	}
	data, err := json.Marshal(task.Executions)
	if err != nil {
		log.Println("error executions:", err)
		return
	}
	h.writingResponse(w, data, "get output [writing into response]:")
}

// Start create all routes and starting server
func (h *HTTP) Start() error {
	mainRoute := mux.NewRouter()
	mainRoute.HandleFunc("/tasks", h.CreateTask).Methods(http.MethodPost)
	mainRoute.HandleFunc("/tasks/action", h.CreateAction).Methods(http.MethodPost)
	mainRoute.HandleFunc("/tasks/{ID}", h.UpdateTask).Methods(http.MethodPut)
	mainRoute.HandleFunc("/tasks/action/{ID}", h.UpdateAction).Methods(http.MethodPut)
	mainRoute.HandleFunc("/tasks/{ID}", h.DeleteTask).Methods(http.MethodDelete)
	mainRoute.HandleFunc("/tasks/{ID}/output", h.GetTaskOutput).Methods(http.MethodGet)
	mainRoute.HandleFunc("/tasks/{ID}/status", h.GetTaskStatus).Methods(http.MethodGet)
	mainRoute.HandleFunc("/tasks/{ID}/history", h.GetTaskHistory).Methods(http.MethodGet)

	fmt.Printf("Server Listening at %s...\n", h.address)
	return http.ListenAndServe(h.address, mainRoute)
}
