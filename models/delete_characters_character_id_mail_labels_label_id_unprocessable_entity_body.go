package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DeleteCharactersCharacterIDMailLabelsLabelIDUnprocessableEntityBody delete_characters_character_id_mail_labels_label_id_unprocessable_entity
//
// Unprocessable entity
// swagger:model deleteCharactersCharacterIdMailLabelsLabelIdUnprocessableEntityBody
type DeleteCharactersCharacterIDMailLabelsLabelIDUnprocessableEntityBody struct {

	// delete_characters_character_id_mail_labels_label_id_422_unprocessable_entity
	//
	// Unprocessable entity message
	Error string `json:"error,omitempty"`
}

// Validate validates this delete characters character Id mail labels label Id unprocessable entity body
func (m *DeleteCharactersCharacterIDMailLabelsLabelIDUnprocessableEntityBody) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *DeleteCharactersCharacterIDMailLabelsLabelIDUnprocessableEntityBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeleteCharactersCharacterIDMailLabelsLabelIDUnprocessableEntityBody) UnmarshalBinary(b []byte) error {
	var res DeleteCharactersCharacterIDMailLabelsLabelIDUnprocessableEntityBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
