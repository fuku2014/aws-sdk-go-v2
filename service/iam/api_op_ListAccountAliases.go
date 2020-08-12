// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListAccountAliasesInput struct {
	_ struct{} `type:"structure"`

	// Use this parameter only when paginating results and only after you receive
	// a response indicating that the results are truncated. Set it to the value
	// of the Marker element in the response that you received to indicate where
	// the next call should start.
	Marker *string `min:"1" type:"string"`

	// Use this only when paginating results to indicate the maximum number of items
	// you want in the response. If additional items exist beyond the maximum you
	// specify, the IsTruncated response element is true.
	//
	// If you do not include this parameter, the number of items defaults to 100.
	// Note that IAM might return fewer results, even when there are more results
	// available. In that case, the IsTruncated response element returns true, and
	// Marker contains a value to include in the subsequent call that tells the
	// service where to continue from.
	MaxItems *int64 `min:"1" type:"integer"`
}

// String returns the string representation
func (s ListAccountAliasesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListAccountAliasesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListAccountAliasesInput"}
	if s.Marker != nil && len(*s.Marker) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Marker", 1))
	}
	if s.MaxItems != nil && *s.MaxItems < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxItems", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains the response to a successful ListAccountAliases request.
type ListAccountAliasesOutput struct {
	_ struct{} `type:"structure"`

	// A list of aliases associated with the account. AWS supports only one alias
	// per account.
	//
	// AccountAliases is a required field
	AccountAliases []string `type:"list" required:"true"`

	// A flag that indicates whether there are more items to return. If your results
	// were truncated, you can make a subsequent pagination request using the Marker
	// request parameter to retrieve more items. Note that IAM might return fewer
	// than the MaxItems number of results even when there are more results available.
	// We recommend that you check IsTruncated after every call to ensure that you
	// receive all your results.
	IsTruncated *bool `type:"boolean"`

	// When IsTruncated is true, this element is present and contains the value
	// to use for the Marker parameter in a subsequent pagination request.
	Marker *string `type:"string"`
}

// String returns the string representation
func (s ListAccountAliasesOutput) String() string {
	return awsutil.Prettify(s)
}

const opListAccountAliases = "ListAccountAliases"

// ListAccountAliasesRequest returns a request value for making API operation for
// AWS Identity and Access Management.
//
// Lists the account alias associated with the AWS account (Note: you can have
// only one). For information about using an AWS account alias, see Using an
// Alias for Your AWS Account ID (https://docs.aws.amazon.com/IAM/latest/UserGuide/AccountAlias.html)
// in the IAM User Guide.
//
//    // Example sending a request using ListAccountAliasesRequest.
//    req := client.ListAccountAliasesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/ListAccountAliases
func (c *Client) ListAccountAliasesRequest(input *ListAccountAliasesInput) ListAccountAliasesRequest {
	op := &aws.Operation{
		Name:       opListAccountAliases,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"Marker"},
			OutputTokens:    []string{"Marker"},
			LimitToken:      "MaxItems",
			TruncationToken: "IsTruncated",
		},
	}

	if input == nil {
		input = &ListAccountAliasesInput{}
	}

	req := c.newRequest(op, input, &ListAccountAliasesOutput{})

	return ListAccountAliasesRequest{Request: req, Input: input, Copy: c.ListAccountAliasesRequest}
}

// ListAccountAliasesRequest is the request type for the
// ListAccountAliases API operation.
type ListAccountAliasesRequest struct {
	*aws.Request
	Input *ListAccountAliasesInput
	Copy  func(*ListAccountAliasesInput) ListAccountAliasesRequest
}

// Send marshals and sends the ListAccountAliases API request.
func (r ListAccountAliasesRequest) Send(ctx context.Context) (*ListAccountAliasesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListAccountAliasesResponse{
		ListAccountAliasesOutput: r.Request.Data.(*ListAccountAliasesOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListAccountAliasesRequestPaginator returns a paginator for ListAccountAliases.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListAccountAliasesRequest(input)
//   p := iam.NewListAccountAliasesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListAccountAliasesPaginator(req ListAccountAliasesRequest) ListAccountAliasesPaginator {
	return ListAccountAliasesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListAccountAliasesInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// ListAccountAliasesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListAccountAliasesPaginator struct {
	aws.Pager
}

func (p *ListAccountAliasesPaginator) CurrentPage() *ListAccountAliasesOutput {
	return p.Pager.CurrentPage().(*ListAccountAliasesOutput)
}

// ListAccountAliasesResponse is the response type for the
// ListAccountAliases API operation.
type ListAccountAliasesResponse struct {
	*ListAccountAliasesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListAccountAliases request.
func (r *ListAccountAliasesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}