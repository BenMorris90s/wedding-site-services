package main

import (
	"context"
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
)

type MockSSMClient struct {
	GetParameterFunc func(ctx context.Context, input *ssm.GetParameterInput, opts ...func(*ssm.Options)) (*ssm.GetParameterOutput, error)
}

func (m *MockSSMClient) GetParameter(ctx context.Context, input *ssm.GetParameterInput, opts ...func(*ssm.Options)) (*ssm.GetParameterOutput, error) {
	return m.GetParameterFunc(ctx, input, opts...)
}


// Basic test to get mocking of AWS clients working and unit tets running in pre push hook. Will be replaced with more comprehensive tests once
// DynamoDB related functions and more complex functions are written.
func TestGetSSMParameter(t *testing.T) {
	const expected_value = "myDbName"
	expected_error := error(nil)

	mockSSM := &MockSSMClient{
		GetParameterFunc: func(ctx context.Context, input *ssm.GetParameterInput, opts ...func(*ssm.Options)) (*ssm.GetParameterOutput, error) {
			return &ssm.GetParameterOutput{
				Parameter: &types.Parameter{
					Value: aws.String(expected_value),
				},
			}, expected_error
		},
	}


	ctx := context.TODO()
	param_name := "testParam"
	with_decryption := true

	value, err := get_ssm_parameter(param_name, with_decryption, mockSSM, ctx)


	assert.Equal(t, value, expected_value)
	assert.Equal(t, expected_error, err)
}
