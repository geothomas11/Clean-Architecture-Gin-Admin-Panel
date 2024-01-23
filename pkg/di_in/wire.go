//go:build wireinject
// +build wireinject

// you can write your dependancy hear ratherthan create wire_gen.go by using google wire. so you need to maintain only single file also get a good idea.
// In my case i genereate wire_gen.go then , i hard core all depedancy rather than user wire auto create.

package di_in

import (
	routes "sample/pkg/api"
	"sample/pkg/api/handler"
	"sample/pkg/config"
	"sample/pkg/db"
	"sample/pkg/repository"
	"sample/pkg/usecase"

	"github.com/google/wire"
)

func InitializeAPI(cfg *config.Config) (*routes.ServerHTTP, error) {

	wire.Build(db.ConnectDatabase, repository.NewUserDataBase, repository.NewAdminRepository, usecase.NewUserCase, usecase.NewAdminUsecase, handler.NewAdminHandler, handler.NewUserHandler, routes.NewServerHttp)

	return &routes.ServerHTTP{}, nil

}
