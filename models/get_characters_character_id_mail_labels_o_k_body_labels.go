package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GetCharactersCharacterIDMailLabelsOKBodyLabels get_characters_character_id_mail_labels_labels
//
// labels array
// swagger:model getCharactersCharacterIdMailLabelsOKBodyLabels
type GetCharactersCharacterIDMailLabelsOKBodyLabels []*GetCharactersCharacterIDMailLabelsOKBodyLabelsItems

// Validate validates this get characters character Id mail labels o k body labels
func (m GetCharactersCharacterIDMailLabelsOKBodyLabels) Validate(formats strfmt.Registry) error {
	var res []error

	iGetCharactersCharacterIDMailLabelsOKBodyLabelsSize := int64(len(m))

	if err := validate.MaxItems("", "body", iGetCharactersCharacterIDMailLabelsOKBodyLabelsSize, 30); err != nil {
		return err
	}

	for i := 0; i < len(m); i++ {

		if swag.IsZero(m[i]) { // not required
			continue
		}

		if m[i] != nil {

			if err := m[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName(strconv.Itoa(i))
				}
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
