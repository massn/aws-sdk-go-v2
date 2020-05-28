// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudsearch

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Container for the parameters to the DescribeSuggester operation. Specifies
// the name of the domain you want to describe. To restrict the response to
// particular suggesters, specify the names of the suggesters you want to describe.
// To show the active configuration and exclude any pending changes, set the
// Deployed option to true.
type DescribeSuggestersInput struct {
	_ struct{} `type:"structure"`

	// Whether to display the deployed configuration (true) or include any pending
	// changes (false). Defaults to false.
	Deployed *bool `type:"boolean"`

	// The name of the domain you want to describe.
	//
	// DomainName is a required field
	DomainName *string `min:"3" type:"string" required:"true"`

	// The suggesters you want to describe.
	SuggesterNames []string `type:"list"`
}

// String returns the string representation
func (s DescribeSuggestersInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeSuggestersInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeSuggestersInput"}

	if s.DomainName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DomainName"))
	}
	if s.DomainName != nil && len(*s.DomainName) < 3 {
		invalidParams.Add(aws.NewErrParamMinLen("DomainName", 3))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The result of a DescribeSuggesters request.
type DescribeSuggestersOutput struct {
	_ struct{} `type:"structure"`

	// The suggesters configured for the domain specified in the request.
	//
	// Suggesters is a required field
	Suggesters []SuggesterStatus `type:"list" required:"true"`
}

// String returns the string representation
func (s DescribeSuggestersOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeSuggesters = "DescribeSuggesters"

// DescribeSuggestersRequest returns a request value for making API operation for
// Amazon CloudSearch.
//
// Gets the suggesters configured for a domain. A suggester enables you to display
// possible matches before users finish typing their queries. Can be limited
// to specific suggesters by name. By default, shows all suggesters and includes
// any pending changes to the configuration. Set the Deployed option to true
// to show the active configuration and exclude pending changes. For more information,
// see Getting Search Suggestions (http://docs.aws.amazon.com/cloudsearch/latest/developerguide/getting-suggestions.html)
// in the Amazon CloudSearch Developer Guide.
//
//    // Example sending a request using DescribeSuggestersRequest.
//    req := client.DescribeSuggestersRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) DescribeSuggestersRequest(input *DescribeSuggestersInput) DescribeSuggestersRequest {
	op := &aws.Operation{
		Name:       opDescribeSuggesters,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeSuggestersInput{}
	}

	req := c.newRequest(op, input, &DescribeSuggestersOutput{})

	return DescribeSuggestersRequest{Request: req, Input: input, Copy: c.DescribeSuggestersRequest}
}

// DescribeSuggestersRequest is the request type for the
// DescribeSuggesters API operation.
type DescribeSuggestersRequest struct {
	*aws.Request
	Input *DescribeSuggestersInput
	Copy  func(*DescribeSuggestersInput) DescribeSuggestersRequest
}

// Send marshals and sends the DescribeSuggesters API request.
func (r DescribeSuggestersRequest) Send(ctx context.Context) (*DescribeSuggestersResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeSuggestersResponse{
		DescribeSuggestersOutput: r.Request.Data.(*DescribeSuggestersOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeSuggestersResponse is the response type for the
// DescribeSuggesters API operation.
type DescribeSuggestersResponse struct {
	*DescribeSuggestersOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeSuggesters request.
func (r *DescribeSuggestersResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
