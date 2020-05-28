// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package personalize

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeRecipeInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the recipe to describe.
	//
	// RecipeArn is a required field
	RecipeArn *string `locationName:"recipeArn" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeRecipeInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeRecipeInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeRecipeInput"}

	if s.RecipeArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("RecipeArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeRecipeOutput struct {
	_ struct{} `type:"structure"`

	// An object that describes the recipe.
	Recipe *Recipe `locationName:"recipe" type:"structure"`
}

// String returns the string representation
func (s DescribeRecipeOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeRecipe = "DescribeRecipe"

// DescribeRecipeRequest returns a request value for making API operation for
// Amazon Personalize.
//
// Describes a recipe.
//
// A recipe contains three items:
//
//    * An algorithm that trains a model.
//
//    * Hyperparameters that govern the training.
//
//    * Feature transformation information for modifying the input data before
//    training.
//
// Amazon Personalize provides a set of predefined recipes. You specify a recipe
// when you create a solution with the CreateSolution API. CreateSolution trains
// a model by using the algorithm in the specified recipe and a training dataset.
// The solution, when deployed as a campaign, can provide recommendations using
// the GetRecommendations (https://docs.aws.amazon.com/personalize/latest/dg/API_RS_GetRecommendations.html)
// API.
//
//    // Example sending a request using DescribeRecipeRequest.
//    req := client.DescribeRecipeRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/personalize-2018-05-22/DescribeRecipe
func (c *Client) DescribeRecipeRequest(input *DescribeRecipeInput) DescribeRecipeRequest {
	op := &aws.Operation{
		Name:       opDescribeRecipe,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeRecipeInput{}
	}

	req := c.newRequest(op, input, &DescribeRecipeOutput{})

	return DescribeRecipeRequest{Request: req, Input: input, Copy: c.DescribeRecipeRequest}
}

// DescribeRecipeRequest is the request type for the
// DescribeRecipe API operation.
type DescribeRecipeRequest struct {
	*aws.Request
	Input *DescribeRecipeInput
	Copy  func(*DescribeRecipeInput) DescribeRecipeRequest
}

// Send marshals and sends the DescribeRecipe API request.
func (r DescribeRecipeRequest) Send(ctx context.Context) (*DescribeRecipeResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeRecipeResponse{
		DescribeRecipeOutput: r.Request.Data.(*DescribeRecipeOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeRecipeResponse is the response type for the
// DescribeRecipe API operation.
type DescribeRecipeResponse struct {
	*DescribeRecipeOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeRecipe request.
func (r *DescribeRecipeResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
