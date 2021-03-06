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
)

// NewDeleteFleetsFleetIDWingsWingIDParams creates a new DeleteFleetsFleetIDWingsWingIDParams object
// with the default values initialized.
func NewDeleteFleetsFleetIDWingsWingIDParams() *DeleteFleetsFleetIDWingsWingIDParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &DeleteFleetsFleetIDWingsWingIDParams{
		Datasource: &datasourceDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteFleetsFleetIDWingsWingIDParamsWithTimeout creates a new DeleteFleetsFleetIDWingsWingIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteFleetsFleetIDWingsWingIDParamsWithTimeout(timeout time.Duration) *DeleteFleetsFleetIDWingsWingIDParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &DeleteFleetsFleetIDWingsWingIDParams{
		Datasource: &datasourceDefault,

		timeout: timeout,
	}
}

// NewDeleteFleetsFleetIDWingsWingIDParamsWithContext creates a new DeleteFleetsFleetIDWingsWingIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteFleetsFleetIDWingsWingIDParamsWithContext(ctx context.Context) *DeleteFleetsFleetIDWingsWingIDParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &DeleteFleetsFleetIDWingsWingIDParams{
		Datasource: &datasourceDefault,

		Context: ctx,
	}
}

// NewDeleteFleetsFleetIDWingsWingIDParamsWithHTTPClient creates a new DeleteFleetsFleetIDWingsWingIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteFleetsFleetIDWingsWingIDParamsWithHTTPClient(client *http.Client) *DeleteFleetsFleetIDWingsWingIDParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &DeleteFleetsFleetIDWingsWingIDParams{
		Datasource: &datasourceDefault,
		HTTPClient: client,
	}
}

/*DeleteFleetsFleetIDWingsWingIDParams contains all the parameters to send to the API endpoint
for the delete fleets fleet id wings wing id operation typically these are written to a http.Request
*/
type DeleteFleetsFleetIDWingsWingIDParams struct {

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
	/*Token
	  Access token to use, if preferred over a header

	*/
	Token *string
	/*UserAgent
	  Client identifier, takes precedence over headers

	*/
	UserAgent *string
	/*WingID
	  The wing to delete

	*/
	WingID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete fleets fleet id wings wing id params
func (o *DeleteFleetsFleetIDWingsWingIDParams) WithTimeout(timeout time.Duration) *DeleteFleetsFleetIDWingsWingIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete fleets fleet id wings wing id params
func (o *DeleteFleetsFleetIDWingsWingIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete fleets fleet id wings wing id params
func (o *DeleteFleetsFleetIDWingsWingIDParams) WithContext(ctx context.Context) *DeleteFleetsFleetIDWingsWingIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete fleets fleet id wings wing id params
func (o *DeleteFleetsFleetIDWingsWingIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete fleets fleet id wings wing id params
func (o *DeleteFleetsFleetIDWingsWingIDParams) WithHTTPClient(client *http.Client) *DeleteFleetsFleetIDWingsWingIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete fleets fleet id wings wing id params
func (o *DeleteFleetsFleetIDWingsWingIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXUserAgent adds the xUserAgent to the delete fleets fleet id wings wing id params
func (o *DeleteFleetsFleetIDWingsWingIDParams) WithXUserAgent(xUserAgent *string) *DeleteFleetsFleetIDWingsWingIDParams {
	o.SetXUserAgent(xUserAgent)
	return o
}

// SetXUserAgent adds the xUserAgent to the delete fleets fleet id wings wing id params
func (o *DeleteFleetsFleetIDWingsWingIDParams) SetXUserAgent(xUserAgent *string) {
	o.XUserAgent = xUserAgent
}

// WithDatasource adds the datasource to the delete fleets fleet id wings wing id params
func (o *DeleteFleetsFleetIDWingsWingIDParams) WithDatasource(datasource *string) *DeleteFleetsFleetIDWingsWingIDParams {
	o.SetDatasource(datasource)
	return o
}

// SetDatasource adds the datasource to the delete fleets fleet id wings wing id params
func (o *DeleteFleetsFleetIDWingsWingIDParams) SetDatasource(datasource *string) {
	o.Datasource = datasource
}

// WithFleetID adds the fleetID to the delete fleets fleet id wings wing id params
func (o *DeleteFleetsFleetIDWingsWingIDParams) WithFleetID(fleetID int64) *DeleteFleetsFleetIDWingsWingIDParams {
	o.SetFleetID(fleetID)
	return o
}

// SetFleetID adds the fleetId to the delete fleets fleet id wings wing id params
func (o *DeleteFleetsFleetIDWingsWingIDParams) SetFleetID(fleetID int64) {
	o.FleetID = fleetID
}

// WithToken adds the token to the delete fleets fleet id wings wing id params
func (o *DeleteFleetsFleetIDWingsWingIDParams) WithToken(token *string) *DeleteFleetsFleetIDWingsWingIDParams {
	o.SetToken(token)
	return o
}

// SetToken adds the token to the delete fleets fleet id wings wing id params
func (o *DeleteFleetsFleetIDWingsWingIDParams) SetToken(token *string) {
	o.Token = token
}

// WithUserAgent adds the userAgent to the delete fleets fleet id wings wing id params
func (o *DeleteFleetsFleetIDWingsWingIDParams) WithUserAgent(userAgent *string) *DeleteFleetsFleetIDWingsWingIDParams {
	o.SetUserAgent(userAgent)
	return o
}

// SetUserAgent adds the userAgent to the delete fleets fleet id wings wing id params
func (o *DeleteFleetsFleetIDWingsWingIDParams) SetUserAgent(userAgent *string) {
	o.UserAgent = userAgent
}

// WithWingID adds the wingID to the delete fleets fleet id wings wing id params
func (o *DeleteFleetsFleetIDWingsWingIDParams) WithWingID(wingID int64) *DeleteFleetsFleetIDWingsWingIDParams {
	o.SetWingID(wingID)
	return o
}

// SetWingID adds the wingId to the delete fleets fleet id wings wing id params
func (o *DeleteFleetsFleetIDWingsWingIDParams) SetWingID(wingID int64) {
	o.WingID = wingID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteFleetsFleetIDWingsWingIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
