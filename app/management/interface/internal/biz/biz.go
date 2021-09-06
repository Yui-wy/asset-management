package biz

import "github.com/google/wire"

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(NewUserUseCase, NewAuthUseCase)

type UserDetails struct {
	Id     uint64
	Power  int32
	AreaId []uint32
}
