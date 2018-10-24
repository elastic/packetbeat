// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lambda

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
)

// Lambda provides the API operation methods for making requests to
// AWS Lambda. See this package's package overview docs
// for details on the service.
//
// Lambda methods are safe to use concurrently. It is not safe to
// modify mutate any of the struct's properties though.
type Lambda struct {
	*aws.Client
}

// Used for custom client initialization logic
var initClient func(*Lambda)

// Used for custom request initialization logic
var initRequest func(*Lambda, *aws.Request)

// Service information constants
const (
	ServiceName = "lambda"    // Service endpoint prefix API calls made to.
	EndpointsID = ServiceName // Service ID for Regions and Endpoints metadata.
)

// New creates a new instance of the Lambda client with a config.
//
// Example:
//     // Create a Lambda client from just a config.
//     svc := lambda.New(myConfig)
func New(config aws.Config) *Lambda {
	var signingName string
	signingRegion := config.Region

	svc := &Lambda{
		Client: aws.NewClient(
			config,
			aws.Metadata{
				ServiceName:   ServiceName,
				SigningName:   signingName,
				SigningRegion: signingRegion,
				APIVersion:    "2015-03-31",
			},
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBackNamed(v4.SignRequestHandler)
	svc.Handlers.Build.PushBackNamed(restjson.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(restjson.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(restjson.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(restjson.UnmarshalErrorHandler)

	// Run custom client initialization if present
	if initClient != nil {
		initClient(svc)
	}

	return svc
}

// newRequest creates a new request for a Lambda operation and runs any
// custom request initialization.
func (c *Lambda) newRequest(op *aws.Operation, params, data interface{}) *aws.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(c, req)
	}

	return req
}
