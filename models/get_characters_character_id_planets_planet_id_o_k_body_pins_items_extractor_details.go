package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GetCharactersCharacterIDPlanetsPlanetIDOKBodyPinsItemsExtractorDetails get_characters_character_id_planets_planet_id_extractor_details
//
// extractor_details object
// swagger:model getCharactersCharacterIdPlanetsPlanetIdOKBodyPinsItemsExtractorDetails
type GetCharactersCharacterIDPlanetsPlanetIDOKBodyPinsItemsExtractorDetails struct {

	// get_characters_character_id_planets_planet_id_cycle_time
	//
	// in seconds
	CycleTime int32 `json:"cycle_time,omitempty"`

	// get_characters_character_id_planets_planet_id_head_radius
	//
	// head_radius number
	HeadRadius float32 `json:"head_radius,omitempty"`

	// heads
	// Required: true
	Heads GetCharactersCharacterIDPlanetsPlanetIDOKBodyPinsItemsExtractorDetailsHeads `json:"heads"`

	// get_characters_character_id_planets_planet_id_product_type_id
	//
	// product_type_id integer
	ProductTypeID int32 `json:"product_type_id,omitempty"`

	// get_characters_character_id_planets_planet_id_qty_per_cycle
	//
	// qty_per_cycle integer
	QtyPerCycle int32 `json:"qty_per_cycle,omitempty"`
}

// Validate validates this get characters character Id planets planet Id o k body pins items extractor details
func (m *GetCharactersCharacterIDPlanetsPlanetIDOKBodyPinsItemsExtractorDetails) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHeads(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetCharactersCharacterIDPlanetsPlanetIDOKBodyPinsItemsExtractorDetails) validateHeads(formats strfmt.Registry) error {

	if err := validate.Required("heads", "body", m.Heads); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetCharactersCharacterIDPlanetsPlanetIDOKBodyPinsItemsExtractorDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetCharactersCharacterIDPlanetsPlanetIDOKBodyPinsItemsExtractorDetails) UnmarshalBinary(b []byte) error {
	var res GetCharactersCharacterIDPlanetsPlanetIDOKBodyPinsItemsExtractorDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}