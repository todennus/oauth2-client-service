package rest

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/todennus/oauth2-client-service/adapter/abstraction"
	"github.com/todennus/oauth2-client-service/adapter/rest/dto"
	"github.com/todennus/shared/errordef"
	"github.com/todennus/shared/middleware"
	"github.com/todennus/shared/response"
	"github.com/todennus/x/xhttp"
)

type OAuth2ClientAdapter struct {
	oauth2ClientUsecase abstraction.OAuth2ClientUsecase
}

func NewOAuth2ClientAdapter(oauth2ClientUsecase abstraction.OAuth2ClientUsecase) *OAuth2ClientAdapter {
	return &OAuth2ClientAdapter{
		oauth2ClientUsecase: oauth2ClientUsecase,
	}
}

func (a *OAuth2ClientAdapter) Router(r chi.Router) {
	r.Get("/{client_id}", middleware.RequireAuthentication(a.Get()))

	r.Post("/", middleware.RequireAuthentication(a.Create()))
}

// @Summary Get oauth2 client by id
// @Description Get OAuth2 Client information by ClientID.
// @Tags OAuth2 Client
// @Produce json
// @Param id path string true "ClientID"
// @Success 200 {object} response.SwaggerSuccessResponse[dto.OAuth2ClientGetByIDResponse] "Get client successfully"
// @Failure 400 {object} response.SwaggerBadRequestErrorResponse "Bad request"
// @Failure 404 {object} response.SwaggerNotFoundErrorResponse "Not found"
// @Router /oauth2_clients/{client_id} [get]
func (a *OAuth2ClientAdapter) Get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		req, err := xhttp.ParseHTTPRequest[dto.OAuth2ClientGetByIDRequest](r)
		if err != nil {
			response.RESTWriteAndLogInvalidRequestError(ctx, w, err)
			return
		}

		resp, err := a.oauth2ClientUsecase.GetByID(ctx, req.To())
		response.NewRESTResponseHandler(ctx, dto.NewOAuth2ClientGetByIDResponse(resp), err).
			Map(http.StatusBadRequest, errordef.ErrRequestInvalid).
			Map(http.StatusNotFound, errordef.ErrOAuth2ClientInvalid).
			WriteHTTPResponse(ctx, w)
	}
}

// @Summary Create oauth2 client
// @Description Create an new OAuth2 Client. <br>
// @Description If the `is_confidential` field is true, a secret is issued. Please carefully store this secret in a confidential place. This secret will never be retrieved by anyway. <br>
// @Description Require `todennus/create:client` or `todennus/admin:create:client` scope.
// @Tags OAuth2 Client
// @Security OAuth2Application[todennus/create:client]
// @Security OAuth2Application[todennus/admin:create:client]
// @Accept json
// @Produce json
// @Param body body dto.OAuth2ClientCreateRequest true "Client Information"
// @Success 201 {object} response.SwaggerSuccessResponse[dto.OAuth2ClientCreateResponse] "Create client successfully"
// @Failure 400 {object} response.SwaggerBadRequestErrorResponse "Bad request"
// @Failure 403 {object} response.SwaggerForbiddenErrorResponse "Forbidden"
// @Router /oauth2_clients [post]
func (a *OAuth2ClientAdapter) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		req, err := xhttp.ParseHTTPRequest[dto.OAuth2ClientCreateRequest](r)
		if err != nil {
			response.RESTWriteAndLogInvalidRequestError(ctx, w, err)
			return
		}

		resp, err := a.oauth2ClientUsecase.Create(ctx, req.To())
		response.NewRESTResponseHandler(ctx, dto.NewOauth2ClientCreateResponse(resp), err).
			Map(http.StatusBadRequest, errordef.ErrRequestInvalid).
			Map(http.StatusForbidden, errordef.ErrForbidden).
			WithDefaultCode(http.StatusCreated).
			WriteHTTPResponse(ctx, w)
	}
}
