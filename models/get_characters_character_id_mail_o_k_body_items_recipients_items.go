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

// GetCharactersCharacterIDMailOKBodyItemsRecipientsItems get_characters_character_id_mail_recipient
//
// recipient object
// swagger:model getCharactersCharacterIdMailOKBodyItemsRecipientsItems
type GetCharactersCharacterIDMailOKBodyItemsRecipientsItems struct {

	// get_characters_character_id_mail_recipient_id
	//
	// recipient_id integer
	// Required: true
	RecipientID *int32 `json:"recipient_id"`

	// get_characters_character_id_mail_recipient_type
	//
	// recipient_type string
	// Required: true
	RecipientType *string `json:"recipient_type"`
}

// Validate validates this get characters character Id mail o k body items recipients items
func (m *GetCharactersCharacterIDMailOKBodyItemsRecipientsItems) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRecipientID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRecipientType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetCharactersCharacterIDMailOKBodyItemsRecipientsItems) validateRecipientID(formats strfmt.Registry) error {

	if err := validate.Required("recipient_id", "body", m.RecipientID); err != nil {
		return err
	}

	return nil
}

var getCharactersCharacterIdMailOKBodyItemsRecipientsItemsTypeRecipientTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["alliance","character","corporation","mailing_list"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getCharactersCharacterIdMailOKBodyItemsRecipientsItemsTypeRecipientTypePropEnum = append(getCharactersCharacterIdMailOKBodyItemsRecipientsItemsTypeRecipientTypePropEnum, v)
	}
}

const (
	// GetCharactersCharacterIDMailOKBodyItemsRecipientsItemsRecipientTypeAlliance captures enum value "alliance"
	GetCharactersCharacterIDMailOKBodyItemsRecipientsItemsRecipientTypeAlliance string = "alliance"
	// GetCharactersCharacterIDMailOKBodyItemsRecipientsItemsRecipientTypeCharacter captures enum value "character"
	GetCharactersCharacterIDMailOKBodyItemsRecipientsItemsRecipientTypeCharacter string = "character"
	// GetCharactersCharacterIDMailOKBodyItemsRecipientsItemsRecipientTypeCorporation captures enum value "corporation"
	GetCharactersCharacterIDMailOKBodyItemsRecipientsItemsRecipientTypeCorporation string = "corporation"
	// GetCharactersCharacterIDMailOKBodyItemsRecipientsItemsRecipientTypeMailingList captures enum value "mailing_list"
	GetCharactersCharacterIDMailOKBodyItemsRecipientsItemsRecipientTypeMailingList string = "mailing_list"
)

// prop value enum
func (m *GetCharactersCharacterIDMailOKBodyItemsRecipientsItems) validateRecipientTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, getCharactersCharacterIdMailOKBodyItemsRecipientsItemsTypeRecipientTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *GetCharactersCharacterIDMailOKBodyItemsRecipientsItems) validateRecipientType(formats strfmt.Registry) error {

	if err := validate.Required("recipient_type", "body", m.RecipientType); err != nil {
		return err
	}

	// value enum
	if err := m.validateRecipientTypeEnum("recipient_type", "body", *m.RecipientType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetCharactersCharacterIDMailOKBodyItemsRecipientsItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetCharactersCharacterIDMailOKBodyItemsRecipientsItems) UnmarshalBinary(b []byte) error {
	var res GetCharactersCharacterIDMailOKBodyItemsRecipientsItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
