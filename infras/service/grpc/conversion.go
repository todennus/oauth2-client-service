package grpc

import (
	"errors"

	"github.com/todennus/oauth2-client-service/domain"
	"github.com/todennus/proto/gen/service/dto/resource"
	"github.com/todennus/shared/enumdef"
	"github.com/xybor-x/enum"
	"github.com/xybor-x/snowflake"
)

func NewUser(user *resource.User) (*domain.User, error) {
	role, ok := enum.FromString[enumdef.UserRole](user.Role)
	if !ok {
		return nil, errors.New("invalid role")
	}

	return &domain.User{
		ID:   snowflake.ID(user.Id),
		Role: role,
	}, nil
}
