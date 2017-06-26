package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GetCharactersCharacterIDBookmarksOKBodyItemsTargetCoordinates get_characters_character_id_bookmarks_coordinates
//
// coordinates object
// swagger:model getCharactersCharacterIdBookmarksOKBodyItemsTargetCoordinates
type GetCharactersCharacterIDBookmarksOKBodyItemsTargetCoordinates struct {

	// get_characters_character_id_bookmarks_x
	//
	// x number
	// Required: true
	X *float64 `json:"x"`

	// get_characters_character_id_bookmarks_y
	//
	// y number
	// Required: true
	Y *float64 `json:"y"`

	// get_characters_character_id_bookmarks_z
	//
	// z number
	// Required: true
	Z *float64 `json:"z"`
}

// Validate validates this get characters character Id bookmarks o k body items target coordinates
func (m *GetCharactersCharacterIDBookmarksOKBodyItemsTargetCoordinates) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateX(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateY(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateZ(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetCharactersCharacterIDBookmarksOKBodyItemsTargetCoordinates) validateX(formats strfmt.Registry) error {

	if err := validate.Required("x", "body", m.X); err != nil {
		return err
	}

	return nil
}

func (m *GetCharactersCharacterIDBookmarksOKBodyItemsTargetCoordinates) validateY(formats strfmt.Registry) error {

	if err := validate.Required("y", "body", m.Y); err != nil {
		return err
	}

	return nil
}

func (m *GetCharactersCharacterIDBookmarksOKBodyItemsTargetCoordinates) validateZ(formats strfmt.Registry) error {

	if err := validate.Required("z", "body", m.Z); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetCharactersCharacterIDBookmarksOKBodyItemsTargetCoordinates) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetCharactersCharacterIDBookmarksOKBodyItemsTargetCoordinates) UnmarshalBinary(b []byte) error {
	var res GetCharactersCharacterIDBookmarksOKBodyItemsTargetCoordinates
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}