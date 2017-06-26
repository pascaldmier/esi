package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GetCharactersCharacterIDBookmarksOKBodyItemsTargetItem get_characters_character_id_bookmarks_item
//
// item object
// swagger:model getCharactersCharacterIdBookmarksOKBodyItemsTargetItem
type GetCharactersCharacterIDBookmarksOKBodyItemsTargetItem struct {

	// get_characters_character_id_bookmarks_item_id
	//
	// item_id integer
	// Required: true
	ItemID *int64 `json:"item_id"`

	// get_characters_character_id_bookmarks_type_id
	//
	// type_id integer
	// Required: true
	TypeID *int32 `json:"type_id"`
}

// Validate validates this get characters character Id bookmarks o k body items target item
func (m *GetCharactersCharacterIDBookmarksOKBodyItemsTargetItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateItemID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTypeID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetCharactersCharacterIDBookmarksOKBodyItemsTargetItem) validateItemID(formats strfmt.Registry) error {

	if err := validate.Required("item_id", "body", m.ItemID); err != nil {
		return err
	}

	return nil
}

func (m *GetCharactersCharacterIDBookmarksOKBodyItemsTargetItem) validateTypeID(formats strfmt.Registry) error {

	if err := validate.Required("type_id", "body", m.TypeID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetCharactersCharacterIDBookmarksOKBodyItemsTargetItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetCharactersCharacterIDBookmarksOKBodyItemsTargetItem) UnmarshalBinary(b []byte) error {
	var res GetCharactersCharacterIDBookmarksOKBodyItemsTargetItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
