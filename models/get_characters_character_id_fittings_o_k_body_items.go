package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GetCharactersCharacterIDFittingsOKBodyItems get_characters_character_id_fittings_200_ok
//
// 200 ok object
// swagger:model getCharactersCharacterIdFittingsOKBodyItems
type GetCharactersCharacterIDFittingsOKBodyItems struct {

	// get_characters_character_id_fittings_description
	//
	// description string
	// Required: true
	Description *string `json:"description"`

	// get_characters_character_id_fittings_fitting_id
	//
	// fitting_id integer
	// Required: true
	FittingID *int32 `json:"fitting_id"`

	// items
	// Required: true
	Items GetCharactersCharacterIDFittingsOKBodyItemsItems `json:"items"`

	// get_characters_character_id_fittings_name
	//
	// name string
	// Required: true
	Name *string `json:"name"`

	// get_characters_character_id_fittings_ship_type_id
	//
	// ship_type_id integer
	// Required: true
	ShipTypeID *int32 `json:"ship_type_id"`
}

// Validate validates this get characters character Id fittings o k body items
func (m *GetCharactersCharacterIDFittingsOKBodyItems) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateFittingID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateItems(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateShipTypeID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetCharactersCharacterIDFittingsOKBodyItems) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *GetCharactersCharacterIDFittingsOKBodyItems) validateFittingID(formats strfmt.Registry) error {

	if err := validate.Required("fitting_id", "body", m.FittingID); err != nil {
		return err
	}

	return nil
}

func (m *GetCharactersCharacterIDFittingsOKBodyItems) validateItems(formats strfmt.Registry) error {

	if err := validate.Required("items", "body", m.Items); err != nil {
		return err
	}

	return nil
}

func (m *GetCharactersCharacterIDFittingsOKBodyItems) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *GetCharactersCharacterIDFittingsOKBodyItems) validateShipTypeID(formats strfmt.Registry) error {

	if err := validate.Required("ship_type_id", "body", m.ShipTypeID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetCharactersCharacterIDFittingsOKBodyItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetCharactersCharacterIDFittingsOKBodyItems) UnmarshalBinary(b []byte) error {
	var res GetCharactersCharacterIDFittingsOKBodyItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
