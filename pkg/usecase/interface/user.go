package interfaceUsecase

import "sample/pkg/utils/models"

type UserUseCase interface {
	UseUserSignup(models.UserDetails) error
	UseUserLogin(models.UserLoginDetails) error
	UseUserName(string) string
}
