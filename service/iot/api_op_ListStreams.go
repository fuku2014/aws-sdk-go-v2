// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type ListStreamsInput struct {
	_ struct{} `type:"structure"`

	// Set to true to return the list of streams in ascending order.
	AscendingOrder *bool `location:"querystring" locationName:"isAscendingOrder" type:"boolean"`

	// The maximum number of results to return at a time.
	MaxResults *int64 `location:"querystring" locationName:"maxResults" min:"1" type:"integer"`

	// A token used to get the next set of results.
	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListStreamsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListStreamsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListStreamsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListStreamsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.AscendingOrder != nil {
		v := *s.AscendingOrder

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "isAscendingOrder", protocol.BoolValue(v), metadata)
	}
	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "maxResults", protocol.Int64Value(v), metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type ListStreamsOutput struct {
	_ struct{} `type:"structure"`

	// A token used to get the next set of results.
	NextToken *string `locationName:"nextToken" type:"string"`

	// A list of streams.
	Streams []StreamSummary `locationName:"streams" type:"list"`
}

// String returns the string representation
func (s ListStreamsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListStreamsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Streams != nil {
		v := s.Streams

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "streams", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opListStreams = "ListStreams"

// ListStreamsRequest returns a request value for making API operation for
// AWS IoT.
//
// Lists all of the streams in your AWS account.
//
//    // Example sending a request using ListStreamsRequest.
//    req := client.ListStreamsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) ListStreamsRequest(input *ListStreamsInput) ListStreamsRequest {
	op := &aws.Operation{
		Name:       opListStreams,
		HTTPMethod: "GET",
		HTTPPath:   "/streams",
	}

	if input == nil {
		input = &ListStreamsInput{}
	}

	req := c.newRequest(op, input, &ListStreamsOutput{})

	return ListStreamsRequest{Request: req, Input: input, Copy: c.ListStreamsRequest}
}

// ListStreamsRequest is the request type for the
// ListStreams API operation.
type ListStreamsRequest struct {
	*aws.Request
	Input *ListStreamsInput
	Copy  func(*ListStreamsInput) ListStreamsRequest
}

// Send marshals and sends the ListStreams API request.
func (r ListStreamsRequest) Send(ctx context.Context) (*ListStreamsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListStreamsResponse{
		ListStreamsOutput: r.Request.Data.(*ListStreamsOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListStreamsResponse is the response type for the
// ListStreams API operation.
type ListStreamsResponse struct {
	*ListStreamsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListStreams request.
func (r *ListStreamsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}