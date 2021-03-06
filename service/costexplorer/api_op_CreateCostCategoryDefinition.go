// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package costexplorer

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateCostCategoryDefinitionInput struct {
	_ struct{} `type:"structure"`

	// The unique name of the Cost Category.
	//
	// Name is a required field
	Name *string `min:"1" type:"string" required:"true"`

	// The rule schema version in this particular Cost Category.
	//
	// RuleVersion is a required field
	RuleVersion CostCategoryRuleVersion `type:"string" required:"true" enum:"true"`

	// The Cost Category rules used to categorize costs. For more information, see
	// CostCategoryRule (https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_CostCategoryRule.html).
	//
	// Rules is a required field
	Rules []CostCategoryRule `min:"1" type:"list" required:"true"`
}

// String returns the string representation
func (s CreateCostCategoryDefinitionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateCostCategoryDefinitionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateCostCategoryDefinitionInput"}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}
	if len(s.RuleVersion) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("RuleVersion"))
	}

	if s.Rules == nil {
		invalidParams.Add(aws.NewErrParamRequired("Rules"))
	}
	if s.Rules != nil && len(s.Rules) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Rules", 1))
	}
	if s.Rules != nil {
		for i, v := range s.Rules {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Rules", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateCostCategoryDefinitionOutput struct {
	_ struct{} `type:"structure"`

	// The unique identifier for your newly created Cost Category.
	CostCategoryArn *string `min:"20" type:"string"`

	// The Cost Category's effective start date.
	EffectiveStart *string `min:"20" type:"string"`
}

// String returns the string representation
func (s CreateCostCategoryDefinitionOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateCostCategoryDefinition = "CreateCostCategoryDefinition"

// CreateCostCategoryDefinitionRequest returns a request value for making API operation for
// AWS Cost Explorer Service.
//
// Creates a new Cost Category with the requested name and rules.
//
//    // Example sending a request using CreateCostCategoryDefinitionRequest.
//    req := client.CreateCostCategoryDefinitionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ce-2017-10-25/CreateCostCategoryDefinition
func (c *Client) CreateCostCategoryDefinitionRequest(input *CreateCostCategoryDefinitionInput) CreateCostCategoryDefinitionRequest {
	op := &aws.Operation{
		Name:       opCreateCostCategoryDefinition,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateCostCategoryDefinitionInput{}
	}

	req := c.newRequest(op, input, &CreateCostCategoryDefinitionOutput{})

	return CreateCostCategoryDefinitionRequest{Request: req, Input: input, Copy: c.CreateCostCategoryDefinitionRequest}
}

// CreateCostCategoryDefinitionRequest is the request type for the
// CreateCostCategoryDefinition API operation.
type CreateCostCategoryDefinitionRequest struct {
	*aws.Request
	Input *CreateCostCategoryDefinitionInput
	Copy  func(*CreateCostCategoryDefinitionInput) CreateCostCategoryDefinitionRequest
}

// Send marshals and sends the CreateCostCategoryDefinition API request.
func (r CreateCostCategoryDefinitionRequest) Send(ctx context.Context) (*CreateCostCategoryDefinitionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateCostCategoryDefinitionResponse{
		CreateCostCategoryDefinitionOutput: r.Request.Data.(*CreateCostCategoryDefinitionOutput),
		response:                           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateCostCategoryDefinitionResponse is the response type for the
// CreateCostCategoryDefinition API operation.
type CreateCostCategoryDefinitionResponse struct {
	*CreateCostCategoryDefinitionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateCostCategoryDefinition request.
func (r *CreateCostCategoryDefinitionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
