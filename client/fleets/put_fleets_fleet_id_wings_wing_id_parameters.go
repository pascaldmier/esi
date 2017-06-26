package fleets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/pascaldmier/esi/models"
)

// NewPutFleetsFleetIDWingsWingIDParams creates a new PutFleetsFleetIDWingsWingIDParams object
// with the default values initialized.
func NewPutFleetsFleetIDWingsWingIDParams() *PutFleetsFleetIDWingsWingIDParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &PutFleetsFleetIDWingsWingIDParams{
		Datasource: &datasourceDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewPutFleetsFleetIDWingsWingIDParamsWithTimeout creates a new PutFleetsFleetIDWingsWingIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutFleetsFleetIDWingsWingIDParamsWithTimeout(timeout time.Duration) *PutFleetsFleetIDWingsWingIDParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &PutFleetsFleetIDWingsWingIDParams{
		Datasource: &datasourceDefault,

		timeout: timeout,
	}
}

// NewPutFleetsFleetIDWingsWingIDParamsWithContext creates a new PutFleetsFleetIDWingsWingIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutFleetsFleetIDWingsWingIDParamsWithContext(ctx context.Context) *PutFleetsFleetIDWingsWingIDParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &PutFleetsFleetIDWingsWingIDParams{
		Datasource: &datasourceDefault,

		Context: ctx,
	}
}

// NewPutFleetsFleetIDWingsWingIDParamsWithHTTPClient creates a new PutFleetsFleetIDWingsWingIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutFleetsFleetIDWingsWingIDParamsWithHTTPClient(client *http.Client) *PutFleetsFleetIDWingsWingIDParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &PutFleetsFleetIDWingsWingIDParams{
		Datasource: &datasourceDefault,
		HTTPClient: client,
	}
}

/*PutFleetsFleetIDWingsWingIDParams contains all the parameters to send to the API endpoint
for the put fleets fleet id wings wing id operation typically these are written to a http.Request
*/
type PutFleetsFleetIDWingsWingIDParams struct {

	/*XUserAgent
	  Client identifier, takes precedence over User-Agent

	*/
	XUserAgent *string
	/*Datasource
	  The server name you would like data from

	*/
	Datasource *string
	/*FleetID
	  ID for a fleet

	*/
	FleetID int64
	/*Naming
	  New name of the wing

	*/
	Naming *models.PutFleetsFleetIDWingsWingIDParamsBody
	/*Token
	  Access token to use, if preferred over a header

	*/
	Token *string
	/*UserAgent
	  Client identifier, takes precedence over headers

	*/
	UserAgent *string
	/*WingID
	  The wing to rename

	*/
	WingID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put fleets fleet id wings wing id params
func (o *PutFleetsFleetIDWingsWingIDParams) WithTimeout(timeout time.Duration) *PutFleetsFleetIDWingsWingIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put fleets fleet id wings wing id params
func (o *PutFleetsFleetIDWingsWingIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put fleets fleet id wings wing id params
func (o *PutFleetsFleetIDWingsWingIDParams) WithContext(ctx context.Context) *PutFleetsFleetIDWingsWingIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put fleets fleet id wings wing id params
func (o *PutFleetsFleetIDWingsWingIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put fleets fleet id wings wing id params
func (o *PutFleetsFleetIDWingsWingIDParams) WithHTTPClient(client *http.Client) *PutFleetsFleetIDWingsWingIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put fleets fleet id wings wing id params
func (o *PutFleetsFleetIDWingsWingIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXUserAgent adds the xUserAgent to the put fleets fleet id wings wing id params
func (o *PutFleetsFleetIDWingsWingIDParams) WithXUserAgent(xUserAgent *string) *PutFleetsFleetIDWingsWingIDParams {
	o.SetXUserAgent(xUserAgent)
	return o
}

// SetXUserAgent adds the xUserAgent to the put fleets fleet id wings wing id params
func (o *PutFleetsFleetIDWingsWingIDParams) SetXUserAgent(xUserAgent *string) {
	o.XUserAgent = xUserAgent
}

// WithDatasource adds the datasource to the put fleets fleet id wings wing id params
func (o *PutFleetsFleetIDWingsWingIDParams) WithDatasource(datasource *string) *PutFleetsFleetIDWingsWingIDParams {
	o.SetDatasource(datasource)
	return o
}

// SetDatasource adds the datasource to the put fleets fleet id wings wing id params
func (o *PutFleetsFleetIDWingsWingIDParams) SetDatasource(datasource *string) {
	o.Datasource = datasource
}

// WithFleetID adds the fleetID to the put fleets fleet id wings wing id params
func (o *PutFleetsFleetIDWingsWingIDParams) WithFleetID(fleetID int64) *PutFleetsFleetIDWingsWingIDParams {
	o.SetFleetID(fleetID)
	return o
}

// SetFleetID adds the fleetId to the put fleets fleet id wings wing id params
func (o *PutFleetsFleetIDWingsWingIDParams) SetFleetID(fleetID int64) {
	o.FleetID = fleetID
}

// WithNaming adds the naming to the put fleets fleet id wings wing id params
func (o *PutFleetsFleetIDWingsWingIDParams) WithNaming(naming *models.PutFleetsFleetIDWingsWingIDParamsBody) *PutFleetsFleetIDWingsWingIDParams {
	o.SetNaming(naming)
	return o
}

// SetNaming adds the naming to the put fleets fleet id wings wing id params
func (o *PutFleetsFleetIDWingsWingIDParams) SetNaming(naming *models.PutFleetsFleetIDWingsWingIDParamsBody) {
	o.Naming = naming
}

// WithToken adds the token to the put fleets fleet id wings wing id params
func (o *PutFleetsFleetIDWingsWingIDParams) WithToken(token *string) *PutFleetsFleetIDWingsWingIDParams {
	o.SetToken(token)
	return o
}

// SetToken adds the token to the put fleets fleet id wings wing id params
func (o *PutFleetsFleetIDWingsWingIDParams) SetToken(token *string) {
	o.Token = token
}

// WithUserAgent adds the userAgent to the put fleets fleet id wings wing id params
func (o *PutFleetsFleetIDWingsWingIDParams) WithUserAgent(userAgent *string) *PutFleetsFleetIDWingsWingIDParams {
	o.SetUserAgent(userAgent)
	return o
}

// SetUserAgent adds the userAgent to the put fleets fleet id wings wing id params
func (o *PutFleetsFleetIDWingsWingIDParams) SetUserAgent(userAgent *string) {
	o.UserAgent = userAgent
}

// WithWingID adds the wingID to the put fleets fleet id wings wing id params
func (o *PutFleetsFleetIDWingsWingIDParams) WithWingID(wingID int64) *PutFleetsFleetIDWingsWingIDParams {
	o.SetWingID(wingID)
	return o
}

// SetWingID adds the wingId to the put fleets fleet id wings wing id params
func (o *PutFleetsFleetIDWingsWingIDParams) SetWingID(wingID int64) {
	o.WingID = wingID
}

// WriteToRequest writes these params to a swagger request
func (o *PutFleetsFleetIDWingsWingIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.XUserAgent != nil {

		// header param X-User-Agent
		if err := r.SetHeaderParam("X-User-Agent", *o.XUserAgent); err != nil {
			return err
		}

	}

	if o.Datasource != nil {

		// query param datasource
		var qrDatasource string
		if o.Datasource != nil {
			qrDatasource = *o.Datasource
		}
		qDatasource := qrDatasource
		if qDatasource != "" {
			if err := r.SetQueryParam("datasource", qDatasource); err != nil {
				return err
			}
		}

	}

	// path param fleet_id
	if err := r.SetPathParam("fleet_id", swag.FormatInt64(o.FleetID)); err != nil {
		return err
	}

	if o.Naming == nil {
		o.Naming = new(models.PutFleetsFleetIDWingsWingIDParamsBody)
	}

	if err := r.SetBodyParam(o.Naming); err != nil {
		return err
	}

	if o.Token != nil {

		// query param token
		var qrToken string
		if o.Token != nil {
			qrToken = *o.Token
		}
		qToken := qrToken
		if qToken != "" {
			if err := r.SetQueryParam("token", qToken); err != nil {
				return err
			}
		}

	}

	if o.UserAgent != nil {

		// query param user_agent
		var qrUserAgent string
		if o.UserAgent != nil {
			qrUserAgent = *o.UserAgent
		}
		qUserAgent := qrUserAgent
		if qUserAgent != "" {
			if err := r.SetQueryParam("user_agent", qUserAgent); err != nil {
				return err
			}
		}

	}

	// path param wing_id
	if err := r.SetPathParam("wing_id", swag.FormatInt64(o.WingID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}