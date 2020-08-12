// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticache

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type RebalanceSlotsInGlobalReplicationGroupInput struct {
	_ struct{} `type:"structure"`

	// If True, redistribution is applied immediately.
	//
	// ApplyImmediately is a required field
	ApplyImmediately *bool `type:"boolean" required:"true"`

	// The name of the Global Datastore
	//
	// GlobalReplicationGroupId is a required field
	GlobalReplicationGroupId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s RebalanceSlotsInGlobalReplicationGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RebalanceSlotsInGlobalReplicationGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RebalanceSlotsInGlobalReplicationGroupInput"}

	if s.ApplyImmediately == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApplyImmediately"))
	}

	if s.GlobalReplicationGroupId == nil {
		invalidParams.Add(aws.NewErrParamRequired("GlobalReplicationGroupId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type RebalanceSlotsInGlobalReplicationGroupOutput struct {
	_ struct{} `type:"structure"`

	// Consists of a primary cluster that accepts writes and an associated secondary
	// cluster that resides in a different AWS region. The secondary cluster accepts
	// only reads. The primary cluster automatically replicates updates to the secondary
	// cluster.
	//
	//    * The GlobalReplicationGroupIdSuffix represents the name of the Global
	//    Datastore, which is what you use to associate a secondary cluster.
	GlobalReplicationGroup *GlobalReplicationGroup `type:"structure"`
}

// String returns the string representation
func (s RebalanceSlotsInGlobalReplicationGroupOutput) String() string {
	return awsutil.Prettify(s)
}

const opRebalanceSlotsInGlobalReplicationGroup = "RebalanceSlotsInGlobalReplicationGroup"

// RebalanceSlotsInGlobalReplicationGroupRequest returns a request value for making API operation for
// Amazon ElastiCache.
//
// Redistribute slots to ensure uniform distribution across existing shards
// in the cluster.
//
//    // Example sending a request using RebalanceSlotsInGlobalReplicationGroupRequest.
//    req := client.RebalanceSlotsInGlobalReplicationGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticache-2015-02-02/RebalanceSlotsInGlobalReplicationGroup
func (c *Client) RebalanceSlotsInGlobalReplicationGroupRequest(input *RebalanceSlotsInGlobalReplicationGroupInput) RebalanceSlotsInGlobalReplicationGroupRequest {
	op := &aws.Operation{
		Name:       opRebalanceSlotsInGlobalReplicationGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &RebalanceSlotsInGlobalReplicationGroupInput{}
	}

	req := c.newRequest(op, input, &RebalanceSlotsInGlobalReplicationGroupOutput{})

	return RebalanceSlotsInGlobalReplicationGroupRequest{Request: req, Input: input, Copy: c.RebalanceSlotsInGlobalReplicationGroupRequest}
}

// RebalanceSlotsInGlobalReplicationGroupRequest is the request type for the
// RebalanceSlotsInGlobalReplicationGroup API operation.
type RebalanceSlotsInGlobalReplicationGroupRequest struct {
	*aws.Request
	Input *RebalanceSlotsInGlobalReplicationGroupInput
	Copy  func(*RebalanceSlotsInGlobalReplicationGroupInput) RebalanceSlotsInGlobalReplicationGroupRequest
}

// Send marshals and sends the RebalanceSlotsInGlobalReplicationGroup API request.
func (r RebalanceSlotsInGlobalReplicationGroupRequest) Send(ctx context.Context) (*RebalanceSlotsInGlobalReplicationGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RebalanceSlotsInGlobalReplicationGroupResponse{
		RebalanceSlotsInGlobalReplicationGroupOutput: r.Request.Data.(*RebalanceSlotsInGlobalReplicationGroupOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RebalanceSlotsInGlobalReplicationGroupResponse is the response type for the
// RebalanceSlotsInGlobalReplicationGroup API operation.
type RebalanceSlotsInGlobalReplicationGroupResponse struct {
	*RebalanceSlotsInGlobalReplicationGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RebalanceSlotsInGlobalReplicationGroup request.
func (r *RebalanceSlotsInGlobalReplicationGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}