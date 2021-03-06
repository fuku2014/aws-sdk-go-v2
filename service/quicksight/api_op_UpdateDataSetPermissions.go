// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package quicksight

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type UpdateDataSetPermissionsInput struct {
	_ struct{} `type:"structure"`

	// The AWS account ID.
	//
	// AwsAccountId is a required field
	AwsAccountId *string `location:"uri" locationName:"AwsAccountId" min:"12" type:"string" required:"true"`

	// The ID for the dataset whose permissions you want to update. This ID is unique
	// per AWS Region for each AWS account.
	//
	// DataSetId is a required field
	DataSetId *string `location:"uri" locationName:"DataSetId" type:"string" required:"true"`

	// The resource permissions that you want to grant to the dataset.
	GrantPermissions []ResourcePermission `min:"1" type:"list"`

	// The resource permissions that you want to revoke from the dataset.
	RevokePermissions []ResourcePermission `min:"1" type:"list"`
}

// String returns the string representation
func (s UpdateDataSetPermissionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateDataSetPermissionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateDataSetPermissionsInput"}

	if s.AwsAccountId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AwsAccountId"))
	}
	if s.AwsAccountId != nil && len(*s.AwsAccountId) < 12 {
		invalidParams.Add(aws.NewErrParamMinLen("AwsAccountId", 12))
	}

	if s.DataSetId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DataSetId"))
	}
	if s.GrantPermissions != nil && len(s.GrantPermissions) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("GrantPermissions", 1))
	}
	if s.RevokePermissions != nil && len(s.RevokePermissions) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("RevokePermissions", 1))
	}
	if s.GrantPermissions != nil {
		for i, v := range s.GrantPermissions {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "GrantPermissions", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.RevokePermissions != nil {
		for i, v := range s.RevokePermissions {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "RevokePermissions", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateDataSetPermissionsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.GrantPermissions != nil {
		v := s.GrantPermissions

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "GrantPermissions", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.RevokePermissions != nil {
		v := s.RevokePermissions

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "RevokePermissions", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.AwsAccountId != nil {
		v := *s.AwsAccountId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "AwsAccountId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.DataSetId != nil {
		v := *s.DataSetId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "DataSetId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type UpdateDataSetPermissionsOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the dataset.
	DataSetArn *string `type:"string"`

	// The ID for the dataset whose permissions you want to update. This ID is unique
	// per AWS Region for each AWS account.
	DataSetId *string `type:"string"`

	// The AWS request ID for this operation.
	RequestId *string `type:"string"`

	// The HTTP status of the request.
	Status *int64 `location:"statusCode" type:"integer"`
}

// String returns the string representation
func (s UpdateDataSetPermissionsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateDataSetPermissionsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.DataSetArn != nil {
		v := *s.DataSetArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "DataSetArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.DataSetId != nil {
		v := *s.DataSetId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "DataSetId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RequestId != nil {
		v := *s.RequestId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "RequestId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	// ignoring invalid encode state, StatusCode. Status
	return nil
}

const opUpdateDataSetPermissions = "UpdateDataSetPermissions"

// UpdateDataSetPermissionsRequest returns a request value for making API operation for
// Amazon QuickSight.
//
// Updates the permissions on a dataset.
//
// The permissions resource is arn:aws:quicksight:region:aws-account-id:dataset/data-set-id.
//
//    // Example sending a request using UpdateDataSetPermissionsRequest.
//    req := client.UpdateDataSetPermissionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/quicksight-2018-04-01/UpdateDataSetPermissions
func (c *Client) UpdateDataSetPermissionsRequest(input *UpdateDataSetPermissionsInput) UpdateDataSetPermissionsRequest {
	op := &aws.Operation{
		Name:       opUpdateDataSetPermissions,
		HTTPMethod: "POST",
		HTTPPath:   "/accounts/{AwsAccountId}/data-sets/{DataSetId}/permissions",
	}

	if input == nil {
		input = &UpdateDataSetPermissionsInput{}
	}

	req := c.newRequest(op, input, &UpdateDataSetPermissionsOutput{})

	return UpdateDataSetPermissionsRequest{Request: req, Input: input, Copy: c.UpdateDataSetPermissionsRequest}
}

// UpdateDataSetPermissionsRequest is the request type for the
// UpdateDataSetPermissions API operation.
type UpdateDataSetPermissionsRequest struct {
	*aws.Request
	Input *UpdateDataSetPermissionsInput
	Copy  func(*UpdateDataSetPermissionsInput) UpdateDataSetPermissionsRequest
}

// Send marshals and sends the UpdateDataSetPermissions API request.
func (r UpdateDataSetPermissionsRequest) Send(ctx context.Context) (*UpdateDataSetPermissionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateDataSetPermissionsResponse{
		UpdateDataSetPermissionsOutput: r.Request.Data.(*UpdateDataSetPermissionsOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateDataSetPermissionsResponse is the response type for the
// UpdateDataSetPermissions API operation.
type UpdateDataSetPermissionsResponse struct {
	*UpdateDataSetPermissionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateDataSetPermissions request.
func (r *UpdateDataSetPermissionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
