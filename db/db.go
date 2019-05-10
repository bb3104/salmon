package db

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
	"os"
)

func GetDB() *dynamo.DB {
	accessKey := os.Getenv("DYNAMO_DB_ACCESS_KEY")
	secretKey := os.Getenv("DYNAMO_DB_SECRET_KEY")

	cred := credentials.NewStaticCredentials(accessKey, secretKey, "")

	db := dynamo.New(session.New(), &aws.Config{
		Credentials: cred,
		Region:      aws.String("ap-northeast-1"),
	})

	return db
}
