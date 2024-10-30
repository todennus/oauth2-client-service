package dto

import (
	"github.com/todennus/oauth2-client-service/adapter/rest/dto/resource"
	"github.com/todennus/oauth2-client-service/usecase/dto"
	"github.com/xybor-x/snowflake"
)

type OAuth2ClientCreateRequest struct {
	Name           string `json:"name" example:"Example Client"`
	IsAdmin        bool   `json:"is_admin" example:"false"`
	IsConfidential bool   `json:"is_confidential" example:"true"`
}

func (req OAuth2ClientCreateRequest) To() *dto.OAuth2ClientCreateRequest {
	return &dto.OAuth2ClientCreateRequest{
		Name:           req.Name,
		IsAdmin:        req.IsAdmin,
		IsConfidential: req.IsConfidential,
	}
}

type OAuth2ClientCreateResponse struct {
	*resource.OAuth2Client
	ClientSecret string `json:"client_secret,omitempty" example:"ElBacv..."`
}

func NewOauth2ClientCreateResponse(resp *dto.OAuth2ClientCreateResponse) *OAuth2ClientCreateResponse {
	if resp == nil {
		return nil
	}

	return &OAuth2ClientCreateResponse{
		OAuth2Client: resource.NewOAuth2Client(resp.Client),
		ClientSecret: resp.ClientSecret,
	}
}

type OAuth2ClientGetByIDRequest struct {
	ClientID string `param:"client_id"`
}

func (req *OAuth2ClientGetByIDRequest) To() *dto.OAuth2ClientGetByIDRequest {
	clientID, err := snowflake.ParseString(req.ClientID)
	if err != nil {
		clientID = 0
	}

	return &dto.OAuth2ClientGetByIDRequest{
		ClientID: clientID,
	}
}

type OAuth2ClientGetByIDResponse struct {
	*resource.OAuth2Client
}

func NewOAuth2ClientGetByIDResponse(resp *dto.OAuth2ClientGetByIDResponse) *OAuth2ClientGetByIDResponse {
	if resp == nil {
		return nil
	}

	return &OAuth2ClientGetByIDResponse{
		OAuth2Client: resource.NewOAuth2Client(resp.Client),
	}
}
