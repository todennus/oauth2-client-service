package abstraction

import (
	"github.com/todennus/oauth2-client-service/domain"
	"github.com/todennus/shared/enumdef"
	"github.com/todennus/x/scope"
	"github.com/xybor-x/snowflake"
)

type OAuth2ClientDomain interface {
	New(ownerID snowflake.ID, name string, isConfidential bool) (*domain.OAuth2Client, string, error)
	ValidateClient(
		client *domain.OAuth2Client,
		clientID snowflake.ID,
		clientSecret string,
		scope scope.Scopes,
		confidentialRequirement enumdef.ConfidentialRequirementType,
	) error
}
