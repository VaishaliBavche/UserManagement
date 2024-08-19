package configs

import "context"

var (
	AppConfig *ApplicationConfig
)

type ApplicationConfig struct {
}

func NewApplicationCnfig(context context.Context) error {
	return nil
}
