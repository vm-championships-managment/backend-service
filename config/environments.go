package config

import (
	"fmt"
	"os"
)

var AwsAccessKeyId string = "access_key"                                                 // os.Getenv("AWS_ACESS_KEY_ID")
var AwsSecretKey string = "secret_key"                                                   // os.Getenv("AWS_SECRET_KEY")
var AwsRegion string = "us-east-1"                                                       // os.Getenv("AWS_REGION")
var AwsEndpoint string = fmt.Sprintf("http://%s:4566", os.Getenv("LOCALSTACK_HOSTNAME")) // os.Getenv("AWS_ENDPOINT")
var ProjectEnvironment string = "development"                                            // os.Getenv("PROJECT_ENVIRONMENT")
var UserTableName string = "users"                                                       // os.Getenv("DYNAMODB_USER_TABLE_NAME")
