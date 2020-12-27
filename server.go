package main

import (
	"crud-gin/controller"
	"crud-gin/middlewares"
	"crud-gin/service"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	gindump "github.com/tpkeeper/gin-dump"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {

	//With custom middleware
	setupLogOutput()

	server := gin.New()
	server.Use(gin.Recovery(), middlewares.Logger(),
		middlewares.BasicAuth(), gindump.Dump())

	//Recovery(): 	get transactions on console
	//Logger(): 	save logs in file
	//BasicAuth():	request basic auth for http request
	//Dump():		get aditional inforamtion for debugging

	//w/o custom middlewware
	//server := gin.Default()

	//load css styles
	server.Static("/css", "./tempaltes/css")

	//load html static pages
	server.LoadHTMLGlob("templates/*.html")

	apiRoutes := server.Group("/api")
	{
		apiRoutes.GET("/videos", func(ctx *gin.Context) {
			ctx.JSON(200, videoController.FindAll())
		})

		apiRoutes.POST("/videos", func(ctx *gin.Context) {
			err := videoController.Save(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": " Input valid!!"})
			}
		})

		apiRoutes.GET("/", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": "OK!!",
			})
		})
	}

	viewRoutes := server.Group("/view")
	{
		viewRoutes.GET("/videos", videoController.ShowAll)
	}

	server.Run(":8080")
}
