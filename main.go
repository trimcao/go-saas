package main

import (
	"log"
	"net/http"
	"time"
)

type Route struct {
	Logger  bool
	Tester  bool
	Handler http.Handler
}

type App struct {
	User *Route
}

func main() {

	app := &App{
		User: &Route{
			Logger: true,
			Tester: true,
		},
		Billing: &Route{
			Logger: true,
			Tester: false,
		},
	}

	http.HandleFunc("/", func(w http.ResponseWriter, h *http.Request) {
		w.Write([]byte("hello world"))
	})

	http.HandleFunc("/time", getTime)
	http.Handle("/iseven", isEven(http.HandlerFunc(getTime)))

	log.Println("web server started at localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("unable to start web server", err)
	}
}

func getTime(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "only GET method are allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte(time.Now().String()))
}

func isEven(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if time.Now().Second()%2 == 0 {
			h.ServeHTTP(w, r)
		} else {
			http.Error(w, "current time second is odd, cannot serve the response",
				http.StatusInternalServerError)
		}
	})
}

func (h *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var next *Route
	var head string
	head, r.URL.Path = shiftPath(r.URL.Path)
	if len(head) == 0 {
		next = &Route{
			Logger: true,
			Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte("home page"))
			}),
		}
	} else if head == "user" {
		var i interface{} = User{}
		next = &Route{
			Logger:  true,
			Tester:  true,
			Handler: i.(http.Handler),
		}
	} else {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	if next.Logger {
		next.Handler = h.log(next.Handler)
	}

	if next.Tester {
		next.Handler = h.test(next.Handler)
	}

	next.Handler.ServeHTTP(w, r)
}
