package middleware

import (
	"suranect_api/database"
	"suranect_api/model"

	"github.com/gin-gonic/gin"
)

func VerifiedEmail(c *gin.Context) {
	is_login := c.MustGet("is_login").(bool)

	if is_login {
		userinfo := c.MustGet("userinfo").(map[string]interface{})
		db, _ := database.ConnectMysql()

		var user = model.User{}
		user.ID = uint(userinfo["User_id"].(float64))
		db.First(&user)

		c.Set("is_verify", user.EmailVerify)
	} else {
		c.Set("is_verify", false)
	}
}
