package router

import (
	"context"
	"fmt"
	"net/http"

	"github.com/pkg/errors"

	httpSwagger "github.com/swaggo/http-swagger"
	svcframework "github.com/tbd54566975/ssi-service/pkg/service/framework"
	"github.com/tbd54566975/ssi-service/pkg/service/swagger"
)

type SwaggerRouter struct {
	service *swagger.Service
}

func NewSwaggerRouter(s svcframework.Service) (*SwaggerRouter, error) {
	if s == nil {
		return nil, errors.New("servce cannot be nil")
	}
	swaggerService, ok := s.(*swagger.Service)
	if !ok {
		return nil, fmt.Errorf("could not create schema router with service type: %s", s.Type())
	}
	return &SwaggerRouter{service: swaggerService}, nil
}

// SwaggerUI godoc
// @Summary      SwaggerUI
// @Description  Swagger UI
// @Tags         Documentation
// @Accept       html
// @Produce      html
// @Success      200  {string}  string  "OK"
// @Failure      400  {string}  string  "Bad request"
// @Failure      500  {string}  string  "Internal server error"
// @Router       /v1/docs [get]
func (sr SwaggerRouter) ServeSpec(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	http.ServeFile(w, r, "doc/swagger.yaml")
	return nil
}

// SwaggerUI godoc
// @Summary      SwaggerUI
// @Description  Swagger UI
// @Tags         Documentation
// @Accept       html
// @Produce      html
// @Success      200  {string}  string  "OK"
// @Failure      400  {string}  string  "Bad request"
// @Failure      500  {string}  string  "Internal server error"
// @Router       /v1/docs [get]
func (sr SwaggerRouter) ServeUI(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	httpSwagger.Handler(
		httpSwagger.URL("http://localhost:3000/v1/swagger/index.yaml"), //The url pointing to API definition
	).ServeHTTP(w, r)
	fmt.Fprintf(w, "hi")
	return nil
}
