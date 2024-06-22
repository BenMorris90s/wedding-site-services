package main

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
	"upsert-rsvp-status/mocks"
)

func TestGetSSMParameterCalledWithCorrectArgs(t *testing.T) {
	const expectedValue = "myDbName"
	expectedError := error(nil)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mocks.NewMockSSMClientInterface(ctrl)

	ctx := context.TODO()
	paramName := "testParam"

	expectedInput := &ssm.GetParameterInput{
		Name:           aws.String(paramName),
		WithDecryption: aws.Bool(true),
	}

	mockClient.EXPECT().GetParameter(ctx, expectedInput).Return(&ssm.GetParameterOutput{
		Parameter: &types.Parameter{Value: aws.String(expectedValue)},
	}, nil)

	value, err := getSsmParameter(paramName, true, mockClient, ctx)

	assert.Equal(t, value, expectedValue)
	assert.Equal(t, expectedError, err)
}
