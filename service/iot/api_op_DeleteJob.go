// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
)

type DeleteJobInput struct {
	_ struct{} `type:"structure"`

	// (Optional) When true, you can delete a job which is "IN_PROGRESS". Otherwise,
	// you can only delete a job which is in a terminal state ("COMPLETED" or "CANCELED")
	// or an exception will occur. The default is false.
	//
	// Deleting a job which is "IN_PROGRESS", will cause a device which is executing
	// the job to be unable to access job information or update the job execution
	// status. Use caution and ensure that each device executing a job which is
	// deleted is able to recover to a valid state.
	Force *bool `location:"querystring" locationName:"force" type:"boolean"`

	// The ID of the job to be deleted.
	//
	// After a job deletion is completed, you may reuse this jobId when you create
	// a new job. However, this is not recommended, and you must ensure that your
	// devices are not using the jobId to refer to the deleted job.
	//
	// JobId is a required field
	JobId *string `location:"uri" locationName:"jobId" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteJobInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteJobInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteJobInput"}

	if s.JobId == nil {
		invalidParams.Add(aws.NewErrParamRequired("JobId"))
	}
	if s.JobId != nil && len(*s.JobId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("JobId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteJobInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.JobId != nil {
		v := *s.JobId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "jobId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Force != nil {
		v := *s.Force

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "force", protocol.BoolValue(v), metadata)
	}
	return nil
}

type DeleteJobOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteJobOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteJobOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDeleteJob = "DeleteJob"

// DeleteJobRequest returns a request value for making API operation for
// AWS IoT.
//
// Deletes a job and its related job executions.
//
// Deleting a job may take time, depending on the number of job executions created
// for the job and various other factors. While the job is being deleted, the
// status of the job will be shown as "DELETION_IN_PROGRESS". Attempting to
// delete or cancel a job whose status is already "DELETION_IN_PROGRESS" will
// result in an error.
//
// Only 10 jobs may have status "DELETION_IN_PROGRESS" at the same time, or
// a LimitExceededException will occur.
//
//    // Example sending a request using DeleteJobRequest.
//    req := client.DeleteJobRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) DeleteJobRequest(input *DeleteJobInput) DeleteJobRequest {
	op := &aws.Operation{
		Name:       opDeleteJob,
		HTTPMethod: "DELETE",
		HTTPPath:   "/jobs/{jobId}",
	}

	if input == nil {
		input = &DeleteJobInput{}
	}

	req := c.newRequest(op, input, &DeleteJobOutput{})
	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return DeleteJobRequest{Request: req, Input: input, Copy: c.DeleteJobRequest}
}

// DeleteJobRequest is the request type for the
// DeleteJob API operation.
type DeleteJobRequest struct {
	*aws.Request
	Input *DeleteJobInput
	Copy  func(*DeleteJobInput) DeleteJobRequest
}

// Send marshals and sends the DeleteJob API request.
func (r DeleteJobRequest) Send(ctx context.Context) (*DeleteJobResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteJobResponse{
		DeleteJobOutput: r.Request.Data.(*DeleteJobOutput),
		response:        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteJobResponse is the response type for the
// DeleteJob API operation.
type DeleteJobResponse struct {
	*DeleteJobOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteJob request.
func (r *DeleteJobResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
