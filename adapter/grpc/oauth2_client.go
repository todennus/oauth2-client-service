package grpc

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/todennus/oauth2-client-service/adapter/abstraction"
	"github.com/todennus/oauth2-client-service/adapter/grpc/conversion"
	"github.com/todennus/proto/gen/service"
	pbdto "github.com/todennus/proto/gen/service/dto"
	"github.com/todennus/shared/errordef"
	"github.com/todennus/shared/response"
	"google.golang.org/grpc/codes"
)

var _ service.OAuth2ClientServer = (*OAuth2ClientServer)(nil)

type OAuth2ClientServer struct {
	service.UnimplementedOAuth2ClientServer

	oauth2ClientUsecase abstraction.OAuth2ClientUsecase
}

func NewOAuth2ClientServer(userUsecase abstraction.OAuth2ClientUsecase) *OAuth2ClientServer {
	return &OAuth2ClientServer{
		oauth2ClientUsecase: userUsecase,
	}
}

func (s *OAuth2ClientServer) GetByID(ctx context.Context, req *pbdto.OAuth2ClientGetByIDRequest) (*pbdto.OAuth2ClientGetByIDResponse, error) {
	ucreq := conversion.NewUsecaseOAuth2GetByIDRequest(req)
	resp, err := s.oauth2ClientUsecase.GetByID(ctx, ucreq)

	return response.NewGRPCResponseHandler(ctx, conversion.NewPbOAuth2GetByIDResponse(resp), err).
		Map(codes.InvalidArgument, errordef.ErrRequestInvalid).
		Map(codes.NotFound, errordef.ErrNotFound).Finalize(ctx)
}

func (s *OAuth2ClientServer) Validate(ctx context.Context, req *pbdto.OAuth2ClientValidateRequest) (*empty.Empty, error) {
	ucreq := conversion.NewUsecaseOAuth2ValidateRequest(req)
	_, err := s.oauth2ClientUsecase.Validate(ctx, ucreq)

	return response.NewGRPCResponseHandler(ctx, &empty.Empty{}, err).
		Map(codes.InvalidArgument, errordef.ErrRequestInvalid).
		Map(codes.PermissionDenied, errordef.ErrCredentialsInvalid).
		Map(codes.PermissionDenied, errordef.ErrOAuth2ScopeInvalid, errordef.ErrOAuth2ClientInvalid).
		Map(codes.NotFound, errordef.ErrNotFound).Finalize(ctx)
}
