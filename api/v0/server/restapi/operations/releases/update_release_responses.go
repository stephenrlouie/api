// Code generated by go-swagger; DO NOT EDIT.

package releases

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/optikon/api/api/v0/models"
)

// UpdateReleaseOKCode is the HTTP code returned for type UpdateReleaseOK
const UpdateReleaseOKCode int = 200

/*UpdateReleaseOK OK

swagger:response updateReleaseOK
*/
type UpdateReleaseOK struct {
}

// NewUpdateReleaseOK creates UpdateReleaseOK with default headers values
func NewUpdateReleaseOK() *UpdateReleaseOK {
	return &UpdateReleaseOK{}
}

// WriteResponse to the client
func (o *UpdateReleaseOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// UpdateReleaseBadRequestCode is the HTTP code returned for type UpdateReleaseBadRequest
const UpdateReleaseBadRequestCode int = 400

/*UpdateReleaseBadRequest Bad Request

swagger:response updateReleaseBadRequest
*/
type UpdateReleaseBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.APIResponse `json:"body,omitempty"`
}

// NewUpdateReleaseBadRequest creates UpdateReleaseBadRequest with default headers values
func NewUpdateReleaseBadRequest() *UpdateReleaseBadRequest {
	return &UpdateReleaseBadRequest{}
}

// WithPayload adds the payload to the update release bad request response
func (o *UpdateReleaseBadRequest) WithPayload(payload *models.APIResponse) *UpdateReleaseBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update release bad request response
func (o *UpdateReleaseBadRequest) SetPayload(payload *models.APIResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateReleaseBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateReleaseUnauthorizedCode is the HTTP code returned for type UpdateReleaseUnauthorized
const UpdateReleaseUnauthorizedCode int = 401

/*UpdateReleaseUnauthorized Unauthorized

swagger:response updateReleaseUnauthorized
*/
type UpdateReleaseUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.APIResponse `json:"body,omitempty"`
}

// NewUpdateReleaseUnauthorized creates UpdateReleaseUnauthorized with default headers values
func NewUpdateReleaseUnauthorized() *UpdateReleaseUnauthorized {
	return &UpdateReleaseUnauthorized{}
}

// WithPayload adds the payload to the update release unauthorized response
func (o *UpdateReleaseUnauthorized) WithPayload(payload *models.APIResponse) *UpdateReleaseUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update release unauthorized response
func (o *UpdateReleaseUnauthorized) SetPayload(payload *models.APIResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateReleaseUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateReleaseNotFoundCode is the HTTP code returned for type UpdateReleaseNotFound
const UpdateReleaseNotFoundCode int = 404

/*UpdateReleaseNotFound The specified resource was not found

swagger:response updateReleaseNotFound
*/
type UpdateReleaseNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.APIResponse `json:"body,omitempty"`
}

// NewUpdateReleaseNotFound creates UpdateReleaseNotFound with default headers values
func NewUpdateReleaseNotFound() *UpdateReleaseNotFound {
	return &UpdateReleaseNotFound{}
}

// WithPayload adds the payload to the update release not found response
func (o *UpdateReleaseNotFound) WithPayload(payload *models.APIResponse) *UpdateReleaseNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update release not found response
func (o *UpdateReleaseNotFound) SetPayload(payload *models.APIResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateReleaseNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateReleaseInternalServerErrorCode is the HTTP code returned for type UpdateReleaseInternalServerError
const UpdateReleaseInternalServerErrorCode int = 500

/*UpdateReleaseInternalServerError Internal Server Error

swagger:response updateReleaseInternalServerError
*/
type UpdateReleaseInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.APIResponse `json:"body,omitempty"`
}

// NewUpdateReleaseInternalServerError creates UpdateReleaseInternalServerError with default headers values
func NewUpdateReleaseInternalServerError() *UpdateReleaseInternalServerError {
	return &UpdateReleaseInternalServerError{}
}

// WithPayload adds the payload to the update release internal server error response
func (o *UpdateReleaseInternalServerError) WithPayload(payload *models.APIResponse) *UpdateReleaseInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update release internal server error response
func (o *UpdateReleaseInternalServerError) SetPayload(payload *models.APIResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateReleaseInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
