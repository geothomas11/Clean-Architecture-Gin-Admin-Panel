package repository

import (
	"errors"
	"fmt"
	interfaces "sample/pkg/repository/interface"
	"sample/pkg/utils/models"

	"gorm.io/gorm"
	// "github.com/gin-contrib/sessions/gorm"
)

type UserDataBase struct {
	DB *gorm.DB
}

func NewUserDataBase(db *gorm.DB) interfaces.UserRepository {
	return &UserDataBase{DB: db}

}
func (c *UserDataBase) SaveuserData(userData models.UserDetails) error {

	var name string

	query1 := "SELECT name FROM users WHERE email=$1"
	row := c.DB.Raw(query1, userData.Email).Row()
	err := row.Scan(&name)

	if err != nil {
		fmt.Println(err, "error at inserting data to databse `SaveUserData`")
	}
	if name != "" {
		fmt.Println(err, "already account exist from email")
		return errors.New("email have an account")
	} else {
		query := `INSERT INTO users (name,email,phone,password) VALUES($1, $2, $3, $4)`
		result := c.DB.Exec(query, userData.Name, userData.Email, userData.Phone, userData.Password)

		if result != nil {
			fmt.Println(result, "Error at inserting data to database`saveUserData`")
		}
	}
	return nil

}

//GET USER DATA

func (c *UserDataBase) GetUserData(LoginData models.UserLoginDetails) (models.UserFeatchData, error) {
	var UserFeatchDeatails models.UserFeatchData
	query := `SELECT email,password FROM users WHERE email$1`

	row := c.DB.Raw(query, LoginData.Email).Row()
	err := row.Scan(&UserFeatchDeatails.Email, &UserFeatchDeatails.Password)
	if err != nil {
		fmt.Println(err, "Error at featching data from database `GetUserData`")
	}
	if UserFeatchDeatails.Email == "" {
		return UserFeatchDeatails, errors.New("no user")
	}
	UserFeatchDeatails.Email = LoginData.Email
	return UserFeatchDeatails, nil
}
func (c *UserDataBase) RepoGetUserName(UserId string) string {
	var name string
	query := `SELECT name FROM users WHERE email=$1`
	row := c.DB.Raw(query, UserId).Row()
	err := row.Scan(&name)
	if err != nil {
		fmt.Println(err, "Error at featchiong data from databse `GetUserData`")
	}
	fmt.Println(name, "----------")
	return name

}
