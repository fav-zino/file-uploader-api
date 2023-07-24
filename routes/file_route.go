package routes

import (
	"file_uploader/handler/file"
	"github.com/gin-gonic/gin"
)

func LoadFileRoutes(router *gin.Engine) {

	authRouter := router.Group("/file")

	authRouter.POST("/upload", file.FileUploadHandler)
	authRouter.POST("/mutiple_upload", file.MutipleFileUploadHandler)
	authRouter.GET("/retrieve/:id", file.FileRetrievalHandler)
	authRouter.GET("/delete/:id", file.FileDeletionHandler)

}
