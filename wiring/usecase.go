package wiring

import (
	"context"
	"time"

	"github.com/todennus/oauth2-client-service/adapter/abstraction"
	"github.com/todennus/oauth2-client-service/usecase"
	"github.com/todennus/shared/config"
	"github.com/todennus/x/lock"
)

type Usecases struct {
	abstraction.OAuth2ClientUsecase
}

func InitializeUsecases(
	ctx context.Context,
	config *config.Config,
	infras *Infras,
	domains *Domains,
	repositories *Repositories,
) (*Usecases, error) {
	uc := &Usecases{}

	uc.OAuth2ClientUsecase = usecase.NewOAuth2ClientUsecase(
		lock.NewRedisLock(infras.Redis, "client-lock", 10*time.Second),
		domains.OAuth2ClientDomain,
		repositories.UserRepository,
		repositories.OAuth2ClientRepository,
	)

	return uc, nil
}
