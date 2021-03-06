package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// GetMarketsStructuresStructureIDForbiddenBody get_markets_structures_structure_id_forbidden
//
// Forbidden
// swagger:model getMarketsStructuresStructureIdForbiddenBody
type GetMarketsStructuresStructureIDForbiddenBody struct {

	// get_markets_structures_structure_id_403_forbidden
	//
	// Forbidden message
	Error string `json:"error,omitempty"`
}

// Validate validates this get markets structures structure Id forbidden body
func (m *GetMarketsStructuresStructureIDForbiddenBody) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *GetMarketsStructuresStructureIDForbiddenBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetMarketsStructuresStructureIDForbiddenBody) UnmarshalBinary(b []byte) error {
	var res GetMarketsStructuresStructureIDForbiddenBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
