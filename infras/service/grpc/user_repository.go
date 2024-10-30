package grpc

import (
	"context"

	"github.com/todennus/oauth2-client-service/domain"
	"github.com/todennus/proto/gen/service"
	"github.com/todennus/proto/gen/service/dto"
	"github.com/todennus/shared/authentication"
	"github.com/todennus/shared/errordef"
	"google.golang.org/grpc"
)

type UserRepository struct {
	client service.UserClient
	auth   *authentication.GrpcAuthorization
}

func NewUserRepository(conn *grpc.ClientConn, authorization *authentication.GrpcAuthorization) *UserRepository {
	return &UserRepository{
		client: service.NewUserClient(conn),
		auth:   authorization,
	}
}

func (repo *UserRepository) Validate(ctx context.Context, username string, password string) (*domain.User, error) {
	req := &dto.UserValidateRequest{Username: username, Password: password}
	resp, err := repo.client.Validate(repo.auth.Context(ctx), req)
	if err != nil {
		return nil, errordef.ConvertGRPCError(err)
	}

	return NewUser(resp.User), nil
}
