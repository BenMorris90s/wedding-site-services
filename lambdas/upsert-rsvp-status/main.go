package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/rs/zerolog/log"
	"os"
)

const GuestInfoParameterName = "GUEST_INFO_PARAMETER_NAME"

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	config, err := config.LoadDefaultConfig(ctx)

	if err != nil {
		log.Error().Msg(err.Error())
		return response(500, "Error when loading AWS SDK")
	}

	dbParameterName, exists := os.LookupEnv(GuestInfoParameterName)

	if !exists {
		return response(500, fmt.Sprintf("Lambda is missing %s environment variable", GuestInfoParameterName))
	}

	client := ssm.NewFromConfig(config)

	// Fetch DB name at runtime to get latest value. AWS options exist to reduce latency for fetching configuration values but this allows free tier usage
	// and the overhead is not a problem for this app currently.
	dbName, err := getSsmParameter(dbParameterName, true, client, ctx)

	if err != nil {
		log.Error().Msg(err.Error())
		return response(500, "Error when fetching DB name")
	}

	return response(200, dbName)
}

func response(statusCode int, body string) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		Body:       body,
		StatusCode: statusCode,
	}, nil
}

func main() {
	lambda.Start(handler)
}
