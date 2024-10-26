package grpc

import (
	"github.com/todennus/oauth2-client-service/domain"
	"github.com/todennus/proto/gen/service/dto/resource"
	"github.com/todennus/shared/enumdef"
	"github.com/todennus/x/enum"
	"github.com/xybor-x/snowflake"
)

func NewUser(user *resource.User) *domain.User {
	return &domain.User{
		ID:   snowflake.ID(user.Id),
		Role: enum.FromStr[enumdef.UserRole](user.Role),
	}
}
