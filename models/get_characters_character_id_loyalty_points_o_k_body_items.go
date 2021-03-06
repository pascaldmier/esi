package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GetCharactersCharacterIDLoyaltyPointsOKBodyItems get_characters_character_id_loyalty_points_200_ok
//
// 200 ok object
// swagger:model getCharactersCharacterIdLoyaltyPointsOKBodyItems
type GetCharactersCharacterIDLoyaltyPointsOKBodyItems struct {

	// get_characters_character_id_loyalty_points_corporation_id
	//
	// corporation_id integer
	// Required: true
	CorporationID *int32 `json:"corporation_id"`

	// get_characters_character_id_loyalty_points_loyalty_points
	//
	// loyalty_points integer
	// Required: true
	LoyaltyPoints *int32 `json:"loyalty_points"`
}

// Validate validates this get characters character Id loyalty points o k body items
func (m *GetCharactersCharacterIDLoyaltyPointsOKBodyItems) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCorporationID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateLoyaltyPoints(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetCharactersCharacterIDLoyaltyPointsOKBodyItems) validateCorporationID(formats strfmt.Registry) error {

	if err := validate.Required("corporation_id", "body", m.CorporationID); err != nil {
		return err
	}

	return nil
}

func (m *GetCharactersCharacterIDLoyaltyPointsOKBodyItems) validateLoyaltyPoints(formats strfmt.Registry) error {

	if err := validate.Required("loyalty_points", "body", m.LoyaltyPoints); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetCharactersCharacterIDLoyaltyPointsOKBodyItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetCharactersCharacterIDLoyaltyPointsOKBodyItems) UnmarshalBinary(b []byte) error {
	var res GetCharactersCharacterIDLoyaltyPointsOKBodyItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
