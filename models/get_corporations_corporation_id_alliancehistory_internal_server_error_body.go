package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// GetCorporationsCorporationIDAlliancehistoryInternalServerErrorBody get_corporations_corporation_id_alliancehistory_internal_server_error
//
// Internal server error
// swagger:model getCorporationsCorporationIdAlliancehistoryInternalServerErrorBody
type GetCorporationsCorporationIDAlliancehistoryInternalServerErrorBody struct {

	// get_corporations_corporation_id_alliancehistory_500_internal_server_error
	//
	// Internal server error message
	Error string `json:"error,omitempty"`
}

// Validate validates this get corporations corporation Id alliancehistory internal server error body
func (m *GetCorporationsCorporationIDAlliancehistoryInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *GetCorporationsCorporationIDAlliancehistoryInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetCorporationsCorporationIDAlliancehistoryInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res GetCorporationsCorporationIDAlliancehistoryInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}