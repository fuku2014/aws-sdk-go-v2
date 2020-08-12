// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package costandusagereportserviceiface provides an interface to enable mocking the AWS Cost and Usage Report Service service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package costandusagereportserviceiface

import (
	"github.com/aws/aws-sdk-go-v2/service/costandusagereportservice"
)

// ClientAPI provides an interface to enable mocking the
// costandusagereportservice.Client methods. This make unit testing your code that
// calls out to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS Cost and Usage Report Service.
//    func myFunc(svc costandusagereportserviceiface.ClientAPI) bool {
//        // Make svc.DeleteReportDefinition request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := costandusagereportservice.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockClientClient struct {
//        costandusagereportserviceiface.ClientPI
//    }
//    func (m *mockClientClient) DeleteReportDefinition(input *costandusagereportservice.DeleteReportDefinitionInput) (*costandusagereportservice.DeleteReportDefinitionOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockClientClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type ClientAPI interface {
	DeleteReportDefinitionRequest(*costandusagereportservice.DeleteReportDefinitionInput) costandusagereportservice.DeleteReportDefinitionRequest

	DescribeReportDefinitionsRequest(*costandusagereportservice.DescribeReportDefinitionsInput) costandusagereportservice.DescribeReportDefinitionsRequest

	ModifyReportDefinitionRequest(*costandusagereportservice.ModifyReportDefinitionInput) costandusagereportservice.ModifyReportDefinitionRequest

	PutReportDefinitionRequest(*costandusagereportservice.PutReportDefinitionInput) costandusagereportservice.PutReportDefinitionRequest
}

var _ ClientAPI = (*costandusagereportservice.Client)(nil)