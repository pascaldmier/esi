package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// GetCharactersCharacterIDBookmarksFoldersOKBodyItems get_characters_character_id_bookmarks_folders_200_ok
//
// 200 ok object
// swagger:model getCharactersCharacterIdBookmarksFoldersOKBodyItems
type GetCharactersCharacterIDBookmarksFoldersOKBodyItems struct {

	// get_characters_character_id_bookmarks_folders_folder_id
	//
	// folder_id integer
	FolderID int32 `json:"folder_id,omitempty"`

	// get_characters_character_id_bookmarks_folders_name
	//
	// name string
	Name string `json:"name,omitempty"`

	// get_characters_character_id_bookmarks_folders_owner_id
	//
	// owner_id integer
	OwnerID int32 `json:"owner_id,omitempty"`
}

// Validate validates this get characters character Id bookmarks folders o k body items
func (m *GetCharactersCharacterIDBookmarksFoldersOKBodyItems) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *GetCharactersCharacterIDBookmarksFoldersOKBodyItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetCharactersCharacterIDBookmarksFoldersOKBodyItems) UnmarshalBinary(b []byte) error {
	var res GetCharactersCharacterIDBookmarksFoldersOKBodyItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
