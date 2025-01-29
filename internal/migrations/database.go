package migrations

import (
	"fmt"
	"log"

	"github.com/Lucasanim/shortly/internal/database"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func Migrate() {
	tableName := database.TableName

	if tableExists(tableName) {
		fmt.Println("Table already exists, skipping creation.")
		return
	}

	input := &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("ID"), // Primary key (HASH)
				AttributeType: aws.String("N"),  // Number type
			},
			{
				AttributeName: aws.String("Hash"), // Sort key (RANGE)
				AttributeType: aws.String("S"),    // String type
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("ID"),
				KeyType:       aws.String("HASH"), // Partition key
			},
			{
				AttributeName: aws.String("Hash"),
				KeyType:       aws.String("RANGE"), // Sort key
			},
		},
		GlobalSecondaryIndexes: []*dynamodb.GlobalSecondaryIndex{
			{
				IndexName: aws.String("HashIndex"),
				KeySchema: []*dynamodb.KeySchemaElement{
					{
						AttributeName: aws.String("Hash"),
						KeyType:       aws.String("HASH"), // Partition key
					},
				},
				Projection: &dynamodb.Projection{
					ProjectionType: aws.String("ALL"),
				},
				ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
					ReadCapacityUnits:  aws.Int64(10),
					WriteCapacityUnits: aws.Int64(10),
				},
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(10),
			WriteCapacityUnits: aws.Int64(10),
		},
		TableName: aws.String(tableName),
	}

	// Create the table
	_, err := database.Instance.CreateTable(input)
	if err != nil {
		log.Fatalf("Failed to create table: %s", err)
	}

	fmt.Println("Created the table", tableName)
}

func tableExists(tableName string) bool {
	_, err := database.Instance.DescribeTable(&dynamodb.DescribeTableInput{
		TableName: aws.String(tableName),
	})

	return err == nil
}
