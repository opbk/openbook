package web

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type request struct {
	req    *http.Request
	params map[string]string
}

func NewRequest(req *http.Request) *request {
	return &request{req, mux.Vars(req)}
}

func (r *request) GetInt64(name string, def ...int64) int64 {
	value, ok := r.params[name]
	if ok {
		intValue, _ := strconv.ParseInt(value, 10, 64)
		return intValue
	}

	return def[0]
}

func (r *request) GetInt(name string, def ...int) int {
	value, ok := r.params[name]
	if ok {
		intValue, _ := strconv.Atoi(value)
		return intValue
	}

	return def[0]
}

func (r *request) GetString(name string, def ...string) string {
	value, ok := r.params[name]
	if ok {
		return value
	}

	return def[0]
}

func (r *request) GetObject(t ...interface{}) error {
	if len(t) > 1 {
		if param, ok := r.params[t[0].(string)]; ok {
			data := []byte(param)
			return json.Unmarshal(data, t[1])
		}

		if file, _, err := r.req.FormFile(t[0].(string)); err == nil {
			decoder := json.NewDecoder(file)
			return decoder.Decode(t[1])
		}
	}

	decoder := json.NewDecoder(r.req.Body)
	return decoder.Decode(t[0])
}

func (r *request) GetObjectFromFile(name, t interface{}) {

}
