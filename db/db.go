package db

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var (
	db  *dynamo.DB
	err error
)

func Init() *dynamo.DB {
	if os.Getenv("ENV") == "development" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	accessKey := os.Getenv("DYNAMO_DB_ACCESS_KEY")
	secretKey := os.Getenv("DYNAMO_DB_SECRET_KEY")

	cred := credentials.NewStaticCredentials(accessKey, secretKey, "")

	db := dynamo.New(session.New(), &aws.Config{
		Credentials: cred,
		Region:      aws.String("ap-northeast-1"),
	})

	return db
}

func GetDB() *dynamo.DB {
	return db
}
