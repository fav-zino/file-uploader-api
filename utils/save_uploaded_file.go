package utils

import (
	"context"
	// "errors"
	"file_uploader/db"
	"file_uploader/model"
	"file_uploader/config"
	"mime/multipart"
	"path/filepath"
	"go.mongodb.org/mongo-driver/bson/primitive"
    "github.com/cloudinary/cloudinary-go/api/uploader"

)

func SaveUploadedFile(file *multipart.FileHeader) (*model.File, error) {

	// Open the uploaded file
	src, err := file.Open()
	if err != nil {
		return nil,err
	}
	defer src.Close()

	uploadResult, err := config.Cld.Upload.Upload(context.Background(), src, uploader.UploadParams{})
    if err != nil {
        return nil,err
    }
	
	var fileInfo model.File
		
	fileInfo.Name    =   file.Filename
	fileInfo.FileUrl = uploadResult.SecureURL
	fileInfo.FileID = uploadResult.PublicID
	fileInfo.Size=        file.Size
	fileInfo.ContentType = filepath.Ext(file.Filename)
	

	// Insert the file details into MongoDB
	result, err := db.FileCollection.InsertOne(context.Background(), &fileInfo)
	if err != nil {
		return nil, err
	}


	fileInfo.ID = result.InsertedID.(primitive.ObjectID)

	return &fileInfo, nil
}
