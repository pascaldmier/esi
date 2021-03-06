package industry

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetIndustrySystemsParams creates a new GetIndustrySystemsParams object
// with the default values initialized.
func NewGetIndustrySystemsParams() *GetIndustrySystemsParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetIndustrySystemsParams{
		Datasource: &datasourceDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetIndustrySystemsParamsWithTimeout creates a new GetIndustrySystemsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetIndustrySystemsParamsWithTimeout(timeout time.Duration) *GetIndustrySystemsParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetIndustrySystemsParams{
		Datasource: &datasourceDefault,

		timeout: timeout,
	}
}

// NewGetIndustrySystemsParamsWithContext creates a new GetIndustrySystemsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetIndustrySystemsParamsWithContext(ctx context.Context) *GetIndustrySystemsParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetIndustrySystemsParams{
		Datasource: &datasourceDefault,

		Context: ctx,
	}
}

// NewGetIndustrySystemsParamsWithHTTPClient creates a new GetIndustrySystemsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetIndustrySystemsParamsWithHTTPClient(client *http.Client) *GetIndustrySystemsParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetIndustrySystemsParams{
		Datasource: &datasourceDefault,
		HTTPClient: client,
	}
}

/*GetIndustrySystemsParams contains all the parameters to send to the API endpoint
for the get industry systems operation typically these are written to a http.Request
*/
type GetIndustrySystemsParams struct {

	/*XUserAgent
	  Client identifier, takes precedence over User-Agent

	*/
	XUserAgent *string
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

// WithTimeout adds the timeout to the get industry systems params
func (o *GetIndustrySystemsParams) WithTimeout(timeout time.Duration) *GetIndustrySystemsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get industry systems params
func (o *GetIndustrySystemsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get industry systems params
func (o *GetIndustrySystemsParams) WithContext(ctx context.Context) *GetIndustrySystemsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get industry systems params
func (o *GetIndustrySystemsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get industry systems params
func (o *GetIndustrySystemsParams) WithHTTPClient(client *http.Client) *GetIndustrySystemsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get industry systems params
func (o *GetIndustrySystemsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXUserAgent adds the xUserAgent to the get industry systems params
func (o *GetIndustrySystemsParams) WithXUserAgent(xUserAgent *string) *GetIndustrySystemsParams {
	o.SetXUserAgent(xUserAgent)
	return o
}

// SetXUserAgent adds the xUserAgent to the get industry systems params
func (o *GetIndustrySystemsParams) SetXUserAgent(xUserAgent *string) {
	o.XUserAgent = xUserAgent
}

// WithDatasource adds the datasource to the get industry systems params
func (o *GetIndustrySystemsParams) WithDatasource(datasource *string) *GetIndustrySystemsParams {
	o.SetDatasource(datasource)
	return o
}

// SetDatasource adds the datasource to the get industry systems params
func (o *GetIndustrySystemsParams) SetDatasource(datasource *string) {
	o.Datasource = datasource
}

// WithUserAgent adds the userAgent to the get industry systems params
func (o *GetIndustrySystemsParams) WithUserAgent(userAgent *string) *GetIndustrySystemsParams {
	o.SetUserAgent(userAgent)
	return o
}

// SetUserAgent adds the userAgent to the get industry systems params
func (o *GetIndustrySystemsParams) SetUserAgent(userAgent *string) {
	o.UserAgent = userAgent
}

// WriteToRequest writes these params to a swagger request
func (o *GetIndustrySystemsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
