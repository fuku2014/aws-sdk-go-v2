// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package secretsmanager

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteResourcePolicyInput struct {
	_ struct{} `type:"structure"`

	// Specifies the secret that you want to delete the attached resource-based
	// policy for. You can specify either the Amazon Resource Name (ARN) or the
	// friendly name of the secret.
	//
	// If you specify an ARN, we generally recommend that you specify a complete
	// ARN. You can specify a partial ARN too—for example, if you don’t include
	// the final hyphen and six random characters that Secrets Manager adds at the
	// end of the ARN when you created the secret. A partial ARN match can work
	// as long as it uniquely matches only one secret. However, if your secret has
	// a name that ends in a hyphen followed by six characters (before Secrets Manager
	// adds the hyphen and six characters to the ARN) and you try to use that as
	// a partial ARN, then those characters cause Secrets Manager to assume that
	// you’re specifying a complete ARN. This confusion can cause unexpected results.
	// To avoid this situation, we recommend that you don’t create secret names
	// ending with a hyphen followed by six characters.
	//
	// If you specify an incomplete ARN without the random suffix, and instead provide
	// the 'friendly name', you must not include the random suffix. If you do include
	// the random suffix added by Secrets Manager, you receive either a ResourceNotFoundException
	// or an AccessDeniedException error, depending on your permissions.
	//
	// SecretId is a required field
	SecretId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteResourcePolicyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteResourcePolicyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteResourcePolicyInput"}

	if s.SecretId == nil {
		invalidParams.Add(aws.NewErrParamRequired("SecretId"))
	}
	if s.SecretId != nil && len(*s.SecretId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("SecretId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteResourcePolicyOutput struct {
	_ struct{} `type:"structure"`

	// The ARN of the secret that the resource-based policy was deleted for.
	ARN *string `min:"20" type:"string"`

	// The friendly name of the secret that the resource-based policy was deleted
	// for.
	Name *string `min:"1" type:"string"`
}

// String returns the string representation
func (s DeleteResourcePolicyOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteResourcePolicy = "DeleteResourcePolicy"

// DeleteResourcePolicyRequest returns a request value for making API operation for
// AWS Secrets Manager.
//
// Deletes the resource-based permission policy attached to the secret.
//
// Minimum permissions
//
// To run this command, you must have the following permissions:
//
//    * secretsmanager:DeleteResourcePolicy
//
// Related operations
//
//    * To attach a resource policy to a secret, use PutResourcePolicy.
//
//    * To retrieve the current resource-based policy that's attached to a secret,
//    use GetResourcePolicy.
//
//    * To list all of the currently available secrets, use ListSecrets.
//
//    // Example sending a request using DeleteResourcePolicyRequest.
//    req := client.DeleteResourcePolicyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/secretsmanager-2017-10-17/DeleteResourcePolicy
func (c *Client) DeleteResourcePolicyRequest(input *DeleteResourcePolicyInput) DeleteResourcePolicyRequest {
	op := &aws.Operation{
		Name:       opDeleteResourcePolicy,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteResourcePolicyInput{}
	}

	req := c.newRequest(op, input, &DeleteResourcePolicyOutput{})

	return DeleteResourcePolicyRequest{Request: req, Input: input, Copy: c.DeleteResourcePolicyRequest}
}

// DeleteResourcePolicyRequest is the request type for the
// DeleteResourcePolicy API operation.
type DeleteResourcePolicyRequest struct {
	*aws.Request
	Input *DeleteResourcePolicyInput
	Copy  func(*DeleteResourcePolicyInput) DeleteResourcePolicyRequest
}

// Send marshals and sends the DeleteResourcePolicy API request.
func (r DeleteResourcePolicyRequest) Send(ctx context.Context) (*DeleteResourcePolicyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteResourcePolicyResponse{
		DeleteResourcePolicyOutput: r.Request.Data.(*DeleteResourcePolicyOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteResourcePolicyResponse is the response type for the
// DeleteResourcePolicy API operation.
type DeleteResourcePolicyResponse struct {
	*DeleteResourcePolicyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteResourcePolicy request.
func (r *DeleteResourcePolicyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
