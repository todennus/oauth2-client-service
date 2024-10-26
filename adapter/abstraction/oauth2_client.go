package abstraction

import (
	"context"

	"github.com/todennus/oauth2-client-service/usecase/dto"
)

type OAuth2ClientUsecase interface {
	GetByID(ctx context.Context, req *dto.OAuth2ClientGetByIDRequest) (*dto.OAuth2ClientGetByIDResponse, error)
	Create(ctx context.Context, req *dto.OAuth2ClientCreateRequest) (*dto.OAuth2ClientCreateResponse, error)
	CreateFirst(ctx context.Context, req *dto.OAuth2ClientCreateFirstRequest) (*dto.OAuth2ClientCreateFirstResponse, error)
	Validate(ctx context.Context, req *dto.OAuth2ClientValidateRequest) (*dto.OAuth2ClientValidateResponse, error)
}
