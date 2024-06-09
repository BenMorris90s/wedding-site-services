## What is this project?

This is a work in progress exercise to create a website capable of receiving wedding guest RSVPs and dietary requirements, as well as the location and date of the
wedding venue.

I am working on this in my spare time primarily in order to gain a familiarity with Go.

## Project Requirements

To run the lambda API locally, you will need to have an AWS account:

https://portal.aws.amazon.com/billing/signup#/start/email

As a minimum, your user will need to have permissions to create the following:

CloudFormation stacks
SSM Parameters
DynamoDB resources
Lambda resources

Download the AWS CLI:

https://docs.aws.amazon.com/cli/latest/userguide/getting-started-install.html

Create programatic access credentials for your IAM user and configure your machine locally with:

```bash
  aws configure
```
And enter your credentials when prompted.

You will also need the aws sam CLI:

https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/install-sam-cli.html

Download Go:

https://go.dev/doc/install

And download golangci-lint:

https://golangci-lint.run/welcome/install/

Install this locally as linting is currently performed locally and not yet as part of a CI process.

Finally, follow the steps in the database/README.md file to deploy the dynamo table and associated SSM parameters. 
The Lambda API will query these when running locally or in the cloud, so they must be deployed in order to develop the Lambda API.
