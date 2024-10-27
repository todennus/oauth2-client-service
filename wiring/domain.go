package wiring

import (
	"context"

	"github.com/todennus/oauth2-client-service/domain"
	"github.com/todennus/oauth2-client-service/usecase/abstraction"
	"github.com/todennus/shared/config"
)

type Domains struct {
	abstraction.OAuth2ClientDomain
}

func InitializeDomains(ctx context.Context, config *config.Config) (*Domains, error) {
	var err error
	domains := &Domains{}

	domains.OAuth2ClientDomain, err = domain.NewOAuth2ClientDomain(
		config.SnowflakeNode,
		config.Variable.OAuth2Client.SecretLength,
	)
	if err != nil {
		return nil, err
	}

	return domains, nil
}
