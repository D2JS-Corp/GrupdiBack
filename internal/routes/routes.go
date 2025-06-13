package routes

import (
	"github.com/D2JS-Corp/GrupdiBack/internal/oauth"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/auth/google", oauth.GoogleLogin)
	r.GET("/auth/google/callback", oauth.GoogleCallBack)
}
