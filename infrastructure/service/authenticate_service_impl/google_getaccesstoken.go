package authenticate_service_impl

import (
	"github.com/noukenolife/mygoproject/application/dto/authenticate_dto"
	"golang.org/x/oauth2"
)

type GoogleGetAccessToken struct {
	Config *oauth2.Config
}

func (self GoogleGetAccessToken) Invoke(code string) (accessToken authenticate_dto.AccessToken, err error) {
	token, err := self.Config.Exchange(oauth2.NoContext, code)
	if err != nil {
		return
	}

	accessToken = authenticate_dto.AccessToken{
		Value: token.AccessToken,
	}
	return
}
