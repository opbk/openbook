package web

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type JsonResponse struct {
	Object interface{}
}

func (r JsonResponse) String() string {
	js, err := json.Marshal(r.Object)
	if err != nil {
		return ""
	}

	return string(js)
}

func WriteJson(rw http.ResponseWriter, object interface{}) {
	rw.Header().Set("Content-Type", "application/json")
	fmt.Fprint(rw, JsonResponse{object})
}

func NotFound(rw http.ResponseWriter) {
	rw.WriteHeader(http.StatusNotFound)
	WriteJson(rw, map[string]string{"error": "not_found"})
}

func BadRequest(rw http.ResponseWriter, message string) {
	rw.WriteHeader(http.StatusBadRequest)
	WriteJson(rw, map[string]string{"error": message})
}

func Forbidden(rw http.ResponseWriter) {
	rw.WriteHeader(http.StatusForbidden)
	WriteJson(rw, map[string]string{"error": "forbidden"})
}
