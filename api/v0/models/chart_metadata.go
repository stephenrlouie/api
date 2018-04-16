// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ChartMetadata chart metadata
// swagger:model chart.Metadata
type ChartMetadata struct {

	// annotations
	Annotations ChartMap `json:"Annotations"`

	// Api version
	APIVersion string `json:"ApiVersion,omitempty"`

	// app version
	AppVersion string `json:"AppVersion,omitempty"`

	// condition
	Condition string `json:"Condition,omitempty"`

	// deprecated
	Deprecated bool `json:"Deprecated,omitempty"`

	// description
	Description string `json:"Description,omitempty"`

	// engine
	Engine string `json:"Engine,omitempty"`

	// home
	Home string `json:"Home,omitempty"`

	// icon
	Icon string `json:"Icon,omitempty"`

	// keywords
	Keywords []string `json:"Keywords"`

	// kube version
	KubeVersion string `json:"KubeVersion,omitempty"`

	// maintainers
	Maintainers ChartMetadataMaintainers `json:"Maintainers"`

	// name
	Name string `json:"Name,omitempty"`

	// sources
	Sources []string `json:"Sources"`

	// tags
	Tags string `json:"Tags,omitempty"`

	// tiller version
	TillerVersion string `json:"TillerVersion,omitempty"`

	// version
	Version string `json:"Version,omitempty"`
}

// Validate validates this chart metadata
func (m *ChartMetadata) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEngine(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateKeywords(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSources(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var chartMetadataTypeEnginePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["UNKNOWN","GOTPL"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		chartMetadataTypeEnginePropEnum = append(chartMetadataTypeEnginePropEnum, v)
	}
}

const (
	// ChartMetadataEngineUNKNOWN captures enum value "UNKNOWN"
	ChartMetadataEngineUNKNOWN string = "UNKNOWN"
	// ChartMetadataEngineGOTPL captures enum value "GOTPL"
	ChartMetadataEngineGOTPL string = "GOTPL"
)

// prop value enum
func (m *ChartMetadata) validateEngineEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, chartMetadataTypeEnginePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ChartMetadata) validateEngine(formats strfmt.Registry) error {

	if swag.IsZero(m.Engine) { // not required
		return nil
	}

	// value enum
	if err := m.validateEngineEnum("Engine", "body", m.Engine); err != nil {
		return err
	}

	return nil
}

func (m *ChartMetadata) validateKeywords(formats strfmt.Registry) error {

	if swag.IsZero(m.Keywords) { // not required
		return nil
	}

	return nil
}

func (m *ChartMetadata) validateSources(formats strfmt.Registry) error {

	if swag.IsZero(m.Sources) { // not required
		return nil
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ChartMetadata) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ChartMetadata) UnmarshalBinary(b []byte) error {
	var res ChartMetadata
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}