package wiring

import (
	"context"

	"github.com/todennus/oauth2-client-service/infras/database/gorm"
	"github.com/todennus/oauth2-client-service/infras/service/grpc"
	"github.com/todennus/oauth2-client-service/usecase/abstraction"
	"github.com/todennus/shared/config"
)

type Repositories struct {
	abstraction.UserRepository
	abstraction.OAuth2ClientRepository
}

func InitializeRepositories(ctx context.Context, config *config.Config, infras *Infras) (*Repositories, error) {
	r := &Repositories{}

	r.UserRepository = grpc.NewUserRepository(infras.UsergRPCConn)
	r.OAuth2ClientRepository = gorm.NewOAuth2ClientRepository(infras.GormPostgres)

	return r, nil
}
