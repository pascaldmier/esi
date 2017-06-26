package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GetUniverseSystemsSystemIDOKBody get_universe_systems_system_id_ok
//
// 200 ok object
// swagger:model getUniverseSystemsSystemIdOKBody
type GetUniverseSystemsSystemIDOKBody struct {

	// get_universe_systems_system_id_constellation_id
	//
	// The constellation this solar system is in
	// Required: true
	ConstellationID *int32 `json:"constellation_id"`

	// get_universe_systems_system_id_name
	//
	// name string
	// Required: true
	Name *string `json:"name"`

	// planets
	// Required: true
	Planets GetUniverseSystemsSystemIDOKBodyPlanets `json:"planets"`

	// position
	// Required: true
	Position *GetUniverseSystemsSystemIDOKBodyPosition `json:"position"`

	// get_universe_systems_system_id_security_class
	//
	// security_class string
	SecurityClass string `json:"security_class,omitempty"`

	// get_universe_systems_system_id_security_status
	//
	// security_status number
	// Required: true
	SecurityStatus *float32 `json:"security_status"`

	// get_universe_systems_system_id_stargates
	//
	// stargates array
	// Required: true
	// Max Items: 25
	Stargates []int32 `json:"stargates"`

	// get_universe_systems_system_id_system_id
	//
	// system_id integer
	// Required: true
	SystemID *int32 `json:"system_id"`
}

// Validate validates this get universe systems system Id o k body
func (m *GetUniverseSystemsSystemIDOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConstellationID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePlanets(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePosition(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSecurityStatus(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStargates(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSystemID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetUniverseSystemsSystemIDOKBody) validateConstellationID(formats strfmt.Registry) error {

	if err := validate.Required("constellation_id", "body", m.ConstellationID); err != nil {
		return err
	}

	return nil
}

func (m *GetUniverseSystemsSystemIDOKBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *GetUniverseSystemsSystemIDOKBody) validatePlanets(formats strfmt.Registry) error {

	if err := validate.Required("planets", "body", m.Planets); err != nil {
		return err
	}

	return nil
}

func (m *GetUniverseSystemsSystemIDOKBody) validatePosition(formats strfmt.Registry) error {

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

func (m *GetUniverseSystemsSystemIDOKBody) validateSecurityStatus(formats strfmt.Registry) error {

	if err := validate.Required("security_status", "body", m.SecurityStatus); err != nil {
		return err
	}

	return nil
}

func (m *GetUniverseSystemsSystemIDOKBody) validateStargates(formats strfmt.Registry) error {

	if err := validate.Required("stargates", "body", m.Stargates); err != nil {
		return err
	}

	iStargatesSize := int64(len(m.Stargates))

	if err := validate.MaxItems("stargates", "body", iStargatesSize, 25); err != nil {
		return err
	}

	return nil
}

func (m *GetUniverseSystemsSystemIDOKBody) validateSystemID(formats strfmt.Registry) error {

	if err := validate.Required("system_id", "body", m.SystemID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetUniverseSystemsSystemIDOKBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetUniverseSystemsSystemIDOKBody) UnmarshalBinary(b []byte) error {
	var res GetUniverseSystemsSystemIDOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}