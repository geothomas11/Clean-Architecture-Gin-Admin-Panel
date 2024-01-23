package routes

import (
	"sample/pkg/api/handler"
	"sample/pkg/api/routes"

	"github.com/gin-gonic/gin"
)

type ServerHTTP struct {
	engin *gin.Engine
}

func NewServerHttp(User *handler.UserHandler, admin *handler.AdminHandler) *ServerHTTP {

	engin := gin.Default()

	engin.LoadHTMLGlob("./templates/*.html")
	engin.Static("/static", "./static")

	engin.Use(gin.Logger())

	routes.UserRoutes(engin.Group("/user"), User)
	routes.AdminRoutes(engin.Group("/admin"), admin)

	return &ServerHTTP{engin: engin}

}

func (sh *ServerHTTP) Start() {
	sh.engin.Run(":8080")

}
