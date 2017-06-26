package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GetStatusOKBody get_status_ok
//
// 200 ok object
// swagger:model getStatusOKBody
type GetStatusOKBody struct {

	// get_status_players
	//
	// Current online player count
	// Required: true
	Players *int64 `json:"players"`

	// get_status_server_version
	//
	// Running version as string
	// Required: true
	ServerVersion *string `json:"server_version"`

	// get_status_start_time
	//
	// Server start timestamp
	// Required: true
	StartTime *strfmt.DateTime `json:"start_time"`

	// get_status_vip
	//
	// If the server is in VIP mode
	Vip bool `json:"vip,omitempty"`
}

// Validate validates this get status o k body
func (m *GetStatusOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePlayers(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateServerVersion(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStartTime(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetStatusOKBody) validatePlayers(formats strfmt.Registry) error {

	if err := validate.Required("players", "body", m.Players); err != nil {
		return err
	}

	return nil
}

func (m *GetStatusOKBody) validateServerVersion(formats strfmt.Registry) error {

	if err := validate.Required("server_version", "body", m.ServerVersion); err != nil {
		return err
	}

	return nil
}

func (m *GetStatusOKBody) validateStartTime(formats strfmt.Registry) error {

	if err := validate.Required("start_time", "body", m.StartTime); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetStatusOKBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetStatusOKBody) UnmarshalBinary(b []byte) error {
	var res GetStatusOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
