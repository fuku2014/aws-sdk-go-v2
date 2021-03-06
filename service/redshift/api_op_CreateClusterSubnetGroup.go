// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package redshift

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateClusterSubnetGroupInput struct {
	_ struct{} `type:"structure"`

	// The name for the subnet group. Amazon Redshift stores the value as a lowercase
	// string.
	//
	// Constraints:
	//
	//    * Must contain no more than 255 alphanumeric characters or hyphens.
	//
	//    * Must not be "Default".
	//
	//    * Must be unique for all subnet groups that are created by your AWS account.
	//
	// Example: examplesubnetgroup
	//
	// ClusterSubnetGroupName is a required field
	ClusterSubnetGroupName *string `type:"string" required:"true"`

	// A description for the subnet group.
	//
	// Description is a required field
	Description *string `type:"string" required:"true"`

	// An array of VPC subnet IDs. A maximum of 20 subnets can be modified in a
	// single request.
	//
	// SubnetIds is a required field
	SubnetIds []string `locationNameList:"SubnetIdentifier" type:"list" required:"true"`

	// A list of tag instances.
	Tags []Tag `locationNameList:"Tag" type:"list"`
}

// String returns the string representation
func (s CreateClusterSubnetGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateClusterSubnetGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateClusterSubnetGroupInput"}

	if s.ClusterSubnetGroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ClusterSubnetGroupName"))
	}

	if s.Description == nil {
		invalidParams.Add(aws.NewErrParamRequired("Description"))
	}

	if s.SubnetIds == nil {
		invalidParams.Add(aws.NewErrParamRequired("SubnetIds"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateClusterSubnetGroupOutput struct {
	_ struct{} `type:"structure"`

	// Describes a subnet group.
	ClusterSubnetGroup *ClusterSubnetGroup `type:"structure"`
}

// String returns the string representation
func (s CreateClusterSubnetGroupOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateClusterSubnetGroup = "CreateClusterSubnetGroup"

// CreateClusterSubnetGroupRequest returns a request value for making API operation for
// Amazon Redshift.
//
// Creates a new Amazon Redshift subnet group. You must provide a list of one
// or more subnets in your existing Amazon Virtual Private Cloud (Amazon VPC)
// when creating Amazon Redshift subnet group.
//
// For information about subnet groups, go to Amazon Redshift Cluster Subnet
// Groups (https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-cluster-subnet-groups.html)
// in the Amazon Redshift Cluster Management Guide.
//
//    // Example sending a request using CreateClusterSubnetGroupRequest.
//    req := client.CreateClusterSubnetGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/redshift-2012-12-01/CreateClusterSubnetGroup
func (c *Client) CreateClusterSubnetGroupRequest(input *CreateClusterSubnetGroupInput) CreateClusterSubnetGroupRequest {
	op := &aws.Operation{
		Name:       opCreateClusterSubnetGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateClusterSubnetGroupInput{}
	}

	req := c.newRequest(op, input, &CreateClusterSubnetGroupOutput{})

	return CreateClusterSubnetGroupRequest{Request: req, Input: input, Copy: c.CreateClusterSubnetGroupRequest}
}

// CreateClusterSubnetGroupRequest is the request type for the
// CreateClusterSubnetGroup API operation.
type CreateClusterSubnetGroupRequest struct {
	*aws.Request
	Input *CreateClusterSubnetGroupInput
	Copy  func(*CreateClusterSubnetGroupInput) CreateClusterSubnetGroupRequest
}

// Send marshals and sends the CreateClusterSubnetGroup API request.
func (r CreateClusterSubnetGroupRequest) Send(ctx context.Context) (*CreateClusterSubnetGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateClusterSubnetGroupResponse{
		CreateClusterSubnetGroupOutput: r.Request.Data.(*CreateClusterSubnetGroupOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateClusterSubnetGroupResponse is the response type for the
// CreateClusterSubnetGroup API operation.
type CreateClusterSubnetGroupResponse struct {
	*CreateClusterSubnetGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateClusterSubnetGroup request.
func (r *CreateClusterSubnetGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
