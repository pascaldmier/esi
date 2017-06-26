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

// NewPostFleetsFleetIDMembersParams creates a new PostFleetsFleetIDMembersParams object
// with the default values initialized.
func NewPostFleetsFleetIDMembersParams() *PostFleetsFleetIDMembersParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &PostFleetsFleetIDMembersParams{
		Datasource: &datasourceDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewPostFleetsFleetIDMembersParamsWithTimeout creates a new PostFleetsFleetIDMembersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostFleetsFleetIDMembersParamsWithTimeout(timeout time.Duration) *PostFleetsFleetIDMembersParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &PostFleetsFleetIDMembersParams{
		Datasource: &datasourceDefault,

		timeout: timeout,
	}
}

// NewPostFleetsFleetIDMembersParamsWithContext creates a new PostFleetsFleetIDMembersParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostFleetsFleetIDMembersParamsWithContext(ctx context.Context) *PostFleetsFleetIDMembersParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &PostFleetsFleetIDMembersParams{
		Datasource: &datasourceDefault,

		Context: ctx,
	}
}

// NewPostFleetsFleetIDMembersParamsWithHTTPClient creates a new PostFleetsFleetIDMembersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostFleetsFleetIDMembersParamsWithHTTPClient(client *http.Client) *PostFleetsFleetIDMembersParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &PostFleetsFleetIDMembersParams{
		Datasource: &datasourceDefault,
		HTTPClient: client,
	}
}

/*PostFleetsFleetIDMembersParams contains all the parameters to send to the API endpoint
for the post fleets fleet id members operation typically these are written to a http.Request
*/
type PostFleetsFleetIDMembersParams struct {

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
	/*Invitation
	  Details of the invitation

	*/
	Invitation *models.PostFleetsFleetIDMembersParamsBody
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

// WithTimeout adds the timeout to the post fleets fleet id members params
func (o *PostFleetsFleetIDMembersParams) WithTimeout(timeout time.Duration) *PostFleetsFleetIDMembersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post fleets fleet id members params
func (o *PostFleetsFleetIDMembersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post fleets fleet id members params
func (o *PostFleetsFleetIDMembersParams) WithContext(ctx context.Context) *PostFleetsFleetIDMembersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post fleets fleet id members params
func (o *PostFleetsFleetIDMembersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post fleets fleet id members params
func (o *PostFleetsFleetIDMembersParams) WithHTTPClient(client *http.Client) *PostFleetsFleetIDMembersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post fleets fleet id members params
func (o *PostFleetsFleetIDMembersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXUserAgent adds the xUserAgent to the post fleets fleet id members params
func (o *PostFleetsFleetIDMembersParams) WithXUserAgent(xUserAgent *string) *PostFleetsFleetIDMembersParams {
	o.SetXUserAgent(xUserAgent)
	return o
}

// SetXUserAgent adds the xUserAgent to the post fleets fleet id members params
func (o *PostFleetsFleetIDMembersParams) SetXUserAgent(xUserAgent *string) {
	o.XUserAgent = xUserAgent
}

// WithDatasource adds the datasource to the post fleets fleet id members params
func (o *PostFleetsFleetIDMembersParams) WithDatasource(datasource *string) *PostFleetsFleetIDMembersParams {
	o.SetDatasource(datasource)
	return o
}

// SetDatasource adds the datasource to the post fleets fleet id members params
func (o *PostFleetsFleetIDMembersParams) SetDatasource(datasource *string) {
	o.Datasource = datasource
}

// WithFleetID adds the fleetID to the post fleets fleet id members params
func (o *PostFleetsFleetIDMembersParams) WithFleetID(fleetID int64) *PostFleetsFleetIDMembersParams {
	o.SetFleetID(fleetID)
	return o
}

// SetFleetID adds the fleetId to the post fleets fleet id members params
func (o *PostFleetsFleetIDMembersParams) SetFleetID(fleetID int64) {
	o.FleetID = fleetID
}

// WithInvitation adds the invitation to the post fleets fleet id members params
func (o *PostFleetsFleetIDMembersParams) WithInvitation(invitation *models.PostFleetsFleetIDMembersParamsBody) *PostFleetsFleetIDMembersParams {
	o.SetInvitation(invitation)
	return o
}

// SetInvitation adds the invitation to the post fleets fleet id members params
func (o *PostFleetsFleetIDMembersParams) SetInvitation(invitation *models.PostFleetsFleetIDMembersParamsBody) {
	o.Invitation = invitation
}

// WithToken adds the token to the post fleets fleet id members params
func (o *PostFleetsFleetIDMembersParams) WithToken(token *string) *PostFleetsFleetIDMembersParams {
	o.SetToken(token)
	return o
}

// SetToken adds the token to the post fleets fleet id members params
func (o *PostFleetsFleetIDMembersParams) SetToken(token *string) {
	o.Token = token
}

// WithUserAgent adds the userAgent to the post fleets fleet id members params
func (o *PostFleetsFleetIDMembersParams) WithUserAgent(userAgent *string) *PostFleetsFleetIDMembersParams {
	o.SetUserAgent(userAgent)
	return o
}

// SetUserAgent adds the userAgent to the post fleets fleet id members params
func (o *PostFleetsFleetIDMembersParams) SetUserAgent(userAgent *string) {
	o.UserAgent = userAgent
}

// WriteToRequest writes these params to a swagger request
func (o *PostFleetsFleetIDMembersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Invitation == nil {
		o.Invitation = new(models.PostFleetsFleetIDMembersParamsBody)
	}

	if err := r.SetBodyParam(o.Invitation); err != nil {
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