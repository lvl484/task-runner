package server

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/lvl484/task-runner/model"
	"github.com/lvl484/task-runner/service"
	"log"
	"net/http"
)

type HTTP struct {
	service *service.Service
	address string
}

func NewHTTP(s *service.Service, addr string) HTTP {
	return HTTP{
		service: s,
		address: addr,
	}
}

func (h HTTP) CreateTask(w http.ResponseWriter, r *http.Request) {
	var task model.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		log.Println("create task:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	id, err := h.service.CreateTask(r.Context(), &task)
	if err != nil {
		log.Println("create task:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, err = w.Write([]byte(id))
	if err != nil {
		log.Println("create task:", err)
	}
}

func (h HTTP) GetTaskStatus(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["ID"]
	task, err := h.service.GetTask(r.Context(), id)
	if err != nil {
		log.Println("get status:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, err = w.Write([]byte(task.Status))
	if err != nil {
		log.Println("get status:", err)
	}
}

func (h HTTP) UpdateTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["ID"]
	var task model.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		log.Println("update task:", err)
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
func (h HTTP) DeleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["ID"]
	err := h.service.DeleteTask(r.Context(), id)
	if err != nil {
		log.Println("delete task:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h HTTP) GetTaskOutput(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["ID"]
	task, err := h.service.GetTask(r.Context(), id)
	if err != nil {
		log.Println("get status:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, err = w.Write([]byte(task.Output))
	if err != nil {
		log.Println("get status:", err)
	}
}

func (h HTTP) Start() error {
	mainRoute := mux.NewRouter()
	mainRoute.HandleFunc("/tasks", h.CreateTask).Methods(http.MethodPost)
	mainRoute.HandleFunc("/tasks/{ID}", h.UpdateTask).Methods(http.MethodPut)
	mainRoute.HandleFunc("/tasks/{ID}", h.DeleteTask).Methods(http.MethodDelete)
	mainRoute.HandleFunc("/tasks/{ID}/output", h.GetTaskOutput).Methods(http.MethodGet)
	mainRoute.HandleFunc("/tasks/{ID}/status", h.GetTaskStatus).Methods(http.MethodGet)

	fmt.Printf("Server Listening...")
	return http.ListenAndServe(h.address, mainRoute)
}
