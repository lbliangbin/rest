package rest

import (
	"net/http"

	"github.com/emicklei/go-restful"
	"github.com/emicklei/go-restful-openapi"
)

// REST Server
type Server struct {
	// Container
	*restful.Container

	// Swagger & Swagger UI
	SwaggerPath   string             // default: /apidocs.json
	SwaggerUIPath string             // default: /apidocs/
	SwaggerUIDir  string             // Swagger UI Dir
	SwaggerConfig restfulspec.Config // swagger Config
}

// Install Swagger & Swagger UI Service
func (s *Server) InstallSwaggerService() {
	if s.Container == nil {
		return
	}

	if s.SwaggerPath == "" {
		s.SwaggerPath = "/apidocs.json"
	}
	if s.SwaggerUIPath == "" {
		s.SwaggerUIPath = "/apidocs/"
	}

	s.SwaggerConfig.WebServices = s.Container.RegisteredWebServices()
	s.SwaggerConfig.APIPath = s.SwaggerPath

	// add swagger service
	s.Container.Add(restfulspec.NewOpenAPIService(s.SwaggerConfig))

	// add swagger ui service
	s.Container.ServeMux.Handle(s.SwaggerUIPath, http.StripPrefix(s.SwaggerUIPath, http.FileServer(http.Dir(s.SwaggerUIDir))))
}

// Create a REST server
func NewServer() *Server {
	return &Server{Container: restful.NewContainer()}
}
