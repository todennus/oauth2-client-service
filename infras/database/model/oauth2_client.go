package model

import (
	"time"

	"github.com/todennus/oauth2-client-service/domain"
	"github.com/xybor-x/snowflake"
)

type OAuth2ClientModel struct {
	ID             int64     `gorm:"id;primaryKey"`
	UserID         int64     `gorm:"user_id"`
	IsAdmin        bool      `gorm:"is_admin"`
	Name           string    `gorm:"name"`
	HashedSecret   string    `gorm:"hashed_secret"`
	IsConfidential bool      `gorm:"is_confidential"`
	UpdatedAt      time.Time `gorm:"updated_at"`
}

func (OAuth2ClientModel) TableName() string {
	return "oauth2_clients"
}

func NewOAuth2Client(domain *domain.OAuth2Client) *OAuth2ClientModel {
	return &OAuth2ClientModel{
		ID:             domain.ID.Int64(),
		UserID:         domain.OwnerUserID.Int64(),
		IsAdmin:        domain.IsAdmin,
		Name:           domain.Name,
		HashedSecret:   domain.HashedSecret,
		IsConfidential: domain.IsConfidential,
		UpdatedAt:      domain.UpdatedAt,
	}
}

func (client OAuth2ClientModel) To() *domain.OAuth2Client {
	return &domain.OAuth2Client{
		ID:             snowflake.ID(client.ID),
		OwnerUserID:    snowflake.ID(client.UserID),
		IsAdmin:        client.IsAdmin,
		Name:           client.Name,
		HashedSecret:   client.HashedSecret,
		IsConfidential: client.IsConfidential,
		UpdatedAt:      client.UpdatedAt,
	}
}
