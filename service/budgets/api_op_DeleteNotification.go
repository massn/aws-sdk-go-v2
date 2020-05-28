// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package budgets

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Request of DeleteNotification
type DeleteNotificationInput struct {
	_ struct{} `type:"structure"`

	// The accountId that is associated with the budget whose notification you want
	// to delete.
	//
	// AccountId is a required field
	AccountId *string `min:"12" type:"string" required:"true"`

	// The name of the budget whose notification you want to delete.
	//
	// BudgetName is a required field
	BudgetName *string `min:"1" type:"string" required:"true"`

	// The notification that you want to delete.
	//
	// Notification is a required field
	Notification *Notification `type:"structure" required:"true"`
}

// String returns the string representation
func (s DeleteNotificationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteNotificationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteNotificationInput"}

	if s.AccountId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AccountId"))
	}
	if s.AccountId != nil && len(*s.AccountId) < 12 {
		invalidParams.Add(aws.NewErrParamMinLen("AccountId", 12))
	}

	if s.BudgetName == nil {
		invalidParams.Add(aws.NewErrParamRequired("BudgetName"))
	}
	if s.BudgetName != nil && len(*s.BudgetName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("BudgetName", 1))
	}

	if s.Notification == nil {
		invalidParams.Add(aws.NewErrParamRequired("Notification"))
	}
	if s.Notification != nil {
		if err := s.Notification.Validate(); err != nil {
			invalidParams.AddNested("Notification", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Response of DeleteNotification
type DeleteNotificationOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteNotificationOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteNotification = "DeleteNotification"

// DeleteNotificationRequest returns a request value for making API operation for
// AWS Budgets.
//
// Deletes a notification.
//
// Deleting a notification also deletes the subscribers that are associated
// with the notification.
//
//    // Example sending a request using DeleteNotificationRequest.
//    req := client.DeleteNotificationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) DeleteNotificationRequest(input *DeleteNotificationInput) DeleteNotificationRequest {
	op := &aws.Operation{
		Name:       opDeleteNotification,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteNotificationInput{}
	}

	req := c.newRequest(op, input, &DeleteNotificationOutput{})

	return DeleteNotificationRequest{Request: req, Input: input, Copy: c.DeleteNotificationRequest}
}

// DeleteNotificationRequest is the request type for the
// DeleteNotification API operation.
type DeleteNotificationRequest struct {
	*aws.Request
	Input *DeleteNotificationInput
	Copy  func(*DeleteNotificationInput) DeleteNotificationRequest
}

// Send marshals and sends the DeleteNotification API request.
func (r DeleteNotificationRequest) Send(ctx context.Context) (*DeleteNotificationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteNotificationResponse{
		DeleteNotificationOutput: r.Request.Data.(*DeleteNotificationOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteNotificationResponse is the response type for the
// DeleteNotification API operation.
type DeleteNotificationResponse struct {
	*DeleteNotificationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteNotification request.
func (r *DeleteNotificationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
