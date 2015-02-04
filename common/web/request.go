package web

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
)

type Request struct {
	req    *http.Request
	params map[string]string
}

func NewRequest(req *http.Request) *Request {
	return &Request{req, mux.Vars(req)}
}

func (r *Request) GetInt64(name string, def ...int64) int64 {
	value, ok := r.params[name]
	if ok {
		intValue, _ := strconv.ParseInt(value, 10, 64)
		return intValue
	}

	value = r.req.FormValue(name)
	if value != "" {
		intValue, _ := strconv.ParseInt(value, 10, 64)
		return intValue
	}

	if len(def) > 0 {
		return def[0]
	}

	return int64(0)
}

func (r *Request) GetInt(name string, def ...int) int {
	value, ok := r.params[name]
	if ok {
		intValue, _ := strconv.Atoi(value)
		return intValue
	}

	value = r.req.FormValue(name)
	if value != "" {
		intValue, _ := strconv.Atoi(value)
		return intValue
	}

	if len(def) > 0 {
		return def[0]
	}

	return 0
}

func (r *Request) GetString(name string, def ...string) string {
	value, ok := r.params[name]
	if ok {
		return value
	}

	value = r.req.FormValue(name)
	if value != "" {
		return value
	}

	if len(def) > 0 {
		return def[0]
	}

	return ""
}

func (r *Request) GetObject(t ...interface{}) error {
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

func (r *Request) GetForm(t interface{}) error {
	decoder := schema.NewDecoder()
	if err := r.req.ParseForm(); err != nil {
		return err
	}

	return decoder.Decode(t, r.req.PostForm)
}
