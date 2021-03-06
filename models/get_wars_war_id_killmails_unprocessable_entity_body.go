package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// GetWarsWarIDKillmailsUnprocessableEntityBody get_wars_war_id_killmails_unprocessable_entity
//
// Unprocessable entity
// swagger:model getWarsWarIdKillmailsUnprocessableEntityBody
type GetWarsWarIDKillmailsUnprocessableEntityBody struct {

	// get_wars_war_id_killmails_422_unprocessable_entity
	//
	// Unprocessable entity message
	Error string `json:"error,omitempty"`
}

// Validate validates this get wars war Id killmails unprocessable entity body
func (m *GetWarsWarIDKillmailsUnprocessableEntityBody) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *GetWarsWarIDKillmailsUnprocessableEntityBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetWarsWarIDKillmailsUnprocessableEntityBody) UnmarshalBinary(b []byte) error {
	var res GetWarsWarIDKillmailsUnprocessableEntityBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
