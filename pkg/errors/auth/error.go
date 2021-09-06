package auth

import "github.com/go-kratos/kratos/v2/errors"

var ErrAuthFail = errors.New(401, "Authentication failed", "Missing token or token incorrect")
var ErrPowerFail = errors.New(401, "Power failed", "Power is not enough")
