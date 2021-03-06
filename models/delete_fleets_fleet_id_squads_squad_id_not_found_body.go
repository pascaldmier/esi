package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DeleteFleetsFleetIDSquadsSquadIDNotFoundBody delete_fleets_fleet_id_squads_squad_id_not_found
//
// Not found
// swagger:model deleteFleetsFleetIdSquadsSquadIdNotFoundBody
type DeleteFleetsFleetIDSquadsSquadIDNotFoundBody struct {

	// delete_fleets_fleet_id_squads_squad_id_404_not_found
	//
	// Not found message
	Error string `json:"error,omitempty"`
}

// Validate validates this delete fleets fleet Id squads squad Id not found body
func (m *DeleteFleetsFleetIDSquadsSquadIDNotFoundBody) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *DeleteFleetsFleetIDSquadsSquadIDNotFoundBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeleteFleetsFleetIDSquadsSquadIDNotFoundBody) UnmarshalBinary(b []byte) error {
	var res DeleteFleetsFleetIDSquadsSquadIDNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
