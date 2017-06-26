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

// NewGetUniverseGroupsParams creates a new GetUniverseGroupsParams object
// with the default values initialized.
func NewGetUniverseGroupsParams() *GetUniverseGroupsParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetUniverseGroupsParams{
		Datasource: &datasourceDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetUniverseGroupsParamsWithTimeout creates a new GetUniverseGroupsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetUniverseGroupsParamsWithTimeout(timeout time.Duration) *GetUniverseGroupsParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetUniverseGroupsParams{
		Datasource: &datasourceDefault,

		timeout: timeout,
	}
}

// NewGetUniverseGroupsParamsWithContext creates a new GetUniverseGroupsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetUniverseGroupsParamsWithContext(ctx context.Context) *GetUniverseGroupsParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetUniverseGroupsParams{
		Datasource: &datasourceDefault,

		Context: ctx,
	}
}

// NewGetUniverseGroupsParamsWithHTTPClient creates a new GetUniverseGroupsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetUniverseGroupsParamsWithHTTPClient(client *http.Client) *GetUniverseGroupsParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetUniverseGroupsParams{
		Datasource: &datasourceDefault,
		HTTPClient: client,
	}
}

/*GetUniverseGroupsParams contains all the parameters to send to the API endpoint
for the get universe groups operation typically these are written to a http.Request
*/
type GetUniverseGroupsParams struct {

	/*XUserAgent
	  Client identifier, takes precedence over User-Agent

	*/
	XUserAgent *string
	/*Datasource
	  The server name you would like data from

	*/
	Datasource *string
	/*Page
	  Which page to query

	*/
	Page *int32
	/*UserAgent
	  Client identifier, takes precedence over headers

	*/
	UserAgent *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get universe groups params
func (o *GetUniverseGroupsParams) WithTimeout(timeout time.Duration) *GetUniverseGroupsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get universe groups params
func (o *GetUniverseGroupsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get universe groups params
func (o *GetUniverseGroupsParams) WithContext(ctx context.Context) *GetUniverseGroupsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get universe groups params
func (o *GetUniverseGroupsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get universe groups params
func (o *GetUniverseGroupsParams) WithHTTPClient(client *http.Client) *GetUniverseGroupsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get universe groups params
func (o *GetUniverseGroupsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXUserAgent adds the xUserAgent to the get universe groups params
func (o *GetUniverseGroupsParams) WithXUserAgent(xUserAgent *string) *GetUniverseGroupsParams {
	o.SetXUserAgent(xUserAgent)
	return o
}

// SetXUserAgent adds the xUserAgent to the get universe groups params
func (o *GetUniverseGroupsParams) SetXUserAgent(xUserAgent *string) {
	o.XUserAgent = xUserAgent
}

// WithDatasource adds the datasource to the get universe groups params
func (o *GetUniverseGroupsParams) WithDatasource(datasource *string) *GetUniverseGroupsParams {
	o.SetDatasource(datasource)
	return o
}

// SetDatasource adds the datasource to the get universe groups params
func (o *GetUniverseGroupsParams) SetDatasource(datasource *string) {
	o.Datasource = datasource
}

// WithPage adds the page to the get universe groups params
func (o *GetUniverseGroupsParams) WithPage(page *int32) *GetUniverseGroupsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get universe groups params
func (o *GetUniverseGroupsParams) SetPage(page *int32) {
	o.Page = page
}

// WithUserAgent adds the userAgent to the get universe groups params
func (o *GetUniverseGroupsParams) WithUserAgent(userAgent *string) *GetUniverseGroupsParams {
	o.SetUserAgent(userAgent)
	return o
}

// SetUserAgent adds the userAgent to the get universe groups params
func (o *GetUniverseGroupsParams) SetUserAgent(userAgent *string) {
	o.UserAgent = userAgent
}

// WriteToRequest writes these params to a swagger request
func (o *GetUniverseGroupsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Page != nil {

		// query param page
		var qrPage int32
		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt32(qrPage)
		if qPage != "" {
			if err := r.SetQueryParam("page", qPage); err != nil {
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
