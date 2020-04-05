package user_service_impl

import (
	"encoding/json"
	"io/ioutil"

	"github.com/noukenolife/mygoproject/application/dto/authenticate_dto"
	"github.com/noukenolife/mygoproject/application/dto/user_dto"
	"golang.org/x/oauth2"
)

type GoogleRetrieveUserProfile struct {
	Config *oauth2.Config
}

func (self GoogleRetrieveUserProfile) Invoke(accessToken authenticate_dto.AccessToken) (userProfile user_dto.UserProfile, err error) {
	client := self.Config.Client(oauth2.NoContext, &oauth2.Token{
		AccessToken: accessToken.Value,
	})
	response, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
	if err != nil {
		return
	}

	defer response.Body.Close()
	data, _ := ioutil.ReadAll(response.Body)

	var profile struct {
		Email string `json:"email"`
	}
	json.Unmarshal(data, &profile)

	userProfile = user_dto.UserProfile{
		Email: user_dto.Email{
			Value: profile.Email,
		},
	}
	return
}
