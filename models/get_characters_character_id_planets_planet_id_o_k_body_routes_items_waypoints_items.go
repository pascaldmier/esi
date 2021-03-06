package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GetCharactersCharacterIDPlanetsPlanetIDOKBodyRoutesItemsWaypointsItems get_characters_character_id_planets_planet_id_waypoint
//
// waypoint object
// swagger:model getCharactersCharacterIdPlanetsPlanetIdOKBodyRoutesItemsWaypointsItems
type GetCharactersCharacterIDPlanetsPlanetIDOKBodyRoutesItemsWaypointsItems struct {

	// get_characters_character_id_planets_planet_id_order
	//
	// order integer
	// Required: true
	// Maximum: 5
	// Minimum: 1
	Order *int32 `json:"order"`

	// get_characters_character_id_planets_planet_id_pin_id
	//
	// pin_id integer
	// Required: true
	PinID *int64 `json:"pin_id"`
}

// Validate validates this get characters character Id planets planet Id o k body routes items waypoints items
func (m *GetCharactersCharacterIDPlanetsPlanetIDOKBodyRoutesItemsWaypointsItems) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOrder(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePinID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetCharactersCharacterIDPlanetsPlanetIDOKBodyRoutesItemsWaypointsItems) validateOrder(formats strfmt.Registry) error {

	if err := validate.Required("order", "body", m.Order); err != nil {
		return err
	}

	if err := validate.MinimumInt("order", "body", int64(*m.Order), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("order", "body", int64(*m.Order), 5, false); err != nil {
		return err
	}

	return nil
}

func (m *GetCharactersCharacterIDPlanetsPlanetIDOKBodyRoutesItemsWaypointsItems) validatePinID(formats strfmt.Registry) error {

	if err := validate.Required("pin_id", "body", m.PinID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetCharactersCharacterIDPlanetsPlanetIDOKBodyRoutesItemsWaypointsItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetCharactersCharacterIDPlanetsPlanetIDOKBodyRoutesItemsWaypointsItems) UnmarshalBinary(b []byte) error {
	var res GetCharactersCharacterIDPlanetsPlanetIDOKBodyRoutesItemsWaypointsItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
