package universe

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

// NewGetUniverseConstellationsConstellationIDParams creates a new GetUniverseConstellationsConstellationIDParams object
// with the default values initialized.
func NewGetUniverseConstellationsConstellationIDParams() GetUniverseConstellationsConstellationIDParams {
	var (
		datasourceDefault = string("tranquility")
		languageDefault   = string("en-us")
	)
	return GetUniverseConstellationsConstellationIDParams{
		Datasource: &datasourceDefault,

		Language: &languageDefault,
	}
}

// GetUniverseConstellationsConstellationIDParams contains all the bound params for the get universe constellations constellation id operation
// typically these are obtained from a http.Request
//
// swagger:parameters get_universe_constellations_constellation_id
type GetUniverseConstellationsConstellationIDParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request

	/*Client identifier, takes precedence over User-Agent
	  In: header
	*/
	XUserAgent *string
	/*constellation_id integer
	  Required: true
	  In: path
	*/
	ConstellationID int32
	/*The server name you would like data from
	  In: query
	  Default: "tranquility"
	*/
	Datasource *string
	/*Language to use in the response
	  In: query
	  Default: "en-us"
	*/
	Language *string
	/*Client identifier, takes precedence over headers
	  In: query
	*/
	UserAgent *string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *GetUniverseConstellationsConstellationIDParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	if err := o.bindXUserAgent(r.Header[http.CanonicalHeaderKey("X-User-Agent")], true, route.Formats); err != nil {
		res = append(res, err)
	}

	rConstellationID, rhkConstellationID, _ := route.Params.GetOK("constellation_id")
	if err := o.bindConstellationID(rConstellationID, rhkConstellationID, route.Formats); err != nil {
		res = append(res, err)
	}

	qDatasource, qhkDatasource, _ := qs.GetOK("datasource")
	if err := o.bindDatasource(qDatasource, qhkDatasource, route.Formats); err != nil {
		res = append(res, err)
	}

	qLanguage, qhkLanguage, _ := qs.GetOK("language")
	if err := o.bindLanguage(qLanguage, qhkLanguage, route.Formats); err != nil {
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

func (o *GetUniverseConstellationsConstellationIDParams) bindXUserAgent(rawData []string, hasKey bool, formats strfmt.Registry) error {
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

func (o *GetUniverseConstellationsConstellationIDParams) bindConstellationID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	value, err := swag.ConvertInt32(raw)
	if err != nil {
		return errors.InvalidType("constellation_id", "path", "int32", raw)
	}
	o.ConstellationID = value

	return nil
}

func (o *GetUniverseConstellationsConstellationIDParams) bindDatasource(rawData []string, hasKey bool, formats strfmt.Registry) error {
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

func (o *GetUniverseConstellationsConstellationIDParams) validateDatasource(formats strfmt.Registry) error {

	if err := validate.Enum("datasource", "query", *o.Datasource, []interface{}{"tranquility", "singularity"}); err != nil {
		return err
	}

	return nil
}

func (o *GetUniverseConstellationsConstellationIDParams) bindLanguage(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		var languageDefault string = string("en-us")
		o.Language = &languageDefault
		return nil
	}

	o.Language = &raw

	if err := o.validateLanguage(formats); err != nil {
		return err
	}

	return nil
}

func (o *GetUniverseConstellationsConstellationIDParams) validateLanguage(formats strfmt.Registry) error {

	if err := validate.Enum("language", "query", *o.Language, []interface{}{"de", "en-us", "fr", "ja", "ru", "zh"}); err != nil {
		return err
	}

	return nil
}

func (o *GetUniverseConstellationsConstellationIDParams) bindUserAgent(rawData []string, hasKey bool, formats strfmt.Registry) error {
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