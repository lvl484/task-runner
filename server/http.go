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

func (h *HTTP) jsonDecoding(w http.ResponseWriter, r *http.Request, task *model.TaskInput, str string) {
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		log.Println(str, err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func (h *HTTP) gettingTask(w http.ResponseWriter, r *http.Request, str string, id string) *model.Task {
	task, err := h.service.GetTask(r.Context(), id)
	if err != nil {
		log.Println(str, err)
		w.WriteHeader(http.StatusInternalServerError)
		return nil
	}
	return task
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
	h.jsonDecoding(w, r, &task, "create task [json decoding]:")

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
	h.jsonDecoding(w, r, &task, "create action [json decoding]")

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
	task := h.gettingTask(w, r, "get status:", id)
	h.writingResponse(w, []byte(task.Status), "get status [writing into response]:")
}

// UpdateTask reads id from request, decodes request to get task, and updates memory
func (h *HTTP) UpdateTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars[TaskID]
	var task model.TaskInput
	h.jsonDecoding(w, r, &task, "create action [json decoding]")

	err := h.service.UpdateTask(r.Context(), id, &task)
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
	task := h.gettingTask(w, r, "get output:", id)
	h.writingResponse(w, []byte(task.Output), "get output [writing into response]:")
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
