package fleets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/pascaldmier/esi/models"
)

// NewPutFleetsFleetIDWingsWingIDParams creates a new PutFleetsFleetIDWingsWingIDParams object
// with the default values initialized.
func NewPutFleetsFleetIDWingsWingIDParams() PutFleetsFleetIDWingsWingIDParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return PutFleetsFleetIDWingsWingIDParams{
		Datasource: &datasourceDefault,
	}
}

// PutFleetsFleetIDWingsWingIDParams contains all the bound params for the put fleets fleet id wings wing id operation
// typically these are obtained from a http.Request
//
// swagger:parameters put_fleets_fleet_id_wings_wing_id
type PutFleetsFleetIDWingsWingIDParams struct {

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
	/*New name of the wing
	  Required: true
	  In: body
	*/
	Naming *models.PutFleetsFleetIDWingsWingIDParamsBody
	/*Access token to use, if preferred over a header
	  In: query
	*/
	Token *string
	/*Client identifier, takes precedence over headers
	  In: query
	*/
	UserAgent *string
	/*The wing to rename
	  Required: true
	  In: path
	*/
	WingID int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *PutFleetsFleetIDWingsWingIDParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
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

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.PutFleetsFleetIDWingsWingIDParamsBody
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("naming", "body"))
			} else {
				res = append(res, errors.NewParseError("naming", "body", "", err))
			}

		} else {
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Naming = &body
			}
		}

	} else {
		res = append(res, errors.Required("naming", "body"))
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

func (o *PutFleetsFleetIDWingsWingIDParams) bindXUserAgent(rawData []string, hasKey bool, formats strfmt.Registry) error {
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

func (o *PutFleetsFleetIDWingsWingIDParams) bindDatasource(rawData []string, hasKey bool, formats strfmt.Registry) error {
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

func (o *PutFleetsFleetIDWingsWingIDParams) validateDatasource(formats strfmt.Registry) error {

	if err := validate.Enum("datasource", "query", *o.Datasource, []interface{}{"tranquility", "singularity"}); err != nil {
		return err
	}

	return nil
}

func (o *PutFleetsFleetIDWingsWingIDParams) bindFleetID(rawData []string, hasKey bool, formats strfmt.Registry) error {
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

func (o *PutFleetsFleetIDWingsWingIDParams) bindToken(rawData []string, hasKey bool, formats strfmt.Registry) error {
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

func (o *PutFleetsFleetIDWingsWingIDParams) bindUserAgent(rawData []string, hasKey bool, formats strfmt.Registry) error {
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

func (o *PutFleetsFleetIDWingsWingIDParams) bindWingID(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
