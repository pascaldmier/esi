package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// GetUniverseRegionsRegionIDNotFoundBody get_universe_regions_region_id_not_found
//
// Not found
// swagger:model getUniverseRegionsRegionIdNotFoundBody
type GetUniverseRegionsRegionIDNotFoundBody struct {

	// get_universe_regions_region_id_404_not_found
	//
	// Not found message
	Error string `json:"error,omitempty"`
}

// Validate validates this get universe regions region Id not found body
func (m *GetUniverseRegionsRegionIDNotFoundBody) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *GetUniverseRegionsRegionIDNotFoundBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetUniverseRegionsRegionIDNotFoundBody) UnmarshalBinary(b []byte) error {
	var res GetUniverseRegionsRegionIDNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
