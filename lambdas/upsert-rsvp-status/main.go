package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/rs/zerolog/log"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
)

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	config, err := config.LoadDefaultConfig(ctx)

	if err != nil {
		log.Error().Msg(err.Error());
		return response(500, "Error when loading AWS SDK");
	}

	client := ssm.NewFromConfig(config);


	if err != nil {
		log.Error().Msg(err.Error());
		return response(500, "Error when fetching DB name");
	}

	return response(200, *param.Parameter.Value);
}

func response(status_code int, body string) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		Body: body,
		StatusCode: status_code,
	}, nil
}

func main() {
	lambda.Start(handler)
}