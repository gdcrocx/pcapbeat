// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/UpdateServiceSpecificCredentialRequest
type UpdateServiceSpecificCredentialInput struct {
	_ struct{} `type:"structure"`

	// The unique identifier of the service-specific credential.
	//
	// This parameter allows (through its regex pattern (http://wikipedia.org/wiki/regex))
	// a string of characters that can consist of any upper or lowercased letter
	// or digit.
	//
	// ServiceSpecificCredentialId is a required field
	ServiceSpecificCredentialId *string `min:"20" type:"string" required:"true"`

	// The status to be assigned to the service-specific credential.
	//
	// Status is a required field
	Status StatusType `type:"string" required:"true" enum:"true"`

	// The name of the IAM user associated with the service-specific credential.
	// If you do not specify this value, then the operation assumes the user whose
	// credentials are used to call the operation.
	//
	// This parameter allows (through its regex pattern (http://wikipedia.org/wiki/regex))
	// a string of characters consisting of upper and lowercase alphanumeric characters
	// with no spaces. You can also include any of the following characters: _+=,.@-
	UserName *string `min:"1" type:"string"`
}

// String returns the string representation
func (s UpdateServiceSpecificCredentialInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateServiceSpecificCredentialInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateServiceSpecificCredentialInput"}

	if s.ServiceSpecificCredentialId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ServiceSpecificCredentialId"))
	}
	if s.ServiceSpecificCredentialId != nil && len(*s.ServiceSpecificCredentialId) < 20 {
		invalidParams.Add(aws.NewErrParamMinLen("ServiceSpecificCredentialId", 20))
	}
	if len(s.Status) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Status"))
	}
	if s.UserName != nil && len(*s.UserName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("UserName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/UpdateServiceSpecificCredentialOutput
type UpdateServiceSpecificCredentialOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UpdateServiceSpecificCredentialOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateServiceSpecificCredential = "UpdateServiceSpecificCredential"

// UpdateServiceSpecificCredentialRequest returns a request value for making API operation for
// AWS Identity and Access Management.
//
// Sets the status of a service-specific credential to Active or Inactive. Service-specific
// credentials that are inactive cannot be used for authentication to the service.
// This operation can be used to disable a user's service-specific credential
// as part of a credential rotation work flow.
//
//    // Example sending a request using UpdateServiceSpecificCredentialRequest.
//    req := client.UpdateServiceSpecificCredentialRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/UpdateServiceSpecificCredential
func (c *Client) UpdateServiceSpecificCredentialRequest(input *UpdateServiceSpecificCredentialInput) UpdateServiceSpecificCredentialRequest {
	op := &aws.Operation{
		Name:       opUpdateServiceSpecificCredential,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateServiceSpecificCredentialInput{}
	}

	req := c.newRequest(op, input, &UpdateServiceSpecificCredentialOutput{})
	req.Handlers.Unmarshal.Remove(query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return UpdateServiceSpecificCredentialRequest{Request: req, Input: input, Copy: c.UpdateServiceSpecificCredentialRequest}
}

// UpdateServiceSpecificCredentialRequest is the request type for the
// UpdateServiceSpecificCredential API operation.
type UpdateServiceSpecificCredentialRequest struct {
	*aws.Request
	Input *UpdateServiceSpecificCredentialInput
	Copy  func(*UpdateServiceSpecificCredentialInput) UpdateServiceSpecificCredentialRequest
}

// Send marshals and sends the UpdateServiceSpecificCredential API request.
func (r UpdateServiceSpecificCredentialRequest) Send(ctx context.Context) (*UpdateServiceSpecificCredentialResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateServiceSpecificCredentialResponse{
		UpdateServiceSpecificCredentialOutput: r.Request.Data.(*UpdateServiceSpecificCredentialOutput),
		response:                              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateServiceSpecificCredentialResponse is the response type for the
// UpdateServiceSpecificCredential API operation.
type UpdateServiceSpecificCredentialResponse struct {
	*UpdateServiceSpecificCredentialOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateServiceSpecificCredential request.
func (r *UpdateServiceSpecificCredentialResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
