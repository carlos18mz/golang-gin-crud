package main

import (
	"crud-gin/controller"
	"crud-gin/docs"
	"crud-gin/middlewares"
	"crud-gin/repository"
	"crud-gin/service"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	gindump "github.com/tpkeeper/gin-dump"
)

//Golang + gin + mysql Crud

//save repos
// go get -u {{repo-link}}

//Run project
//	go guild
//	.\crud-gin

//deploy to heroku
//	heroku create {{project-name}}
//	heroku mod init
//	**change go version to 1.14 in go.mod**
//	git push heroku master

var (
	videoRepository repository.VideoRepository = repository.NewVideoRepository()
	videoService    service.VideoService       = service.New(videoRepository)
	loginService    service.LoginService       = service.NewLoginService()
	jwtService      service.JWTService         = service.NewJWTService()

	videoController controller.VideoController = controller.New(videoService)
	loginController controller.LoginController = controller.NewLoginController(loginService, jwtService)
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func index(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Listen and Serve")
}

func main() {

	//Swagger documentation
	docs.SwaggerInfo.Title = "FaceChat V1"
	docs.SwaggerInfo.Description = "A chat project with sockets"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "golang-gin-be.herokuapp.com"
	docs.SwaggerInfo.BasePath = "api/v1"
	docs.SwaggerInfo.Schemes = []string{"https"}

	//With custom middleware
	setupLogOutput()

	server := gin.New()
	server.Use(gin.Recovery(), middlewares.Logger(),
		gindump.Dump())

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

	server.GET("/", func(ctx *gin.Context) {
		fmt.Println("execute /")
		ctx.JSON(200, "listen and serve")
	})

	// Login Endpoint: Authentication + Token creation
	server.POST("/login", func(ctx *gin.Context) {
		token := loginController.Login(ctx)
		if token != "" {
			ctx.JSON(http.StatusOK, gin.H{
				"token": token,
			})
		} else {
			ctx.JSON(http.StatusUnauthorized, nil)
		}
	})

	apiRoutes := server.Group("/api", middlewares.AuthorizeJWT())
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

		apiRoutes.PUT("/videos/:id", func(ctx *gin.Context) {
			err := videoController.Update(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "Video Input is Valid!!"})
			}
		})

		apiRoutes.DELETE("/videos/:id", func(ctx *gin.Context) {
			err := videoController.Delete(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "Video Input is Valid!!"})
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

	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	port := os.Getenv("PORT")
	server.Run(":" + port)
}
