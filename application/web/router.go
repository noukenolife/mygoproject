package web

import (
	"github.com/gin-gonic/gin"
	"github.com/noukenolife/mygoproject/application/middleware/checkauth"
	"github.com/noukenolife/mygoproject/application/web/action/google/oauth"
	"github.com/noukenolife/mygoproject/application/web/action/google/oauth/oauth_callback"
	"github.com/noukenolife/mygoproject/application/web/action/index"
	"github.com/noukenolife/mygoproject/application/web/action/signin"
	"github.com/noukenolife/mygoproject/application/web/action/signout"
)

type WebRouter struct {
	GetIndex               index.GetIndex
	GetSignin              signin.GetSignin
	GetSignout             signout.GetSignout
	GetGoogleOAuth         oauth.GetGoogleOAuth
	GetGoogleOAuthCallback oauth_callback.GetGoogleOAuthCallback
}

func (self WebRouter) InitRouter(r *gin.Engine) {
	r.LoadHTMLGlob("application/web/template/*")

	r.GET("/signin", self.GetSignin.Invoke)
	r.GET("/signout", self.GetSignout.Invoke)
	r.GET("/google/oauth", self.GetGoogleOAuth.Invoke)
	r.GET("/google/oauth/callback", self.GetGoogleOAuthCallback.Invoke)

	authorized := r.Group("/", checkauth.CheckAuth())
	authorized.GET("/", self.GetIndex.Invoke)
}
