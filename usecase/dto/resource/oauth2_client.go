package resource

import (
	"context"

	"github.com/todennus/oauth2-client-service/domain"
	"github.com/todennus/shared/scopedef"
	"github.com/todennus/shared/xcontext"
	"github.com/xybor-x/snowflake"
)

type OAuth2Client struct {
	ClientID       snowflake.ID
	OwnerID        *snowflake.ID
	Name           *string
	IsAdmin        *bool
	IsConfidential *bool
}

func NewOAuth2ClientWithFilter(ctx context.Context, client *domain.OAuth2Client) *OAuth2Client {
	c := NewOAuth2ClientWithoutFilter(client)

	scopedef.Eval(xcontext.Scope(ctx)).
		RequireAdmin(scopedef.AdminReadClientProfile).
		RequireUser(ctx, scopedef.UserReadClientProfile, client.OwnerUserID).
		RequireApp(ctx, scopedef.AppReadClientOwner, client.ID).
		FilterIfUnsatisfied(&c.OwnerID)

	scopedef.Eval(xcontext.Scope(ctx)).
		RequireAdmin(scopedef.AdminReadClientProfile).
		RequireUser(ctx, scopedef.UserReadClientProfile, client.OwnerUserID).
		RequireApp(ctx, scopedef.AppReadClientProfile, client.ID).
		FilterIfUnsatisfied(&c.Name, &c.IsAdmin, &c.IsConfidential)

	return c
}

func NewOAuth2ClientWithoutFilter(client *domain.OAuth2Client) *OAuth2Client {
	usecaseClient := &OAuth2Client{
		ClientID:       client.ID,
		OwnerID:        &client.OwnerUserID,
		Name:           &client.Name,
		IsAdmin:        &client.IsAdmin,
		IsConfidential: &client.IsConfidential,
	}

	return usecaseClient
}
