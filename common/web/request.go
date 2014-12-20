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

func (r *request) GetInt64(name string) int64 {
	param, _ := strconv.ParseInt(r.params[name], 10, 64)
	return param
}

func (r *request) GetString(name string) string {
	return r.params[name]
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
