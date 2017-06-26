package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GetUniverseRegionsRegionIDOKBody get_universe_regions_region_id_ok
//
// 200 ok object
// swagger:model getUniverseRegionsRegionIdOKBody
type GetUniverseRegionsRegionIDOKBody struct {

	// get_universe_regions_region_id_constellations
	//
	// constellations array
	// Required: true
	// Max Items: 1000
	Constellations []int32 `json:"constellations"`

	// get_universe_regions_region_id_description
	//
	// description string
	Description string `json:"description,omitempty"`

	// get_universe_regions_region_id_name
	//
	// name string
	// Required: true
	Name *string `json:"name"`

	// get_universe_regions_region_id_region_id
	//
	// region_id integer
	// Required: true
	RegionID *int32 `json:"region_id"`
}

// Validate validates this get universe regions region Id o k body
func (m *GetUniverseRegionsRegionIDOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConstellations(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRegionID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetUniverseRegionsRegionIDOKBody) validateConstellations(formats strfmt.Registry) error {

	if err := validate.Required("constellations", "body", m.Constellations); err != nil {
		return err
	}

	iConstellationsSize := int64(len(m.Constellations))

	if err := validate.MaxItems("constellations", "body", iConstellationsSize, 1000); err != nil {
		return err
	}

	return nil
}

func (m *GetUniverseRegionsRegionIDOKBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *GetUniverseRegionsRegionIDOKBody) validateRegionID(formats strfmt.Registry) error {

	if err := validate.Required("region_id", "body", m.RegionID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetUniverseRegionsRegionIDOKBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetUniverseRegionsRegionIDOKBody) UnmarshalBinary(b []byte) error {
	var res GetUniverseRegionsRegionIDOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
