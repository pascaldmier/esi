package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GetCharactersCharacterIDChatChannelsOKBodyItemsAllowedItems get_characters_character_id_chat_channels_allowed
//
// allowed object
// swagger:model getCharactersCharacterIdChatChannelsOKBodyItemsAllowedItems
type GetCharactersCharacterIDChatChannelsOKBodyItemsAllowedItems struct {

	// get_characters_character_id_chat_channels_accessor_id
	//
	// ID of an allowed channel member
	// Required: true
	AccessorID *int32 `json:"accessor_id"`

	// get_characters_character_id_chat_channels_accessor_type
	//
	// accessor_type string
	// Required: true
	AccessorType *string `json:"accessor_type"`
}

// Validate validates this get characters character Id chat channels o k body items allowed items
func (m *GetCharactersCharacterIDChatChannelsOKBodyItemsAllowedItems) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccessorID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateAccessorType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetCharactersCharacterIDChatChannelsOKBodyItemsAllowedItems) validateAccessorID(formats strfmt.Registry) error {

	if err := validate.Required("accessor_id", "body", m.AccessorID); err != nil {
		return err
	}

	return nil
}

var getCharactersCharacterIdChatChannelsOKBodyItemsAllowedItemsTypeAccessorTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["character","corporation","alliance"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getCharactersCharacterIdChatChannelsOKBodyItemsAllowedItemsTypeAccessorTypePropEnum = append(getCharactersCharacterIdChatChannelsOKBodyItemsAllowedItemsTypeAccessorTypePropEnum, v)
	}
}

const (
	// GetCharactersCharacterIDChatChannelsOKBodyItemsAllowedItemsAccessorTypeCharacter captures enum value "character"
	GetCharactersCharacterIDChatChannelsOKBodyItemsAllowedItemsAccessorTypeCharacter string = "character"
	// GetCharactersCharacterIDChatChannelsOKBodyItemsAllowedItemsAccessorTypeCorporation captures enum value "corporation"
	GetCharactersCharacterIDChatChannelsOKBodyItemsAllowedItemsAccessorTypeCorporation string = "corporation"
	// GetCharactersCharacterIDChatChannelsOKBodyItemsAllowedItemsAccessorTypeAlliance captures enum value "alliance"
	GetCharactersCharacterIDChatChannelsOKBodyItemsAllowedItemsAccessorTypeAlliance string = "alliance"
)

// prop value enum
func (m *GetCharactersCharacterIDChatChannelsOKBodyItemsAllowedItems) validateAccessorTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, getCharactersCharacterIdChatChannelsOKBodyItemsAllowedItemsTypeAccessorTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *GetCharactersCharacterIDChatChannelsOKBodyItemsAllowedItems) validateAccessorType(formats strfmt.Registry) error {

	if err := validate.Required("accessor_type", "body", m.AccessorType); err != nil {
		return err
	}

	// value enum
	if err := m.validateAccessorTypeEnum("accessor_type", "body", *m.AccessorType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetCharactersCharacterIDChatChannelsOKBodyItemsAllowedItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetCharactersCharacterIDChatChannelsOKBodyItemsAllowedItems) UnmarshalBinary(b []byte) error {
	var res GetCharactersCharacterIDChatChannelsOKBodyItemsAllowedItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
