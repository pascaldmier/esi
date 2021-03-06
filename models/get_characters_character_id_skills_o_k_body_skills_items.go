package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// GetCharactersCharacterIDSkillsOKBodySkillsItems get_characters_character_id_skills_skill
//
// skill object
// swagger:model getCharactersCharacterIdSkillsOKBodySkillsItems
type GetCharactersCharacterIDSkillsOKBodySkillsItems struct {

	// get_characters_character_id_skills_current_skill_level
	//
	// current_skill_level integer
	CurrentSkillLevel int32 `json:"current_skill_level,omitempty"`

	// get_characters_character_id_skills_skill_id
	//
	// skill_id integer
	SkillID int32 `json:"skill_id,omitempty"`

	// get_characters_character_id_skills_skillpoints_in_skill
	//
	// skillpoints_in_skill integer
	SkillpointsInSkill int64 `json:"skillpoints_in_skill,omitempty"`
}

// Validate validates this get characters character Id skills o k body skills items
func (m *GetCharactersCharacterIDSkillsOKBodySkillsItems) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *GetCharactersCharacterIDSkillsOKBodySkillsItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetCharactersCharacterIDSkillsOKBodySkillsItems) UnmarshalBinary(b []byte) error {
	var res GetCharactersCharacterIDSkillsOKBodySkillsItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
