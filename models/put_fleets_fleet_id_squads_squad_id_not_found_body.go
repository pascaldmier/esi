package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// PutFleetsFleetIDSquadsSquadIDNotFoundBody put_fleets_fleet_id_squads_squad_id_not_found
//
// Not found
// swagger:model putFleetsFleetIdSquadsSquadIdNotFoundBody
type PutFleetsFleetIDSquadsSquadIDNotFoundBody struct {

	// put_fleets_fleet_id_squads_squad_id_404_not_found
	//
	// Not found message
	Error string `json:"error,omitempty"`
}

// Validate validates this put fleets fleet Id squads squad Id not found body
func (m *PutFleetsFleetIDSquadsSquadIDNotFoundBody) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *PutFleetsFleetIDSquadsSquadIDNotFoundBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PutFleetsFleetIDSquadsSquadIDNotFoundBody) UnmarshalBinary(b []byte) error {
	var res PutFleetsFleetIDSquadsSquadIDNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}