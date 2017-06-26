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

// PostCharactersCharacterIDFittingsParamsBodyItems post_characters_character_id_fittings_items
//
// items array
// swagger:model postCharactersCharacterIdFittingsParamsBodyItems
type PostCharactersCharacterIDFittingsParamsBodyItems []*PostCharactersCharacterIDFittingsParamsBodyItemsItems

// Validate validates this post characters character Id fittings params body items
func (m PostCharactersCharacterIDFittingsParamsBodyItems) Validate(formats strfmt.Registry) error {
	var res []error

	iPostCharactersCharacterIDFittingsParamsBodyItemsSize := int64(len(m))

	if err := validate.MinItems("", "body", iPostCharactersCharacterIDFittingsParamsBodyItemsSize, 1); err != nil {
		return err
	}

	if err := validate.MaxItems("", "body", iPostCharactersCharacterIDFittingsParamsBodyItemsSize, 255); err != nil {
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
