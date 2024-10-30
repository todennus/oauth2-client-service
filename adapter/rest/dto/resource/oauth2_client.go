package resource

import (
	"github.com/todennus/oauth2-client-service/usecase/dto/resource"
	"github.com/todennus/x/conversion"
)

type OAuth2Client struct {
	ClientID       string  `json:"client_id,omitempty" example:"332974701238012989"`
	OwnerID        *string `json:"owner_id,omitempty" example:"330559330522759168"`
	Name           *string `json:"name,omitempty" example:"Example Client"`
	IsAdmin        *bool   `json:"is_admin,omitempty" example:"false"`
	IsConfidential *bool   `json:"is_confidential,omitempty" example:"false"`
}

func NewOAuth2Client(client *resource.OAuth2Client) *OAuth2Client {
	return &OAuth2Client{
		ClientID:       client.ClientID.String(),
		OwnerID:        conversion.MakeSnowflakePointerString(client.OwnerID),
		Name:           client.Name,
		IsAdmin:        client.IsAdmin,
		IsConfidential: client.IsConfidential,
	}
}
