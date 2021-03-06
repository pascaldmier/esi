package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GetCorporationsCorporationIDStructuresOKBodyItemsCurrentVulItems get_corporations_corporation_id_structures_current_vul
//
// current_vul object
// swagger:model getCorporationsCorporationIdStructuresOKBodyItemsCurrentVulItems
type GetCorporationsCorporationIDStructuresOKBodyItemsCurrentVulItems struct {

	// get_corporations_corporation_id_structures_day
	//
	// day integer
	// Required: true
	Day *int32 `json:"day"`

	// get_corporations_corporation_id_structures_hour
	//
	// hour integer
	// Required: true
	Hour *int32 `json:"hour"`
}

// Validate validates this get corporations corporation Id structures o k body items current vul items
func (m *GetCorporationsCorporationIDStructuresOKBodyItemsCurrentVulItems) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDay(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateHour(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetCorporationsCorporationIDStructuresOKBodyItemsCurrentVulItems) validateDay(formats strfmt.Registry) error {

	if err := validate.Required("day", "body", m.Day); err != nil {
		return err
	}

	return nil
}

func (m *GetCorporationsCorporationIDStructuresOKBodyItemsCurrentVulItems) validateHour(formats strfmt.Registry) error {

	if err := validate.Required("hour", "body", m.Hour); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetCorporationsCorporationIDStructuresOKBodyItemsCurrentVulItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetCorporationsCorporationIDStructuresOKBodyItemsCurrentVulItems) UnmarshalBinary(b []byte) error {
	var res GetCorporationsCorporationIDStructuresOKBodyItemsCurrentVulItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
