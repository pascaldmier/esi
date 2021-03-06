package killmails

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetCharactersCharacterIDKillmailsRecentParams creates a new GetCharactersCharacterIDKillmailsRecentParams object
// with the default values initialized.
func NewGetCharactersCharacterIDKillmailsRecentParams() GetCharactersCharacterIDKillmailsRecentParams {
	var (
		datasourceDefault = string("tranquility")
		maxCountDefault   = int32(50)
	)
	return GetCharactersCharacterIDKillmailsRecentParams{
		Datasource: &datasourceDefault,

		MaxCount: &maxCountDefault,
	}
}

// GetCharactersCharacterIDKillmailsRecentParams contains all the bound params for the get characters character id killmails recent operation
// typically these are obtained from a http.Request
//
// swagger:parameters get_characters_character_id_killmails_recent
type GetCharactersCharacterIDKillmailsRecentParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request

	/*Client identifier, takes precedence over User-Agent
	  In: header
	*/
	XUserAgent *string
	/*An EVE character ID
	  Required: true
	  In: path
	*/
	CharacterID int32
	/*The server name you would like data from
	  In: query
	  Default: "tranquility"
	*/
	Datasource *string
	/*How many killmails to return at maximum
	  Maximum: 5000
	  In: query
	  Default: 50
	*/
	MaxCount *int32
	/*Only return killmails with ID smaller than this.

	  In: query
	*/
	MaxKillID *int32
	/*Access token to use, if preferred over a header
	  In: query
	*/
	Token *string
	/*Client identifier, takes precedence over headers
	  In: query
	*/
	UserAgent *string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *GetCharactersCharacterIDKillmailsRecentParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	if err := o.bindXUserAgent(r.Header[http.CanonicalHeaderKey("X-User-Agent")], true, route.Formats); err != nil {
		res = append(res, err)
	}

	rCharacterID, rhkCharacterID, _ := route.Params.GetOK("character_id")
	if err := o.bindCharacterID(rCharacterID, rhkCharacterID, route.Formats); err != nil {
		res = append(res, err)
	}

	qDatasource, qhkDatasource, _ := qs.GetOK("datasource")
	if err := o.bindDatasource(qDatasource, qhkDatasource, route.Formats); err != nil {
		res = append(res, err)
	}

	qMaxCount, qhkMaxCount, _ := qs.GetOK("max_count")
	if err := o.bindMaxCount(qMaxCount, qhkMaxCount, route.Formats); err != nil {
		res = append(res, err)
	}

	qMaxKillID, qhkMaxKillID, _ := qs.GetOK("max_kill_id")
	if err := o.bindMaxKillID(qMaxKillID, qhkMaxKillID, route.Formats); err != nil {
		res = append(res, err)
	}

	qToken, qhkToken, _ := qs.GetOK("token")
	if err := o.bindToken(qToken, qhkToken, route.Formats); err != nil {
		res = append(res, err)
	}

	qUserAgent, qhkUserAgent, _ := qs.GetOK("user_agent")
	if err := o.bindUserAgent(qUserAgent, qhkUserAgent, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetCharactersCharacterIDKillmailsRecentParams) bindXUserAgent(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.XUserAgent = &raw

	return nil
}

func (o *GetCharactersCharacterIDKillmailsRecentParams) bindCharacterID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	value, err := swag.ConvertInt32(raw)
	if err != nil {
		return errors.InvalidType("character_id", "path", "int32", raw)
	}
	o.CharacterID = value

	return nil
}

func (o *GetCharactersCharacterIDKillmailsRecentParams) bindDatasource(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		var datasourceDefault string = string("tranquility")
		o.Datasource = &datasourceDefault
		return nil
	}

	o.Datasource = &raw

	if err := o.validateDatasource(formats); err != nil {
		return err
	}

	return nil
}

func (o *GetCharactersCharacterIDKillmailsRecentParams) validateDatasource(formats strfmt.Registry) error {

	if err := validate.Enum("datasource", "query", *o.Datasource, []interface{}{"tranquility", "singularity"}); err != nil {
		return err
	}

	return nil
}

func (o *GetCharactersCharacterIDKillmailsRecentParams) bindMaxCount(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		var maxCountDefault int32 = int32(50)
		o.MaxCount = &maxCountDefault
		return nil
	}

	value, err := swag.ConvertInt32(raw)
	if err != nil {
		return errors.InvalidType("max_count", "query", "int32", raw)
	}
	o.MaxCount = &value

	if err := o.validateMaxCount(formats); err != nil {
		return err
	}

	return nil
}

func (o *GetCharactersCharacterIDKillmailsRecentParams) validateMaxCount(formats strfmt.Registry) error {

	if err := validate.MaximumInt("max_count", "query", int64(*o.MaxCount), 5000, false); err != nil {
		return err
	}

	return nil
}

func (o *GetCharactersCharacterIDKillmailsRecentParams) bindMaxKillID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt32(raw)
	if err != nil {
		return errors.InvalidType("max_kill_id", "query", "int32", raw)
	}
	o.MaxKillID = &value

	return nil
}

func (o *GetCharactersCharacterIDKillmailsRecentParams) bindToken(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Token = &raw

	return nil
}

func (o *GetCharactersCharacterIDKillmailsRecentParams) bindUserAgent(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.UserAgent = &raw

	return nil
}
