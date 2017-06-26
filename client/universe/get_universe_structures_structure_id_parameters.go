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

// NewGetUniverseStructuresStructureIDParams creates a new GetUniverseStructuresStructureIDParams object
// with the default values initialized.
func NewGetUniverseStructuresStructureIDParams() *GetUniverseStructuresStructureIDParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetUniverseStructuresStructureIDParams{
		Datasource: &datasourceDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetUniverseStructuresStructureIDParamsWithTimeout creates a new GetUniverseStructuresStructureIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetUniverseStructuresStructureIDParamsWithTimeout(timeout time.Duration) *GetUniverseStructuresStructureIDParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetUniverseStructuresStructureIDParams{
		Datasource: &datasourceDefault,

		timeout: timeout,
	}
}

// NewGetUniverseStructuresStructureIDParamsWithContext creates a new GetUniverseStructuresStructureIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetUniverseStructuresStructureIDParamsWithContext(ctx context.Context) *GetUniverseStructuresStructureIDParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetUniverseStructuresStructureIDParams{
		Datasource: &datasourceDefault,

		Context: ctx,
	}
}

// NewGetUniverseStructuresStructureIDParamsWithHTTPClient creates a new GetUniverseStructuresStructureIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetUniverseStructuresStructureIDParamsWithHTTPClient(client *http.Client) *GetUniverseStructuresStructureIDParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetUniverseStructuresStructureIDParams{
		Datasource: &datasourceDefault,
		HTTPClient: client,
	}
}

/*GetUniverseStructuresStructureIDParams contains all the parameters to send to the API endpoint
for the get universe structures structure id operation typically these are written to a http.Request
*/
type GetUniverseStructuresStructureIDParams struct {

	/*XUserAgent
	  Client identifier, takes precedence over User-Agent

	*/
	XUserAgent *string
	/*Datasource
	  The server name you would like data from

	*/
	Datasource *string
	/*StructureID
	  An Eve structure ID

	*/
	StructureID int64
	/*Token
	  Access token to use, if preferred over a header

	*/
	Token *string
	/*UserAgent
	  Client identifier, takes precedence over headers

	*/
	UserAgent *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get universe structures structure id params
func (o *GetUniverseStructuresStructureIDParams) WithTimeout(timeout time.Duration) *GetUniverseStructuresStructureIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get universe structures structure id params
func (o *GetUniverseStructuresStructureIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get universe structures structure id params
func (o *GetUniverseStructuresStructureIDParams) WithContext(ctx context.Context) *GetUniverseStructuresStructureIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get universe structures structure id params
func (o *GetUniverseStructuresStructureIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get universe structures structure id params
func (o *GetUniverseStructuresStructureIDParams) WithHTTPClient(client *http.Client) *GetUniverseStructuresStructureIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get universe structures structure id params
func (o *GetUniverseStructuresStructureIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXUserAgent adds the xUserAgent to the get universe structures structure id params
func (o *GetUniverseStructuresStructureIDParams) WithXUserAgent(xUserAgent *string) *GetUniverseStructuresStructureIDParams {
	o.SetXUserAgent(xUserAgent)
	return o
}

// SetXUserAgent adds the xUserAgent to the get universe structures structure id params
func (o *GetUniverseStructuresStructureIDParams) SetXUserAgent(xUserAgent *string) {
	o.XUserAgent = xUserAgent
}

// WithDatasource adds the datasource to the get universe structures structure id params
func (o *GetUniverseStructuresStructureIDParams) WithDatasource(datasource *string) *GetUniverseStructuresStructureIDParams {
	o.SetDatasource(datasource)
	return o
}

// SetDatasource adds the datasource to the get universe structures structure id params
func (o *GetUniverseStructuresStructureIDParams) SetDatasource(datasource *string) {
	o.Datasource = datasource
}

// WithStructureID adds the structureID to the get universe structures structure id params
func (o *GetUniverseStructuresStructureIDParams) WithStructureID(structureID int64) *GetUniverseStructuresStructureIDParams {
	o.SetStructureID(structureID)
	return o
}

// SetStructureID adds the structureId to the get universe structures structure id params
func (o *GetUniverseStructuresStructureIDParams) SetStructureID(structureID int64) {
	o.StructureID = structureID
}

// WithToken adds the token to the get universe structures structure id params
func (o *GetUniverseStructuresStructureIDParams) WithToken(token *string) *GetUniverseStructuresStructureIDParams {
	o.SetToken(token)
	return o
}

// SetToken adds the token to the get universe structures structure id params
func (o *GetUniverseStructuresStructureIDParams) SetToken(token *string) {
	o.Token = token
}

// WithUserAgent adds the userAgent to the get universe structures structure id params
func (o *GetUniverseStructuresStructureIDParams) WithUserAgent(userAgent *string) *GetUniverseStructuresStructureIDParams {
	o.SetUserAgent(userAgent)
	return o
}

// SetUserAgent adds the userAgent to the get universe structures structure id params
func (o *GetUniverseStructuresStructureIDParams) SetUserAgent(userAgent *string) {
	o.UserAgent = userAgent
}

// WriteToRequest writes these params to a swagger request
func (o *GetUniverseStructuresStructureIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param structure_id
	if err := r.SetPathParam("structure_id", swag.FormatInt64(o.StructureID)); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}