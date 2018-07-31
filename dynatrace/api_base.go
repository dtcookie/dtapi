package dynatrace

import (
	"github.com/dtcookie/dtapi/config"
	"github.com/dtcookie/dtapi/environment"
)

type service struct {
	client *environment.APIClient
}

type configService struct {
	client *config.APIClient
}
