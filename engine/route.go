package engine

import "net/http"

type Route struct {
	Logger  bool
	Handler http.Handler
}
