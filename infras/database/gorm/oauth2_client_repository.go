package gorm

import (
	"context"

	"github.com/todennus/oauth2-client-service/domain"
	"github.com/todennus/oauth2-client-service/infras/database/model"
	"github.com/todennus/shared/errordef"
	"github.com/todennus/shared/xcontext"
	"github.com/xybor-x/snowflake"
	"gorm.io/gorm"
)

type OAuth2ClientRepository struct {
	db *gorm.DB
}

func NewOAuth2ClientRepository(db *gorm.DB) *OAuth2ClientRepository {
	return &OAuth2ClientRepository{db: db}
}

func (repo *OAuth2ClientRepository) Create(ctx context.Context, client *domain.OAuth2Client) error {
	model := model.NewOAuth2Client(client)
	return errordef.ConvertGormError(xcontext.DB(ctx, repo.db).Create(&model).Error)
}

func (repo *OAuth2ClientRepository) GetByID(ctx context.Context, clientID snowflake.ID) (*domain.OAuth2Client, error) {
	model := model.OAuth2ClientModel{}
	if err := xcontext.DB(ctx, repo.db).Take(&model, "id=?", clientID).Error; err != nil {
		return nil, errordef.ConvertGormError(err)
	}

	return model.To(), nil
}

func (repo *OAuth2ClientRepository) Count(ctx context.Context) (int64, error) {
	var n int64
	err := xcontext.DB(ctx, repo.db).Model(&model.OAuth2ClientModel{}).Count(&n).Error
	return n, errordef.ConvertGormError(err)
}
