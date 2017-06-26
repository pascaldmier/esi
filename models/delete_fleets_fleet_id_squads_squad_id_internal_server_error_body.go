package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DeleteFleetsFleetIDSquadsSquadIDInternalServerErrorBody delete_fleets_fleet_id_squads_squad_id_internal_server_error
//
// Internal server error
// swagger:model deleteFleetsFleetIdSquadsSquadIdInternalServerErrorBody
type DeleteFleetsFleetIDSquadsSquadIDInternalServerErrorBody struct {

	// delete_fleets_fleet_id_squads_squad_id_500_internal_server_error
	//
	// Internal server error message
	Error string `json:"error,omitempty"`
}

// Validate validates this delete fleets fleet Id squads squad Id internal server error body
func (m *DeleteFleetsFleetIDSquadsSquadIDInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *DeleteFleetsFleetIDSquadsSquadIDInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeleteFleetsFleetIDSquadsSquadIDInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res DeleteFleetsFleetIDSquadsSquadIDInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}