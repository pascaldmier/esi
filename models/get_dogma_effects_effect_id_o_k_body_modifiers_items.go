package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GetDogmaEffectsEffectIDOKBodyModifiersItems get_dogma_effects_effect_id_modifier
//
// modifier object
// swagger:model getDogmaEffectsEffectIdOKBodyModifiersItems
type GetDogmaEffectsEffectIDOKBodyModifiersItems struct {

	// get_dogma_effects_effect_id_domain
	//
	// domain string
	// Required: true
	Domain *string `json:"domain"`

	// get_dogma_effects_effect_id_func
	//
	// func string
	// Required: true
	Func *string `json:"func"`

	// get_dogma_effects_effect_id_modified_attribute_id
	//
	// modified_attribute_id integer
	// Required: true
	ModifiedAttributeID *int32 `json:"modified_attribute_id"`

	// get_dogma_effects_effect_id_modifying_attribute_id
	//
	// modifying_attribute_id integer
	// Required: true
	ModifyingAttributeID *int32 `json:"modifying_attribute_id"`

	// get_dogma_effects_effect_id_operator
	//
	// operator integer
	// Required: true
	Operator *int32 `json:"operator"`
}

// Validate validates this get dogma effects effect Id o k body modifiers items
func (m *GetDogmaEffectsEffectIDOKBodyModifiersItems) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDomain(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateFunc(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateModifiedAttributeID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateModifyingAttributeID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateOperator(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetDogmaEffectsEffectIDOKBodyModifiersItems) validateDomain(formats strfmt.Registry) error {

	if err := validate.Required("domain", "body", m.Domain); err != nil {
		return err
	}

	return nil
}

func (m *GetDogmaEffectsEffectIDOKBodyModifiersItems) validateFunc(formats strfmt.Registry) error {

	if err := validate.Required("func", "body", m.Func); err != nil {
		return err
	}

	return nil
}

func (m *GetDogmaEffectsEffectIDOKBodyModifiersItems) validateModifiedAttributeID(formats strfmt.Registry) error {

	if err := validate.Required("modified_attribute_id", "body", m.ModifiedAttributeID); err != nil {
		return err
	}

	return nil
}

func (m *GetDogmaEffectsEffectIDOKBodyModifiersItems) validateModifyingAttributeID(formats strfmt.Registry) error {

	if err := validate.Required("modifying_attribute_id", "body", m.ModifyingAttributeID); err != nil {
		return err
	}

	return nil
}

func (m *GetDogmaEffectsEffectIDOKBodyModifiersItems) validateOperator(formats strfmt.Registry) error {

	if err := validate.Required("operator", "body", m.Operator); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetDogmaEffectsEffectIDOKBodyModifiersItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetDogmaEffectsEffectIDOKBodyModifiersItems) UnmarshalBinary(b []byte) error {
	var res GetDogmaEffectsEffectIDOKBodyModifiersItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}