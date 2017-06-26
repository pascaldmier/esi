package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GetUniverseTypesTypeIDOKBody get_universe_types_type_id_ok
//
// 200 ok object
// swagger:model getUniverseTypesTypeIdOKBody
type GetUniverseTypesTypeIDOKBody struct {

	// get_universe_types_type_id_capacity
	//
	// capacity number
	Capacity float32 `json:"capacity,omitempty"`

	// get_universe_types_type_id_description
	//
	// description string
	// Required: true
	Description *string `json:"description"`

	// dogma attributes
	DogmaAttributes GetUniverseTypesTypeIDOKBodyDogmaAttributes `json:"dogma_attributes"`

	// dogma effects
	DogmaEffects GetUniverseTypesTypeIDOKBodyDogmaEffects `json:"dogma_effects"`

	// get_universe_types_type_id_graphic_id
	//
	// graphic_id integer
	GraphicID int32 `json:"graphic_id,omitempty"`

	// get_universe_types_type_id_group_id
	//
	// group_id integer
	// Required: true
	GroupID *int32 `json:"group_id"`

	// get_universe_types_type_id_icon_id
	//
	// icon_id integer
	IconID int32 `json:"icon_id,omitempty"`

	// get_universe_types_type_id_mass
	//
	// mass number
	Mass float32 `json:"mass,omitempty"`

	// get_universe_types_type_id_name
	//
	// name string
	// Required: true
	Name *string `json:"name"`

	// get_universe_types_type_id_portion_size
	//
	// portion_size integer
	PortionSize int32 `json:"portion_size,omitempty"`

	// get_universe_types_type_id_published
	//
	// published boolean
	// Required: true
	Published *bool `json:"published"`

	// get_universe_types_type_id_radius
	//
	// radius number
	Radius float32 `json:"radius,omitempty"`

	// get_universe_types_type_id_type_id
	//
	// type_id integer
	// Required: true
	TypeID *int32 `json:"type_id"`

	// get_universe_types_type_id_volume
	//
	// volume number
	Volume float32 `json:"volume,omitempty"`
}

// Validate validates this get universe types type Id o k body
func (m *GetUniverseTypesTypeIDOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDogmaAttributes(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDogmaEffects(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateGroupID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePublished(formats); err != nil {
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

func (m *GetUniverseTypesTypeIDOKBody) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *GetUniverseTypesTypeIDOKBody) validateDogmaAttributes(formats strfmt.Registry) error {

	if swag.IsZero(m.DogmaAttributes) { // not required
		return nil
	}

	return nil
}

func (m *GetUniverseTypesTypeIDOKBody) validateDogmaEffects(formats strfmt.Registry) error {

	if swag.IsZero(m.DogmaEffects) { // not required
		return nil
	}

	return nil
}

func (m *GetUniverseTypesTypeIDOKBody) validateGroupID(formats strfmt.Registry) error {

	if err := validate.Required("group_id", "body", m.GroupID); err != nil {
		return err
	}

	return nil
}

func (m *GetUniverseTypesTypeIDOKBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *GetUniverseTypesTypeIDOKBody) validatePublished(formats strfmt.Registry) error {

	if err := validate.Required("published", "body", m.Published); err != nil {
		return err
	}

	return nil
}

func (m *GetUniverseTypesTypeIDOKBody) validateTypeID(formats strfmt.Registry) error {

	if err := validate.Required("type_id", "body", m.TypeID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetUniverseTypesTypeIDOKBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetUniverseTypesTypeIDOKBody) UnmarshalBinary(b []byte) error {
	var res GetUniverseTypesTypeIDOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
