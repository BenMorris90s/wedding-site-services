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
		log.Error().Msg(err.Error())
		return response(500, "Error when loading AWS SDK")
	}

	client := ssm.NewFromConfig(config);

	// Fetch DB name at runtime to get latest value. AWS options exist to reduce latency for fetching configuration values but this allows free tier usage
	// and the overhead is not a problem for this app currently.
	dbName, err := get_ssm_parameter("/database/guest_info/table_name", client, ctx)

	if err != nil {
		log.Error().Msg(err.Error());
		return response(500, "Error when fetching DB name")
	}

	return response(200, dbName);
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