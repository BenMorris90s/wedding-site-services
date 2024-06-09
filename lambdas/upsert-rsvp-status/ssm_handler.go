package main

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
)


func get_ssm_parameter(param_name string, ssm_client *ssm.Client, ctx context.Context) (string, error) {
	with_decryption := true

	get_param_command := ssm.GetParameterInput{
		Name: &param_name,
		WithDecryption: &with_decryption,
	}

	param, err := ssm_client.GetParameter(ctx, &get_param_command)

	return *param.Parameter.Value, err
}
