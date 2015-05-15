package web

import (
	"encoding/json"
	"encoding/xml"
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

type XmlResponse struct {
	Object interface{}
}

func (r XmlResponse) String() string {
	x, err := xml.MarshalIndent(r.Object, "", "  ")
	if err != nil {
		return ""
	}

	return string(x)
}

func WriteXml(rw http.ResponseWriter, object interface{}) {
	rw.Header().Add("Content-Type", "text/xml")
	rw.Write([]byte(xml.Header))
	fmt.Fprint(rw, XmlResponse{object})
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
