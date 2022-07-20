package aws

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var (
	s3session *s3.S3
)

var region string
var secretKey string
var accessKey string

func init() {
	region = os.Getenv("REGION")
	secretKey = os.Getenv("AWS_SECRET_KEY")
	accessKey = os.Getenv("AWS_ACCESS_KEY")

	fmt.Println(region)
}

func ConnectS3() *s3.S3 {
	return s3.New(session.Must(session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials(accessKey, secretKey, ""),
	})))

}