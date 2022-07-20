package aws

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/joho/godotenv"
)

var (
	s3session *s3.S3
)

var region string
var secretKey string
var accessKey string

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	region = os.Getenv("REGION")
	secretKey = os.Getenv("AWS_SECRET_KEY")
	accessKey = os.Getenv("AWS_ACCESS_KEY")
}

func ConnectS3() *s3.S3 {
	return s3.New(session.Must(session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials(accessKey, secretKey, ""),
	})))

}