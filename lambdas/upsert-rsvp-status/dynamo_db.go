package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/rs/zerolog/log"
)

type RsvpItem struct {
	CognitoSub string
	RsvpStatus bool
}

type DynamoDbClientInterface interface {
	PutItem(ctx context.Context, params *dynamodb.PutItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.PutItemOutput, error)
}

func insertRsvpItem(rsvpItem RsvpItem, dbName string, client DynamoDbClientInterface, ctx context.Context) error {

	item := map[string]types.AttributeValue{
		"cognito_guid": &types.AttributeValueMemberS{Value: rsvpItem.CognitoSub},
		"rsvp_status":  &types.AttributeValueMemberBOOL{Value: rsvpItem.RsvpStatus},
	}

	_, err := client.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(dbName),
		Item:      item,
	})

	if err != nil {
		log.Info().Msg(fmt.Sprintf("Successfully added %s with RSVP status %t", rsvpItem.CognitoSub, rsvpItem.RsvpStatus))
	}

	return err
}
