package Routes

import (
	"GoPractice/Controllers"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetUpRouter() *gin.Engine {
	r := gin.Default()

	grpNonToken := r.Group("auth-api")
	{
		grpNonToken.POST("register", Controllers.CreateUser)
		grpNonToken.POST("login", Controllers.Login)
	}

	grpToken := r.Group("user-api", authorizationMiddleware)
	{
		grpToken.GET("user", Controllers.GetUsers)
		grpToken.GET("user/:id", Controllers.GetUserById)
		grpToken.PUT("user/:id", Controllers.UpdateUser)
		grpToken.DELETE("user/:id", Controllers.DeleteUser)
	}

	grpUpload := r.Group("/upload-api", authorizationMiddleware)
	{
		grpUpload.GET("getUploadLists", Controllers.GetUploadLists)
		grpUpload.POST("uploadFile", Controllers.UploadFile)
	}

	grpDownload := r.Group("/downloadFile")
	{
		grpDownload.GET("byId/:fileId", Controllers.DownloadFileById)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
