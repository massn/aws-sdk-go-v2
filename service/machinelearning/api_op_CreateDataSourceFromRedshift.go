// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package machinelearning

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateDataSourceFromRedshiftInput struct {
	_ struct{} `type:"structure"`

	// The compute statistics for a DataSource. The statistics are generated from
	// the observation data referenced by a DataSource. Amazon ML uses the statistics
	// internally during MLModel training. This parameter must be set to true if
	// the DataSource needs to be used for MLModel training.
	ComputeStatistics *bool `type:"boolean"`

	// A user-supplied ID that uniquely identifies the DataSource.
	//
	// DataSourceId is a required field
	DataSourceId *string `min:"1" type:"string" required:"true"`

	// A user-supplied name or description of the DataSource.
	DataSourceName *string `type:"string"`

	// The data specification of an Amazon Redshift DataSource:
	//
	//    * DatabaseInformation -
	//    * DatabaseName - The name of the Amazon Redshift database.
	//
	//    * ClusterIdentifier - The unique ID for the Amazon Redshift cluster.
	//
	//    * DatabaseCredentials - The AWS Identity and Access Management (IAM) credentials
	//    that are used to connect to the Amazon Redshift database.
	//
	//    * SelectSqlQuery - The query that is used to retrieve the observation
	//    data for the Datasource.
	//
	//    * S3StagingLocation - The Amazon Simple Storage Service (Amazon S3) location
	//    for staging Amazon Redshift data. The data retrieved from Amazon Redshift
	//    using the SelectSqlQuery query is stored in this location.
	//
	//    * DataSchemaUri - The Amazon S3 location of the DataSchema.
	//
	//    * DataSchema - A JSON string representing the schema. This is not required
	//    if DataSchemaUri is specified.
	//
	//    * DataRearrangement - A JSON string that represents the splitting and
	//    rearrangement requirements for the DataSource. Sample - "{\"splitting\":{\"percentBegin\":10,\"percentEnd\":60}}"
	//
	// DataSpec is a required field
	DataSpec *RedshiftDataSpec `type:"structure" required:"true"`

	// A fully specified role Amazon Resource Name (ARN). Amazon ML assumes the
	// role on behalf of the user to create the following:
	//
	//    * A security group to allow Amazon ML to execute the SelectSqlQuery query
	//    on an Amazon Redshift cluster
	//
	//    * An Amazon S3 bucket policy to grant Amazon ML read/write permissions
	//    on the S3StagingLocation
	//
	// RoleARN is a required field
	RoleARN *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s CreateDataSourceFromRedshiftInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateDataSourceFromRedshiftInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateDataSourceFromRedshiftInput"}

	if s.DataSourceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DataSourceId"))
	}
	if s.DataSourceId != nil && len(*s.DataSourceId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DataSourceId", 1))
	}

	if s.DataSpec == nil {
		invalidParams.Add(aws.NewErrParamRequired("DataSpec"))
	}

	if s.RoleARN == nil {
		invalidParams.Add(aws.NewErrParamRequired("RoleARN"))
	}
	if s.RoleARN != nil && len(*s.RoleARN) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("RoleARN", 1))
	}
	if s.DataSpec != nil {
		if err := s.DataSpec.Validate(); err != nil {
			invalidParams.AddNested("DataSpec", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the output of a CreateDataSourceFromRedshift operation, and is
// an acknowledgement that Amazon ML received the request.
//
// The CreateDataSourceFromRedshift operation is asynchronous. You can poll
// for updates by using the GetBatchPrediction operation and checking the Status
// parameter.
type CreateDataSourceFromRedshiftOutput struct {
	_ struct{} `type:"structure"`

	// A user-supplied ID that uniquely identifies the datasource. This value should
	// be identical to the value of the DataSourceID in the request.
	DataSourceId *string `min:"1" type:"string"`
}

// String returns the string representation
func (s CreateDataSourceFromRedshiftOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateDataSourceFromRedshift = "CreateDataSourceFromRedshift"

// CreateDataSourceFromRedshiftRequest returns a request value for making API operation for
// Amazon Machine Learning.
//
// Creates a DataSource from a database hosted on an Amazon Redshift cluster.
// A DataSource references data that can be used to perform either CreateMLModel,
// CreateEvaluation, or CreateBatchPrediction operations.
//
// CreateDataSourceFromRedshift is an asynchronous operation. In response to
// CreateDataSourceFromRedshift, Amazon Machine Learning (Amazon ML) immediately
// returns and sets the DataSource status to PENDING. After the DataSource is
// created and ready for use, Amazon ML sets the Status parameter to COMPLETED.
// DataSource in COMPLETED or PENDING states can be used to perform only CreateMLModel,
// CreateEvaluation, or CreateBatchPrediction operations.
//
// If Amazon ML can't accept the input source, it sets the Status parameter
// to FAILED and includes an error message in the Message attribute of the GetDataSource
// operation response.
//
// The observations should be contained in the database hosted on an Amazon
// Redshift cluster and should be specified by a SelectSqlQuery query. Amazon
// ML executes an Unload command in Amazon Redshift to transfer the result set
// of the SelectSqlQuery query to S3StagingLocation.
//
// After the DataSource has been created, it's ready for use in evaluations
// and batch predictions. If you plan to use the DataSource to train an MLModel,
// the DataSource also requires a recipe. A recipe describes how each input
// variable will be used in training an MLModel. Will the variable be included
// or excluded from training? Will the variable be manipulated; for example,
// will it be combined with another variable or will it be split apart into
// word combinations? The recipe provides answers to these questions.
//
// You can't change an existing datasource, but you can copy and modify the
// settings from an existing Amazon Redshift datasource to create a new datasource.
// To do so, call GetDataSource for an existing datasource and copy the values
// to a CreateDataSource call. Change the settings that you want to change and
// make sure that all required fields have the appropriate values.
//
//    // Example sending a request using CreateDataSourceFromRedshiftRequest.
//    req := client.CreateDataSourceFromRedshiftRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) CreateDataSourceFromRedshiftRequest(input *CreateDataSourceFromRedshiftInput) CreateDataSourceFromRedshiftRequest {
	op := &aws.Operation{
		Name:       opCreateDataSourceFromRedshift,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateDataSourceFromRedshiftInput{}
	}

	req := c.newRequest(op, input, &CreateDataSourceFromRedshiftOutput{})

	return CreateDataSourceFromRedshiftRequest{Request: req, Input: input, Copy: c.CreateDataSourceFromRedshiftRequest}
}

// CreateDataSourceFromRedshiftRequest is the request type for the
// CreateDataSourceFromRedshift API operation.
type CreateDataSourceFromRedshiftRequest struct {
	*aws.Request
	Input *CreateDataSourceFromRedshiftInput
	Copy  func(*CreateDataSourceFromRedshiftInput) CreateDataSourceFromRedshiftRequest
}

// Send marshals and sends the CreateDataSourceFromRedshift API request.
func (r CreateDataSourceFromRedshiftRequest) Send(ctx context.Context) (*CreateDataSourceFromRedshiftResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateDataSourceFromRedshiftResponse{
		CreateDataSourceFromRedshiftOutput: r.Request.Data.(*CreateDataSourceFromRedshiftOutput),
		response:                           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateDataSourceFromRedshiftResponse is the response type for the
// CreateDataSourceFromRedshift API operation.
type CreateDataSourceFromRedshiftResponse struct {
	*CreateDataSourceFromRedshiftOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateDataSourceFromRedshift request.
func (r *CreateDataSourceFromRedshiftResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
