package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"linux-host/config"
	"net/http"
)

func RouterInit() {
	router := gin.Default()
	router.Use(cors.Default())
	// Set a lower memory limit for multipart forms (default is 32 MiB)
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	router.Static("/fonts", "./static/fonts")
	router.Static("/css", "./static/css")
	router.Static("/js", "./static/js")
	router.Static("/img", "./static/img")
	router.LoadHTMLGlob("./static/index.html")
	router.GET("/index.html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	app := router.Group("/app/v1")
	{
		app.POST("/login", nil)
		//app.GET("/userPermission", controller.UserPermissionController)
	}
	con := router.Group("/config/v1")
	{
		con.POST("/login1", nil)
		//config.GET("/userPermission", controller.UserPermissionController)
	}
	router.Run(":" + config.ParamsConfig.GetString("server.port"))
}
