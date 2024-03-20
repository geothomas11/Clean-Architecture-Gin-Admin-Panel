package routes

import (
	"sample/pkg/api/handler"
	"sample/pkg/api/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(engin *gin.RouterGroup, user *handler.UserHandler) {

	engin.Use(middleware.ClearCache)

	engin.GET("/", user.HandlerGetHome)

	engin.GET("/signup", user.HandlerGetUserSignup)
	engin.POST("/signup", user.HandlerUserSignup)

	engin.GET("/login", user.HandlerGetLogin)
	engin.POST("/login", user.HandlerPostLogin)

	engin.GET("/logout", user.HandlerPostLogout)

}
