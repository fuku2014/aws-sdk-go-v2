// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mturk

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateHITWithHITTypeInput struct {
	_ struct{} `type:"structure"`

	// The Assignment-level Review Policy applies to the assignments under the HIT.
	// You can specify for Mechanical Turk to take various actions based on the
	// policy.
	AssignmentReviewPolicy *ReviewPolicy `type:"structure"`

	// The HITLayoutId allows you to use a pre-existing HIT design with placeholder
	// values and create an additional HIT by providing those values as HITLayoutParameters.
	//
	// Constraints: Either a Question parameter or a HITLayoutId parameter must
	// be provided.
	HITLayoutId *string `min:"1" type:"string"`

	// If the HITLayoutId is provided, any placeholder values must be filled in
	// with values using the HITLayoutParameter structure. For more information,
	// see HITLayout.
	HITLayoutParameters []HITLayoutParameter `type:"list"`

	// The HIT-level Review Policy applies to the HIT. You can specify for Mechanical
	// Turk to take various actions based on the policy.
	HITReviewPolicy *ReviewPolicy `type:"structure"`

	// The HIT type ID you want to create this HIT with.
	//
	// HITTypeId is a required field
	HITTypeId *string `min:"1" type:"string" required:"true"`

	// An amount of time, in seconds, after which the HIT is no longer available
	// for users to accept. After the lifetime of the HIT elapses, the HIT no longer
	// appears in HIT searches, even if not all of the assignments for the HIT have
	// been accepted.
	//
	// LifetimeInSeconds is a required field
	LifetimeInSeconds *int64 `type:"long" required:"true"`

	// The number of times the HIT can be accepted and completed before the HIT
	// becomes unavailable.
	MaxAssignments *int64 `type:"integer"`

	// The data the person completing the HIT uses to produce the results.
	//
	// Constraints: Must be a QuestionForm data structure, an ExternalQuestion data
	// structure, or an HTMLQuestion data structure. The XML question data must
	// not be larger than 64 kilobytes (65,535 bytes) in size, including whitespace.
	//
	// Either a Question parameter or a HITLayoutId parameter must be provided.
	Question *string `type:"string"`

	// An arbitrary data field. The RequesterAnnotation parameter lets your application
	// attach arbitrary data to the HIT for tracking purposes. For example, this
	// parameter could be an identifier internal to the Requester's application
	// that corresponds with the HIT.
	//
	// The RequesterAnnotation parameter for a HIT is only visible to the Requester
	// who created the HIT. It is not shown to the Worker, or any other Requester.
	//
	// The RequesterAnnotation parameter may be different for each HIT you submit.
	// It does not affect how your HITs are grouped.
	RequesterAnnotation *string `type:"string"`

	// A unique identifier for this request which allows you to retry the call on
	// error without creating duplicate HITs. This is useful in cases such as network
	// timeouts where it is unclear whether or not the call succeeded on the server.
	// If the HIT already exists in the system from a previous call using the same
	// UniqueRequestToken, subsequent calls will return a AWS.MechanicalTurk.HitAlreadyExists
	// error with a message containing the HITId.
	//
	// Note: It is your responsibility to ensure uniqueness of the token. The unique
	// token expires after 24 hours. Subsequent calls using the same UniqueRequestToken
	// made after the 24 hour limit could create duplicate HITs.
	UniqueRequestToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s CreateHITWithHITTypeInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateHITWithHITTypeInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateHITWithHITTypeInput"}
	if s.HITLayoutId != nil && len(*s.HITLayoutId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("HITLayoutId", 1))
	}

	if s.HITTypeId == nil {
		invalidParams.Add(aws.NewErrParamRequired("HITTypeId"))
	}
	if s.HITTypeId != nil && len(*s.HITTypeId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("HITTypeId", 1))
	}

	if s.LifetimeInSeconds == nil {
		invalidParams.Add(aws.NewErrParamRequired("LifetimeInSeconds"))
	}
	if s.UniqueRequestToken != nil && len(*s.UniqueRequestToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("UniqueRequestToken", 1))
	}
	if s.AssignmentReviewPolicy != nil {
		if err := s.AssignmentReviewPolicy.Validate(); err != nil {
			invalidParams.AddNested("AssignmentReviewPolicy", err.(aws.ErrInvalidParams))
		}
	}
	if s.HITLayoutParameters != nil {
		for i, v := range s.HITLayoutParameters {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "HITLayoutParameters", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.HITReviewPolicy != nil {
		if err := s.HITReviewPolicy.Validate(); err != nil {
			invalidParams.AddNested("HITReviewPolicy", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateHITWithHITTypeOutput struct {
	_ struct{} `type:"structure"`

	// Contains the newly created HIT data. For a description of the HIT data structure
	// as it appears in responses, see the HIT Data Structure documentation.
	HIT *HIT `type:"structure"`
}

// String returns the string representation
func (s CreateHITWithHITTypeOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateHITWithHITType = "CreateHITWithHITType"

// CreateHITWithHITTypeRequest returns a request value for making API operation for
// Amazon Mechanical Turk.
//
// The CreateHITWithHITType operation creates a new Human Intelligence Task
// (HIT) using an existing HITTypeID generated by the CreateHITType operation.
//
// This is an alternative way to create HITs from the CreateHIT operation. This
// is the recommended best practice for Requesters who are creating large numbers
// of HITs.
//
// CreateHITWithHITType also supports several ways to provide question data:
// by providing a value for the Question parameter that fully specifies the
// contents of the HIT, or by providing a HitLayoutId and associated HitLayoutParameters.
//
// If a HIT is created with 10 or more maximum assignments, there is an additional
// fee. For more information, see Amazon Mechanical Turk Pricing (https://requester.mturk.com/pricing).
//
//    // Example sending a request using CreateHITWithHITTypeRequest.
//    req := client.CreateHITWithHITTypeRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mturk-requester-2017-01-17/CreateHITWithHITType
func (c *Client) CreateHITWithHITTypeRequest(input *CreateHITWithHITTypeInput) CreateHITWithHITTypeRequest {
	op := &aws.Operation{
		Name:       opCreateHITWithHITType,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateHITWithHITTypeInput{}
	}

	req := c.newRequest(op, input, &CreateHITWithHITTypeOutput{})

	return CreateHITWithHITTypeRequest{Request: req, Input: input, Copy: c.CreateHITWithHITTypeRequest}
}

// CreateHITWithHITTypeRequest is the request type for the
// CreateHITWithHITType API operation.
type CreateHITWithHITTypeRequest struct {
	*aws.Request
	Input *CreateHITWithHITTypeInput
	Copy  func(*CreateHITWithHITTypeInput) CreateHITWithHITTypeRequest
}

// Send marshals and sends the CreateHITWithHITType API request.
func (r CreateHITWithHITTypeRequest) Send(ctx context.Context) (*CreateHITWithHITTypeResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateHITWithHITTypeResponse{
		CreateHITWithHITTypeOutput: r.Request.Data.(*CreateHITWithHITTypeOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateHITWithHITTypeResponse is the response type for the
// CreateHITWithHITType API operation.
type CreateHITWithHITTypeResponse struct {
	*CreateHITWithHITTypeOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateHITWithHITType request.
func (r *CreateHITWithHITTypeResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
