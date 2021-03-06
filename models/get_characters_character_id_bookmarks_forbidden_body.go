package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// GetCharactersCharacterIDBookmarksForbiddenBody get_characters_character_id_bookmarks_forbidden
//
// Forbidden
// swagger:model getCharactersCharacterIdBookmarksForbiddenBody
type GetCharactersCharacterIDBookmarksForbiddenBody struct {

	// get_characters_character_id_bookmarks_403_forbidden
	//
	// Forbidden message
	Error string `json:"error,omitempty"`
}

// Validate validates this get characters character Id bookmarks forbidden body
func (m *GetCharactersCharacterIDBookmarksForbiddenBody) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *GetCharactersCharacterIDBookmarksForbiddenBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetCharactersCharacterIDBookmarksForbiddenBody) UnmarshalBinary(b []byte) error {
	var res GetCharactersCharacterIDBookmarksForbiddenBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
