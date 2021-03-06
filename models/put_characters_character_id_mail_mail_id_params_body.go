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

// PutCharactersCharacterIDMailMailIDParamsBody put_characters_character_id_mail_mail_id_contents
//
// contents object
// swagger:model putCharactersCharacterIdMailMailIdParamsBody
type PutCharactersCharacterIDMailMailIDParamsBody struct {

	// put_characters_character_id_mail_mail_id_labels
	//
	// Labels to assign to the mail. Pre-existing labels are unassigned.
	// Max Items: 25
	Labels []*int64 `json:"labels"`

	// put_characters_character_id_mail_mail_id_read
	//
	// Whether the mail is flagged as read
	Read bool `json:"read,omitempty"`
}

// Validate validates this put characters character Id mail mail Id params body
func (m *PutCharactersCharacterIDMailMailIDParamsBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLabels(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PutCharactersCharacterIDMailMailIDParamsBody) validateLabels(formats strfmt.Registry) error {

	if swag.IsZero(m.Labels) { // not required
		return nil
	}

	iLabelsSize := int64(len(m.Labels))

	if err := validate.MaxItems("labels", "body", iLabelsSize, 25); err != nil {
		return err
	}

	for i := 0; i < len(m.Labels); i++ {

		if swag.IsZero(m.Labels[i]) { // not required
			continue
		}

		if err := validate.MinimumInt("labels"+"."+strconv.Itoa(i), "body", int64(*m.Labels[i]), 0, false); err != nil {
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PutCharactersCharacterIDMailMailIDParamsBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PutCharactersCharacterIDMailMailIDParamsBody) UnmarshalBinary(b []byte) error {
	var res PutCharactersCharacterIDMailMailIDParamsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
