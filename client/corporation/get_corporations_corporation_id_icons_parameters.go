package corporation

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

// NewGetCorporationsCorporationIDIconsParams creates a new GetCorporationsCorporationIDIconsParams object
// with the default values initialized.
func NewGetCorporationsCorporationIDIconsParams() *GetCorporationsCorporationIDIconsParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetCorporationsCorporationIDIconsParams{
		Datasource: &datasourceDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCorporationsCorporationIDIconsParamsWithTimeout creates a new GetCorporationsCorporationIDIconsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCorporationsCorporationIDIconsParamsWithTimeout(timeout time.Duration) *GetCorporationsCorporationIDIconsParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetCorporationsCorporationIDIconsParams{
		Datasource: &datasourceDefault,

		timeout: timeout,
	}
}

// NewGetCorporationsCorporationIDIconsParamsWithContext creates a new GetCorporationsCorporationIDIconsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCorporationsCorporationIDIconsParamsWithContext(ctx context.Context) *GetCorporationsCorporationIDIconsParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetCorporationsCorporationIDIconsParams{
		Datasource: &datasourceDefault,

		Context: ctx,
	}
}

// NewGetCorporationsCorporationIDIconsParamsWithHTTPClient creates a new GetCorporationsCorporationIDIconsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCorporationsCorporationIDIconsParamsWithHTTPClient(client *http.Client) *GetCorporationsCorporationIDIconsParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetCorporationsCorporationIDIconsParams{
		Datasource: &datasourceDefault,
		HTTPClient: client,
	}
}

/*GetCorporationsCorporationIDIconsParams contains all the parameters to send to the API endpoint
for the get corporations corporation id icons operation typically these are written to a http.Request
*/
type GetCorporationsCorporationIDIconsParams struct {

	/*XUserAgent
	  Client identifier, takes precedence over User-Agent

	*/
	XUserAgent *string
	/*CorporationID
	  An EVE corporation ID

	*/
	CorporationID int32
	/*Datasource
	  The server name you would like data from

	*/
	Datasource *string
	/*UserAgent
	  Client identifier, takes precedence over headers

	*/
	UserAgent *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get corporations corporation id icons params
func (o *GetCorporationsCorporationIDIconsParams) WithTimeout(timeout time.Duration) *GetCorporationsCorporationIDIconsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get corporations corporation id icons params
func (o *GetCorporationsCorporationIDIconsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get corporations corporation id icons params
func (o *GetCorporationsCorporationIDIconsParams) WithContext(ctx context.Context) *GetCorporationsCorporationIDIconsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get corporations corporation id icons params
func (o *GetCorporationsCorporationIDIconsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get corporations corporation id icons params
func (o *GetCorporationsCorporationIDIconsParams) WithHTTPClient(client *http.Client) *GetCorporationsCorporationIDIconsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get corporations corporation id icons params
func (o *GetCorporationsCorporationIDIconsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXUserAgent adds the xUserAgent to the get corporations corporation id icons params
func (o *GetCorporationsCorporationIDIconsParams) WithXUserAgent(xUserAgent *string) *GetCorporationsCorporationIDIconsParams {
	o.SetXUserAgent(xUserAgent)
	return o
}

// SetXUserAgent adds the xUserAgent to the get corporations corporation id icons params
func (o *GetCorporationsCorporationIDIconsParams) SetXUserAgent(xUserAgent *string) {
	o.XUserAgent = xUserAgent
}

// WithCorporationID adds the corporationID to the get corporations corporation id icons params
func (o *GetCorporationsCorporationIDIconsParams) WithCorporationID(corporationID int32) *GetCorporationsCorporationIDIconsParams {
	o.SetCorporationID(corporationID)
	return o
}

// SetCorporationID adds the corporationId to the get corporations corporation id icons params
func (o *GetCorporationsCorporationIDIconsParams) SetCorporationID(corporationID int32) {
	o.CorporationID = corporationID
}

// WithDatasource adds the datasource to the get corporations corporation id icons params
func (o *GetCorporationsCorporationIDIconsParams) WithDatasource(datasource *string) *GetCorporationsCorporationIDIconsParams {
	o.SetDatasource(datasource)
	return o
}

// SetDatasource adds the datasource to the get corporations corporation id icons params
func (o *GetCorporationsCorporationIDIconsParams) SetDatasource(datasource *string) {
	o.Datasource = datasource
}

// WithUserAgent adds the userAgent to the get corporations corporation id icons params
func (o *GetCorporationsCorporationIDIconsParams) WithUserAgent(userAgent *string) *GetCorporationsCorporationIDIconsParams {
	o.SetUserAgent(userAgent)
	return o
}

// SetUserAgent adds the userAgent to the get corporations corporation id icons params
func (o *GetCorporationsCorporationIDIconsParams) SetUserAgent(userAgent *string) {
	o.UserAgent = userAgent
}

// WriteToRequest writes these params to a swagger request
func (o *GetCorporationsCorporationIDIconsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param corporation_id
	if err := r.SetPathParam("corporation_id", swag.FormatInt32(o.CorporationID)); err != nil {
		return err
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
