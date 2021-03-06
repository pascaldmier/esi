package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GetInsurancePricesOKBodyItemsLevelsItems get_insurance_prices_level
//
// level object
// swagger:model getInsurancePricesOKBodyItemsLevelsItems
type GetInsurancePricesOKBodyItemsLevelsItems struct {

	// get_insurance_prices_cost
	//
	// cost number
	// Required: true
	Cost *float32 `json:"cost"`

	// get_insurance_prices_name
	//
	// Localized insurance level
	// Required: true
	Name *string `json:"name"`

	// get_insurance_prices_payout
	//
	// payout number
	// Required: true
	Payout *float32 `json:"payout"`
}

// Validate validates this get insurance prices o k body items levels items
func (m *GetInsurancePricesOKBodyItemsLevelsItems) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCost(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePayout(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetInsurancePricesOKBodyItemsLevelsItems) validateCost(formats strfmt.Registry) error {

	if err := validate.Required("cost", "body", m.Cost); err != nil {
		return err
	}

	return nil
}

func (m *GetInsurancePricesOKBodyItemsLevelsItems) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *GetInsurancePricesOKBodyItemsLevelsItems) validatePayout(formats strfmt.Registry) error {

	if err := validate.Required("payout", "body", m.Payout); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetInsurancePricesOKBodyItemsLevelsItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetInsurancePricesOKBodyItemsLevelsItems) UnmarshalBinary(b []byte) error {
	var res GetInsurancePricesOKBodyItemsLevelsItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
