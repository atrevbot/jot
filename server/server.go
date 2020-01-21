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
}

func New(a string, r store.Repo) *Server {
	h := mux.NewRouter()
	s := Server{
		http.Server{Addr: ":8080", Handler: h},
		&r,
	}

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
		n := "index"
		t := "Home"

		// Handle 404 routes
		if r.URL.Path != "/" {
			w.WriteHeader(http.StatusNotFound)
			n = "404"
			t = "Uh oh"
		}

		s.getTemplate(n, nil).Execute(w, map[string]interface{}{
			"title": t,
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
func (s *Server) getTemplate(n string, fm template.FuncMap) *template.Template {
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
		fmt.Sprintf("templates/%s.html", n),
	)
	if err != nil {
		fmt.Printf("Unable to load template %s: \n", n, err)
	}

	return t
}
