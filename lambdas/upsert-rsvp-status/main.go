package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/rs/zerolog/log"
	"os"
)

const GuestInfoParameterName = "GUEST_INFO_PARAMETER_NAME"

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	rsvpItem := RsvpItem{}
	err := json.Unmarshal([]byte(request.Body), &rsvpItem)

	if err != nil {
		return response(404, "Request body is not a valid RSVP item", err)
	}

	sdkConfig, err := config.LoadDefaultConfig(ctx)

	if err != nil {
		return response(500, "Error when loading AWS SDK", err)
	}

	dbParameterName, exists := os.LookupEnv(GuestInfoParameterName)

	if !exists {
		err := errors.New(fmt.Sprintf("missing environment variable %s", GuestInfoParameterName))
		return response(500, fmt.Sprintf("Lambda is missing %s environment variable", GuestInfoParameterName), err)
	}

	ssmClient := ssm.NewFromConfig(sdkConfig)

	// Fetch DB name at runtime to get latest value. AWS options exist to reduce latency for fetching configuration values but this allows free tier usage
	// and the overhead is not a problem for this app currently.
	dbName, err := getSsmParameter(dbParameterName, true, ssmClient, ctx)

	if err != nil {
		return response(500, "Error when fetching DB name", err)
	}

	dynamoClient := dynamodb.NewFromConfig(sdkConfig)

	err = insertRsvpItem(rsvpItem, dbName, dynamoClient, ctx)

	if err != nil {
		return response(500, "Error when inserting rsvp item", err)
	}

	return response(200, "Execution completed successfully", nil)
}

func response(statusCode int, body string, err error) (events.APIGatewayProxyResponse, error) {
	if err != nil {
		log.Error().Msg(err.Error())
	}

	return events.APIGatewayProxyResponse{
		Body:       body,
		StatusCode: statusCode,
	}, nil
}

func main() {
	lambda.Start(handler)
}
