package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type File struct {
    ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
    Name      string    `json:"name" bson:"name"`
    Size      int64     `json:"size" bson:"size"`
	FileUrl string    `json:"file_url" bson:"file_url"`
	FileID string    `json:"file_id" bson:"file_id"`
    ContentType string  `json:"content_type" bson:"content_type"`
}
