// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ReleaseTestRun release test run
// swagger:model release.TestRun
type ReleaseTestRun struct {

	// completed at
	CompletedAt string `json:"CompletedAt,omitempty"`

	// info
	Info string `json:"Info,omitempty"`

	// name
	Name string `json:"Name,omitempty"`

	// started at
	StartedAt string `json:"StartedAt,omitempty"`

	// status
	Status string `json:"Status,omitempty"`
}

// Validate validates this release test run
func (m *ReleaseTestRun) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *ReleaseTestRun) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReleaseTestRun) UnmarshalBinary(b []byte) error {
	var res ReleaseTestRun
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
