package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// GetUniverseTypesTypeIDInternalServerErrorBody get_universe_types_type_id_internal_server_error
//
// Internal server error
// swagger:model getUniverseTypesTypeIdInternalServerErrorBody
type GetUniverseTypesTypeIDInternalServerErrorBody struct {

	// get_universe_types_type_id_500_internal_server_error
	//
	// Internal server error message
	Error string `json:"error,omitempty"`
}

// Validate validates this get universe types type Id internal server error body
func (m *GetUniverseTypesTypeIDInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *GetUniverseTypesTypeIDInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetUniverseTypesTypeIDInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res GetUniverseTypesTypeIDInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}