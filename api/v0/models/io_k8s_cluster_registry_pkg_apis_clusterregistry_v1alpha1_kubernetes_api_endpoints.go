// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// IoK8sClusterRegistryPkgApisClusterregistryV1alpha1KubernetesAPIEndpoints KubernetesAPIEndpoints represents the endpoints for one and only one Kubernetes API server.
// swagger:model io.k8s.cluster-registry.pkg.apis.clusterregistry.v1alpha1.KubernetesAPIEndpoints
type IoK8sClusterRegistryPkgApisClusterregistryV1alpha1KubernetesAPIEndpoints struct {

	// CABundle contains the certificate authority information.
	CaBundle strfmt.Base64 `json:"caBundle,omitempty"`

	// server endpoints
	ServerEndpoints IoK8sClusterRegistryPkgApisClusterregistryV1alpha1KubernetesAPIEndpointsServerEndpoints `json:"serverEndpoints"`
}

// Validate validates this io k8s cluster registry pkg apis clusterregistry v1alpha1 kubernetes API endpoints
func (m *IoK8sClusterRegistryPkgApisClusterregistryV1alpha1KubernetesAPIEndpoints) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sClusterRegistryPkgApisClusterregistryV1alpha1KubernetesAPIEndpoints) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sClusterRegistryPkgApisClusterregistryV1alpha1KubernetesAPIEndpoints) UnmarshalBinary(b []byte) error {
	var res IoK8sClusterRegistryPkgApisClusterregistryV1alpha1KubernetesAPIEndpoints
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
