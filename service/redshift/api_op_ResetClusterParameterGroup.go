// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package redshift

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ResetClusterParameterGroupInput struct {
	_ struct{} `type:"structure"`

	// The name of the cluster parameter group to be reset.
	//
	// ParameterGroupName is a required field
	ParameterGroupName *string `type:"string" required:"true"`

	// An array of names of parameters to be reset. If ResetAllParameters option
	// is not used, then at least one parameter name must be supplied.
	//
	// Constraints: A maximum of 20 parameters can be reset in a single request.
	Parameters []Parameter `locationNameList:"Parameter" type:"list"`

	// If true, all parameters in the specified parameter group will be reset to
	// their default values.
	//
	// Default: true
	ResetAllParameters *bool `type:"boolean"`
}

// String returns the string representation
func (s ResetClusterParameterGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ResetClusterParameterGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ResetClusterParameterGroupInput"}

	if s.ParameterGroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ParameterGroupName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ResetClusterParameterGroupOutput struct {
	_ struct{} `type:"structure"`

	// The name of the cluster parameter group.
	ParameterGroupName *string `type:"string"`

	// The status of the parameter group. For example, if you made a change to a
	// parameter group name-value pair, then the change could be pending a reboot
	// of an associated cluster.
	ParameterGroupStatus *string `type:"string"`
}

// String returns the string representation
func (s ResetClusterParameterGroupOutput) String() string {
	return awsutil.Prettify(s)
}

const opResetClusterParameterGroup = "ResetClusterParameterGroup"

// ResetClusterParameterGroupRequest returns a request value for making API operation for
// Amazon Redshift.
//
// Sets one or more parameters of the specified parameter group to their default
// values and sets the source values of the parameters to "engine-default".
// To reset the entire parameter group specify the ResetAllParameters parameter.
// For parameter changes to take effect you must reboot any associated clusters.
//
//    // Example sending a request using ResetClusterParameterGroupRequest.
//    req := client.ResetClusterParameterGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/redshift-2012-12-01/ResetClusterParameterGroup
func (c *Client) ResetClusterParameterGroupRequest(input *ResetClusterParameterGroupInput) ResetClusterParameterGroupRequest {
	op := &aws.Operation{
		Name:       opResetClusterParameterGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ResetClusterParameterGroupInput{}
	}

	req := c.newRequest(op, input, &ResetClusterParameterGroupOutput{})

	return ResetClusterParameterGroupRequest{Request: req, Input: input, Copy: c.ResetClusterParameterGroupRequest}
}

// ResetClusterParameterGroupRequest is the request type for the
// ResetClusterParameterGroup API operation.
type ResetClusterParameterGroupRequest struct {
	*aws.Request
	Input *ResetClusterParameterGroupInput
	Copy  func(*ResetClusterParameterGroupInput) ResetClusterParameterGroupRequest
}

// Send marshals and sends the ResetClusterParameterGroup API request.
func (r ResetClusterParameterGroupRequest) Send(ctx context.Context) (*ResetClusterParameterGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ResetClusterParameterGroupResponse{
		ResetClusterParameterGroupOutput: r.Request.Data.(*ResetClusterParameterGroupOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ResetClusterParameterGroupResponse is the response type for the
// ResetClusterParameterGroup API operation.
type ResetClusterParameterGroupResponse struct {
	*ResetClusterParameterGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ResetClusterParameterGroup request.
func (r *ResetClusterParameterGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
