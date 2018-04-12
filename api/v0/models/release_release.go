// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ReleaseRelease release release
// swagger:model release.Release
type ReleaseRelease struct {

	// chart
	Chart *ChartChart `json:"Chart,omitempty"`

	// config
	Config *ChartConfig `json:"Config,omitempty"`

	// hooks
	Hooks ReleaseReleaseHooks `json:"Hooks"`

	// info
	Info *ReleaseInfo `json:"Info,omitempty"`

	// manifest
	Manifest string `json:"Manifest,omitempty"`

	// name
	Name string `json:"Name,omitempty"`

	// namespace
	Namespace string `json:"Namespace,omitempty"`

	// version
	Version int32 `json:"Version,omitempty"`
}

// Validate validates this release release
func (m *ReleaseRelease) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChart(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateConfig(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateInfo(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReleaseRelease) validateChart(formats strfmt.Registry) error {

	if swag.IsZero(m.Chart) { // not required
		return nil
	}

	if m.Chart != nil {

		if err := m.Chart.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Chart")
			}
			return err
		}
	}

	return nil
}

func (m *ReleaseRelease) validateConfig(formats strfmt.Registry) error {

	if swag.IsZero(m.Config) { // not required
		return nil
	}

	if m.Config != nil {

		if err := m.Config.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Config")
			}
			return err
		}
	}

	return nil
}

func (m *ReleaseRelease) validateInfo(formats strfmt.Registry) error {

	if swag.IsZero(m.Info) { // not required
		return nil
	}

	if m.Info != nil {

		if err := m.Info.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Info")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReleaseRelease) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReleaseRelease) UnmarshalBinary(b []byte) error {
	var res ReleaseRelease
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
