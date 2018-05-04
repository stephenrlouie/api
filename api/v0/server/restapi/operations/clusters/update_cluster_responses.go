// Code generated by go-swagger; DO NOT EDIT.

package clusters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/optikon/api/api/v0/models"
)

// UpdateClusterOKCode is the HTTP code returned for type UpdateClusterOK
const UpdateClusterOKCode int = 200

/*UpdateClusterOK OK

swagger:response updateClusterOK
*/
type UpdateClusterOK struct {
}

// NewUpdateClusterOK creates UpdateClusterOK with default headers values
func NewUpdateClusterOK() *UpdateClusterOK {
	return &UpdateClusterOK{}
}

// WriteResponse to the client
func (o *UpdateClusterOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// UpdateClusterBadRequestCode is the HTTP code returned for type UpdateClusterBadRequest
const UpdateClusterBadRequestCode int = 400

/*UpdateClusterBadRequest Bad Request

swagger:response updateClusterBadRequest
*/
type UpdateClusterBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.APIResponse `json:"body,omitempty"`
}

// NewUpdateClusterBadRequest creates UpdateClusterBadRequest with default headers values
func NewUpdateClusterBadRequest() *UpdateClusterBadRequest {
	return &UpdateClusterBadRequest{}
}

// WithPayload adds the payload to the update cluster bad request response
func (o *UpdateClusterBadRequest) WithPayload(payload *models.APIResponse) *UpdateClusterBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update cluster bad request response
func (o *UpdateClusterBadRequest) SetPayload(payload *models.APIResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateClusterBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateClusterUnauthorizedCode is the HTTP code returned for type UpdateClusterUnauthorized
const UpdateClusterUnauthorizedCode int = 401

/*UpdateClusterUnauthorized Unauthorized

swagger:response updateClusterUnauthorized
*/
type UpdateClusterUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.APIResponse `json:"body,omitempty"`
}

// NewUpdateClusterUnauthorized creates UpdateClusterUnauthorized with default headers values
func NewUpdateClusterUnauthorized() *UpdateClusterUnauthorized {
	return &UpdateClusterUnauthorized{}
}

// WithPayload adds the payload to the update cluster unauthorized response
func (o *UpdateClusterUnauthorized) WithPayload(payload *models.APIResponse) *UpdateClusterUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update cluster unauthorized response
func (o *UpdateClusterUnauthorized) SetPayload(payload *models.APIResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateClusterUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateClusterNotFoundCode is the HTTP code returned for type UpdateClusterNotFound
const UpdateClusterNotFoundCode int = 404

/*UpdateClusterNotFound The specified resource was not found

swagger:response updateClusterNotFound
*/
type UpdateClusterNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.APIResponse `json:"body,omitempty"`
}

// NewUpdateClusterNotFound creates UpdateClusterNotFound with default headers values
func NewUpdateClusterNotFound() *UpdateClusterNotFound {
	return &UpdateClusterNotFound{}
}

// WithPayload adds the payload to the update cluster not found response
func (o *UpdateClusterNotFound) WithPayload(payload *models.APIResponse) *UpdateClusterNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update cluster not found response
func (o *UpdateClusterNotFound) SetPayload(payload *models.APIResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateClusterNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateClusterInternalServerErrorCode is the HTTP code returned for type UpdateClusterInternalServerError
const UpdateClusterInternalServerErrorCode int = 500

/*UpdateClusterInternalServerError Internal Server Error

swagger:response updateClusterInternalServerError
*/
type UpdateClusterInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.APIResponse `json:"body,omitempty"`
}

// NewUpdateClusterInternalServerError creates UpdateClusterInternalServerError with default headers values
func NewUpdateClusterInternalServerError() *UpdateClusterInternalServerError {
	return &UpdateClusterInternalServerError{}
}

// WithPayload adds the payload to the update cluster internal server error response
func (o *UpdateClusterInternalServerError) WithPayload(payload *models.APIResponse) *UpdateClusterInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update cluster internal server error response
func (o *UpdateClusterInternalServerError) SetPayload(payload *models.APIResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateClusterInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
