package repository

import (
	"errors"
	"fmt"
	interfaces "sample/pkg/repository/interface"
	"sample/pkg/utils/models"

	// "github.com/gin-contrib/sessions/gorm"
	// 	"github.com/gin-contrib/sessions/gorm"
	"gorm.io/gorm"
)

type AdminDatabase struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) interfaces.AdminRepository {
	return &AdminDatabase{db}
}

//ADMIN LOGIN

func (c *AdminDatabase) GetAdminData(LoginData models.AdminLoginDetails) (models.AdminLoginDetails, error) {
	var AdminFeatchDetails models.AdminLoginDetails
	query := `SELECT name,password FROM admins WHERE name=$1`

	row := c.db.Raw(query, LoginData.Name).Row()
	err := row.Scan(&AdminFeatchDetails.Name, &AdminFeatchDetails.Password)
	if err != nil {
		fmt.Println(err, "Error at featching data from database `GetUserData`")
	}
	if AdminFeatchDetails.Name == "" {
		return AdminFeatchDetails, errors.New("no admin")
	}
	return AdminFeatchDetails, nil
}

// CREATE USER
func (c *AdminDatabase) SaveuserData(userData models.UserDetails) error {
	var name string
	query1 := "SELECT name FROM users WHERE email=&1"

	row := c.db.Raw(query1, userData.Email).Row()
	err := row.Scan(&name)

	if err != nil {
		fmt.Println(err, "Error inserting of data to database `SaveUserData`")
	}

	if name != "" {
		fmt.Println("Already Account exists from email ")
		return errors.New("Email contain a account")
	} else {
		query := `INSERT INTO Users(name,email,phone,password) VALUES($1,$2,$3,$4)`
		result := c.db.Exec(query, userData.Name, userData.Email, userData.Phone, userData.Password)
		if result != nil {
			fmt.Println(result, "Error at inserting of data to database `SaveUserData`")
		}
	}
	return nil

}

func (c *AdminDatabase) AllUserData() *[]models.UserData {

	var user []models.UserData

	query := "SELECT name, email, phone FROM users"
	rows, err := c.db.Raw(query).Rows()
	if err != nil {
		fmt.Println(err, "error at fetchin user data")
	}

	for rows.Next() {
		var u models.UserData
		err := rows.Scan(&u.Name, &u.Email, &u.Phone)
		if err != nil {
			fmt.Println(err, "error at rows scan")
		}
		user = append(user, u)
	}
	return &user
}

///// delete user

func (c *AdminDatabase) UserDelete(UserMail models.UserDelete) {
	query := "DELETE FROM users WHERE email=?"
	c.db.Raw(query, UserMail.Email).Row()
}

//get single user data

func (c *AdminDatabase) SingleUserData(UserMai models.UserMail) models.UserData {
	var userData models.UserData
	query := "SELECT name,email,phone FROM users WHERE email=?"
	row := c.db.Raw(query, UserMai.Email).Row()
	err := row.Scan(&userData.Name, &userData.Email, &userData.Phone)
	if err != nil {
		fmt.Println(err, "error at rows scan single user data")
	}
	fmt.Println(userData)
	return userData

}

//EDIT USER

func (c *AdminDatabase) UserEdit(EditUser models.UserData) {

	query := "UPDATE users SET name=?, phone=?,WHERE email=?"
	c.db.Exec(query, EditUser.Name, EditUser.Phone, EditUser.Email)
}
