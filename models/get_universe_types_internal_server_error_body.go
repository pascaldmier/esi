package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// GetUniverseTypesInternalServerErrorBody get_universe_types_internal_server_error
//
// Internal server error
// swagger:model getUniverseTypesInternalServerErrorBody
type GetUniverseTypesInternalServerErrorBody struct {

	// get_universe_types_500_internal_server_error
	//
	// Internal server error message
	Error string `json:"error,omitempty"`
}

// Validate validates this get universe types internal server error body
func (m *GetUniverseTypesInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *GetUniverseTypesInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetUniverseTypesInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res GetUniverseTypesInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
