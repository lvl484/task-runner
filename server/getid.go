package server

import (
	"github.com/gorilla/mux"
	"net/http"
)

// GetIdFromRequest gets variable ID from request url
// and returns it as string
func GetIdFromRequest(r *http.Request) string {
	vars := mux.Vars(r)
	return vars[TaskID]
}
