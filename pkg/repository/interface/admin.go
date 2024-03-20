package interfaces

import "sample/pkg/utils/models"

type AdminRepository interface {
	GetAdminData(models.AdminLoginDetails) (models.AdminLoginDetails, error)
	SaveuserData(models.UserDetails) error
	AllUserData() *[]models.UserData
	SingleUserData(models.UserMail) models.UserData
	UserEdit(models.UserData)
	UserDelete(models.UserDelete)
}
