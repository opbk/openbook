package auth

import (
	"net/http"
)

func UserRequired(f func(rw http.ResponseWriter, req *http.Request)) func(rw http.ResponseWriter, req *http.Request) {
	return func(rw http.ResponseWriter, req *http.Request) {
		user := Get(req)
		if user != nil {
			f(rw, req)
			return
		}

		http.Redirect(rw, req, "/signup?from="+req.URL.String(), http.StatusFound)
	}
}

func GuestRequired(f func(rw http.ResponseWriter, req *http.Request)) func(rw http.ResponseWriter, req *http.Request) {
	return func(rw http.ResponseWriter, req *http.Request) {
		user := Get(req)
		if user == nil {
			f(rw, req)
			return
		}

		http.Redirect(rw, req, "/", http.StatusFound)
	}
}
