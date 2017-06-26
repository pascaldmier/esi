package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DeleteFleetsFleetIDWingsWingIDForbiddenBody delete_fleets_fleet_id_wings_wing_id_forbidden
//
// Forbidden
// swagger:model deleteFleetsFleetIdWingsWingIdForbiddenBody
type DeleteFleetsFleetIDWingsWingIDForbiddenBody struct {

	// delete_fleets_fleet_id_wings_wing_id_403_forbidden
	//
	// Forbidden message
	Error string `json:"error,omitempty"`
}

// Validate validates this delete fleets fleet Id wings wing Id forbidden body
func (m *DeleteFleetsFleetIDWingsWingIDForbiddenBody) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *DeleteFleetsFleetIDWingsWingIDForbiddenBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeleteFleetsFleetIDWingsWingIDForbiddenBody) UnmarshalBinary(b []byte) error {
	var res DeleteFleetsFleetIDWingsWingIDForbiddenBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
