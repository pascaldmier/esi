package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GetUniverseStargatesStargateIDOKBody get_universe_stargates_stargate_id_ok
//
// 200 ok object
// swagger:model getUniverseStargatesStargateIdOKBody
type GetUniverseStargatesStargateIDOKBody struct {

	// destination
	// Required: true
	Destination *GetUniverseStargatesStargateIDOKBodyDestination `json:"destination"`

	// get_universe_stargates_stargate_id_name
	//
	// name string
	// Required: true
	Name *string `json:"name"`

	// position
	// Required: true
	Position *GetUniverseStargatesStargateIDOKBodyPosition `json:"position"`

	// get_universe_stargates_stargate_id_stargate_id
	//
	// stargate_id integer
	// Required: true
	StargateID *int32 `json:"stargate_id"`

	// get_universe_stargates_stargate_id_system_id
	//
	// The solar system this stargate is in
	// Required: true
	SystemID *int32 `json:"system_id"`

	// get_universe_stargates_stargate_id_type_id
	//
	// type_id integer
	// Required: true
	TypeID *int32 `json:"type_id"`
}

// Validate validates this get universe stargates stargate Id o k body
func (m *GetUniverseStargatesStargateIDOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDestination(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePosition(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStargateID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSystemID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTypeID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetUniverseStargatesStargateIDOKBody) validateDestination(formats strfmt.Registry) error {

	if err := validate.Required("destination", "body", m.Destination); err != nil {
		return err
	}

	if m.Destination != nil {

		if err := m.Destination.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("destination")
			}
			return err
		}
	}

	return nil
}

func (m *GetUniverseStargatesStargateIDOKBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *GetUniverseStargatesStargateIDOKBody) validatePosition(formats strfmt.Registry) error {

	if err := validate.Required("position", "body", m.Position); err != nil {
		return err
	}

	if m.Position != nil {

		if err := m.Position.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("position")
			}
			return err
		}
	}

	return nil
}

func (m *GetUniverseStargatesStargateIDOKBody) validateStargateID(formats strfmt.Registry) error {

	if err := validate.Required("stargate_id", "body", m.StargateID); err != nil {
		return err
	}

	return nil
}

func (m *GetUniverseStargatesStargateIDOKBody) validateSystemID(formats strfmt.Registry) error {

	if err := validate.Required("system_id", "body", m.SystemID); err != nil {
		return err
	}

	return nil
}

func (m *GetUniverseStargatesStargateIDOKBody) validateTypeID(formats strfmt.Registry) error {

	if err := validate.Required("type_id", "body", m.TypeID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetUniverseStargatesStargateIDOKBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetUniverseStargatesStargateIDOKBody) UnmarshalBinary(b []byte) error {
	var res GetUniverseStargatesStargateIDOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
