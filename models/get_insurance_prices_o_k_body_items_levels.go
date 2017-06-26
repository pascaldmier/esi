package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GetInsurancePricesOKBodyItemsLevels get_insurance_prices_levels
//
// A list of a available insurance levels for this ship type
// swagger:model getInsurancePricesOKBodyItemsLevels
type GetInsurancePricesOKBodyItemsLevels []*GetInsurancePricesOKBodyItemsLevelsItems

// Validate validates this get insurance prices o k body items levels
func (m GetInsurancePricesOKBodyItemsLevels) Validate(formats strfmt.Registry) error {
	var res []error

	iGetInsurancePricesOKBodyItemsLevelsSize := int64(len(m))

	if err := validate.MaxItems("", "body", iGetInsurancePricesOKBodyItemsLevelsSize, 6); err != nil {
		return err
	}

	for i := 0; i < len(m); i++ {

		if swag.IsZero(m[i]) { // not required
			continue
		}

		if m[i] != nil {

			if err := m[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName(strconv.Itoa(i))
				}
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}