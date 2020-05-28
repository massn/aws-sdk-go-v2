// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kms

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GenerateDataKeyPairInput struct {
	_ struct{} `type:"structure"`

	// Specifies the encryption context that will be used when encrypting the private
	// key in the data key pair.
	//
	// An encryption context is a collection of non-secret key-value pairs that
	// represents additional authenticated data. When you use an encryption context
	// to encrypt data, you must specify the same (an exact case-sensitive match)
	// encryption context to decrypt the data. An encryption context is optional
	// when encrypting with a symmetric CMK, but it is highly recommended.
	//
	// For more information, see Encryption Context (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#encrypt_context)
	// in the AWS Key Management Service Developer Guide.
	EncryptionContext map[string]string `type:"map"`

	// A list of grant tokens.
	//
	// For more information, see Grant Tokens (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#grant_token)
	// in the AWS Key Management Service Developer Guide.
	GrantTokens []string `type:"list"`

	// Specifies the symmetric CMK that encrypts the private key in the data key
	// pair. You cannot specify an asymmetric CMKs.
	//
	// To specify a CMK, use its key ID, Amazon Resource Name (ARN), alias name,
	// or alias ARN. When using an alias name, prefix it with "alias/". To specify
	// a CMK in a different AWS account, you must use the key ARN or alias ARN.
	//
	// For example:
	//
	//    * Key ID: 1234abcd-12ab-34cd-56ef-1234567890ab
	//
	//    * Key ARN: arn:aws:kms:us-east-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab
	//
	//    * Alias name: alias/ExampleAlias
	//
	//    * Alias ARN: arn:aws:kms:us-east-2:111122223333:alias/ExampleAlias
	//
	// To get the key ID and key ARN for a CMK, use ListKeys or DescribeKey. To
	// get the alias name and alias ARN, use ListAliases.
	//
	// KeyId is a required field
	KeyId *string `min:"1" type:"string" required:"true"`

	// Determines the type of data key pair that is generated.
	//
	// The AWS KMS rule that restricts the use of asymmetric RSA CMKs to encrypt
	// and decrypt or to sign and verify (but not both), and the rule that permits
	// you to use ECC CMKs only to sign and verify, are not effective outside of
	// AWS KMS.
	//
	// KeyPairSpec is a required field
	KeyPairSpec DataKeyPairSpec `type:"string" required:"true" enum:"true"`
}

// String returns the string representation
func (s GenerateDataKeyPairInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GenerateDataKeyPairInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GenerateDataKeyPairInput"}

	if s.KeyId == nil {
		invalidParams.Add(aws.NewErrParamRequired("KeyId"))
	}
	if s.KeyId != nil && len(*s.KeyId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("KeyId", 1))
	}
	if len(s.KeyPairSpec) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("KeyPairSpec"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GenerateDataKeyPairOutput struct {
	_ struct{} `type:"structure"`

	// The identifier of the CMK that encrypted the private key.
	KeyId *string `min:"1" type:"string"`

	// The type of data key pair that was generated.
	KeyPairSpec DataKeyPairSpec `type:"string" enum:"true"`

	// The encrypted copy of the private key. When you use the HTTP API or the AWS
	// CLI, the value is Base64-encoded. Otherwise, it is not Base64-encoded.
	//
	// PrivateKeyCiphertextBlob is automatically base64 encoded/decoded by the SDK.
	PrivateKeyCiphertextBlob []byte `min:"1" type:"blob"`

	// The plaintext copy of the private key. When you use the HTTP API or the AWS
	// CLI, the value is Base64-encoded. Otherwise, it is not Base64-encoded.
	//
	// PrivateKeyPlaintext is automatically base64 encoded/decoded by the SDK.
	PrivateKeyPlaintext []byte `min:"1" type:"blob" sensitive:"true"`

	// The public key (in plaintext).
	//
	// PublicKey is automatically base64 encoded/decoded by the SDK.
	PublicKey []byte `min:"1" type:"blob"`
}

// String returns the string representation
func (s GenerateDataKeyPairOutput) String() string {
	return awsutil.Prettify(s)
}

const opGenerateDataKeyPair = "GenerateDataKeyPair"

// GenerateDataKeyPairRequest returns a request value for making API operation for
// AWS Key Management Service.
//
// Generates a unique asymmetric data key pair. The GenerateDataKeyPair operation
// returns a plaintext public key, a plaintext private key, and a copy of the
// private key that is encrypted under the symmetric CMK you specify. You can
// use the data key pair to perform asymmetric cryptography outside of AWS KMS.
//
// GenerateDataKeyPair returns a unique data key pair for each request. The
// bytes in the keys are not related to the caller or the CMK that is used to
// encrypt the private key.
//
// You can use the public key that GenerateDataKeyPair returns to encrypt data
// or verify a signature outside of AWS KMS. Then, store the encrypted private
// key with the data. When you are ready to decrypt data or sign a message,
// you can use the Decrypt operation to decrypt the encrypted private key.
//
// To generate a data key pair, you must specify a symmetric customer master
// key (CMK) to encrypt the private key in a data key pair. You cannot use an
// asymmetric CMK. To get the type of your CMK, use the DescribeKey operation.
//
// If you are using the data key pair to encrypt data, or for any operation
// where you don't immediately need a private key, consider using the GenerateDataKeyPairWithoutPlaintext
// operation. GenerateDataKeyPairWithoutPlaintext returns a plaintext public
// key and an encrypted private key, but omits the plaintext private key that
// you need only to decrypt ciphertext or sign a message. Later, when you need
// to decrypt the data or sign a message, use the Decrypt operation to decrypt
// the encrypted private key in the data key pair.
//
// You can use the optional encryption context to add additional security to
// the encryption operation. If you specify an EncryptionContext, you must specify
// the same encryption context (a case-sensitive exact match) when decrypting
// the encrypted data key. Otherwise, the request to decrypt fails with an InvalidCiphertextException.
// For more information, see Encryption Context (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#encrypt_context)
// in the AWS Key Management Service Developer Guide.
//
// The CMK that you use for this operation must be in a compatible key state.
// For details, see How Key State Affects Use of a Customer Master Key (https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html)
// in the AWS Key Management Service Developer Guide.
//
//    // Example sending a request using GenerateDataKeyPairRequest.
//    req := client.GenerateDataKeyPairRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kms-2014-11-01/GenerateDataKeyPair
func (c *Client) GenerateDataKeyPairRequest(input *GenerateDataKeyPairInput) GenerateDataKeyPairRequest {
	op := &aws.Operation{
		Name:       opGenerateDataKeyPair,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GenerateDataKeyPairInput{}
	}

	req := c.newRequest(op, input, &GenerateDataKeyPairOutput{})

	return GenerateDataKeyPairRequest{Request: req, Input: input, Copy: c.GenerateDataKeyPairRequest}
}

// GenerateDataKeyPairRequest is the request type for the
// GenerateDataKeyPair API operation.
type GenerateDataKeyPairRequest struct {
	*aws.Request
	Input *GenerateDataKeyPairInput
	Copy  func(*GenerateDataKeyPairInput) GenerateDataKeyPairRequest
}

// Send marshals and sends the GenerateDataKeyPair API request.
func (r GenerateDataKeyPairRequest) Send(ctx context.Context) (*GenerateDataKeyPairResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GenerateDataKeyPairResponse{
		GenerateDataKeyPairOutput: r.Request.Data.(*GenerateDataKeyPairOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GenerateDataKeyPairResponse is the response type for the
// GenerateDataKeyPair API operation.
type GenerateDataKeyPairResponse struct {
	*GenerateDataKeyPairOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GenerateDataKeyPair request.
func (r *GenerateDataKeyPairResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
