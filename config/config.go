package config

import (
	"os"

	"github.com/cloudinary/cloudinary-go"
)


var Cld *cloudinary.Cloudinary

func SetupCloudinary() {
    Cld, _ = cloudinary.NewFromParams(os.Getenv("CLOUDINARY_CLOUD_NAME"), os.Getenv("CLOUDINARY_API_KEY"), os.Getenv("CLOUDINARY_API_SECRET"))

   
   

 
}



