// Code generated by go-swagger; DO NOT EDIT.

package clusters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"wwwin-github.cisco.com/edge/optikon/api/v0/models"
)

// AddClusterReader is a Reader for the AddCluster structure.
type AddClusterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddClusterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewAddClusterCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewAddClusterBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewAddClusterUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewAddClusterInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAddClusterCreated creates a AddClusterCreated with default headers values
func NewAddClusterCreated() *AddClusterCreated {
	return &AddClusterCreated{}
}

/*AddClusterCreated handles this case with default header values.

Created
*/
type AddClusterCreated struct {
	Payload *models.APIResponse
}

func (o *AddClusterCreated) Error() string {
	return fmt.Sprintf("[POST /clusters][%d] addClusterCreated  %+v", 201, o.Payload)
}

func (o *AddClusterCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddClusterBadRequest creates a AddClusterBadRequest with default headers values
func NewAddClusterBadRequest() *AddClusterBadRequest {
	return &AddClusterBadRequest{}
}

/*AddClusterBadRequest handles this case with default header values.

Bad Request
*/
type AddClusterBadRequest struct {
	Payload *models.APIResponse
}

func (o *AddClusterBadRequest) Error() string {
	return fmt.Sprintf("[POST /clusters][%d] addClusterBadRequest  %+v", 400, o.Payload)
}

func (o *AddClusterBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddClusterUnauthorized creates a AddClusterUnauthorized with default headers values
func NewAddClusterUnauthorized() *AddClusterUnauthorized {
	return &AddClusterUnauthorized{}
}

/*AddClusterUnauthorized handles this case with default header values.

Unauthorized
*/
type AddClusterUnauthorized struct {
	Payload *models.APIResponse
}

func (o *AddClusterUnauthorized) Error() string {
	return fmt.Sprintf("[POST /clusters][%d] addClusterUnauthorized  %+v", 401, o.Payload)
}

func (o *AddClusterUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddClusterInternalServerError creates a AddClusterInternalServerError with default headers values
func NewAddClusterInternalServerError() *AddClusterInternalServerError {
	return &AddClusterInternalServerError{}
}

/*AddClusterInternalServerError handles this case with default header values.

Internal Server Error
*/
type AddClusterInternalServerError struct {
	Payload *models.APIResponse
}

func (o *AddClusterInternalServerError) Error() string {
	return fmt.Sprintf("[POST /clusters][%d] addClusterInternalServerError  %+v", 500, o.Payload)
}

func (o *AddClusterInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
