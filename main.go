package main

import (
	"fmt"
	"os"
	"strings"
	"suranect_api/database"
	"suranect_api/model"
	"suranect_api/route"
	"suranect_api/utils"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type headerIsAdmin struct {
	Testing string
}

func is_admin(c *gin.Context) {
	fmt.Println("on is_admin")
	var t headerIsAdmin
	c.ShouldBindHeader(&t)

	fmt.Println(t.Testing)

}

func main() {
	godotenv.Load(".env")

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	route.Init(r)

	r.GET("/migration_up", func(c *gin.Context) {

		db, err := database.ConnectMysql()

		if err != nil {
			c.JSON(200, gin.H{
				"status":  "error",
				"message": err.Error(),
			})
			return
		}

		db.AutoMigrate(&model.User{})
		c.JSON(200, gin.H{
			"status":  "success",
			"message": "Success Migrate All Database",
		})
		return
	})

	r.GET("/test_pollution", func(c *gin.Context) {

		url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/air_pollution?lat=-7.250445&lon=112.768845&appid=%s", os.Getenv("openweather_api"))

		c.JSON(200, gin.H{
			"Pollution": utils.FetchApi(url),
		})
	})

	type userInfoJwt struct {
		User_id string
	}

	type headerJWT struct {
		Authorization string
	}

	r.GET("/test_jwt", func(c *gin.Context) {

		var userInfo userInfoJwt
		userInfo.User_id = "1"

		jwt, err := utils.CreateToken("Hello World", userInfo)

		if err != nil {
			c.JSON(200, gin.H{
				"Message": "Gagal",
			})
		}

		c.JSON(200, gin.H{
			"Jwt": jwt,
		})
	})

	r.GET("/test_jwt_parser", func(c *gin.Context) {

		var formjwt headerJWT
		c.ShouldBindHeader(&formjwt)

		lah := strings.Split(formjwt.Authorization, " ")[1]

		test, err := utils.GetClaimsFromToken(lah)

		if err != nil {
			c.JSON(200, gin.H{
				"Message": "Gagal",
			})
		}

		userinfo := test["UserInfo"].(map[string]interface{})

		c.JSON(200, gin.H{
			"valid": lah,
			"da":    userinfo["User_id"],
		})
	})

	r.GET("/test_firebase", func(c *gin.Context) {
		testing := database.ConnectFirebase("Testing")

		c.JSON(200, gin.H{
			"data": testing.Get(),
		})
	})

	r.GET("/test_middleware", is_admin, func(c *gin.Context) {

		fmt.Println("on controller")

		c.JSON(200, gin.H{
			"testing": "loalah",
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}
