package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GetIndustrySystemsOKBodyItems get_industry_systems_200_ok
//
// 200 ok object
// swagger:model getIndustrySystemsOKBodyItems
type GetIndustrySystemsOKBodyItems struct {

	// cost indices
	// Required: true
	CostIndices GetIndustrySystemsOKBodyItemsCostIndices `json:"cost_indices"`

	// get_industry_systems_solar_system_id
	//
	// solar_system_id integer
	// Required: true
	SolarSystemID *int32 `json:"solar_system_id"`
}

// Validate validates this get industry systems o k body items
func (m *GetIndustrySystemsOKBodyItems) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCostIndices(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSolarSystemID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetIndustrySystemsOKBodyItems) validateCostIndices(formats strfmt.Registry) error {

	if err := validate.Required("cost_indices", "body", m.CostIndices); err != nil {
		return err
	}

	return nil
}

func (m *GetIndustrySystemsOKBodyItems) validateSolarSystemID(formats strfmt.Registry) error {

	if err := validate.Required("solar_system_id", "body", m.SolarSystemID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetIndustrySystemsOKBodyItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetIndustrySystemsOKBodyItems) UnmarshalBinary(b []byte) error {
	var res GetIndustrySystemsOKBodyItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}