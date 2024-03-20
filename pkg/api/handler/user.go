package handler

import (
	"fmt"
	"net/http"
	"sample/pkg/helper"
	interfaces "sample/pkg/usecase/interface"
	"sample/pkg/utils/models"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userUseCase interfaces.UserUseCase
}

func NewUserHandler(useCase interfaces.UserUseCase) *UserHandler {
	return &UserHandler{userUseCase: useCase}
}

// get signup

func (u *UserHandler) HandlerGetUserSignup(c *gin.Context) {
	_, TokenExist := helper.CheckCookie(c)
	if !TokenExist {
		c.HTML(http.StatusOK, "signup.html", nil)
	} else {
		c.Redirect(http.StatusFound, "/user/")
	}
}

// post singup

func (u *UserHandler) HandlerUserSignup(c *gin.Context) {
	_, TokenExist := helper.CheckCookie(c)
	if TokenExist {
		c.Redirect(http.StatusFound, "/user/")
	} else {
		err := c.Request.ParseForm()
		if err != nil {
			fmt.Println(err, "at userSignup")
		}

		name := c.Request.FormValue("name")
		email := c.Request.FormValue("email")
		phone := c.Request.FormValue("phone")
		password := c.Request.FormValue("password")
		confirmPassword := c.Request.FormValue("confirmpassword")

		SignupData := models.UserDetails{Name: name, Email: email, Phone: phone, Password: password, ConfirmPassword: confirmPassword}
		TokenData := models.GenerateToken{Email: email}

		fmt.Println(SignupData)
		IsMatch := u.userUseCase.UseUserSignup(SignupData)
		if IsMatch != nil {

			fmt.Println(err)
			c.HTML(http.StatusOK, "signup.html", gin.H{
				"Error": "Email already used",
			})

		}
		helper.SetToken(TokenData, c)
		c.Redirect(http.StatusFound, "/user/")
		return

	}
}

// Get Login
func (u *UserHandler) HandlerGetLogin(c *gin.Context) {

	_, TokenExist := helper.CheckCookie(c)
	if TokenExist {
		c.Redirect(http.StatusFound, "/user/")
	} else {
		c.HTML(http.StatusOK, "login.html", nil)
	}
}

// Post Login
func (u *UserHandler) HandlerPostLogin(c *gin.Context) {

	_, TokenExist := helper.CheckCookie(c)
	if TokenExist {
		c.Redirect(http.StatusFound, "/user/")
	} else {
		err := c.Request.ParseForm()
		if err != nil {
			fmt.Println(err, "at UserPostLogin")
		}

		email := c.Request.FormValue("email")
		password := c.Request.FormValue("password")
		LoginData := models.UserLoginDetails{Email: email, Password: password}

		TokenData := models.GenerateToken{Email: email}

		error := u.userUseCase.UseUserLogin(LoginData)
		if error != nil {
			c.HTML(http.StatusOK, "login.html", error)
		} else {
			helper.SetToken(TokenData, c)
			c.Redirect(http.StatusFound, "/user/")
		}
	}
}

// home page

func (u *UserHandler) HandlerGetHome(c *gin.Context) {
	email, TokenExist := helper.CheckCookie(c)
	if TokenExist {
		name := u.userUseCase.UseUserName(email)
		c.HTML(http.StatusOK, "index.html", name)
	} else {
		c.Redirect(http.StatusFound, "/user/login")
	}
}

//Logout

func (u *UserHandler) HandlerPostLogout(c *gin.Context) {
	helper.DeleteToken(c)
	c.Redirect(http.StatusFound, "/user/login")
}
