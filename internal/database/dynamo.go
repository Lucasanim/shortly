package database

import (
	"github.com/Lucasanim/shortly/config"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var Instance *dynamodb.DynamoDB
var TableName = "Link"

func InitializeDb() {
	// Initialize a session that the SDK will use to load
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
		Config: aws.Config{
			Region:   aws.String("us-east-1"),
			Endpoint: aws.String(config.Env.DynamoDbEndpoint),
		},
	}))

	// Create DynamoDB client
	Instance = dynamodb.New(sess)
}
