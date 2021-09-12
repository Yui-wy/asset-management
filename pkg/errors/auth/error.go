package auth

import "github.com/go-kratos/kratos/v2/errors"

var (
	ErrAuthFail               = errors.New(401, "Authentication failed", "Missing token or token incorrect")
	ErrPowerFail              = errors.New(403, "Power failed", "Power is not enough")
	ErrAreaFail               = errors.New(403, "Area failed", "Please check areaId")
	ErrMissingJwtToken        = errors.Unauthorized("UNAUTHORIZED", "JWT token is missing")
	ErrMissingKeyFunc         = errors.Unauthorized("UNAUTHORIZED", "keyFunc is missing")
	ErrTokenInvalid           = errors.Unauthorized("UNAUTHORIZED", "Token is invalid")
	ErrTokenExpired           = errors.Unauthorized("UNAUTHORIZED", "JWT token has expired")
	ErrTokenParseFail         = errors.Unauthorized("UNAUTHORIZED", "Fail to parse JWT token ")
	ErrUnSupportSigningMethod = errors.Unauthorized("UNAUTHORIZED", "Wrong signing method")
	ErrWrongContext           = errors.Unauthorized("UNAUTHORIZED", "Wrong context for middelware")
	ErrNeedTokenProvider      = errors.Unauthorized("UNAUTHORIZED", "Token provider is missing")
	ErrSignToken              = errors.Unauthorized("UNAUTHORIZED", "Can not sign token.Is the key correct?")
	ErrGetKey                 = errors.Unauthorized("UNAUTHORIZED", "Can not get key while signing token")
)
