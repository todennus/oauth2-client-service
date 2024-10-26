package abstraction

import (
	"context"

	"github.com/todennus/oauth2-client-service/domain"
)

type UserRepository interface {
	GetByID(ctx context.Context, userID int64) (*domain.User, error)
	Validate(ctx context.Context, username string, password string) (*domain.User, error)
}

type OAuth2ClientRepository interface {
	Create(ctx context.Context, client *domain.OAuth2Client) error
	GetByID(ctx context.Context, clientID int64) (*domain.OAuth2Client, error)
	Count(ctx context.Context) (int64, error)
}
