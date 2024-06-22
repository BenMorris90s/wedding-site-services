package main

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
	"upsert-rsvp-status/mocks"
)

func TestDynamoPassesCorrectDataToPutItem(t *testing.T) {
	rsvpItem := RsvpItem{RsvpStatus: true, CognitoSub: "1-2-3-4"}
	ctx := context.TODO()
	dbName := "test-db"

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mocks.NewMockDynamoDbClientInterface(ctrl)

	expectedItem := map[string]types.AttributeValue{
		"cognito_guid": &types.AttributeValueMemberS{Value: rsvpItem.CognitoSub},
		"rsvp_status":  &types.AttributeValueMemberBOOL{Value: rsvpItem.RsvpStatus},
	}

	mockClient.EXPECT().PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(dbName),
		Item:      expectedItem,
	}).Return(&dynamodb.PutItemOutput{}, nil)

	err := insertRsvpItem(rsvpItem, dbName, mockClient, ctx)

	assert.NoError(t, err)
}
