## Lambda API

API for connecting to the Guest Info table.

## Authentication

The API is under development and currently unauthenticated. The lambda stack should be torn down once work is complete for the day
to avoid leaving a public endpoint floating around.

## Running Locally

```bash
  ./run_local.sh
```

Trigger lambda once running with:

```bash
    curl -X POST http://127.0.0.1:3000/rsvp \
         -H "Content-Type: application/json" \
         -d '{"CognitoSub": "1-2-3-4", "RsvpStatus": true}'
```

## Deployment

The settings in samconfig.toml are used by the deploy script. Change these if needed.

```bash
    ./deploy.sh
```

Upsert RSVP endpoint can be hit with

```bash
    curl -X POST https://{MY_API_GATEWAY_ID}.execute-api.eu-west-2.amazonaws.com/Prod/rsvp \                                         
             -H "Content-Type: application/json" \
             -d '{"CognitoSub": "1-2-3-4", "RsvpStatus": true}'
```

## Unit Tests

```bash
  cd ./upsert-rsvp-status 
  go test
```

Generate mocks with mockgen as needed. An example is the mock generated for the DynamoDB interface in dynamo.go:

```bash
  mockgen -destination=mocks/dynamodb_client_mock.go -package=mocks -source dynamo.go DynamoDbClientInterface
```

Yet to be added to a CI step in Github, so can be run on a pre push git hook locally.
