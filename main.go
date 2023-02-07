package main

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

const awsAccessKeyID = "AKIA2IUTP2UD55QMFPVC"
const awsSecretAccessKey = "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"

func main() {
	_, err := session.NewSession(&aws.Config{
		Region:      aws.String("us-east-1"),
		Credentials: credentials.NewStaticCredentials(awsAccessKeyID, awsSecretAccessKey, ""),
	})

	if err != nil {
		log.Fatalln(err)
	}
}
