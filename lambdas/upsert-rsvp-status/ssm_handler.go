package main

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
)

type SSMClientInterface interface {
	GetParameter(ctx context.Context, params *ssm.GetParameterInput, optFns ...func(*ssm.Options)) (*ssm.GetParameterOutput, error)
}

func getSsmParameter(paramName string, withDecryption bool, ssmClient SSMClientInterface, ctx context.Context) (string, error) {
	getParamCommand := ssm.GetParameterInput{
		Name:           &paramName,
		WithDecryption: &withDecryption,
	}

	param, err := ssmClient.GetParameter(ctx, &getParamCommand)

	if err != nil {
		return "", err
	}

	return *param.Parameter.Value, err
}
