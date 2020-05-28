// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ssm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type UpdateMaintenanceWindowInput struct {
	_ struct{} `type:"structure"`

	// Whether targets must be registered with the maintenance window before tasks
	// can be defined for those targets.
	AllowUnassociatedTargets *bool `type:"boolean"`

	// The number of hours before the end of the maintenance window that Systems
	// Manager stops scheduling new tasks for execution.
	Cutoff *int64 `type:"integer"`

	// An optional description for the update request.
	Description *string `min:"1" type:"string" sensitive:"true"`

	// The duration of the maintenance window in hours.
	Duration *int64 `min:"1" type:"integer"`

	// Whether the maintenance window is enabled.
	Enabled *bool `type:"boolean"`

	// The date and time, in ISO-8601 Extended format, for when you want the maintenance
	// window to become inactive. EndDate allows you to set a date and time in the
	// future when the maintenance window will no longer run.
	EndDate *string `type:"string"`

	// The name of the maintenance window.
	Name *string `min:"3" type:"string"`

	// If True, then all fields that are required by the CreateMaintenanceWindow
	// action are also required for this API request. Optional fields that are not
	// specified are set to null.
	Replace *bool `type:"boolean"`

	// The schedule of the maintenance window in the form of a cron or rate expression.
	Schedule *string `min:"1" type:"string"`

	// The time zone that the scheduled maintenance window executions are based
	// on, in Internet Assigned Numbers Authority (IANA) format. For example: "America/Los_Angeles",
	// "etc/UTC", or "Asia/Seoul". For more information, see the Time Zone Database
	// (https://www.iana.org/time-zones) on the IANA website.
	ScheduleTimezone *string `type:"string"`

	// The time zone that the scheduled maintenance window executions are based
	// on, in Internet Assigned Numbers Authority (IANA) format. For example: "America/Los_Angeles",
	// "etc/UTC", or "Asia/Seoul". For more information, see the Time Zone Database
	// (https://www.iana.org/time-zones) on the IANA website.
	StartDate *string `type:"string"`

	// The ID of the maintenance window to update.
	//
	// WindowId is a required field
	WindowId *string `min:"20" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateMaintenanceWindowInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateMaintenanceWindowInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateMaintenanceWindowInput"}
	if s.Description != nil && len(*s.Description) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Description", 1))
	}
	if s.Duration != nil && *s.Duration < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("Duration", 1))
	}
	if s.Name != nil && len(*s.Name) < 3 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 3))
	}
	if s.Schedule != nil && len(*s.Schedule) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Schedule", 1))
	}

	if s.WindowId == nil {
		invalidParams.Add(aws.NewErrParamRequired("WindowId"))
	}
	if s.WindowId != nil && len(*s.WindowId) < 20 {
		invalidParams.Add(aws.NewErrParamMinLen("WindowId", 20))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdateMaintenanceWindowOutput struct {
	_ struct{} `type:"structure"`

	// Whether targets must be registered with the maintenance window before tasks
	// can be defined for those targets.
	AllowUnassociatedTargets *bool `type:"boolean"`

	// The number of hours before the end of the maintenance window that Systems
	// Manager stops scheduling new tasks for execution.
	Cutoff *int64 `type:"integer"`

	// An optional description of the update.
	Description *string `min:"1" type:"string" sensitive:"true"`

	// The duration of the maintenance window in hours.
	Duration *int64 `min:"1" type:"integer"`

	// Whether the maintenance window is enabled.
	Enabled *bool `type:"boolean"`

	// The date and time, in ISO-8601 Extended format, for when the maintenance
	// window is scheduled to become inactive. The maintenance window will not run
	// after this specified time.
	EndDate *string `type:"string"`

	// The name of the maintenance window.
	Name *string `min:"3" type:"string"`

	// The schedule of the maintenance window in the form of a cron or rate expression.
	Schedule *string `min:"1" type:"string"`

	// The time zone that the scheduled maintenance window executions are based
	// on, in Internet Assigned Numbers Authority (IANA) format. For example: "America/Los_Angeles",
	// "etc/UTC", or "Asia/Seoul". For more information, see the Time Zone Database
	// (https://www.iana.org/time-zones) on the IANA website.
	ScheduleTimezone *string `type:"string"`

	// The date and time, in ISO-8601 Extended format, for when the maintenance
	// window is scheduled to become active. The maintenance window will not run
	// before this specified time.
	StartDate *string `type:"string"`

	// The ID of the created maintenance window.
	WindowId *string `min:"20" type:"string"`
}

// String returns the string representation
func (s UpdateMaintenanceWindowOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateMaintenanceWindow = "UpdateMaintenanceWindow"

// UpdateMaintenanceWindowRequest returns a request value for making API operation for
// Amazon Simple Systems Manager (SSM).
//
// Updates an existing maintenance window. Only specified parameters are modified.
//
// The value you specify for Duration determines the specific end time for the
// maintenance window based on the time it begins. No maintenance window tasks
// are permitted to start after the resulting endtime minus the number of hours
// you specify for Cutoff. For example, if the maintenance window starts at
// 3 PM, the duration is three hours, and the value you specify for Cutoff is
// one hour, no maintenance window tasks can start after 5 PM.
//
//    // Example sending a request using UpdateMaintenanceWindowRequest.
//    req := client.UpdateMaintenanceWindowRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/UpdateMaintenanceWindow
func (c *Client) UpdateMaintenanceWindowRequest(input *UpdateMaintenanceWindowInput) UpdateMaintenanceWindowRequest {
	op := &aws.Operation{
		Name:       opUpdateMaintenanceWindow,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateMaintenanceWindowInput{}
	}

	req := c.newRequest(op, input, &UpdateMaintenanceWindowOutput{})

	return UpdateMaintenanceWindowRequest{Request: req, Input: input, Copy: c.UpdateMaintenanceWindowRequest}
}

// UpdateMaintenanceWindowRequest is the request type for the
// UpdateMaintenanceWindow API operation.
type UpdateMaintenanceWindowRequest struct {
	*aws.Request
	Input *UpdateMaintenanceWindowInput
	Copy  func(*UpdateMaintenanceWindowInput) UpdateMaintenanceWindowRequest
}

// Send marshals and sends the UpdateMaintenanceWindow API request.
func (r UpdateMaintenanceWindowRequest) Send(ctx context.Context) (*UpdateMaintenanceWindowResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateMaintenanceWindowResponse{
		UpdateMaintenanceWindowOutput: r.Request.Data.(*UpdateMaintenanceWindowOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateMaintenanceWindowResponse is the response type for the
// UpdateMaintenanceWindow API operation.
type UpdateMaintenanceWindowResponse struct {
	*UpdateMaintenanceWindowOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateMaintenanceWindow request.
func (r *UpdateMaintenanceWindowResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
