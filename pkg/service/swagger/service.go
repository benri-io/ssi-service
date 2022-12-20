package swagger

import (
	"fmt"

	sdkutil "github.com/TBD54566975/ssi-sdk/util"
	"github.com/pkg/errors"

	"github.com/tbd54566975/ssi-service/config"
	"github.com/tbd54566975/ssi-service/pkg/service/framework"
)

type Service struct {
	config config.SwaggerServiceConfig
}

func (s Service) Type() framework.Type {
	return framework.Swagger
}

func (s Service) Status() framework.Status {
	ae := sdkutil.NewAppendError()
	if !ae.IsEmpty() {
		return framework.Status{
			Status:  framework.StatusNotReady,
			Message: fmt.Sprintf("schema service is not ready: %s", ae.Error().Error()),
		}
	}
	return framework.Status{Status: framework.StatusReady}
}

func (s Service) Config() config.SwaggerServiceConfig {
	return s.config
}

func NewSwaggerService(config config.SwaggerServiceConfig) (*Service, error) {
	service := Service{
		config: config,
	}
	if !service.Status().IsReady() {
		return nil, errors.New(service.Status().Message)
	}
	return &service, nil
}
