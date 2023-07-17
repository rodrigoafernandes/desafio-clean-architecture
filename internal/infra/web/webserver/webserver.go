package webserver

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type WebServer struct {
	Router        chi.Router
	Handlers      []routeHandler
	WebServerPort string
}

func NewWebServer(serverPort string) *WebServer {
	return &WebServer{
		Router:        chi.NewRouter(),
		WebServerPort: serverPort,
	}
}

func (s *WebServer) AddHandler(method, path string, handler http.HandlerFunc) {
	s.Handlers = append(s.Handlers, routeHandler{method: method, path: path, handlerFunc: handler})
}

// loop through the handlers and add them to the router
// register middeleware logger
// start the server
func (s *WebServer) Start() {
	s.Router.Use(middleware.Logger)
	for _, h := range s.Handlers {
		s.Router.Method(h.method, h.path, h.handlerFunc)
	}
	http.ListenAndServe(s.WebServerPort, s.Router)
}

type routeHandler struct {
	method      string
	path        string
	handlerFunc http.HandlerFunc
}
