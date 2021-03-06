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

// GetCharactersCharacterIDChatChannelsOKBodyItemsMutedItems get_characters_character_id_chat_channels_muted
//
// muted object
// swagger:model getCharactersCharacterIdChatChannelsOKBodyItemsMutedItems
type GetCharactersCharacterIDChatChannelsOKBodyItemsMutedItems struct {

	// get_characters_character_id_chat_channels_accessor_id
	//
	// ID of a muted channel member
	// Required: true
	AccessorID *int32 `json:"accessor_id"`

	// get_characters_character_id_chat_channels_accessor_type
	//
	// accessor_type string
	// Required: true
	AccessorType *string `json:"accessor_type"`

	// get_characters_character_id_chat_channels_end_at
	//
	// Time at which this accessor will no longer be muted
	EndAt strfmt.DateTime `json:"end_at,omitempty"`

	// get_characters_character_id_chat_channels_reason
	//
	// Reason this accessor is muted
	Reason string `json:"reason,omitempty"`
}

// Validate validates this get characters character Id chat channels o k body items muted items
func (m *GetCharactersCharacterIDChatChannelsOKBodyItemsMutedItems) Validate(formats strfmt.Registry) error {
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

func (m *GetCharactersCharacterIDChatChannelsOKBodyItemsMutedItems) validateAccessorID(formats strfmt.Registry) error {

	if err := validate.Required("accessor_id", "body", m.AccessorID); err != nil {
		return err
	}

	return nil
}

var getCharactersCharacterIdChatChannelsOKBodyItemsMutedItemsTypeAccessorTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["character","corporation","alliance"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getCharactersCharacterIdChatChannelsOKBodyItemsMutedItemsTypeAccessorTypePropEnum = append(getCharactersCharacterIdChatChannelsOKBodyItemsMutedItemsTypeAccessorTypePropEnum, v)
	}
}

const (
	// GetCharactersCharacterIDChatChannelsOKBodyItemsMutedItemsAccessorTypeCharacter captures enum value "character"
	GetCharactersCharacterIDChatChannelsOKBodyItemsMutedItemsAccessorTypeCharacter string = "character"
	// GetCharactersCharacterIDChatChannelsOKBodyItemsMutedItemsAccessorTypeCorporation captures enum value "corporation"
	GetCharactersCharacterIDChatChannelsOKBodyItemsMutedItemsAccessorTypeCorporation string = "corporation"
	// GetCharactersCharacterIDChatChannelsOKBodyItemsMutedItemsAccessorTypeAlliance captures enum value "alliance"
	GetCharactersCharacterIDChatChannelsOKBodyItemsMutedItemsAccessorTypeAlliance string = "alliance"
)

// prop value enum
func (m *GetCharactersCharacterIDChatChannelsOKBodyItemsMutedItems) validateAccessorTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, getCharactersCharacterIdChatChannelsOKBodyItemsMutedItemsTypeAccessorTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *GetCharactersCharacterIDChatChannelsOKBodyItemsMutedItems) validateAccessorType(formats strfmt.Registry) error {

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
func (m *GetCharactersCharacterIDChatChannelsOKBodyItemsMutedItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetCharactersCharacterIDChatChannelsOKBodyItemsMutedItems) UnmarshalBinary(b []byte) error {
	var res GetCharactersCharacterIDChatChannelsOKBodyItemsMutedItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
