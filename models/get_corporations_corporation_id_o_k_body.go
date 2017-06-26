package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GetCorporationsCorporationIDOKBody get_corporations_corporation_id_ok
//
// 200 ok object
// swagger:model getCorporationsCorporationIdOKBody
type GetCorporationsCorporationIDOKBody struct {

	// get_corporations_corporation_id_alliance_id
	//
	// id of alliance that corporation is a member of, if any
	AllianceID int32 `json:"alliance_id,omitempty"`

	// get_corporations_corporation_id_ceo_id
	//
	// ceo_id integer
	// Required: true
	CeoID *int32 `json:"ceo_id"`

	// get_corporations_corporation_id_corporation_description
	//
	// corporation_description string
	// Required: true
	CorporationDescription *string `json:"corporation_description"`

	// get_corporations_corporation_id_corporation_name
	//
	// the full name of the corporation
	// Required: true
	CorporationName *string `json:"corporation_name"`

	// get_corporations_corporation_id_creation_date
	//
	// creation_date string
	CreationDate strfmt.DateTime `json:"creation_date,omitempty"`

	// get_corporations_corporation_id_creator_id
	//
	// creator_id integer
	// Required: true
	CreatorID *int32 `json:"creator_id"`

	// get_corporations_corporation_id_faction
	//
	// faction string
	Faction string `json:"faction,omitempty"`

	// get_corporations_corporation_id_member_count
	//
	// member_count integer
	// Required: true
	MemberCount *int32 `json:"member_count"`

	// get_corporations_corporation_id_tax_rate
	//
	// tax_rate number
	// Required: true
	// Maximum: 1
	// Minimum: 0
	TaxRate *float32 `json:"tax_rate"`

	// get_corporations_corporation_id_ticker
	//
	// the short name of the corporation
	// Required: true
	Ticker *string `json:"ticker"`

	// get_corporations_corporation_id_url
	//
	// url string
	// Required: true
	URL *string `json:"url"`
}

// Validate validates this get corporations corporation Id o k body
func (m *GetCorporationsCorporationIDOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCeoID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCorporationDescription(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCorporationName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCreatorID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateFaction(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateMemberCount(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTaxRate(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTicker(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateURL(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetCorporationsCorporationIDOKBody) validateCeoID(formats strfmt.Registry) error {

	if err := validate.Required("ceo_id", "body", m.CeoID); err != nil {
		return err
	}

	return nil
}

func (m *GetCorporationsCorporationIDOKBody) validateCorporationDescription(formats strfmt.Registry) error {

	if err := validate.Required("corporation_description", "body", m.CorporationDescription); err != nil {
		return err
	}

	return nil
}

func (m *GetCorporationsCorporationIDOKBody) validateCorporationName(formats strfmt.Registry) error {

	if err := validate.Required("corporation_name", "body", m.CorporationName); err != nil {
		return err
	}

	return nil
}

func (m *GetCorporationsCorporationIDOKBody) validateCreatorID(formats strfmt.Registry) error {

	if err := validate.Required("creator_id", "body", m.CreatorID); err != nil {
		return err
	}

	return nil
}

var getCorporationsCorporationIdOKBodyTypeFactionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Minmatar","Gallente","Caldari","Amarr"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getCorporationsCorporationIdOKBodyTypeFactionPropEnum = append(getCorporationsCorporationIdOKBodyTypeFactionPropEnum, v)
	}
}

const (
	// GetCorporationsCorporationIDOKBodyFactionMinmatar captures enum value "Minmatar"
	GetCorporationsCorporationIDOKBodyFactionMinmatar string = "Minmatar"
	// GetCorporationsCorporationIDOKBodyFactionGallente captures enum value "Gallente"
	GetCorporationsCorporationIDOKBodyFactionGallente string = "Gallente"
	// GetCorporationsCorporationIDOKBodyFactionCaldari captures enum value "Caldari"
	GetCorporationsCorporationIDOKBodyFactionCaldari string = "Caldari"
	// GetCorporationsCorporationIDOKBodyFactionAmarr captures enum value "Amarr"
	GetCorporationsCorporationIDOKBodyFactionAmarr string = "Amarr"
)

// prop value enum
func (m *GetCorporationsCorporationIDOKBody) validateFactionEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, getCorporationsCorporationIdOKBodyTypeFactionPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *GetCorporationsCorporationIDOKBody) validateFaction(formats strfmt.Registry) error {

	if swag.IsZero(m.Faction) { // not required
		return nil
	}

	// value enum
	if err := m.validateFactionEnum("faction", "body", m.Faction); err != nil {
		return err
	}

	return nil
}

func (m *GetCorporationsCorporationIDOKBody) validateMemberCount(formats strfmt.Registry) error {

	if err := validate.Required("member_count", "body", m.MemberCount); err != nil {
		return err
	}

	return nil
}

func (m *GetCorporationsCorporationIDOKBody) validateTaxRate(formats strfmt.Registry) error {

	if err := validate.Required("tax_rate", "body", m.TaxRate); err != nil {
		return err
	}

	if err := validate.Minimum("tax_rate", "body", float64(*m.TaxRate), 0, false); err != nil {
		return err
	}

	if err := validate.Maximum("tax_rate", "body", float64(*m.TaxRate), 1, false); err != nil {
		return err
	}

	return nil
}

func (m *GetCorporationsCorporationIDOKBody) validateTicker(formats strfmt.Registry) error {

	if err := validate.Required("ticker", "body", m.Ticker); err != nil {
		return err
	}

	return nil
}

func (m *GetCorporationsCorporationIDOKBody) validateURL(formats strfmt.Registry) error {

	if err := validate.Required("url", "body", m.URL); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetCorporationsCorporationIDOKBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetCorporationsCorporationIDOKBody) UnmarshalBinary(b []byte) error {
	var res GetCorporationsCorporationIDOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}