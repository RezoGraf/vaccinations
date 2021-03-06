package main

import (
	"./controllers"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/contrib/jwt"
	"github.com/gin-contrib/cors"
)

var (
	mysupersecretpassword = "mySecret"
)

func main() {

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowAllOrigins: 	true,
		AllowMethods: 		[]string{"PUT", "GET", "POST","OPTIONS", "PATH"},
		AllowHeaders: 		[]string{"Origin", "Authorization"},
		ExposeHeaders: 		[]string{"Content-Length", "Access-Control-Allow-Headers", "Access-Control-Request-Headers", "Access-Control-Request-Method"},
		AllowCredentials: 	true,
	}))

	jwtGroup := router.Group("/jwt")
	jwtGroup.Use(jwt.Auth(mysupersecretpassword))

	jwtGroup.GET("/main", controllers.MainJwt)
	jwtGroup.POST("/FindPatientInArena", controllers.FindPatientInArena)
	jwtGroup.POST("checktoken", controllers.CheckToken)

	jwtGroup.GET("encodeJwt", controllers.GetCurentUser)

	jwtGroup.POST("addPrivivka", controllers.AddPrivivka)

	jwtGroup.POST("getPrivivka", controllers.GetPrivivka)


	router.POST("/login", controllers.Auth)
	router.Use(gin.Logger())
	router.Run(":8084")
}
