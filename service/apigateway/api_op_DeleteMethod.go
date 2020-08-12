// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package apigateway

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
)

// Request to delete an existing Method resource.
type DeleteMethodInput struct {
	_ struct{} `type:"structure"`

	// [Required] The HTTP verb of the Method resource.
	//
	// HttpMethod is a required field
	HttpMethod *string `location:"uri" locationName:"http_method" type:"string" required:"true"`

	// [Required] The Resource identifier for the Method resource.
	//
	// ResourceId is a required field
	ResourceId *string `location:"uri" locationName:"resource_id" type:"string" required:"true"`

	// [Required] The string identifier of the associated RestApi.
	//
	// RestApiId is a required field
	RestApiId *string `location:"uri" locationName:"restapi_id" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteMethodInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteMethodInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteMethodInput"}

	if s.HttpMethod == nil {
		invalidParams.Add(aws.NewErrParamRequired("HttpMethod"))
	}

	if s.ResourceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceId"))
	}

	if s.RestApiId == nil {
		invalidParams.Add(aws.NewErrParamRequired("RestApiId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteMethodInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.HttpMethod != nil {
		v := *s.HttpMethod

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "http_method", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ResourceId != nil {
		v := *s.ResourceId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "resource_id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RestApiId != nil {
		v := *s.RestApiId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "restapi_id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DeleteMethodOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteMethodOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteMethodOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDeleteMethod = "DeleteMethod"

// DeleteMethodRequest returns a request value for making API operation for
// Amazon API Gateway.
//
// Deletes an existing Method resource.
//
//    // Example sending a request using DeleteMethodRequest.
//    req := client.DeleteMethodRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) DeleteMethodRequest(input *DeleteMethodInput) DeleteMethodRequest {
	op := &aws.Operation{
		Name:       opDeleteMethod,
		HTTPMethod: "DELETE",
		HTTPPath:   "/restapis/{restapi_id}/resources/{resource_id}/methods/{http_method}",
	}

	if input == nil {
		input = &DeleteMethodInput{}
	}

	req := c.newRequest(op, input, &DeleteMethodOutput{})
	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return DeleteMethodRequest{Request: req, Input: input, Copy: c.DeleteMethodRequest}
}

// DeleteMethodRequest is the request type for the
// DeleteMethod API operation.
type DeleteMethodRequest struct {
	*aws.Request
	Input *DeleteMethodInput
	Copy  func(*DeleteMethodInput) DeleteMethodRequest
}

// Send marshals and sends the DeleteMethod API request.
func (r DeleteMethodRequest) Send(ctx context.Context) (*DeleteMethodResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteMethodResponse{
		DeleteMethodOutput: r.Request.Data.(*DeleteMethodOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteMethodResponse is the response type for the
// DeleteMethod API operation.
type DeleteMethodResponse struct {
	*DeleteMethodOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteMethod request.
func (r *DeleteMethodResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}