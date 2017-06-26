package universe

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

// NewGetUniverseMoonsMoonIDParams creates a new GetUniverseMoonsMoonIDParams object
// with the default values initialized.
func NewGetUniverseMoonsMoonIDParams() *GetUniverseMoonsMoonIDParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetUniverseMoonsMoonIDParams{
		Datasource: &datasourceDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetUniverseMoonsMoonIDParamsWithTimeout creates a new GetUniverseMoonsMoonIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetUniverseMoonsMoonIDParamsWithTimeout(timeout time.Duration) *GetUniverseMoonsMoonIDParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetUniverseMoonsMoonIDParams{
		Datasource: &datasourceDefault,

		timeout: timeout,
	}
}

// NewGetUniverseMoonsMoonIDParamsWithContext creates a new GetUniverseMoonsMoonIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetUniverseMoonsMoonIDParamsWithContext(ctx context.Context) *GetUniverseMoonsMoonIDParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetUniverseMoonsMoonIDParams{
		Datasource: &datasourceDefault,

		Context: ctx,
	}
}

// NewGetUniverseMoonsMoonIDParamsWithHTTPClient creates a new GetUniverseMoonsMoonIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetUniverseMoonsMoonIDParamsWithHTTPClient(client *http.Client) *GetUniverseMoonsMoonIDParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetUniverseMoonsMoonIDParams{
		Datasource: &datasourceDefault,
		HTTPClient: client,
	}
}

/*GetUniverseMoonsMoonIDParams contains all the parameters to send to the API endpoint
for the get universe moons moon id operation typically these are written to a http.Request
*/
type GetUniverseMoonsMoonIDParams struct {

	/*XUserAgent
	  Client identifier, takes precedence over User-Agent

	*/
	XUserAgent *string
	/*Datasource
	  The server name you would like data from

	*/
	Datasource *string
	/*MoonID
	  moon_id integer

	*/
	MoonID int32
	/*UserAgent
	  Client identifier, takes precedence over headers

	*/
	UserAgent *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get universe moons moon id params
func (o *GetUniverseMoonsMoonIDParams) WithTimeout(timeout time.Duration) *GetUniverseMoonsMoonIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get universe moons moon id params
func (o *GetUniverseMoonsMoonIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get universe moons moon id params
func (o *GetUniverseMoonsMoonIDParams) WithContext(ctx context.Context) *GetUniverseMoonsMoonIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get universe moons moon id params
func (o *GetUniverseMoonsMoonIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get universe moons moon id params
func (o *GetUniverseMoonsMoonIDParams) WithHTTPClient(client *http.Client) *GetUniverseMoonsMoonIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get universe moons moon id params
func (o *GetUniverseMoonsMoonIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXUserAgent adds the xUserAgent to the get universe moons moon id params
func (o *GetUniverseMoonsMoonIDParams) WithXUserAgent(xUserAgent *string) *GetUniverseMoonsMoonIDParams {
	o.SetXUserAgent(xUserAgent)
	return o
}

// SetXUserAgent adds the xUserAgent to the get universe moons moon id params
func (o *GetUniverseMoonsMoonIDParams) SetXUserAgent(xUserAgent *string) {
	o.XUserAgent = xUserAgent
}

// WithDatasource adds the datasource to the get universe moons moon id params
func (o *GetUniverseMoonsMoonIDParams) WithDatasource(datasource *string) *GetUniverseMoonsMoonIDParams {
	o.SetDatasource(datasource)
	return o
}

// SetDatasource adds the datasource to the get universe moons moon id params
func (o *GetUniverseMoonsMoonIDParams) SetDatasource(datasource *string) {
	o.Datasource = datasource
}

// WithMoonID adds the moonID to the get universe moons moon id params
func (o *GetUniverseMoonsMoonIDParams) WithMoonID(moonID int32) *GetUniverseMoonsMoonIDParams {
	o.SetMoonID(moonID)
	return o
}

// SetMoonID adds the moonId to the get universe moons moon id params
func (o *GetUniverseMoonsMoonIDParams) SetMoonID(moonID int32) {
	o.MoonID = moonID
}

// WithUserAgent adds the userAgent to the get universe moons moon id params
func (o *GetUniverseMoonsMoonIDParams) WithUserAgent(userAgent *string) *GetUniverseMoonsMoonIDParams {
	o.SetUserAgent(userAgent)
	return o
}

// SetUserAgent adds the userAgent to the get universe moons moon id params
func (o *GetUniverseMoonsMoonIDParams) SetUserAgent(userAgent *string) {
	o.UserAgent = userAgent
}

// WriteToRequest writes these params to a swagger request
func (o *GetUniverseMoonsMoonIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param moon_id
	if err := r.SetPathParam("moon_id", swag.FormatInt32(o.MoonID)); err != nil {
		return err
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
