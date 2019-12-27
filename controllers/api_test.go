package controllers

import (
	"net/http"
	"net/http/httptest"
)

func logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	api := &API{
		Logger: logger,
	}
	rec := httptest.NewRecorder()
	api.ServeHTTP(rec, req)
	return rec
}
