package abstraction

import (
	"github.com/todennus/oauth2-client-service/domain"
	"github.com/todennus/shared/enumdef"
	"github.com/xybor-x/snowflake"
)

type OAuth2ClientDomain interface {
	New(ownerID snowflake.ID, name string, isAdmin, isConfidential bool) (*domain.OAuth2Client, string, error)
	NewFirst(ownerID snowflake.ID, name string) (*domain.OAuth2Client, string, error)
	ValidateClient(
		client *domain.OAuth2Client,
		clientID snowflake.ID,
		clientSecret string,
		requirement enumdef.OAuth2ClientConfidentialRequirement,
	) error
}
