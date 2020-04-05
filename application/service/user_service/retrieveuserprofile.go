package user_service

import (
	"github.com/noukenolife/mygoproject/application/dto/authenticate_dto"
	"github.com/noukenolife/mygoproject/application/dto/user_dto"
)

type RetrieveUserProfile interface {
	Invoke(accessToken authenticate_dto.AccessToken) (user_dto.UserProfile, error)
}
