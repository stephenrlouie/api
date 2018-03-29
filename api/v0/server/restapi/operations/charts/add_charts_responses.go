// Code generated by go-swagger; DO NOT EDIT.

package charts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"wwwin-github.cisco.com/edge/optikon/api/v0/models"
)

// AddChartsCreatedCode is the HTTP code returned for type AddChartsCreated
const AddChartsCreatedCode int = 201

/*AddChartsCreated Created

swagger:response addChartsCreated
*/
type AddChartsCreated struct {

	/*
	  In: Body
	*/
	Payload *models.APIResponse `json:"body,omitempty"`
}

// NewAddChartsCreated creates AddChartsCreated with default headers values
func NewAddChartsCreated() *AddChartsCreated {
	return &AddChartsCreated{}
}

// WithPayload adds the payload to the add charts created response
func (o *AddChartsCreated) WithPayload(payload *models.APIResponse) *AddChartsCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add charts created response
func (o *AddChartsCreated) SetPayload(payload *models.APIResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddChartsCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AddChartsBadRequestCode is the HTTP code returned for type AddChartsBadRequest
const AddChartsBadRequestCode int = 400

/*AddChartsBadRequest Bad Request

swagger:response addChartsBadRequest
*/
type AddChartsBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.APIResponse `json:"body,omitempty"`
}

// NewAddChartsBadRequest creates AddChartsBadRequest with default headers values
func NewAddChartsBadRequest() *AddChartsBadRequest {
	return &AddChartsBadRequest{}
}

// WithPayload adds the payload to the add charts bad request response
func (o *AddChartsBadRequest) WithPayload(payload *models.APIResponse) *AddChartsBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add charts bad request response
func (o *AddChartsBadRequest) SetPayload(payload *models.APIResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddChartsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AddChartsUnauthorizedCode is the HTTP code returned for type AddChartsUnauthorized
const AddChartsUnauthorizedCode int = 401

/*AddChartsUnauthorized Unauthorized

swagger:response addChartsUnauthorized
*/
type AddChartsUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.APIResponse `json:"body,omitempty"`
}

// NewAddChartsUnauthorized creates AddChartsUnauthorized with default headers values
func NewAddChartsUnauthorized() *AddChartsUnauthorized {
	return &AddChartsUnauthorized{}
}

// WithPayload adds the payload to the add charts unauthorized response
func (o *AddChartsUnauthorized) WithPayload(payload *models.APIResponse) *AddChartsUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add charts unauthorized response
func (o *AddChartsUnauthorized) SetPayload(payload *models.APIResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddChartsUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AddChartsInternalServerErrorCode is the HTTP code returned for type AddChartsInternalServerError
const AddChartsInternalServerErrorCode int = 500

/*AddChartsInternalServerError Internal Server Error

swagger:response addChartsInternalServerError
*/
type AddChartsInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.APIResponse `json:"body,omitempty"`
}

// NewAddChartsInternalServerError creates AddChartsInternalServerError with default headers values
func NewAddChartsInternalServerError() *AddChartsInternalServerError {
	return &AddChartsInternalServerError{}
}

// WithPayload adds the payload to the add charts internal server error response
func (o *AddChartsInternalServerError) WithPayload(payload *models.APIResponse) *AddChartsInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add charts internal server error response
func (o *AddChartsInternalServerError) SetPayload(payload *models.APIResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddChartsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
