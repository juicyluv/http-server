package repository

import (
	"context"
	"log"
	"os"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
	"github.com/ellywynn/http-server/internal/app/models/interfaces"
)

type CloudinaryService struct {
	cld *cloudinary.Cloudinary
}

// Returns cloudinary service instance with config from .env file
// Fatals if cloudinary cannot be created
func NewCloudinaryService() interfaces.CloudinaryService {
	cloud := os.Getenv("CLD_CLOUD")
	key := os.Getenv("CLD_KEY")
	secret := os.Getenv("CLD_SECRET")
	cld, err := cloudinary.NewFromParams(cloud, key, secret)
	if err != nil {
		log.Fatalf(err.Error())
	}

	return &CloudinaryService{
		cld: cld,
	}
}

// Uploads an image and returns image URL or error
func (c *CloudinaryService) UploadImage(image, name, folder string) (string, error) {
	resp, err := c.cld.Upload.Upload(context.TODO(), image, uploader.UploadParams{Folder: "travels", PublicID: name})
	if err != nil {
		return "", err
	}
	return resp.URL, nil
}
