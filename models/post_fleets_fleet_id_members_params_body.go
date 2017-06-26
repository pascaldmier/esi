package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PostFleetsFleetIDMembersParamsBody post_fleets_fleet_id_members_invitation
//
// invitation object
// swagger:model postFleetsFleetIdMembersParamsBody
type PostFleetsFleetIDMembersParamsBody struct {

	// post_fleets_fleet_id_members_character_id
	//
	// The character you want to invite
	// Required: true
	CharacterID *int32 `json:"character_id"`

	// post_fleets_fleet_id_members_role
	//
	// - If a character is invited with the `fleet_commander` role, neither `wing_id` or `squad_id` should be specified - If a character is invited with the `wing_commander` role, only `wing_id` should be specified - If a character is invited with the `squad_commander` role, both `wing_id` and `squad_id` should be specified - If a character is invited with the `squad_member` role, `wing_id` and `squad_id` should either both be specified or not specified at all. If they aren’t specified, the invited character will join any squad with available positions
	//
	// Required: true
	Role *string `json:"role"`

	// post_fleets_fleet_id_members_squad_id
	//
	// squad_id integer
	// Minimum: 0
	SquadID *int64 `json:"squad_id,omitempty"`

	// post_fleets_fleet_id_members_wing_id
	//
	// wing_id integer
	// Minimum: 0
	WingID *int64 `json:"wing_id,omitempty"`
}

// Validate validates this post fleets fleet Id members params body
func (m *PostFleetsFleetIDMembersParamsBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCharacterID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRole(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSquadID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateWingID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostFleetsFleetIDMembersParamsBody) validateCharacterID(formats strfmt.Registry) error {

	if err := validate.Required("character_id", "body", m.CharacterID); err != nil {
		return err
	}

	return nil
}

var postFleetsFleetIdMembersParamsBodyTypeRolePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["fleet_commander","wing_commander","squad_commander","squad_member"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		postFleetsFleetIdMembersParamsBodyTypeRolePropEnum = append(postFleetsFleetIdMembersParamsBodyTypeRolePropEnum, v)
	}
}

const (
	// PostFleetsFleetIDMembersParamsBodyRoleFleetCommander captures enum value "fleet_commander"
	PostFleetsFleetIDMembersParamsBodyRoleFleetCommander string = "fleet_commander"
	// PostFleetsFleetIDMembersParamsBodyRoleWingCommander captures enum value "wing_commander"
	PostFleetsFleetIDMembersParamsBodyRoleWingCommander string = "wing_commander"
	// PostFleetsFleetIDMembersParamsBodyRoleSquadCommander captures enum value "squad_commander"
	PostFleetsFleetIDMembersParamsBodyRoleSquadCommander string = "squad_commander"
	// PostFleetsFleetIDMembersParamsBodyRoleSquadMember captures enum value "squad_member"
	PostFleetsFleetIDMembersParamsBodyRoleSquadMember string = "squad_member"
)

// prop value enum
func (m *PostFleetsFleetIDMembersParamsBody) validateRoleEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, postFleetsFleetIdMembersParamsBodyTypeRolePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *PostFleetsFleetIDMembersParamsBody) validateRole(formats strfmt.Registry) error {

	if err := validate.Required("role", "body", m.Role); err != nil {
		return err
	}

	// value enum
	if err := m.validateRoleEnum("role", "body", *m.Role); err != nil {
		return err
	}

	return nil
}

func (m *PostFleetsFleetIDMembersParamsBody) validateSquadID(formats strfmt.Registry) error {

	if swag.IsZero(m.SquadID) { // not required
		return nil
	}

	if err := validate.MinimumInt("squad_id", "body", int64(*m.SquadID), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *PostFleetsFleetIDMembersParamsBody) validateWingID(formats strfmt.Registry) error {

	if swag.IsZero(m.WingID) { // not required
		return nil
	}

	if err := validate.MinimumInt("wing_id", "body", int64(*m.WingID), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PostFleetsFleetIDMembersParamsBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostFleetsFleetIDMembersParamsBody) UnmarshalBinary(b []byte) error {
	var res PostFleetsFleetIDMembersParamsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
