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

// GetCharactersCharacterIDOrdersOKBodyItems get_characters_character_id_orders_200_ok
//
// 200 ok object
// swagger:model getCharactersCharacterIdOrdersOKBodyItems
type GetCharactersCharacterIDOrdersOKBodyItems struct {

	// get_characters_character_id_orders_account_id
	//
	// Wallet division for the buyer or seller of this order. Always 1000 for characters. Currently 1000 through 1006 for corporations
	// Required: true
	AccountID *int32 `json:"account_id"`

	// get_characters_character_id_orders_duration
	//
	// Numer of days for which order is valid (starting from the issued date). An order expires at time issued + duration
	// Required: true
	Duration *int32 `json:"duration"`

	// get_characters_character_id_orders_escrow
	//
	// For buy orders, the amount of ISK in escrow
	// Required: true
	Escrow *float32 `json:"escrow"`

	// get_characters_character_id_orders_is_buy_order
	//
	// True for a bid (buy) order. False for an offer (sell) order
	// Required: true
	IsBuyOrder *bool `json:"is_buy_order"`

	// get_characters_character_id_orders_is_corp
	//
	// is_corp boolean
	// Required: true
	IsCorp *bool `json:"is_corp"`

	// get_characters_character_id_orders_issued
	//
	// Date and time when this order was issued
	// Required: true
	Issued *strfmt.DateTime `json:"issued"`

	// get_characters_character_id_orders_location_id
	//
	// ID of the location where order was placed
	// Required: true
	LocationID *int64 `json:"location_id"`

	// get_characters_character_id_orders_min_volume
	//
	// For bids (buy orders), the minimum quantity that will be accepted in a matching offer (sell order)
	// Required: true
	MinVolume *int32 `json:"min_volume"`

	// get_characters_character_id_orders_order_id
	//
	// Unique order ID
	// Required: true
	OrderID *int64 `json:"order_id"`

	// get_characters_character_id_orders_price
	//
	// Cost per unit for this order
	// Required: true
	Price *float32 `json:"price"`

	// get_characters_character_id_orders_range
	//
	// Valid order range, numbers are ranges in jumps
	// Required: true
	Range *string `json:"range"`

	// get_characters_character_id_orders_region_id
	//
	// ID of the region where order was placed
	// Required: true
	RegionID *int32 `json:"region_id"`

	// get_characters_character_id_orders_state
	//
	// Current order state
	// Required: true
	State *string `json:"state"`

	// get_characters_character_id_orders_type_id
	//
	// The type ID of the item transacted in this order
	// Required: true
	TypeID *int32 `json:"type_id"`

	// get_characters_character_id_orders_volume_remain
	//
	// Quantity of items still required or offered
	// Required: true
	VolumeRemain *int32 `json:"volume_remain"`

	// get_characters_character_id_orders_volume_total
	//
	// Quantity of items required or offered at time order was placed
	// Required: true
	VolumeTotal *int32 `json:"volume_total"`
}

// Validate validates this get characters character Id orders o k body items
func (m *GetCharactersCharacterIDOrdersOKBodyItems) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccountID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDuration(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateEscrow(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateIsBuyOrder(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateIsCorp(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateIssued(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateLocationID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateMinVolume(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateOrderID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePrice(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRange(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRegionID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTypeID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateVolumeRemain(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateVolumeTotal(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetCharactersCharacterIDOrdersOKBodyItems) validateAccountID(formats strfmt.Registry) error {

	if err := validate.Required("account_id", "body", m.AccountID); err != nil {
		return err
	}

	return nil
}

func (m *GetCharactersCharacterIDOrdersOKBodyItems) validateDuration(formats strfmt.Registry) error {

	if err := validate.Required("duration", "body", m.Duration); err != nil {
		return err
	}

	return nil
}

func (m *GetCharactersCharacterIDOrdersOKBodyItems) validateEscrow(formats strfmt.Registry) error {

	if err := validate.Required("escrow", "body", m.Escrow); err != nil {
		return err
	}

	return nil
}

func (m *GetCharactersCharacterIDOrdersOKBodyItems) validateIsBuyOrder(formats strfmt.Registry) error {

	if err := validate.Required("is_buy_order", "body", m.IsBuyOrder); err != nil {
		return err
	}

	return nil
}

func (m *GetCharactersCharacterIDOrdersOKBodyItems) validateIsCorp(formats strfmt.Registry) error {

	if err := validate.Required("is_corp", "body", m.IsCorp); err != nil {
		return err
	}

	return nil
}

func (m *GetCharactersCharacterIDOrdersOKBodyItems) validateIssued(formats strfmt.Registry) error {

	if err := validate.Required("issued", "body", m.Issued); err != nil {
		return err
	}

	return nil
}

func (m *GetCharactersCharacterIDOrdersOKBodyItems) validateLocationID(formats strfmt.Registry) error {

	if err := validate.Required("location_id", "body", m.LocationID); err != nil {
		return err
	}

	return nil
}

func (m *GetCharactersCharacterIDOrdersOKBodyItems) validateMinVolume(formats strfmt.Registry) error {

	if err := validate.Required("min_volume", "body", m.MinVolume); err != nil {
		return err
	}

	return nil
}

func (m *GetCharactersCharacterIDOrdersOKBodyItems) validateOrderID(formats strfmt.Registry) error {

	if err := validate.Required("order_id", "body", m.OrderID); err != nil {
		return err
	}

	return nil
}

func (m *GetCharactersCharacterIDOrdersOKBodyItems) validatePrice(formats strfmt.Registry) error {

	if err := validate.Required("price", "body", m.Price); err != nil {
		return err
	}

	return nil
}

var getCharactersCharacterIdOrdersOKBodyItemsTypeRangePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["station","region","solarsystem","1","2","3","4","5","10","20","30","40"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getCharactersCharacterIdOrdersOKBodyItemsTypeRangePropEnum = append(getCharactersCharacterIdOrdersOKBodyItemsTypeRangePropEnum, v)
	}
}

const (
	// GetCharactersCharacterIDOrdersOKBodyItemsRangeStation captures enum value "station"
	GetCharactersCharacterIDOrdersOKBodyItemsRangeStation string = "station"
	// GetCharactersCharacterIDOrdersOKBodyItemsRangeRegion captures enum value "region"
	GetCharactersCharacterIDOrdersOKBodyItemsRangeRegion string = "region"
	// GetCharactersCharacterIDOrdersOKBodyItemsRangeSolarsystem captures enum value "solarsystem"
	GetCharactersCharacterIDOrdersOKBodyItemsRangeSolarsystem string = "solarsystem"
	// GetCharactersCharacterIDOrdersOKBodyItemsRangeNr1 captures enum value "1"
	GetCharactersCharacterIDOrdersOKBodyItemsRangeNr1 string = "1"
	// GetCharactersCharacterIDOrdersOKBodyItemsRangeNr2 captures enum value "2"
	GetCharactersCharacterIDOrdersOKBodyItemsRangeNr2 string = "2"
	// GetCharactersCharacterIDOrdersOKBodyItemsRangeNr3 captures enum value "3"
	GetCharactersCharacterIDOrdersOKBodyItemsRangeNr3 string = "3"
	// GetCharactersCharacterIDOrdersOKBodyItemsRangeNr4 captures enum value "4"
	GetCharactersCharacterIDOrdersOKBodyItemsRangeNr4 string = "4"
	// GetCharactersCharacterIDOrdersOKBodyItemsRangeNr5 captures enum value "5"
	GetCharactersCharacterIDOrdersOKBodyItemsRangeNr5 string = "5"
	// GetCharactersCharacterIDOrdersOKBodyItemsRangeNr10 captures enum value "10"
	GetCharactersCharacterIDOrdersOKBodyItemsRangeNr10 string = "10"
	// GetCharactersCharacterIDOrdersOKBodyItemsRangeNr20 captures enum value "20"
	GetCharactersCharacterIDOrdersOKBodyItemsRangeNr20 string = "20"
	// GetCharactersCharacterIDOrdersOKBodyItemsRangeNr30 captures enum value "30"
	GetCharactersCharacterIDOrdersOKBodyItemsRangeNr30 string = "30"
	// GetCharactersCharacterIDOrdersOKBodyItemsRangeNr40 captures enum value "40"
	GetCharactersCharacterIDOrdersOKBodyItemsRangeNr40 string = "40"
)

// prop value enum
func (m *GetCharactersCharacterIDOrdersOKBodyItems) validateRangeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, getCharactersCharacterIdOrdersOKBodyItemsTypeRangePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *GetCharactersCharacterIDOrdersOKBodyItems) validateRange(formats strfmt.Registry) error {

	if err := validate.Required("range", "body", m.Range); err != nil {
		return err
	}

	// value enum
	if err := m.validateRangeEnum("range", "body", *m.Range); err != nil {
		return err
	}

	return nil
}

func (m *GetCharactersCharacterIDOrdersOKBodyItems) validateRegionID(formats strfmt.Registry) error {

	if err := validate.Required("region_id", "body", m.RegionID); err != nil {
		return err
	}

	return nil
}

var getCharactersCharacterIdOrdersOKBodyItemsTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["open","closed","expired","cancelled","pending","character_deleted"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getCharactersCharacterIdOrdersOKBodyItemsTypeStatePropEnum = append(getCharactersCharacterIdOrdersOKBodyItemsTypeStatePropEnum, v)
	}
}

const (
	// GetCharactersCharacterIDOrdersOKBodyItemsStateOpen captures enum value "open"
	GetCharactersCharacterIDOrdersOKBodyItemsStateOpen string = "open"
	// GetCharactersCharacterIDOrdersOKBodyItemsStateClosed captures enum value "closed"
	GetCharactersCharacterIDOrdersOKBodyItemsStateClosed string = "closed"
	// GetCharactersCharacterIDOrdersOKBodyItemsStateExpired captures enum value "expired"
	GetCharactersCharacterIDOrdersOKBodyItemsStateExpired string = "expired"
	// GetCharactersCharacterIDOrdersOKBodyItemsStateCancelled captures enum value "cancelled"
	GetCharactersCharacterIDOrdersOKBodyItemsStateCancelled string = "cancelled"
	// GetCharactersCharacterIDOrdersOKBodyItemsStatePending captures enum value "pending"
	GetCharactersCharacterIDOrdersOKBodyItemsStatePending string = "pending"
	// GetCharactersCharacterIDOrdersOKBodyItemsStateCharacterDeleted captures enum value "character_deleted"
	GetCharactersCharacterIDOrdersOKBodyItemsStateCharacterDeleted string = "character_deleted"
)

// prop value enum
func (m *GetCharactersCharacterIDOrdersOKBodyItems) validateStateEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, getCharactersCharacterIdOrdersOKBodyItemsTypeStatePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *GetCharactersCharacterIDOrdersOKBodyItems) validateState(formats strfmt.Registry) error {

	if err := validate.Required("state", "body", m.State); err != nil {
		return err
	}

	// value enum
	if err := m.validateStateEnum("state", "body", *m.State); err != nil {
		return err
	}

	return nil
}

func (m *GetCharactersCharacterIDOrdersOKBodyItems) validateTypeID(formats strfmt.Registry) error {

	if err := validate.Required("type_id", "body", m.TypeID); err != nil {
		return err
	}

	return nil
}

func (m *GetCharactersCharacterIDOrdersOKBodyItems) validateVolumeRemain(formats strfmt.Registry) error {

	if err := validate.Required("volume_remain", "body", m.VolumeRemain); err != nil {
		return err
	}

	return nil
}

func (m *GetCharactersCharacterIDOrdersOKBodyItems) validateVolumeTotal(formats strfmt.Registry) error {

	if err := validate.Required("volume_total", "body", m.VolumeTotal); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetCharactersCharacterIDOrdersOKBodyItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetCharactersCharacterIDOrdersOKBodyItems) UnmarshalBinary(b []byte) error {
	var res GetCharactersCharacterIDOrdersOKBodyItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}