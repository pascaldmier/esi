package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GetKillmailsKillmailIDKillmailHashOKBodyVictimPosition get_killmails_killmail_id_killmail_hash_position
//
// Coordinates of the victim in Cartesian space relative to the Sun
//
// swagger:model getKillmailsKillmailIdKillmailHashOKBodyVictimPosition
type GetKillmailsKillmailIDKillmailHashOKBodyVictimPosition struct {

	// get_killmails_killmail_id_killmail_hash_x
	//
	// x number
	// Required: true
	X *float32 `json:"x"`

	// get_killmails_killmail_id_killmail_hash_y
	//
	// y number
	// Required: true
	Y *float32 `json:"y"`

	// get_killmails_killmail_id_killmail_hash_z
	//
	// z number
	// Required: true
	Z *float32 `json:"z"`
}

// Validate validates this get killmails killmail Id killmail hash o k body victim position
func (m *GetKillmailsKillmailIDKillmailHashOKBodyVictimPosition) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateX(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateY(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateZ(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetKillmailsKillmailIDKillmailHashOKBodyVictimPosition) validateX(formats strfmt.Registry) error {

	if err := validate.Required("x", "body", m.X); err != nil {
		return err
	}

	return nil
}

func (m *GetKillmailsKillmailIDKillmailHashOKBodyVictimPosition) validateY(formats strfmt.Registry) error {

	if err := validate.Required("y", "body", m.Y); err != nil {
		return err
	}

	return nil
}

func (m *GetKillmailsKillmailIDKillmailHashOKBodyVictimPosition) validateZ(formats strfmt.Registry) error {

	if err := validate.Required("z", "body", m.Z); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetKillmailsKillmailIDKillmailHashOKBodyVictimPosition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetKillmailsKillmailIDKillmailHashOKBodyVictimPosition) UnmarshalBinary(b []byte) error {
	var res GetKillmailsKillmailIDKillmailHashOKBodyVictimPosition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
