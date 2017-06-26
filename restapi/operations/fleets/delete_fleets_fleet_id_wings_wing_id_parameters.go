package fleets

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

// NewDeleteFleetsFleetIDWingsWingIDParams creates a new DeleteFleetsFleetIDWingsWingIDParams object
// with the default values initialized.
func NewDeleteFleetsFleetIDWingsWingIDParams() DeleteFleetsFleetIDWingsWingIDParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return DeleteFleetsFleetIDWingsWingIDParams{
		Datasource: &datasourceDefault,
	}
}

// DeleteFleetsFleetIDWingsWingIDParams contains all the bound params for the delete fleets fleet id wings wing id operation
// typically these are obtained from a http.Request
//
// swagger:parameters delete_fleets_fleet_id_wings_wing_id
type DeleteFleetsFleetIDWingsWingIDParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request

	/*Client identifier, takes precedence over User-Agent
	  In: header
	*/
	XUserAgent *string
	/*The server name you would like data from
	  In: query
	  Default: "tranquility"
	*/
	Datasource *string
	/*ID for a fleet
	  Required: true
	  In: path
	*/
	FleetID int64
	/*Access token to use, if preferred over a header
	  In: query
	*/
	Token *string
	/*Client identifier, takes precedence over headers
	  In: query
	*/
	UserAgent *string
	/*The wing to delete
	  Required: true
	  In: path
	*/
	WingID int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *DeleteFleetsFleetIDWingsWingIDParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	if err := o.bindXUserAgent(r.Header[http.CanonicalHeaderKey("X-User-Agent")], true, route.Formats); err != nil {
		res = append(res, err)
	}

	qDatasource, qhkDatasource, _ := qs.GetOK("datasource")
	if err := o.bindDatasource(qDatasource, qhkDatasource, route.Formats); err != nil {
		res = append(res, err)
	}

	rFleetID, rhkFleetID, _ := route.Params.GetOK("fleet_id")
	if err := o.bindFleetID(rFleetID, rhkFleetID, route.Formats); err != nil {
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

	rWingID, rhkWingID, _ := route.Params.GetOK("wing_id")
	if err := o.bindWingID(rWingID, rhkWingID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DeleteFleetsFleetIDWingsWingIDParams) bindXUserAgent(rawData []string, hasKey bool, formats strfmt.Registry) error {
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

func (o *DeleteFleetsFleetIDWingsWingIDParams) bindDatasource(rawData []string, hasKey bool, formats strfmt.Registry) error {
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

func (o *DeleteFleetsFleetIDWingsWingIDParams) validateDatasource(formats strfmt.Registry) error {

	if err := validate.Enum("datasource", "query", *o.Datasource, []interface{}{"tranquility", "singularity"}); err != nil {
		return err
	}

	return nil
}

func (o *DeleteFleetsFleetIDWingsWingIDParams) bindFleetID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("fleet_id", "path", "int64", raw)
	}
	o.FleetID = value

	return nil
}

func (o *DeleteFleetsFleetIDWingsWingIDParams) bindToken(rawData []string, hasKey bool, formats strfmt.Registry) error {
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

func (o *DeleteFleetsFleetIDWingsWingIDParams) bindUserAgent(rawData []string, hasKey bool, formats strfmt.Registry) error {
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

func (o *DeleteFleetsFleetIDWingsWingIDParams) bindWingID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("wing_id", "path", "int64", raw)
	}
	o.WingID = value

	return nil
}
