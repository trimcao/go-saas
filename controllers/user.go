package controllers

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/trimcao/go-saas/engine"
)

type User struct{}

func (u User) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var head string
	head, r.URL.Path = engine.ShiftPath(r.URL.Path)
	if head == "profile" {
		u.Profile(w, r)
		return
	} else if head == "detail" {
		head, _ := engine.ShiftPath(r.URL.Path)
		i, err := strconv.Atoi(head)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		ctx := r.Context()
		ctx = context.WithValue(ctx, engine.ContextTestKey, "pulisic")
		ctx = context.WithValue(ctx, engine.ContextUserID, i)
		u.Detail(w, r.WithContext(ctx))
		return
	}
	engine.Respond(w, r, http.StatusNotFound, "user not found")
}

func (u User) Detail(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	v := ctx.Value(engine.ContextTestKey)
	id := ctx.Value(engine.ContextUserID)
	engine.Respond(w, r, http.StatusOK, fmt.Sprintf("value of context is %s for user id %d", v, id))
}

func (u User) Profile(w http.ResponseWriter, r *http.Request) {
	engine.Respond(w, r, http.StatusOK, "profile was called")
}
