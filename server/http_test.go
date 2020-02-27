package server

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHTTP_CreateTask(t *testing.T) {
	tests := []struct {
		want  int
		input string
		err   error
	}{
		{
			want:  200,
			input: `{"script":"pwd", "schedule": {"count":6}}`,
			err:   nil,
		},
		{
			want:  500,
			input: `{"script":"pwd", "schedule": {"count":5}}`,
			err:   errors.New("incorrect"),
		},
	}

	for _, tt := range tests {
		h := NewHTTP(mockService{err: tt.err}, ":8099")
		rr := httptest.NewRecorder()

		req, err := http.NewRequest("POST", "/tasks", strings.NewReader(tt.input))
		if err != nil {
			t.Fatal(err)
		}

		h.CreateTask(rr, req)

		if rr.Code != tt.want {
			t.Fatalf("want %v got %v", tt.want, rr.Code)
		}
	}
}

func TestHTTP_CreateAction(t *testing.T) {
	tests := []struct {
		want  int
		input string
		err   error
	}{
		{
			want:  200,
			input: `{"script":"CurrentTime", "schedule": {"count":2}}`,
			err:   nil,
		},
		{
			want:  500,
			input: `{"script":"CurrentTime", "schedule": {"count":2}}`,
			err:   errors.New("incorrect"),
		},
	}

	for _, tt := range tests {
		h := NewHTTP(mockService{err: tt.err}, ":8099")
		rr := httptest.NewRecorder()

		req, err := http.NewRequest("POST", "/tasks/action", strings.NewReader(tt.input))
		if err != nil {
			t.Fatal(err)
		}

		h.CreateAction(rr, req)

		if rr.Code != tt.want {
			t.Fatalf("want %v got %v", tt.want, rr.Code)
		}
	}
}

func TestHTTP_GetTaskStatus(t *testing.T) {
	tests := []struct {
		want  int
		input string
		err   error
	}{
		{
			want:  200,
			input: "GetTask",
			err:   nil,
		},
		{
			want:  500,
			input: "GetTask",
			err:   errors.New("error incorrect"),
		},
	}

	for _, tt := range tests {
		h := NewHTTP(mockService{err: tt.err}, ":8099")
		rr := httptest.NewRecorder()
		req, err := http.NewRequest("GET", "/tasks/"+tt.input+"/status", nil)
		if err != nil {
			t.Fatal(err)
		}

		h.GetTaskStatus(rr, req)
		if rr.Code != tt.want {
			t.Fatalf("want %v got %v", tt.want, rr.Code)
		}
	}
}

func TestHTTP_GetTaskOutput(t *testing.T) {
	tests := []struct {
		want  int
		input string
		err   error
	}{
		{
			want:  200,
			input: "GetTask",
			err:   nil,
		},
		{
			want:  500,
			input: "GetTask",
			err:   errors.New("error incorrect"),
		},
	}

	for _, tt := range tests {
		h := NewHTTP(mockService{err: tt.err}, ":8099")
		rr := httptest.NewRecorder()
		req, err := http.NewRequest("GET", "/tasks/"+tt.input+"/output", nil)
		if err != nil {
			t.Fatal(err)
		}

		h.GetTaskOutput(rr, req)
		if rr.Code != tt.want {
			t.Fatalf("want %v got %v", tt.want, rr.Code)
		}
	}
}

func TestHTTP_DeleteTask(t *testing.T) {
	tests := []struct {
		want  int
		input string
		err   error
	}{
		{
			want:  200,
			input: "GetTask",
			err:   nil,
		},
		{
			want:  500,
			input: "GetTask",
			err:   errors.New("error incorrect"),
		},
	}

	for _, tt := range tests {
		h := NewHTTP(mockService{err: tt.err}, ":8099")
		rr := httptest.NewRecorder()
		req, err := http.NewRequest("GET", "/tasks/"+tt.input, nil)
		if err != nil {
			t.Fatal(err)
		}

		h.DeleteTask(rr, req)
		if rr.Code != tt.want {
			t.Fatalf("want %v got %v", tt.want, rr.Code)
		}
	}
}

func TestHTTP_GetTaskHistory(t *testing.T) {
	tests := []struct {
		want  int
		input string
		err   error
	}{
		{
			want:  200,
			input: "GetTask",
			err:   nil,
		},
		{
			want:  500,
			input: "GetTask",
			err:   errors.New("error incorrect"),
		},
	}

	for _, tt := range tests {
		h := NewHTTP(mockService{err: tt.err}, ":8099")
		rr := httptest.NewRecorder()
		req, err := http.NewRequest("GET", "/tasks/"+tt.input+"/history", nil)
		if err != nil {
			t.Fatal(err)
		}

		h.GetTaskHistory(rr, req)
		if rr.Code != tt.want {
			t.Fatalf("want %v got %v", tt.want, rr.Code)
		}
	}
}

func TestHTTP_UpdateTask(t *testing.T) {
	tests := []struct {
		want  int
		input string
		err   error
	}{
		{
			want:  200,
			input: `{"script":"pwd", "schedule": {"count":5}}`,
			err:   nil,
		},
		{
			want:  500,
			input: `{"script":"pwd", "schedule": {"count":4}}`,
			err:   errors.New("error incorrect"),
		},
	}

	for _, tt := range tests {
		h := NewHTTP(mockService{err: tt.err}, ":8099")
		rr := httptest.NewRecorder()
		req, err := http.NewRequest("PUT", "/tasks/GetTask", strings.NewReader(tt.input))
		if err != nil {
			t.Fatal(err)
		}

		h.UpdateTask(rr, req)
		if rr.Code != tt.want {
			t.Fatalf("want %v got %v", tt.want, rr.Code)
		}
	}
}

func TestHTTP_UpdateAction(t *testing.T) {
	tests := []struct {
		want  int
		input string
		err   error
	}{
		{
			want:  200,
			input: `{"script":"CurrentTime", "schedule": {"count":5}}`,
			err:   nil,
		},
		{
			want:  500,
			input: `{"script":"CurrentTime", "schedule": {"count":4}}`,
			err:   errors.New("error incorrect"),
		},
	}

	for _, tt := range tests {
		h := NewHTTP(mockService{err: tt.err}, ":8099")
		rr := httptest.NewRecorder()
		req, err := http.NewRequest("PUT", "/tasks/action/GetTask", strings.NewReader(tt.input))
		if err != nil {
			t.Fatal(err)
		}

		h.UpdateAction(rr, req)
		if rr.Code != tt.want {
			t.Fatalf("want %v got %v", tt.want, rr.Code)
		}
	}
}

func TestHTTP_Start(t *testing.T) {
	h := NewHTTP(mockService{}, ":aaaaa")
	err := h.Start()
	if err == nil {
		t.Errorf("want error got nil")
	}
}
