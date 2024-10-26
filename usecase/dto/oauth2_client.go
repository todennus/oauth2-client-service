package dto

import (
	"context"

	"github.com/todennus/oauth2-client-service/domain"
	"github.com/todennus/oauth2-client-service/usecase/dto/resource"
	"github.com/todennus/shared/enumdef"
	"github.com/xybor-x/snowflake"
)

type OAuth2ClientCreateRequest struct {
	Name           string
	IsConfidential bool
}

type OAuth2ClientCreateResponse struct {
	Client       *resource.OAuth2Client
	ClientSecret string
}

func NewOAuth2ClientCreateResponse(client *domain.OAuth2Client, secret string) *OAuth2ClientCreateResponse {
	return &OAuth2ClientCreateResponse{
		Client:       resource.NewOAuth2ClientWithoutFilter(client),
		ClientSecret: secret,
	}
}

type OAuth2ClientCreateFirstRequest struct {
	Username string
	Password string

	Name string
}

type OAuth2ClientCreateFirstResponse struct {
	Client       *resource.OAuth2Client
	ClientSecret string
}

func NewOAuth2ClientCreateFirstResponse(ctx context.Context, client *domain.OAuth2Client, secret string) *OAuth2ClientCreateFirstResponse {
	return &OAuth2ClientCreateFirstResponse{
		Client:       resource.NewOAuth2ClientWithoutFilter(client),
		ClientSecret: secret,
	}
}

type OAuth2ClientGetByIDRequest struct {
	ClientID snowflake.ID
}

type OAuth2ClientGetByIDResponse struct {
	Client *resource.OAuth2Client
}

func NewOAuth2ClientGetResponse(ctx context.Context, client *domain.OAuth2Client) *OAuth2ClientGetByIDResponse {
	return &OAuth2ClientGetByIDResponse{
		Client: resource.NewOAuth2Client(ctx, client),
	}
}

type OAuth2ClientValidateRequest struct {
	ClientID                snowflake.ID
	ClientSecret            string
	Scope                   string
	ConfidentialRequirement enumdef.ConfidentialRequirementType
}

type OAuth2ClientValidateResponse struct{}

func NewOAuth2ClientValidateResponse() *OAuth2ClientValidateResponse {
	return &OAuth2ClientValidateResponse{}
}
