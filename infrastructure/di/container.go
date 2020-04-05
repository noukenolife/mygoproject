package di

import (
	"github.com/noukenolife/mygoproject/application/api"
	"github.com/noukenolife/mygoproject/application/api/action/helloworld"
	"github.com/noukenolife/mygoproject/application/service/authenticate_service"
	"github.com/noukenolife/mygoproject/application/web"
	"github.com/noukenolife/mygoproject/application/web/action/google/oauth"
	"github.com/noukenolife/mygoproject/application/web/action/google/oauth/oauth_callback"
	"github.com/noukenolife/mygoproject/application/web/action/index"
	"github.com/noukenolife/mygoproject/application/web/action/signin"
	"github.com/noukenolife/mygoproject/application/web/action/signout"
	"github.com/noukenolife/mygoproject/helper"
	"github.com/noukenolife/mygoproject/infrastructure/service/authenticate_service_impl"
	"github.com/noukenolife/mygoproject/infrastructure/service/user_service_impl"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type Container struct {
	ApiRouter  api.ApiRouter
	WebRouter  web.WebRouter
	GetAuthUrl authenticate_service.GetAuthUrl
}

func CreateContainer() Container {
	GoogleClientID := helper.Getenv("GOOGLE_OAUTH_CLIENT_ID")
	ClientSecret := helper.Getenv("GOOGLE_OAUTH_CLIENT_SECRET")
	GoogleRedirectURL := helper.Getenv("GOOGLE_OAUTH_REDIRECT_URL")

	googleOAuthConfig := &oauth2.Config{
		ClientID:     GoogleClientID,
		ClientSecret: ClientSecret,
		RedirectURL:  GoogleRedirectURL,
		Scopes:       []string{"email", "profile"},
		Endpoint:     google.Endpoint,
	}

	return Container{
		ApiRouter: api.ApiRouter{
			GetHelloWorld:  helloworld.GetHelloWorld{},
			PostHelloWorld: helloworld.PostHelloWorld{},
		},
		WebRouter: web.WebRouter{
			GetIndex:   index.GetIndex{},
			GetSignin:  signin.GetSignin{},
			GetSignout: signout.GetSignout{},
			GetGoogleOAuth: oauth.GetGoogleOAuth{
				GetAuthURL: authenticate_service_impl.GoogleGetAuthURL{
					Config: googleOAuthConfig,
				},
			},
			GetGoogleOAuthCallback: oauth_callback.GetGoogleOAuthCallback{
				GetAccessToken: authenticate_service_impl.GoogleGetAccessToken{
					Config: googleOAuthConfig,
				},
				RetrieveUserProfile: user_service_impl.GoogleRetrieveUserProfile{
					Config: googleOAuthConfig,
				},
			},
		},
	}
}
