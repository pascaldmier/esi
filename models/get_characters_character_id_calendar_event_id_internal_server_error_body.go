package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// GetCharactersCharacterIDCalendarEventIDInternalServerErrorBody get_characters_character_id_calendar_event_id_internal_server_error
//
// Internal server error
// swagger:model getCharactersCharacterIdCalendarEventIdInternalServerErrorBody
type GetCharactersCharacterIDCalendarEventIDInternalServerErrorBody struct {

	// get_characters_character_id_calendar_event_id_500_internal_server_error
	//
	// Internal server error message
	Error string `json:"error,omitempty"`
}

// Validate validates this get characters character Id calendar event Id internal server error body
func (m *GetCharactersCharacterIDCalendarEventIDInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *GetCharactersCharacterIDCalendarEventIDInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetCharactersCharacterIDCalendarEventIDInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res GetCharactersCharacterIDCalendarEventIDInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
