package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// GetKillmailsKillmailIDKillmailHashUnprocessableEntityBody get_killmails_killmail_id_killmail_hash_unprocessable_entity
//
// Unprocessable entity
// swagger:model getKillmailsKillmailIdKillmailHashUnprocessableEntityBody
type GetKillmailsKillmailIDKillmailHashUnprocessableEntityBody struct {

	// get_killmails_killmail_id_killmail_hash_422_unprocessable_entity
	//
	// Unprocessable entity message
	Error string `json:"error,omitempty"`
}

// Validate validates this get killmails killmail Id killmail hash unprocessable entity body
func (m *GetKillmailsKillmailIDKillmailHashUnprocessableEntityBody) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *GetKillmailsKillmailIDKillmailHashUnprocessableEntityBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetKillmailsKillmailIDKillmailHashUnprocessableEntityBody) UnmarshalBinary(b []byte) error {
	var res GetKillmailsKillmailIDKillmailHashUnprocessableEntityBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}