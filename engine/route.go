package engine

import (
	"net/http"

	"github.com/trimcao/go-saas/data/model"
)

type Route struct {
	Logger  bool
	Handler http.Handler

	// authorization
	MinimumRole model.Roles
}
