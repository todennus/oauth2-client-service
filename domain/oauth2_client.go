package domain

import (
	"errors"
	"fmt"
	"time"

	"github.com/todennus/shared/enumdef"
	"github.com/todennus/x/xcrypto"
	"github.com/todennus/x/xstring"
	"github.com/xybor-x/snowflake"
)

const (
	MaximumClientNameLength int = 64
	MinimumClientNameLength int = 3
)

type OAuth2Client struct {
	ID             snowflake.ID
	OwnerUserID    snowflake.ID
	IsAdmin        bool
	Name           string
	HashedSecret   string
	IsConfidential bool
	UpdatedAt      time.Time
}

type OAuth2ClientDomain struct {
	Snowflake          *snowflake.Node
	ClientSecretLength int
}

func NewOAuth2ClientDomain(
	snowflake *snowflake.Node,
	clientSecretLength int,
) (*OAuth2ClientDomain, error) {
	return &OAuth2ClientDomain{Snowflake: snowflake, ClientSecretLength: clientSecretLength}, nil
}

func (domain *OAuth2ClientDomain) New(ownerID snowflake.ID, name string, isAdmin, isConfidential bool) (*OAuth2Client, string, error) {
	err := domain.validateClientName(name)
	if err != nil {
		return nil, "", err
	}

	secret := ""
	hashedSecret := []byte{}
	if isConfidential {
		secret = xcrypto.RandString(domain.ClientSecretLength)
		if hashedSecret, err = HashPassword(secret); err != nil {
			return nil, "", err
		}
	}

	return &OAuth2Client{
		ID:             domain.Snowflake.Generate(),
		Name:           name,
		IsAdmin:        isAdmin,
		OwnerUserID:    ownerID,
		IsConfidential: isConfidential,
		HashedSecret:   string(hashedSecret),
	}, secret, nil
}

func (domain *OAuth2ClientDomain) NewFirst(ownerID snowflake.ID, name string) (*OAuth2Client, string, error) {
	client, secret, err := domain.New(ownerID, name, true, true)
	if err != nil {
		return nil, "", err
	}

	return client, secret, nil
}

func (domain *OAuth2ClientDomain) ValidateClient(
	client *OAuth2Client,
	clientID snowflake.ID,
	clientSecret string,
	confidentialRequirement enumdef.OAuth2ClientConfidentialRequirement,
) error {
	if client.ID != clientID {
		return errors.New("mismatched client id")
	}

	switch confidentialRequirement {
	case enumdef.CRTRequire:
		if !client.IsConfidential {
			return fmt.Errorf("%w: require a confidential client", ErrClientInvalid)
		}

		if err := ValidatePassword(client.HashedSecret, clientSecret); err != nil {
			return err
		}

	case enumdef.CRTDependOnType:
		if client.IsConfidential {
			if err := ValidatePassword(client.HashedSecret, clientSecret); err != nil {
				return err
			}
		}
	}

	return nil
}

func (domain *OAuth2ClientDomain) validateClientName(clientName string) error {
	if len(clientName) > MaximumClientNameLength {
		return fmt.Errorf("%w: require at most %d characters", ErrClientNameInvalid, MaximumClientNameLength)
	}

	if len(clientName) < MinimumClientNameLength {
		return fmt.Errorf("%w: require at least %d characters", ErrClientNameInvalid, MinimumClientNameLength)
	}

	for _, c := range clientName {
		if !xstring.IsNumber(c) && !xstring.IsLetter(c) && !xstring.IsUnderscore(c) && !xstring.IsSpace(c) {
			return fmt.Errorf("%w: got an invalid character %c", ErrClientNameInvalid, c)
		}
	}

	return nil
}
