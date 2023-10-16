package controller

import (
	"math/rand"
	"suranect_api/database"
	"suranect_api/model"
	"suranect_api/utils"

	"github.com/gin-gonic/gin"
)

type registerForm struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

func Register(c *gin.Context) {
	db, _ := database.ConnectMysql()

	var request registerForm

	err := c.ShouldBindJSON(&request)

	if err != nil {
		c.JSON(200, gin.H{
			"Status":         "error",
			"Message":        "Please do not leave the form blank",
			"Message_server": err.Error(),
		})
		return
	}

	hashPassword, _ := utils.HashPassword(request.Password)

	user := db.Create(&model.User{Username: request.Username, Password: hashPassword, Email: request.Email})

	if user.Error != nil {
		c.JSON(200, gin.H{
			"Status":  "error",
			"Message": user.Error.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"Status":  "success",
		"Message": "Berhasil membuat akun",
	})
	return
}

type loginForm struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type authUserInfo struct {
	User_id uint
}

func Login(c *gin.Context) {
	db, _ := database.ConnectMysql()

	var request loginForm

	err := c.ShouldBindJSON(&request)

	if err != nil {
		c.JSON(200, gin.H{
			"Status":         "error",
			"Message":        "Please do not leave the form blank",
			"Message_server": err.Error(),
		})
		return
	}

	var user model.User

	db.Where("username = ?", request.Username).Find(&user)

	if user.ID == 0 || !utils.CheckPasswordHash(request.Password, user.Password) {
		c.JSON(200, gin.H{
			"Status":  "error",
			"Message": "Password atau Username Salah",
		})
		return
	}

	// Generate Token
	token, _ := utils.CreateToken("authToken", authUserInfo{User_id: user.ID})

	c.JSON(200, gin.H{
		"Status":  "euccess",
		"Message": "Berhasil Login",
		"Token":   token,
	})
	return
}

func SendVerifyEmail(c *gin.Context) {
	is_login := c.MustGet("is_login").(bool)

	if !is_login {
		return
	}

	is_verify := c.MustGet("is_verify").(bool)
	if is_verify {
		c.JSON(200, gin.H{
			"Status":  "error",
			"Message": "Account has been verified",
		})
		return
	}

	userinfo := c.MustGet("userinfo").(map[string]interface{})
	db, _ := database.ConnectMysql()

	var user = model.User{}
	user.ID = uint(userinfo["User_id"].(float64))
	db.First(&user)

	kode := rand.Intn(8999) + 1000

	err := utils.SendVerifyEmail(user.Email, kode)

	if err != nil {
		c.JSON(200, gin.H{
			"Status":  "error",
			"Message": err.Error(),
		})
		return
	}

	user.EmailCode = kode
	db.Save(&user)

	c.JSON(200, gin.H{
		"Status":  "success",
		"Message": "Berhasil Mengirim Verifikasi Email",
	})
	return
}

type formVerifyEmail struct {
	Code int `binding:"required" json:"code"`
}

func VerifyEmail(c *gin.Context) {
	is_login := c.MustGet("is_login").(bool)
	if !is_login {
		return
	}

	is_verify := c.MustGet("is_verify").(bool)
	if is_verify {
		c.JSON(200, gin.H{
			"Status":  "error",
			"Message": "Account has been verified",
		})
		return
	}

	var request formVerifyEmail
	err := c.ShouldBindJSON(&request)

	if err != nil {
		c.JSON(200, gin.H{
			"Status":         "error",
			"Message":        "Please do not leave the form blank",
			"Message_server": err.Error(),
		})
		return
	}

	userinfo := c.MustGet("userinfo").(map[string]interface{})
	db, _ := database.ConnectMysql()

	var user = model.User{}
	user.ID = uint(userinfo["User_id"].(float64))
	db.First(&user)

	if request.Code == user.EmailCode {
		c.JSON(200, gin.H{
			"Status":  "success",
			"Message": "Account Has Been Verified",
		})

		user.EmailVerify = true
		user.EmailCode = 0
		db.Save(&user)

		return
	} else {
		c.JSON(200, gin.H{
			"Status":  "error",
			"Message": "The code entered is incorrect",
		})
		return
	}
}
