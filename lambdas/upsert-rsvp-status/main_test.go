package main

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

type MockSSMClient struct {
	GetParameterFunc func(ctx context.Context, input *ssm.GetParameterInput, opts ...func(*ssm.Options)) (*ssm.GetParameterOutput, error)
}

func (m *MockSSMClient) GetParameter(ctx context.Context, input *ssm.GetParameterInput, opts ...func(*ssm.Options)) (*ssm.GetParameterOutput, error) {
	return m.GetParameterFunc(ctx, input, opts...)
}

func TestGetSSMParameter(t *testing.T) {
	const expectedValue = "myDbName"
	expectedError := error(nil)

	mockSSM := &MockSSMClient{
		GetParameterFunc: func(ctx context.Context, input *ssm.GetParameterInput, opts ...func(*ssm.Options)) (*ssm.GetParameterOutput, error) {
			return &ssm.GetParameterOutput{
				Parameter: &types.Parameter{
					Value: aws.String(expectedValue),
				},
			}, expectedError
		},
	}

	ctx := context.TODO()
	paramName := "testParam"

	value, err := getSsmParameter(paramName, true, mockSSM, ctx)

	assert.Equal(t, value, expectedValue)
	assert.Equal(t, expectedError, err)
}
