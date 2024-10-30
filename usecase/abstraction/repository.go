package abstraction

import (
	"context"

	"github.com/todennus/oauth2-client-service/domain"
	"github.com/xybor-x/snowflake"
)

type UserRepository interface {
	Validate(ctx context.Context, username string, password string) (*domain.User, error)
}

type OAuth2ClientRepository interface {
	Create(ctx context.Context, client *domain.OAuth2Client) error
	GetByID(ctx context.Context, clientID snowflake.ID) (*domain.OAuth2Client, error)
	Count(ctx context.Context) (int64, error)
}
