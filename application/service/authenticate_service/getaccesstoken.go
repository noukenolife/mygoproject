package authenticate_service

import "github.com/noukenolife/mygoproject/application/dto/authenticate_dto"

type GetAccessToken interface {
	Invoke(code string) (authenticate_dto.AccessToken, error)
}
