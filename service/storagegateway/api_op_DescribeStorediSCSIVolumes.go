// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package storagegateway

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// A JSON object containing a list of DescribeStorediSCSIVolumesInput$VolumeARNs.
type DescribeStorediSCSIVolumesInput struct {
	_ struct{} `type:"structure"`

	// An array of strings where each string represents the Amazon Resource Name
	// (ARN) of a stored volume. All of the specified stored volumes must be from
	// the same gateway. Use ListVolumes to get volume ARNs for a gateway.
	//
	// VolumeARNs is a required field
	VolumeARNs []string `type:"list" required:"true"`
}

// String returns the string representation
func (s DescribeStorediSCSIVolumesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeStorediSCSIVolumesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeStorediSCSIVolumesInput"}

	if s.VolumeARNs == nil {
		invalidParams.Add(aws.NewErrParamRequired("VolumeARNs"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeStorediSCSIVolumesOutput struct {
	_ struct{} `type:"structure"`

	// Describes a single unit of output from DescribeStorediSCSIVolumes. The following
	// fields are returned:
	//
	//    * ChapEnabled: Indicates whether mutual CHAP is enabled for the iSCSI
	//    target.
	//
	//    * LunNumber: The logical disk number.
	//
	//    * NetworkInterfaceId: The network interface ID of the stored volume that
	//    initiator use to map the stored volume as an iSCSI target.
	//
	//    * NetworkInterfacePort: The port used to communicate with iSCSI targets.
	//
	//    * PreservedExistingData: Indicates if when the stored volume was created,
	//    existing data on the underlying local disk was preserved.
	//
	//    * SourceSnapshotId: If the stored volume was created from a snapshot,
	//    this field contains the snapshot ID used, e.g. snap-1122aabb. Otherwise,
	//    this field is not included.
	//
	//    * StorediSCSIVolumes: An array of StorediSCSIVolume objects where each
	//    object contains metadata about one stored volume.
	//
	//    * TargetARN: The Amazon Resource Name (ARN) of the volume target.
	//
	//    * VolumeARN: The Amazon Resource Name (ARN) of the stored volume.
	//
	//    * VolumeDiskId: The disk ID of the local disk that was specified in the
	//    CreateStorediSCSIVolume operation.
	//
	//    * VolumeId: The unique identifier of the storage volume, e.g. vol-1122AABB.
	//
	//    * VolumeiSCSIAttributes: An VolumeiSCSIAttributes object that represents
	//    a collection of iSCSI attributes for one stored volume.
	//
	//    * VolumeProgress: Represents the percentage complete if the volume is
	//    restoring or bootstrapping that represents the percent of data transferred.
	//    This field does not appear in the response if the stored volume is not
	//    restoring or bootstrapping.
	//
	//    * VolumeSizeInBytes: The size of the volume in bytes.
	//
	//    * VolumeStatus: One of the VolumeStatus values that indicates the state
	//    of the volume.
	//
	//    * VolumeType: One of the enumeration values describing the type of the
	//    volume. Currently, on STORED volumes are supported.
	StorediSCSIVolumes []StorediSCSIVolume `type:"list"`
}

// String returns the string representation
func (s DescribeStorediSCSIVolumesOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeStorediSCSIVolumes = "DescribeStorediSCSIVolumes"

// DescribeStorediSCSIVolumesRequest returns a request value for making API operation for
// AWS Storage Gateway.
//
// Returns the description of the gateway volumes specified in the request.
// The list of gateway volumes in the request must be from one gateway. In the
// response AWS Storage Gateway returns volume information sorted by volume
// ARNs. This operation is only supported in stored volume gateway type.
//
//    // Example sending a request using DescribeStorediSCSIVolumesRequest.
//    req := client.DescribeStorediSCSIVolumesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/storagegateway-2013-06-30/DescribeStorediSCSIVolumes
func (c *Client) DescribeStorediSCSIVolumesRequest(input *DescribeStorediSCSIVolumesInput) DescribeStorediSCSIVolumesRequest {
	op := &aws.Operation{
		Name:       opDescribeStorediSCSIVolumes,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeStorediSCSIVolumesInput{}
	}

	req := c.newRequest(op, input, &DescribeStorediSCSIVolumesOutput{})

	return DescribeStorediSCSIVolumesRequest{Request: req, Input: input, Copy: c.DescribeStorediSCSIVolumesRequest}
}

// DescribeStorediSCSIVolumesRequest is the request type for the
// DescribeStorediSCSIVolumes API operation.
type DescribeStorediSCSIVolumesRequest struct {
	*aws.Request
	Input *DescribeStorediSCSIVolumesInput
	Copy  func(*DescribeStorediSCSIVolumesInput) DescribeStorediSCSIVolumesRequest
}

// Send marshals and sends the DescribeStorediSCSIVolumes API request.
func (r DescribeStorediSCSIVolumesRequest) Send(ctx context.Context) (*DescribeStorediSCSIVolumesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeStorediSCSIVolumesResponse{
		DescribeStorediSCSIVolumesOutput: r.Request.Data.(*DescribeStorediSCSIVolumesOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeStorediSCSIVolumesResponse is the response type for the
// DescribeStorediSCSIVolumes API operation.
type DescribeStorediSCSIVolumesResponse struct {
	*DescribeStorediSCSIVolumesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeStorediSCSIVolumes request.
func (r *DescribeStorediSCSIVolumesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
