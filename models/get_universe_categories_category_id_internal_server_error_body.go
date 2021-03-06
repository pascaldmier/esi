package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// GetUniverseCategoriesCategoryIDInternalServerErrorBody get_universe_categories_category_id_internal_server_error
//
// Internal server error
// swagger:model getUniverseCategoriesCategoryIdInternalServerErrorBody
type GetUniverseCategoriesCategoryIDInternalServerErrorBody struct {

	// get_universe_categories_category_id_500_internal_server_error
	//
	// Internal server error message
	Error string `json:"error,omitempty"`
}

// Validate validates this get universe categories category Id internal server error body
func (m *GetUniverseCategoriesCategoryIDInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *GetUniverseCategoriesCategoryIDInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetUniverseCategoriesCategoryIDInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res GetUniverseCategoriesCategoryIDInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
