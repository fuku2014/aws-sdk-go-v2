// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codecommit

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/codecommit-2015-04-13/UpdatePullRequestDescriptionInput
type UpdatePullRequestDescriptionInput struct {
	_ struct{} `type:"structure"`

	// The updated content of the description for the pull request. This content
	// will replace the existing description.
	//
	// Description is a required field
	Description *string `locationName:"description" type:"string" required:"true"`

	// The system-generated ID of the pull request. To get this ID, use ListPullRequests.
	//
	// PullRequestId is a required field
	PullRequestId *string `locationName:"pullRequestId" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdatePullRequestDescriptionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdatePullRequestDescriptionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdatePullRequestDescriptionInput"}

	if s.Description == nil {
		invalidParams.Add(aws.NewErrParamRequired("Description"))
	}

	if s.PullRequestId == nil {
		invalidParams.Add(aws.NewErrParamRequired("PullRequestId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/codecommit-2015-04-13/UpdatePullRequestDescriptionOutput
type UpdatePullRequestDescriptionOutput struct {
	_ struct{} `type:"structure"`

	// Information about the updated pull request.
	//
	// PullRequest is a required field
	PullRequest *PullRequest `locationName:"pullRequest" type:"structure" required:"true"`
}

// String returns the string representation
func (s UpdatePullRequestDescriptionOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdatePullRequestDescription = "UpdatePullRequestDescription"

// UpdatePullRequestDescriptionRequest returns a request value for making API operation for
// AWS CodeCommit.
//
// Replaces the contents of the description of a pull request.
//
//    // Example sending a request using UpdatePullRequestDescriptionRequest.
//    req := client.UpdatePullRequestDescriptionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codecommit-2015-04-13/UpdatePullRequestDescription
func (c *Client) UpdatePullRequestDescriptionRequest(input *UpdatePullRequestDescriptionInput) UpdatePullRequestDescriptionRequest {
	op := &aws.Operation{
		Name:       opUpdatePullRequestDescription,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdatePullRequestDescriptionInput{}
	}

	req := c.newRequest(op, input, &UpdatePullRequestDescriptionOutput{})
	return UpdatePullRequestDescriptionRequest{Request: req, Input: input, Copy: c.UpdatePullRequestDescriptionRequest}
}

// UpdatePullRequestDescriptionRequest is the request type for the
// UpdatePullRequestDescription API operation.
type UpdatePullRequestDescriptionRequest struct {
	*aws.Request
	Input *UpdatePullRequestDescriptionInput
	Copy  func(*UpdatePullRequestDescriptionInput) UpdatePullRequestDescriptionRequest
}

// Send marshals and sends the UpdatePullRequestDescription API request.
func (r UpdatePullRequestDescriptionRequest) Send(ctx context.Context) (*UpdatePullRequestDescriptionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdatePullRequestDescriptionResponse{
		UpdatePullRequestDescriptionOutput: r.Request.Data.(*UpdatePullRequestDescriptionOutput),
		response:                           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdatePullRequestDescriptionResponse is the response type for the
// UpdatePullRequestDescription API operation.
type UpdatePullRequestDescriptionResponse struct {
	*UpdatePullRequestDescriptionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdatePullRequestDescription request.
func (r *UpdatePullRequestDescriptionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}