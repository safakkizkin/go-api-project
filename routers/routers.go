package routers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"safakkizkin/controllers"
)

func SetupRouters() *gin.Engine {
	r := gin.Default()
	r.Use(cors.Default())
	v1 := r.Group("/api/v1")
	{
		v1users := v1.Group("users")
		{
			v1users.GET("", controllers.GetUsers)
			v1users.GET(":mail", controllers.GetOneUserByMail)
			v1users.DELETE(":mail", controllers.DeleteUser)
			v1users.POST("", controllers.AddNewUser)
		}
	}

	return r
}
