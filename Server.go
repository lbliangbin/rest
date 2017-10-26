package rest

import (
	"net/http"

	"github.com/emicklei/go-restful"
	"github.com/emicklei/go-restful-openapi"
)

// Server REST服务器
type Server struct {
	// Container
	*restful.Container

	// Swagger & Swagger UI
	SwaggerPath   string             // default: /apidocs.json
	SwaggerUIPath string             // default: /apidocs/
	SwaggerUIDir  string             // Swagger UI地址
	SwaggerConfig restfulspec.Config // swagger UI额外配置
}

// InstallSwaggerService 安装Swagger & SwaggerUI服务
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

	// 注册swagger描述文件
	s.Container.Add(restfulspec.NewOpenAPIService(s.SwaggerConfig))

	// 注册swagger ui
	s.Container.ServeMux.Handle(s.SwaggerUIPath, http.StripPrefix(s.SwaggerUIPath, http.FileServer(http.Dir(s.SwaggerUIDir))))
}

// NewServer 创建新的Server
func NewServer() *Server {
	return &Server{Container: restful.NewContainer()}
}
