package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// GetFleetsFleetIDWingsForbiddenBody get_fleets_fleet_id_wings_forbidden
//
// Forbidden
// swagger:model getFleetsFleetIdWingsForbiddenBody
type GetFleetsFleetIDWingsForbiddenBody struct {

	// get_fleets_fleet_id_wings_403_forbidden
	//
	// Forbidden message
	Error string `json:"error,omitempty"`
}

// Validate validates this get fleets fleet Id wings forbidden body
func (m *GetFleetsFleetIDWingsForbiddenBody) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *GetFleetsFleetIDWingsForbiddenBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetFleetsFleetIDWingsForbiddenBody) UnmarshalBinary(b []byte) error {
	var res GetFleetsFleetIDWingsForbiddenBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
