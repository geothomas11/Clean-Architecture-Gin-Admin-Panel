package usecase

import (
	"errors"
	"fmt"

	interfaces "sample/pkg/repository/interface"
	interfacesUseCase "sample/pkg/usecase/interface"
	"sample/pkg/utils/models"

	"golang.org/x/crypto/bcrypt"
)

type userUseCase struct {
	userRepo interfaces.UserRepository
}

func NewUserCase(repo interfaces.UserRepository) interfacesUseCase.UserUseCase {
	return &userUseCase{userRepo: repo}
}

// USER SIGNUP
func (c *userUseCase) UseUserSignup(userData models.UserDetails) error {
	if userData.ConfirmPassword == userData.Password {

		HashedPassword, err := bcrypt.GenerateFromPassword([]byte(userData.Password), bcrypt.DefaultCost)
		if err != nil {
			fmt.Println(err, "Problem at hashing Signup")

		}

		userData.Password = string(HashedPassword)
		exist := c.userRepo.SaveuserData(userData)
		if exist != nil {
			fmt.Println(exist, "at in usecase Exist")
			return exist
		}
	} else {
		return errors.New("Confirm password is not match")
	}
	return nil
}

//USER LOGIN

func (c *userUseCase) UseUserLogin(LoginData models.UserLoginDetails) error {
	LoginFeatchData, err := c.userRepo.GetUserData(LoginData)
	if err != nil {
		return errors.New("No user Exist")
	} else {
		err := bcrypt.CompareHashAndPassword([]byte(LoginFeatchData.Password), []byte(LoginData.Password))
		if err != nil {
			return errors.New("Password is not match")
		} else {
			return nil
		}
	}
}

// USER HOME
func (c *userUseCase) UseUserName(UserId string) string {
	name := c.userRepo.RepoGetUserName(UserId)
	return name

}
