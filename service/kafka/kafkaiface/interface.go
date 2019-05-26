// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package kafkaiface provides an interface to enable mocking the Managed Streaming for Kafka service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package kafkaiface

import (
	"github.com/aws/aws-sdk-go-v2/service/kafka"
)

// ClientAPI provides an interface to enable mocking the
// kafka.Client methods. This make unit testing your code that
// calls out to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Kafka.
//    func myFunc(svc kafkaiface.ClientAPI) bool {
//        // Make svc.CreateCluster request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := kafka.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockClientClient struct {
//        kafkaiface.ClientPI
//    }
//    func (m *mockClientClient) CreateCluster(input *kafka.CreateClusterInput) (*kafka.CreateClusterOutput, error) {
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
	CreateClusterRequest(*kafka.CreateClusterInput) kafka.CreateClusterRequest

	CreateConfigurationRequest(*kafka.CreateConfigurationInput) kafka.CreateConfigurationRequest

	DeleteClusterRequest(*kafka.DeleteClusterInput) kafka.DeleteClusterRequest

	DescribeClusterRequest(*kafka.DescribeClusterInput) kafka.DescribeClusterRequest

	DescribeConfigurationRequest(*kafka.DescribeConfigurationInput) kafka.DescribeConfigurationRequest

	DescribeConfigurationRevisionRequest(*kafka.DescribeConfigurationRevisionInput) kafka.DescribeConfigurationRevisionRequest

	GetBootstrapBrokersRequest(*kafka.GetBootstrapBrokersInput) kafka.GetBootstrapBrokersRequest

	ListClustersRequest(*kafka.ListClustersInput) kafka.ListClustersRequest

	ListConfigurationsRequest(*kafka.ListConfigurationsInput) kafka.ListConfigurationsRequest

	ListNodesRequest(*kafka.ListNodesInput) kafka.ListNodesRequest

	ListTagsForResourceRequest(*kafka.ListTagsForResourceInput) kafka.ListTagsForResourceRequest

	TagResourceRequest(*kafka.TagResourceInput) kafka.TagResourceRequest

	UntagResourceRequest(*kafka.UntagResourceInput) kafka.UntagResourceRequest
}

var _ ClientAPI = (*kafka.Client)(nil)