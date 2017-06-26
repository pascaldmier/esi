package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// GetCorporationsNamesInternalServerErrorBody get_corporations_names_internal_server_error
//
// Internal server error
// swagger:model getCorporationsNamesInternalServerErrorBody
type GetCorporationsNamesInternalServerErrorBody struct {

	// get_corporations_names_500_internal_server_error
	//
	// Internal server error message
	Error string `json:"error,omitempty"`
}

// Validate validates this get corporations names internal server error body
func (m *GetCorporationsNamesInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *GetCorporationsNamesInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetCorporationsNamesInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res GetCorporationsNamesInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
