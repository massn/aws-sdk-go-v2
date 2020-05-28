// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package backup

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type ListCopyJobsInput struct {
	_ struct{} `type:"structure"`

	// Returns only copy jobs that were created after the specified date.
	ByCreatedAfter *time.Time `location:"querystring" locationName:"createdAfter" type:"timestamp"`

	// Returns only copy jobs that were created before the specified date.
	ByCreatedBefore *time.Time `location:"querystring" locationName:"createdBefore" type:"timestamp"`

	// An Amazon Resource Name (ARN) that uniquely identifies a source backup vault
	// to copy from; for example, arn:aws:backup:us-east-1:123456789012:vault:aBackupVault.
	ByDestinationVaultArn *string `location:"querystring" locationName:"destinationVaultArn" type:"string"`

	// Returns only copy jobs that match the specified resource Amazon Resource
	// Name (ARN).
	ByResourceArn *string `location:"querystring" locationName:"resourceArn" type:"string"`

	// Returns only backup jobs for the specified resources:
	//
	//    * EBS for Amazon Elastic Block Store
	//
	//    * EFS for Amazon Elastic File System
	//
	//    * RDS for Amazon Relational Database Service
	//
	//    * Storage Gateway for AWS Storage Gateway
	ByResourceType *string `location:"querystring" locationName:"resourceType" type:"string"`

	// Returns only copy jobs that are in the specified state.
	ByState CopyJobState `location:"querystring" locationName:"state" type:"string" enum:"true"`

	// The maximum number of items to be returned.
	MaxResults *int64 `location:"querystring" locationName:"maxResults" min:"1" type:"integer"`

	// The next item following a partial list of returned items. For example, if
	// a request is made to return maxResults number of items, NextToken allows
	// you to return more items in your list starting at the location pointed to
	// by the next token.
	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListCopyJobsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListCopyJobsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListCopyJobsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListCopyJobsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ByCreatedAfter != nil {
		v := *s.ByCreatedAfter

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "createdAfter",
			protocol.TimeValue{V: v, Format: protocol.ISO8601TimeFormatName, QuotedFormatTime: false}, metadata)
	}
	if s.ByCreatedBefore != nil {
		v := *s.ByCreatedBefore

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "createdBefore",
			protocol.TimeValue{V: v, Format: protocol.ISO8601TimeFormatName, QuotedFormatTime: false}, metadata)
	}
	if s.ByDestinationVaultArn != nil {
		v := *s.ByDestinationVaultArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "destinationVaultArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ByResourceArn != nil {
		v := *s.ByResourceArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "resourceArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ByResourceType != nil {
		v := *s.ByResourceType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "resourceType", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.ByState) > 0 {
		v := s.ByState

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "state", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "maxResults", protocol.Int64Value(v), metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type ListCopyJobsOutput struct {
	_ struct{} `type:"structure"`

	// An array of structures containing metadata about your copy jobs returned
	// in JSON format.
	CopyJobs []CopyJob `type:"list"`

	// The next item following a partial list of returned items. For example, if
	// a request is made to return maxResults number of items, NextToken allows
	// you to return more items in your list starting at the location pointed to
	// by the next token.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListCopyJobsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListCopyJobsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.CopyJobs != nil {
		v := s.CopyJobs

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "CopyJobs", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opListCopyJobs = "ListCopyJobs"

// ListCopyJobsRequest returns a request value for making API operation for
// AWS Backup.
//
// Returns metadata about your copy jobs.
//
//    // Example sending a request using ListCopyJobsRequest.
//    req := client.ListCopyJobsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/backup-2018-11-15/ListCopyJobs
func (c *Client) ListCopyJobsRequest(input *ListCopyJobsInput) ListCopyJobsRequest {
	op := &aws.Operation{
		Name:       opListCopyJobs,
		HTTPMethod: "GET",
		HTTPPath:   "/copy-jobs/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListCopyJobsInput{}
	}

	req := c.newRequest(op, input, &ListCopyJobsOutput{})

	return ListCopyJobsRequest{Request: req, Input: input, Copy: c.ListCopyJobsRequest}
}

// ListCopyJobsRequest is the request type for the
// ListCopyJobs API operation.
type ListCopyJobsRequest struct {
	*aws.Request
	Input *ListCopyJobsInput
	Copy  func(*ListCopyJobsInput) ListCopyJobsRequest
}

// Send marshals and sends the ListCopyJobs API request.
func (r ListCopyJobsRequest) Send(ctx context.Context) (*ListCopyJobsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListCopyJobsResponse{
		ListCopyJobsOutput: r.Request.Data.(*ListCopyJobsOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListCopyJobsRequestPaginator returns a paginator for ListCopyJobs.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListCopyJobsRequest(input)
//   p := backup.NewListCopyJobsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListCopyJobsPaginator(req ListCopyJobsRequest) ListCopyJobsPaginator {
	return ListCopyJobsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListCopyJobsInput
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

// ListCopyJobsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListCopyJobsPaginator struct {
	aws.Pager
}

func (p *ListCopyJobsPaginator) CurrentPage() *ListCopyJobsOutput {
	return p.Pager.CurrentPage().(*ListCopyJobsOutput)
}

// ListCopyJobsResponse is the response type for the
// ListCopyJobs API operation.
type ListCopyJobsResponse struct {
	*ListCopyJobsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListCopyJobs request.
func (r *ListCopyJobsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
