package file

import (
	"file_uploader/model"
	"file_uploader/utils"
	"net/http"
	"github.com/gin-gonic/gin"
	_ "file_uploader/docs"
)

// @Description Upload mutiple files
// @Tags 		file
// @Accept 		multipart/form-data
// @Produce 	json
// @Param       files formData file true "File to upload"
// @Router 		/file/mutiple_upload [post]
func MutipleFileUploadHandler(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err})
		return
	}

	files := form.File["files"]

	if len(files) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "No file was added"})
		return
	}

	fileInfos := make([]*model.File, 0, len(files))




	for _, file := range files {
		fileInfo, err := utils.SaveUploadedFile(file)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err})
			return
		}

		fileInfos = append(fileInfos, fileInfo)
	}

	

	c.JSON(http.StatusOK, gin.H{"status": "ok","message": "files uploaded successfully", "file_infos": fileInfos})
}

