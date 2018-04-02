// Code generated by go-swagger; DO NOT EDIT.

package charts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"wwwin-github.cisco.com/edge/optikon/api/v0/models"
)

// GetChartByIDReader is a Reader for the GetChartByID structure.
type GetChartByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetChartByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetChartByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetChartByIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetChartByIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetChartByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetChartByIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetChartByIDOK creates a GetChartByIDOK with default headers values
func NewGetChartByIDOK() *GetChartByIDOK {
	return &GetChartByIDOK{}
}

/*GetChartByIDOK handles this case with default header values.

OK
*/
type GetChartByIDOK struct {
	Payload *models.ChartChart
}

func (o *GetChartByIDOK) Error() string {
	return fmt.Sprintf("[GET /charts/{chartId}][%d] getChartByIdOK  %+v", 200, o.Payload)
}

func (o *GetChartByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ChartChart)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetChartByIDBadRequest creates a GetChartByIDBadRequest with default headers values
func NewGetChartByIDBadRequest() *GetChartByIDBadRequest {
	return &GetChartByIDBadRequest{}
}

/*GetChartByIDBadRequest handles this case with default header values.

Bad Request
*/
type GetChartByIDBadRequest struct {
	Payload *models.APIResponse
}

func (o *GetChartByIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /charts/{chartId}][%d] getChartByIdBadRequest  %+v", 400, o.Payload)
}

func (o *GetChartByIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetChartByIDUnauthorized creates a GetChartByIDUnauthorized with default headers values
func NewGetChartByIDUnauthorized() *GetChartByIDUnauthorized {
	return &GetChartByIDUnauthorized{}
}

/*GetChartByIDUnauthorized handles this case with default header values.

Unauthorized
*/
type GetChartByIDUnauthorized struct {
	Payload *models.APIResponse
}

func (o *GetChartByIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /charts/{chartId}][%d] getChartByIdUnauthorized  %+v", 401, o.Payload)
}

func (o *GetChartByIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetChartByIDNotFound creates a GetChartByIDNotFound with default headers values
func NewGetChartByIDNotFound() *GetChartByIDNotFound {
	return &GetChartByIDNotFound{}
}

/*GetChartByIDNotFound handles this case with default header values.

The specified resource was not found
*/
type GetChartByIDNotFound struct {
	Payload *models.APIResponse
}

func (o *GetChartByIDNotFound) Error() string {
	return fmt.Sprintf("[GET /charts/{chartId}][%d] getChartByIdNotFound  %+v", 404, o.Payload)
}

func (o *GetChartByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetChartByIDInternalServerError creates a GetChartByIDInternalServerError with default headers values
func NewGetChartByIDInternalServerError() *GetChartByIDInternalServerError {
	return &GetChartByIDInternalServerError{}
}

/*GetChartByIDInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetChartByIDInternalServerError struct {
	Payload *models.APIResponse
}

func (o *GetChartByIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /charts/{chartId}][%d] getChartByIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetChartByIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}