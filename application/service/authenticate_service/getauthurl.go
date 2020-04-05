package authenticate_service

import (
	"github.com/noukenolife/mygoproject/application/dto/authenticate_dto"
)

type GetAuthUrl interface {
	Invoke(state string) authenticate_dto.AuthURL
}
