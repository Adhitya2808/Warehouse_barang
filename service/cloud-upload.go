package service

import (
	"context"
	"SaveHouse/config"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func UploadImage(folder string, file any) (string, error) {
	cfg := config.InitDB()
	cld, err := cloudinary.NewFromURL(cfg.CLOUD_URL)

	if err != nil {
		return "", err
	}

	ctx := context.Background()

	upload, err := cld.Upload.Upload(ctx, file, uploader.UploadParams{
		Folder: folder,
	})

	return upload.SecureURL, nil
}