package file

import (
	"context"
	"file_uploader/db"
	"file_uploader/model"
	"net/http"
	"github.com/cloudinary/cloudinary-go/api/uploader"
	"file_uploader/config"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	_ "file_uploader/docs"
)

// @Description Delete a file
// @Tags 		file
// @Accept 		json
// @Produce 	json
// @Param 		id path string true "File ID"
// @Router 		/file/delete/{id} [get]
func FileDeletionHandler(c *gin.Context){
	fileID := c.Param("id")
	fileIDFormatted,_ := primitive.ObjectIDFromHex(fileID)

	//retrieve cloudinary_file_id
	file := model.File{}
	filter := bson.M{"_id":fileIDFormatted }
	err := db.FileCollection.FindOne(context.Background(), filter).Decode(&file)
	 if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err})
		return
	}

	//delete mongo document
	res, err := db.FileCollection.DeleteOne(context.Background(), filter)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err})
		return
	}

	if res.DeletedCount == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"status": "error", "message": "File with this id not found"})
		return
	}
	//delete file from cloudinary
	_,err = config.Cld.Upload.Destroy(context.Background(), uploader.DestroyParams{PublicID:file.FileID})
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"status": "ok", "message": "File deleted successful"})
}