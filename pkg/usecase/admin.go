package usecase

import (
	"errors"
	"fmt"

	interfaces "sample/pkg/repository/interface"
	interfaceUsecase "sample/pkg/usecase/interface"
	"sample/pkg/utils/models"

	"golang.org/x/crypto/bcrypt"
)

type AdminUseCase struct {
	adminRepository interfaces.AdminRepository
}

func NewAdminUsecase(repo interfaces.AdminRepository) interfaceUsecase.AdminUseCase {
	return &AdminUseCase{adminRepository: repo}

}

// ADMIN LOGIN
func (c *AdminUseCase) UseAdminLogin(LoginData models.AdminLoginDetails) error {

	LoginFeatchData, err := c.adminRepository.GetAdminData(LoginData)

	if err != nil {
		return errors.New("no Admin exist")
	} else {
		if LoginData.Password != LoginFeatchData.Password {
			return errors.New("Password is not matched")
		} else {
			return nil
		}
	}

}

//CREATE USER

func (c *AdminUseCase) CreateUser(userData models.UserDetails) error {

	if userData.ConfirmPassword == userData.Password {

		HashedPassword, err := bcrypt.GenerateFromPassword([]byte(userData.Password), bcrypt.DefaultCost)
		if err != nil {
			fmt.Println(err, "Problem at hashing signup", HashedPassword)
			return err

		}

	} else {
		return errors.New("Confirm password is not match")
	}
	return nil
}

// FULL USERDATA
func (c *AdminUseCase) FullUserData() *[]models.UserData {
	UserCollection := c.adminRepository.AllUserData()
	return UserCollection
}

// DELETE USER
func (c *AdminUseCase) DeleteUser(UserMail models.UserDelete) {
	c.adminRepository.UserDelete(UserMail)
}

// DELETE USER
func (c *AdminUseCase) SingleUserDelete(userMail models.UserMail) models.UserData {
	return c.adminRepository.SingleUserData(userMail)
}

// EDIT USER
func (c *AdminUseCase) EditUser(UserData models.UserData) {
	c.adminRepository.UserEdit(UserData)

}
