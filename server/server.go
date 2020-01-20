package server

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/atrevbot/jot/store"
	"github.com/gorilla/mux"
	"github.com/rs/xid"
)

type Server struct {
	http.Server
	Repo *store.Repo
	Env  map[string]string
}

func New(a string, r store.Repo, e map[string]string) *Server {
	h := mux.NewRouter()
	s := Server{
		http.Server{Addr: ":8080", Handler: h},
		&r,
		e,
	}

	// Define router, static files, and middleware
	h.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	// Routes to handle requests from browsers and HTML forms
	h.HandleFunc("/", s.handleIndex()).Methods("GET")
	h.HandleFunc("/close", s.handleClose()).Methods("GET")

	return &s
}

/**
 * HTTP handler for get requests to view the homepage
 */
func (s *Server) handleIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := "index"
		title := "Home"

		// Handle 404 routes
		if r.URL.Path != "/" {
			w.WriteHeader(http.StatusNotFound)
			name = "404"
			title = "Uh oh"
		}

		s.getTemplate(name, nil).Execute(w, map[string]interface{}{
			"title": title,
			"env":   s.Env,
		})
	}
}

/**
 * HTTP handler for closing the server
 */
func (s *Server) handleClose() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.Close()
	}
}

/**
 * Helper function to load all required layouts and partials to load template
 */
func (s *Server) getTemplate(name string, fm template.FuncMap) *template.Template {
	funcMap := template.FuncMap{
		"now": func() int {
			return time.Now().Year()
		},
		"uniqueID": func() string {
			return xid.New().String()
		},
	}
	// Merge custom funcMap
	for k, v := range fm {
		funcMap[k] = v
	}

	t, err := template.New("main.html").Funcs(funcMap).ParseFiles(
		"templates/_layouts/main.html",
		"templates/_meta/data.html",
		"templates/_meta/favicons.html",
		fmt.Sprintf("templates/%s.html", name),
	)
	if err != nil {
		fmt.Printf("Unable to load template %s: \n", name, err)
	}

	return t
}
