package file

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"file_uploader/utils"
	_ "file_uploader/docs"
	
)

	
// @Description Upload a single file
// @Tags 		file
// @Accept 		multipart/form-data
// @Produce 	json
// @Param       file formData file true "File to upload"
// @Router 		/file/upload [post]
func FileUploadHandler(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error","message": "No file uploaded"})
		return
	}

	fileInfo, err := utils.SaveUploadedFile(file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error","message": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok","message": "file uploaded successfully", "file_info": fileInfo})
}


