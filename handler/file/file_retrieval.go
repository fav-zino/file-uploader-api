package file

import (
	"context"
	"net/http"
"file_uploader/model"
"file_uploader/db"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson"
	_ "file_uploader/docs"
)

// @Description Retrieve a file
// @Tags 		file
// @Accept 		json
// @Produce 	json
// @Param 		id path string true "File ID"
// @Router 		/file/retrieve/{id} [get]
func FileRetrievalHandler(c *gin.Context){
	fileID := c.Param("id")
	fileIDFormatted,_ := primitive.ObjectIDFromHex(fileID)
	file := model.File{}
	filter := bson.M{"_id":fileIDFormatted }
	err := db.FileCollection.FindOne(context.Background(), filter).Decode(&file)
	 if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"status": "ok", "message": "success", "file_info": file})
}