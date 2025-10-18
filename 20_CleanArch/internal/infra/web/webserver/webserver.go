package webserver

import (
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type WebServer struct {
	Router        chi.Router
	Handlers      map[string]http.HandlerFunc
	WebServerPort string
}

func NewWebServer(serverPort string) *WebServer {
	return &WebServer{
		Router:        chi.NewRouter(),
		Handlers:      make(map[string]http.HandlerFunc),
		WebServerPort: serverPort,
	}
}

func (s *WebServer) AddHandler(path string, handler http.HandlerFunc) {
	s.Handlers[path] = handler
}

func (s *WebServer) AddHandlerWithMethod(method, path string, handler http.HandlerFunc) {
	key := method + " " + path
	s.Handlers[key] = handler
}

// loop through the handlers and add them to the router
// register middeleware logger
// start the server
func (s *WebServer) Start() {
	s.Router.Use(middleware.Logger)
	for pathMethod, handler := range s.Handlers {
		// Se contém método HTTP (GET, POST, etc)
		if strings.Contains(pathMethod, " ") {
			parts := strings.SplitN(pathMethod, " ", 2)
			method := parts[0]
			path := parts[1]
			s.Router.Method(method, path, http.HandlerFunc(handler))
		} else {
			// Rota sem método específico (aceita todos)
			s.Router.HandleFunc(pathMethod, handler)
		}
	}
	// Adiciona : se não tiver
	port := s.WebServerPort
	if port[0] != ':' {
		port = ":" + port
	}
	if err := http.ListenAndServe(port, s.Router); err != nil {
		panic(err)
	}
}
