package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GetCorporationsCorporationIDStructuresOKBodyItems get_corporations_corporation_id_structures_200_ok
//
// 200 ok object
// swagger:model getCorporationsCorporationIdStructuresOKBodyItems
type GetCorporationsCorporationIDStructuresOKBodyItems struct {

	// get_corporations_corporation_id_structures_corporation_id
	//
	// ID of the corporation that owns the structure
	// Required: true
	CorporationID *int32 `json:"corporation_id"`

	// current vul
	// Required: true
	CurrentVul GetCorporationsCorporationIDStructuresOKBodyItemsCurrentVul `json:"current_vul"`

	// get_corporations_corporation_id_structures_fuel_expires
	//
	// Date on which the structure will run out of fuel
	FuelExpires strfmt.Date `json:"fuel_expires,omitempty"`

	// next vul
	// Required: true
	NextVul GetCorporationsCorporationIDStructuresOKBodyItemsNextVul `json:"next_vul"`

	// get_corporations_corporation_id_structures_profile_id
	//
	// The id of the ACL profile for this citadel
	// Required: true
	ProfileID *int32 `json:"profile_id"`

	// services
	Services GetCorporationsCorporationIDStructuresOKBodyItemsServices `json:"services"`

	// get_corporations_corporation_id_structures_state_timer_end
	//
	// Date at which the structure will move to it's next state
	StateTimerEnd strfmt.Date `json:"state_timer_end,omitempty"`

	// get_corporations_corporation_id_structures_state_timer_start
	//
	// Date at which the structure entered it's current state
	StateTimerStart strfmt.Date `json:"state_timer_start,omitempty"`

	// get_corporations_corporation_id_structures_structure_id
	//
	// The Item ID of the structure
	// Required: true
	StructureID *int64 `json:"structure_id"`

	// get_corporations_corporation_id_structures_system_id
	//
	// The solar system the structure is in
	// Required: true
	SystemID *int32 `json:"system_id"`

	// get_corporations_corporation_id_structures_type_id
	//
	// The type id of the structure
	// Required: true
	TypeID *int32 `json:"type_id"`

	// get_corporations_corporation_id_structures_unanchors_at
	//
	// Date at which the structure will unanchor
	UnanchorsAt strfmt.Date `json:"unanchors_at,omitempty"`
}

// Validate validates this get corporations corporation Id structures o k body items
func (m *GetCorporationsCorporationIDStructuresOKBodyItems) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCorporationID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCurrentVul(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateNextVul(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateProfileID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateServices(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStructureID(formats); err != nil {
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

func (m *GetCorporationsCorporationIDStructuresOKBodyItems) validateCorporationID(formats strfmt.Registry) error {

	if err := validate.Required("corporation_id", "body", m.CorporationID); err != nil {
		return err
	}

	return nil
}

func (m *GetCorporationsCorporationIDStructuresOKBodyItems) validateCurrentVul(formats strfmt.Registry) error {

	if err := validate.Required("current_vul", "body", m.CurrentVul); err != nil {
		return err
	}

	return nil
}

func (m *GetCorporationsCorporationIDStructuresOKBodyItems) validateNextVul(formats strfmt.Registry) error {

	if err := validate.Required("next_vul", "body", m.NextVul); err != nil {
		return err
	}

	return nil
}

func (m *GetCorporationsCorporationIDStructuresOKBodyItems) validateProfileID(formats strfmt.Registry) error {

	if err := validate.Required("profile_id", "body", m.ProfileID); err != nil {
		return err
	}

	return nil
}

func (m *GetCorporationsCorporationIDStructuresOKBodyItems) validateServices(formats strfmt.Registry) error {

	if swag.IsZero(m.Services) { // not required
		return nil
	}

	return nil
}

func (m *GetCorporationsCorporationIDStructuresOKBodyItems) validateStructureID(formats strfmt.Registry) error {

	if err := validate.Required("structure_id", "body", m.StructureID); err != nil {
		return err
	}

	return nil
}

func (m *GetCorporationsCorporationIDStructuresOKBodyItems) validateSystemID(formats strfmt.Registry) error {

	if err := validate.Required("system_id", "body", m.SystemID); err != nil {
		return err
	}

	return nil
}

func (m *GetCorporationsCorporationIDStructuresOKBodyItems) validateTypeID(formats strfmt.Registry) error {

	if err := validate.Required("type_id", "body", m.TypeID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetCorporationsCorporationIDStructuresOKBodyItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetCorporationsCorporationIDStructuresOKBodyItems) UnmarshalBinary(b []byte) error {
	var res GetCorporationsCorporationIDStructuresOKBodyItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
