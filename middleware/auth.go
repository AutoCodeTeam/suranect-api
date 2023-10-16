package middleware

import (
	"fmt"
	"strings"
	"suranect_api/utils"

	"github.com/gin-gonic/gin"
)

type headerJWT struct {
	Authorization string `binding:"required"`
}

func Auth(c *gin.Context) {
	var formjwt headerJWT
	errors := c.ShouldBindHeader(&formjwt)

	if errors != nil {
		c.JSON(401, gin.H{
			"Status":  "error",
			"Message": errors.Error(),
		})

		c.Set("is_login", false)
		return
	}

	stringToken := strings.Split(formjwt.Authorization, " ")[1]

	claims, err := utils.GetClaimsFromToken(stringToken)

	if err != nil {
		c.JSON(401, gin.H{
			"Status":  "error",
			"Message": err.Error(),
		})
		c.Set("is_login", false)
		return
	}

	userinfo := claims["UserInfo"].(map[string]interface{})

	c.Set("is_login", true)
	c.Set("userinfo", userinfo)

	fmt.Print(userinfo["User_id"])
}
