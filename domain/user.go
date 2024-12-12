package domain

import (
	"github.com/todennus/shared/enumdef"
	"github.com/xybor-x/snowflake"
)

type User struct {
	ID   snowflake.ID
	Role enumdef.UserRole
}
