package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// PostUIOpenwindowContractForbiddenBody post_ui_openwindow_contract_forbidden
//
// Forbidden
// swagger:model postUiOpenwindowContractForbiddenBody
type PostUIOpenwindowContractForbiddenBody struct {

	// post_ui_openwindow_contract_403_forbidden
	//
	// Forbidden message
	Error string `json:"error,omitempty"`
}

// Validate validates this post Ui openwindow contract forbidden body
func (m *PostUIOpenwindowContractForbiddenBody) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *PostUIOpenwindowContractForbiddenBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostUIOpenwindowContractForbiddenBody) UnmarshalBinary(b []byte) error {
	var res PostUIOpenwindowContractForbiddenBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
