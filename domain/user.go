package domain

import (
	"github.com/todennus/shared/enumdef"
	"github.com/todennus/x/enum"
	"github.com/xybor-x/snowflake"
)

type User struct {
	ID   snowflake.ID
	Role enum.Enum[enumdef.UserRole]
}
