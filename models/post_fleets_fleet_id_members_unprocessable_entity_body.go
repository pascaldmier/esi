package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// PostFleetsFleetIDMembersUnprocessableEntityBody post_fleets_fleet_id_members_unprocessable_entity
//
// 422 unprocessable entity object
// swagger:model postFleetsFleetIdMembersUnprocessableEntityBody
type PostFleetsFleetIDMembersUnprocessableEntityBody struct {

	// post_fleets_fleet_id_members_error
	//
	// error message
	Error string `json:"error,omitempty"`
}

// Validate validates this post fleets fleet Id members unprocessable entity body
func (m *PostFleetsFleetIDMembersUnprocessableEntityBody) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *PostFleetsFleetIDMembersUnprocessableEntityBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostFleetsFleetIDMembersUnprocessableEntityBody) UnmarshalBinary(b []byte) error {
	var res PostFleetsFleetIDMembersUnprocessableEntityBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
