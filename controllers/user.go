package controllers

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/trimcao/go-saas/data"
	"github.com/trimcao/go-saas/engine"
)

type User struct{}

func newUser() *engine.Route {
	var u interface{} = User{}
	return &engine.Route{
		Logger:  true,
		Handler: u.(http.Handler),
	}
}

func (u User) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var head string
	head, r.URL.Path = engine.ShiftPath(r.URL.Path)
	if head == "profile" {
		u.Profile(w, r)
		return
	} else if head == "detail" {
		head, _ := engine.ShiftPath(r.URL.Path)
		i, err := strconv.Atoi(head)
		i64 := int64(i)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		ctx := r.Context()
		// ctx = context.WithValue(ctx, engine.ContextTestKey, "pulisic")
		ctx = context.WithValue(ctx, engine.ContextUserID, i64)
		u.detail(w, r.WithContext(ctx))
		return
	}
	engine.Respond(w, r, http.StatusNotFound, "user not found")
}

func (u User) detail(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	// v := ctx.Value(engine.ContextTestKey)
	id := ctx.Value(engine.ContextUserID).(int64)
	db := ctx.Value(engine.ContextDatabase).(*data.DB)

	var result = new(struct {
		ID    int64     `json:"userId"`
		Email string    `json:"email"`
		Time  time.Time `json:"time"`
	})

	user, err := db.Users.GetDetail(id)
	if err != nil {
		engine.Respond(w, r, http.StatusInternalServerError, err)
		return
	}

	result.ID = user.ID
	result.Email = user.Email
	result.Time = time.Now()

	engine.Respond(w, r, http.StatusOK, result)
}

func (u User) Profile(w http.ResponseWriter, r *http.Request) {
	engine.Respond(w, r, http.StatusOK, "profile was called")
}
