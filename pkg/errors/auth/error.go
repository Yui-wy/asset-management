package auth

import "github.com/go-kratos/kratos/v2/errors"

var ErrAuthFail = errors.New(401, "Authentication failed", "Missing token or token incorrect")
var ErrPowerFail = errors.New(403, "Power failed", "Power is not enough")
var ErrAreaFail = errors.New(403, "Area failed", "Please check areaId")
