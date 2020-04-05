package authenticate_service_impl

import (
	"github.com/noukenolife/mygoproject/application/dto/authenticate_dto"
	"golang.org/x/oauth2"
)

type GoogleGetAuthURL struct {
	Config *oauth2.Config
}

func (self GoogleGetAuthURL) Invoke(state string) authenticate_dto.AuthURL {
	authURL := self.Config.AuthCodeURL(state, oauth2.ApprovalForce)

	return authenticate_dto.AuthURL{
		Value: authURL,
	}
}
