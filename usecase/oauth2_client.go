package usecase

import (
	"context"
	"errors"

	"github.com/todennus/oauth2-client-service/domain"
	"github.com/todennus/oauth2-client-service/usecase/abstraction"
	"github.com/todennus/oauth2-client-service/usecase/dto"
	"github.com/todennus/shared/enumdef"
	"github.com/todennus/shared/errordef"
	"github.com/todennus/shared/scopedef"
	"github.com/todennus/shared/xcontext"
	"github.com/todennus/x/lock"
	"github.com/todennus/x/xerror"
)

type OAuth2ClientUsecase struct {
	isNoClient      bool
	firstClientLock lock.Locker

	oauth2ClientDomain abstraction.OAuth2ClientDomain

	userRepo         abstraction.UserRepository
	oauth2ClientRepo abstraction.OAuth2ClientRepository
}

func NewOAuth2ClientUsecase(
	locker lock.Locker,
	oauth2ClientDomain abstraction.OAuth2ClientDomain,
	userRepo abstraction.UserRepository,
	oauth2ClientRepo abstraction.OAuth2ClientRepository,
) *OAuth2ClientUsecase {
	return &OAuth2ClientUsecase{
		isNoClient:         true,
		firstClientLock:    locker,
		oauth2ClientDomain: oauth2ClientDomain,
		userRepo:           userRepo,
		oauth2ClientRepo:   oauth2ClientRepo,
	}
}

func (usecase *OAuth2ClientUsecase) Create(
	ctx context.Context,
	req *dto.OAuth2ClientCreateRequest,
) (*dto.OAuth2ClientCreateResponse, error) {
	if xcontext.RequestSubjectType(ctx) != enumdef.SubjectUser {
		return nil, xerror.Enrich(errordef.ErrForbidden, "only allow creating client by user token")
	}

	scopeRequirement := scopedef.Eval(xcontext.Scope(ctx)).RequireAdmin(scopedef.AdminCreateClient)
	if !req.IsAdmin {
		scopeRequirement = scopeRequirement.RequireAnyUser(scopedef.UserCreateClient)
	}

	if scopeRequirement.IsUnsatisfied() {
		return nil, xerror.Enrich(errordef.ErrForbidden, "insufficient scope")
	}

	userID := xcontext.RequestSubjectID(ctx)
	client, secret, err := usecase.oauth2ClientDomain.New(userID, req.Name, req.IsAdmin, req.IsConfidential)
	if err != nil {
		return nil, errordef.DomainWrapper.Event(err, "failed-to-new-client").
			Enrich(errordef.ErrRequestInvalid).Error()
	}

	if err = usecase.oauth2ClientRepo.Create(ctx, client); err != nil {
		return nil, errordef.ErrServer.Hide(err, "failed-to-create-client")
	}

	return dto.NewOAuth2ClientCreateResponse(client, secret), nil
}

func (usecase *OAuth2ClientUsecase) CreateFirst(
	ctx context.Context,
	req *dto.OAuth2ClientCreateFirstRequest,
) (*dto.OAuth2ClientCreateFirstResponse, error) {
	if !usecase.isNoClient {
		return nil, xerror.Enrich(errordef.ErrNotFound, "this api is only openned for creating the first client")
	}

	if err := usecase.firstClientLock.Lock(ctx); err != nil {
		return nil, errordef.ErrServer.Hide(err, "failed-to-lock-first-client-flow")
	}
	defer usecase.firstClientLock.Unlock(ctx)

	count, err := usecase.oauth2ClientRepo.Count(ctx)
	if err != nil {
		return nil, errordef.ErrServer.Hide(err, "failed-to-count-client")
	}

	if count > 0 {
		usecase.isNoClient = false
		return nil, xerror.Enrich(errordef.ErrNotFound, "this api is only openned for creating the first client")
	}

	client, secret, err := usecase.oauth2ClientDomain.NewFirst(req.UserID, req.ClientName)
	if err != nil {
		return nil, errordef.DomainWrapper.Event(err, "failed-to-new-client").
			Enrich(errordef.ErrRequestInvalid).Error()
	}

	if err = usecase.oauth2ClientRepo.Create(ctx, client); err != nil {
		return nil, errordef.ErrServer.Hide(err, "failed-to-create-first-client")
	}

	usecase.isNoClient = false
	return dto.NewOAuth2ClientCreateFirstResponse(client, secret), nil
}

func (usecase *OAuth2ClientUsecase) GetByID(
	ctx context.Context,
	req *dto.OAuth2ClientGetByIDRequest,
) (*dto.OAuth2ClientGetByIDResponse, error) {
	client, err := usecase.oauth2ClientRepo.GetByID(ctx, req.ClientID)
	if err != nil {
		if errors.Is(err, errordef.ErrNotFound) {
			return nil, xerror.Enrich(errordef.ErrNotFound, "not found client")
		}

		return nil, errordef.ErrServer.Hide(err, "failed-to-get-client", "cid", req.ClientID)
	}

	return dto.NewOAuth2ClientGetResponse(ctx, client), nil
}

func (usecase *OAuth2ClientUsecase) Validate(
	ctx context.Context,
	req *dto.OAuth2ClientValidateRequest,
) (*dto.OAuth2ClientValidateResponse, error) {
	if scopedef.Eval(xcontext.Scope(ctx)).RequireAdmin(scopedef.AdminValidateClient).IsUnsatisfied() {
		return nil, xerror.Enrich(errordef.ErrForbidden, "insufficient scope")
	}

	if req.ClientID == 0 {
		return nil, xerror.Enrich(errordef.ErrRequestInvalid, "require client id")
	}

	client, err := usecase.oauth2ClientRepo.GetByID(ctx, req.ClientID)
	if err != nil {
		if errors.Is(err, errordef.ErrNotFound) {
			return nil, xerror.Enrich(errordef.ErrNotFound, "not found client")
		}

		return nil, errordef.ErrServer.Hide(err, "failed-to-get-client", "cid", req.ClientID)
	}

	err = usecase.oauth2ClientDomain.ValidateClient(
		client, req.ClientID, req.ClientSecret, req.ConfidentialRequirement)
	if err != nil {
		return nil, errordef.DomainWrapper.Event(err, "failed-to-validate-client").
			Enrich(errordef.ErrCredentialsInvalid).If(domain.ErrMismatchedPassword).
			Enrich(errordef.ErrOAuth2ClientInvalid).If(domain.ErrClientInvalid).
			Error()
	}

	return dto.NewOAuth2ClientValidateResponse(ctx, client), nil
}
