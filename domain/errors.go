package domain

import (
	"fmt"

	"github.com/todennus/shared/errordef"
)

var (
	ErrMismatchedPassword = fmt.Errorf("%wmismatched password", errordef.ErrDomainKnown)
	ErrClientInvalid      = fmt.Errorf("%winvalid client", errordef.ErrDomainKnown)
	ErrClientNameInvalid  = fmt.Errorf("%winvalid client name", errordef.ErrDomainKnown)
)
