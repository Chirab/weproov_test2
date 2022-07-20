package images

import (
	"mime/multipart"
	"time"

	pkgAws "weproov/pkg"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/cip8/autoname"
)

type ImageServices struct {}

func NewImagesServices() *ImageServices {
	return &ImageServices{}
}

func (i *ImageServices) CreateImage(filename, headerContentTtype string, file multipart.File,) (string, error) {
	keyName := autoname.Generate("")
	option := &s3.PutObjectInput{
			Bucket: aws.String("weproov"),
			Key:    aws.String(keyName),
			Body:   file,
	}

	respS3 := pkgAws.ConnectS3()
	_, err := respS3.PutObject(option)
	if err != nil {
		return "",err
	}

	return keyName, nil
}

func (i *ImageServices) GetImageByKey(keyName string) (string,error) {
	respS3 := pkgAws.ConnectS3()
	req, _ := respS3.GetObjectRequest(&s3.GetObjectInput{
			Bucket: aws.String("weproov"),
			Key:    aws.String(keyName),
	})

	imgUrl, err := req.Presign(15 * time.Minute)
	if err != nil {
		return "", err
	}

	return imgUrl, nil
}

func (i *ImageServices) DeleteImage(keyName string) error {
	respS3 := pkgAws.ConnectS3()
	_, err := respS3.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String("weproov"),
		Key:    aws.String(keyName),
	})
	if err != nil {		
		return err
	}

	return nil
}
