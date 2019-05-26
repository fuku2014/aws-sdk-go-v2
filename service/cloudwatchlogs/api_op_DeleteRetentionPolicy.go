// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudwatchlogs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/logs-2014-03-28/DeleteRetentionPolicyRequest
type DeleteRetentionPolicyInput struct {
	_ struct{} `type:"structure"`

	// The name of the log group.
	//
	// LogGroupName is a required field
	LogGroupName *string `locationName:"logGroupName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteRetentionPolicyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteRetentionPolicyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteRetentionPolicyInput"}

	if s.LogGroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("LogGroupName"))
	}
	if s.LogGroupName != nil && len(*s.LogGroupName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("LogGroupName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/logs-2014-03-28/DeleteRetentionPolicyOutput
type DeleteRetentionPolicyOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteRetentionPolicyOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteRetentionPolicy = "DeleteRetentionPolicy"

// DeleteRetentionPolicyRequest returns a request value for making API operation for
// Amazon CloudWatch Logs.
//
// Deletes the specified retention policy.
//
// Log events do not expire if they belong to log groups without a retention
// policy.
//
//    // Example sending a request using DeleteRetentionPolicyRequest.
//    req := client.DeleteRetentionPolicyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/logs-2014-03-28/DeleteRetentionPolicy
func (c *Client) DeleteRetentionPolicyRequest(input *DeleteRetentionPolicyInput) DeleteRetentionPolicyRequest {
	op := &aws.Operation{
		Name:       opDeleteRetentionPolicy,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteRetentionPolicyInput{}
	}

	req := c.newRequest(op, input, &DeleteRetentionPolicyOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteRetentionPolicyRequest{Request: req, Input: input, Copy: c.DeleteRetentionPolicyRequest}
}

// DeleteRetentionPolicyRequest is the request type for the
// DeleteRetentionPolicy API operation.
type DeleteRetentionPolicyRequest struct {
	*aws.Request
	Input *DeleteRetentionPolicyInput
	Copy  func(*DeleteRetentionPolicyInput) DeleteRetentionPolicyRequest
}

// Send marshals and sends the DeleteRetentionPolicy API request.
func (r DeleteRetentionPolicyRequest) Send(ctx context.Context) (*DeleteRetentionPolicyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteRetentionPolicyResponse{
		DeleteRetentionPolicyOutput: r.Request.Data.(*DeleteRetentionPolicyOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteRetentionPolicyResponse is the response type for the
// DeleteRetentionPolicy API operation.
type DeleteRetentionPolicyResponse struct {
	*DeleteRetentionPolicyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteRetentionPolicy request.
func (r *DeleteRetentionPolicyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}