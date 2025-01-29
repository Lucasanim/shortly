package repository

import (
	"fmt"
	"log"

	"github.com/Lucasanim/shortly/internal/database"
	"github.com/Lucasanim/shortly/internal/models"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type LinkRepository struct{}

func (lr *LinkRepository) Create(link models.Link) {
	av, err := dynamodbattribute.MarshalMap(link)
	if err != nil {
		log.Fatalf("Got error marshalling link: %s", err)
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(database.TableName),
	}

	_, err = database.Instance.PutItem(input)
	if err != nil {
		log.Fatalf("Got error calling PutItem: %s", err)
	}

	fmt.Println("Link created successfuly")
}

func (lr *LinkRepository) Get(hash string) models.Link {
	input := &dynamodb.QueryInput{
		TableName:              aws.String(database.TableName),
		IndexName:              aws.String("HashIndex"),
		KeyConditionExpression: aws.String("#H = :hash"),
		ExpressionAttributeNames: map[string]*string{
			"#H": aws.String("Hash"),
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":hash": {S: aws.String(hash)},
		},
	}

	result, err := database.Instance.Query(input)
	if err != nil {
		log.Fatalf("failed to get item, %v", err)
	}

	link := models.Link{}
	if len(result.Items) == 0 {
		fmt.Println("No item found")
		return link
	}

	err = dynamodbattribute.UnmarshalMap(result.Items[0], &link)
	if err != nil {
		log.Fatalf("failed to unmarshal item, %v", err)
	}

	return link
}
