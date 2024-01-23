package routes

import (
	"sample/pkg/api/handler"
	"sample/pkg/api/middleware"

	"github.com/gin-gonic/gin"
)

func AdminRoutes(engin *gin.RouterGroup, admin *handler.AdminHandler) {

	engin.Use(middleware.ClearCache)

	engin.GET("/", admin.HandlerAdminPage)

	engin.GET("/login", admin.AdminLogin)
	engin.POST("/login", admin.HandlerPostAdminLogin)

	engin.GET("/createUser", admin.CreateUser)
	engin.POST("/createUser", admin.CreateUserPost)

	engin.POST("/userDelete", admin.UserDelete)

	engin.POST("/update", admin.SingleUserData)
	engin.POST("/updateUser", admin.EditUser)

	engin.GET("/logout", admin.AdminLogout)

}
