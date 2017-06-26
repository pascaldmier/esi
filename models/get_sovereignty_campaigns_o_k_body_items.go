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

// GetSovereigntyCampaignsOKBodyItems get_sovereignty_campaigns_200_ok
//
// 200 ok object
// swagger:model getSovereigntyCampaignsOKBodyItems
type GetSovereigntyCampaignsOKBodyItems struct {

	// get_sovereignty_campaigns_attackers_score
	//
	// Score for all attacking parties, only present in Defense Events.
	//
	AttackersScore float32 `json:"attackers_score,omitempty"`

	// get_sovereignty_campaigns_campaign_id
	//
	// Unique ID for this campaign.
	// Required: true
	CampaignID *int32 `json:"campaign_id"`

	// get_sovereignty_campaigns_constellation_id
	//
	// The constellation in which the campaign will take place.
	//
	// Required: true
	ConstellationID *int32 `json:"constellation_id"`

	// get_sovereignty_campaigns_defender_id
	//
	// Defending alliance, only present in Defense Events
	//
	DefenderID int32 `json:"defender_id,omitempty"`

	// get_sovereignty_campaigns_defender_score
	//
	// Score for the defending alliance, only present in Defense Events.
	//
	DefenderScore float32 `json:"defender_score,omitempty"`

	// get_sovereignty_campaigns_event_type
	//
	// Type of event this campaign is for. tcu_defense, ihub_defense and station_defense are referred to as "Defense Events", station_freeport as "Freeport Events".
	//
	// Required: true
	EventType *string `json:"event_type"`

	// participants
	Participants GetSovereigntyCampaignsOKBodyItemsParticipants `json:"participants"`

	// get_sovereignty_campaigns_solar_system_id
	//
	// The solar system the structure is located in.
	//
	// Required: true
	SolarSystemID *int32 `json:"solar_system_id"`

	// get_sovereignty_campaigns_start_time
	//
	// Time the event is scheduled to start.
	//
	// Required: true
	StartTime *strfmt.DateTime `json:"start_time"`

	// get_sovereignty_campaigns_structure_id
	//
	// The structure item ID that is related to this campaign.
	//
	// Required: true
	StructureID *int64 `json:"structure_id"`
}

// Validate validates this get sovereignty campaigns o k body items
func (m *GetSovereigntyCampaignsOKBodyItems) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCampaignID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateConstellationID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateEventType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateParticipants(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSolarSystemID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStartTime(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStructureID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetSovereigntyCampaignsOKBodyItems) validateCampaignID(formats strfmt.Registry) error {

	if err := validate.Required("campaign_id", "body", m.CampaignID); err != nil {
		return err
	}

	return nil
}

func (m *GetSovereigntyCampaignsOKBodyItems) validateConstellationID(formats strfmt.Registry) error {

	if err := validate.Required("constellation_id", "body", m.ConstellationID); err != nil {
		return err
	}

	return nil
}

var getSovereigntyCampaignsOKBodyItemsTypeEventTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["tcu_defense","ihub_defense","station_defense","station_freeport"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getSovereigntyCampaignsOKBodyItemsTypeEventTypePropEnum = append(getSovereigntyCampaignsOKBodyItemsTypeEventTypePropEnum, v)
	}
}

const (
	// GetSovereigntyCampaignsOKBodyItemsEventTypeTcuDefense captures enum value "tcu_defense"
	GetSovereigntyCampaignsOKBodyItemsEventTypeTcuDefense string = "tcu_defense"
	// GetSovereigntyCampaignsOKBodyItemsEventTypeIhubDefense captures enum value "ihub_defense"
	GetSovereigntyCampaignsOKBodyItemsEventTypeIhubDefense string = "ihub_defense"
	// GetSovereigntyCampaignsOKBodyItemsEventTypeStationDefense captures enum value "station_defense"
	GetSovereigntyCampaignsOKBodyItemsEventTypeStationDefense string = "station_defense"
	// GetSovereigntyCampaignsOKBodyItemsEventTypeStationFreeport captures enum value "station_freeport"
	GetSovereigntyCampaignsOKBodyItemsEventTypeStationFreeport string = "station_freeport"
)

// prop value enum
func (m *GetSovereigntyCampaignsOKBodyItems) validateEventTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, getSovereigntyCampaignsOKBodyItemsTypeEventTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *GetSovereigntyCampaignsOKBodyItems) validateEventType(formats strfmt.Registry) error {

	if err := validate.Required("event_type", "body", m.EventType); err != nil {
		return err
	}

	// value enum
	if err := m.validateEventTypeEnum("event_type", "body", *m.EventType); err != nil {
		return err
	}

	return nil
}

func (m *GetSovereigntyCampaignsOKBodyItems) validateParticipants(formats strfmt.Registry) error {

	if swag.IsZero(m.Participants) { // not required
		return nil
	}

	return nil
}

func (m *GetSovereigntyCampaignsOKBodyItems) validateSolarSystemID(formats strfmt.Registry) error {

	if err := validate.Required("solar_system_id", "body", m.SolarSystemID); err != nil {
		return err
	}

	return nil
}

func (m *GetSovereigntyCampaignsOKBodyItems) validateStartTime(formats strfmt.Registry) error {

	if err := validate.Required("start_time", "body", m.StartTime); err != nil {
		return err
	}

	return nil
}

func (m *GetSovereigntyCampaignsOKBodyItems) validateStructureID(formats strfmt.Registry) error {

	if err := validate.Required("structure_id", "body", m.StructureID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetSovereigntyCampaignsOKBodyItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetSovereigntyCampaignsOKBodyItems) UnmarshalBinary(b []byte) error {
	var res GetSovereigntyCampaignsOKBodyItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
