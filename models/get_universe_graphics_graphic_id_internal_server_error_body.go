package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// GetUniverseGraphicsGraphicIDInternalServerErrorBody get_universe_graphics_graphic_id_internal_server_error
//
// Internal server error
// swagger:model getUniverseGraphicsGraphicIdInternalServerErrorBody
type GetUniverseGraphicsGraphicIDInternalServerErrorBody struct {

	// get_universe_graphics_graphic_id_500_internal_server_error
	//
	// Internal server error message
	Error string `json:"error,omitempty"`
}

// Validate validates this get universe graphics graphic Id internal server error body
func (m *GetUniverseGraphicsGraphicIDInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *GetUniverseGraphicsGraphicIDInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetUniverseGraphicsGraphicIDInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res GetUniverseGraphicsGraphicIDInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
