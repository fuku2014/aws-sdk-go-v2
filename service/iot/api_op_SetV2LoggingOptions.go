// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
)

type SetV2LoggingOptionsInput struct {
	_ struct{} `type:"structure"`

	// The default logging level.
	DefaultLogLevel LogLevel `locationName:"defaultLogLevel" type:"string" enum:"true"`

	// If true all logs are disabled. The default is false.
	DisableAllLogs *bool `locationName:"disableAllLogs" type:"boolean"`

	// The ARN of the role that allows IoT to write to Cloudwatch logs.
	RoleArn *string `locationName:"roleArn" type:"string"`
}

// String returns the string representation
func (s SetV2LoggingOptionsInput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s SetV2LoggingOptionsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if len(s.DefaultLogLevel) > 0 {
		v := s.DefaultLogLevel

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "defaultLogLevel", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.DisableAllLogs != nil {
		v := *s.DisableAllLogs

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "disableAllLogs", protocol.BoolValue(v), metadata)
	}
	if s.RoleArn != nil {
		v := *s.RoleArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "roleArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type SetV2LoggingOptionsOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s SetV2LoggingOptionsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s SetV2LoggingOptionsOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opSetV2LoggingOptions = "SetV2LoggingOptions"

// SetV2LoggingOptionsRequest returns a request value for making API operation for
// AWS IoT.
//
// Sets the logging options for the V2 logging service.
//
//    // Example sending a request using SetV2LoggingOptionsRequest.
//    req := client.SetV2LoggingOptionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) SetV2LoggingOptionsRequest(input *SetV2LoggingOptionsInput) SetV2LoggingOptionsRequest {
	op := &aws.Operation{
		Name:       opSetV2LoggingOptions,
		HTTPMethod: "POST",
		HTTPPath:   "/v2LoggingOptions",
	}

	if input == nil {
		input = &SetV2LoggingOptionsInput{}
	}

	req := c.newRequest(op, input, &SetV2LoggingOptionsOutput{})
	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return SetV2LoggingOptionsRequest{Request: req, Input: input, Copy: c.SetV2LoggingOptionsRequest}
}

// SetV2LoggingOptionsRequest is the request type for the
// SetV2LoggingOptions API operation.
type SetV2LoggingOptionsRequest struct {
	*aws.Request
	Input *SetV2LoggingOptionsInput
	Copy  func(*SetV2LoggingOptionsInput) SetV2LoggingOptionsRequest
}

// Send marshals and sends the SetV2LoggingOptions API request.
func (r SetV2LoggingOptionsRequest) Send(ctx context.Context) (*SetV2LoggingOptionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &SetV2LoggingOptionsResponse{
		SetV2LoggingOptionsOutput: r.Request.Data.(*SetV2LoggingOptionsOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// SetV2LoggingOptionsResponse is the response type for the
// SetV2LoggingOptions API operation.
type SetV2LoggingOptionsResponse struct {
	*SetV2LoggingOptionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// SetV2LoggingOptions request.
func (r *SetV2LoggingOptionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}