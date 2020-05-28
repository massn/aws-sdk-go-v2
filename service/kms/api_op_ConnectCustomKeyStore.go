// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kms

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ConnectCustomKeyStoreInput struct {
	_ struct{} `type:"structure"`

	// Enter the key store ID of the custom key store that you want to connect.
	// To find the ID of a custom key store, use the DescribeCustomKeyStores operation.
	//
	// CustomKeyStoreId is a required field
	CustomKeyStoreId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s ConnectCustomKeyStoreInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ConnectCustomKeyStoreInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ConnectCustomKeyStoreInput"}

	if s.CustomKeyStoreId == nil {
		invalidParams.Add(aws.NewErrParamRequired("CustomKeyStoreId"))
	}
	if s.CustomKeyStoreId != nil && len(*s.CustomKeyStoreId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("CustomKeyStoreId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ConnectCustomKeyStoreOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s ConnectCustomKeyStoreOutput) String() string {
	return awsutil.Prettify(s)
}

const opConnectCustomKeyStore = "ConnectCustomKeyStore"

// ConnectCustomKeyStoreRequest returns a request value for making API operation for
// AWS Key Management Service.
//
// Connects or reconnects a custom key store (https://docs.aws.amazon.com/kms/latest/developerguide/custom-key-store-overview.html)
// to its associated AWS CloudHSM cluster.
//
// The custom key store must be connected before you can create customer master
// keys (CMKs) in the key store or use the CMKs it contains. You can disconnect
// and reconnect a custom key store at any time.
//
// To connect a custom key store, its associated AWS CloudHSM cluster must have
// at least one active HSM. To get the number of active HSMs in a cluster, use
// the DescribeClusters (https://docs.aws.amazon.com/cloudhsm/latest/APIReference/API_DescribeClusters.html)
// operation. To add HSMs to the cluster, use the CreateHsm (https://docs.aws.amazon.com/cloudhsm/latest/APIReference/API_CreateHsm.html)
// operation. Also, the kmsuser crypto user (https://docs.aws.amazon.com/kms/latest/developerguide/key-store-concepts.html#concept-kmsuser)
// (CU) must not be logged into the cluster. This prevents AWS KMS from using
// this account to log in.
//
// The connection process can take an extended amount of time to complete; up
// to 20 minutes. This operation starts the connection process, but it does
// not wait for it to complete. When it succeeds, this operation quickly returns
// an HTTP 200 response and a JSON object with no properties. However, this
// response does not indicate that the custom key store is connected. To get
// the connection state of the custom key store, use the DescribeCustomKeyStores
// operation.
//
// During the connection process, AWS KMS finds the AWS CloudHSM cluster that
// is associated with the custom key store, creates the connection infrastructure,
// connects to the cluster, logs into the AWS CloudHSM client as the kmsuser
// CU, and rotates its password.
//
// The ConnectCustomKeyStore operation might fail for various reasons. To find
// the reason, use the DescribeCustomKeyStores operation and see the ConnectionErrorCode
// in the response. For help interpreting the ConnectionErrorCode, see CustomKeyStoresListEntry.
//
// To fix the failure, use the DisconnectCustomKeyStore operation to disconnect
// the custom key store, correct the error, use the UpdateCustomKeyStore operation
// if necessary, and then use ConnectCustomKeyStore again.
//
// If you are having trouble connecting or disconnecting a custom key store,
// see Troubleshooting a Custom Key Store (https://docs.aws.amazon.com/kms/latest/developerguide/fix-keystore.html)
// in the AWS Key Management Service Developer Guide.
//
//    // Example sending a request using ConnectCustomKeyStoreRequest.
//    req := client.ConnectCustomKeyStoreRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kms-2014-11-01/ConnectCustomKeyStore
func (c *Client) ConnectCustomKeyStoreRequest(input *ConnectCustomKeyStoreInput) ConnectCustomKeyStoreRequest {
	op := &aws.Operation{
		Name:       opConnectCustomKeyStore,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ConnectCustomKeyStoreInput{}
	}

	req := c.newRequest(op, input, &ConnectCustomKeyStoreOutput{})

	return ConnectCustomKeyStoreRequest{Request: req, Input: input, Copy: c.ConnectCustomKeyStoreRequest}
}

// ConnectCustomKeyStoreRequest is the request type for the
// ConnectCustomKeyStore API operation.
type ConnectCustomKeyStoreRequest struct {
	*aws.Request
	Input *ConnectCustomKeyStoreInput
	Copy  func(*ConnectCustomKeyStoreInput) ConnectCustomKeyStoreRequest
}

// Send marshals and sends the ConnectCustomKeyStore API request.
func (r ConnectCustomKeyStoreRequest) Send(ctx context.Context) (*ConnectCustomKeyStoreResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ConnectCustomKeyStoreResponse{
		ConnectCustomKeyStoreOutput: r.Request.Data.(*ConnectCustomKeyStoreOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ConnectCustomKeyStoreResponse is the response type for the
// ConnectCustomKeyStore API operation.
type ConnectCustomKeyStoreResponse struct {
	*ConnectCustomKeyStoreOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ConnectCustomKeyStore request.
func (r *ConnectCustomKeyStoreResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
