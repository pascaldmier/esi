package character

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

// NewGetCharactersCharacterIDRolesParams creates a new GetCharactersCharacterIDRolesParams object
// with the default values initialized.
func NewGetCharactersCharacterIDRolesParams() *GetCharactersCharacterIDRolesParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetCharactersCharacterIDRolesParams{
		Datasource: &datasourceDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCharactersCharacterIDRolesParamsWithTimeout creates a new GetCharactersCharacterIDRolesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCharactersCharacterIDRolesParamsWithTimeout(timeout time.Duration) *GetCharactersCharacterIDRolesParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetCharactersCharacterIDRolesParams{
		Datasource: &datasourceDefault,

		timeout: timeout,
	}
}

// NewGetCharactersCharacterIDRolesParamsWithContext creates a new GetCharactersCharacterIDRolesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCharactersCharacterIDRolesParamsWithContext(ctx context.Context) *GetCharactersCharacterIDRolesParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetCharactersCharacterIDRolesParams{
		Datasource: &datasourceDefault,

		Context: ctx,
	}
}

// NewGetCharactersCharacterIDRolesParamsWithHTTPClient creates a new GetCharactersCharacterIDRolesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCharactersCharacterIDRolesParamsWithHTTPClient(client *http.Client) *GetCharactersCharacterIDRolesParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetCharactersCharacterIDRolesParams{
		Datasource: &datasourceDefault,
		HTTPClient: client,
	}
}

/*GetCharactersCharacterIDRolesParams contains all the parameters to send to the API endpoint
for the get characters character id roles operation typically these are written to a http.Request
*/
type GetCharactersCharacterIDRolesParams struct {

	/*XUserAgent
	  Client identifier, takes precedence over User-Agent

	*/
	XUserAgent *string
	/*CharacterID
	  A character ID

	*/
	CharacterID int32
	/*Datasource
	  The server name you would like data from

	*/
	Datasource *string
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

// WithTimeout adds the timeout to the get characters character id roles params
func (o *GetCharactersCharacterIDRolesParams) WithTimeout(timeout time.Duration) *GetCharactersCharacterIDRolesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get characters character id roles params
func (o *GetCharactersCharacterIDRolesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get characters character id roles params
func (o *GetCharactersCharacterIDRolesParams) WithContext(ctx context.Context) *GetCharactersCharacterIDRolesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get characters character id roles params
func (o *GetCharactersCharacterIDRolesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get characters character id roles params
func (o *GetCharactersCharacterIDRolesParams) WithHTTPClient(client *http.Client) *GetCharactersCharacterIDRolesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get characters character id roles params
func (o *GetCharactersCharacterIDRolesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXUserAgent adds the xUserAgent to the get characters character id roles params
func (o *GetCharactersCharacterIDRolesParams) WithXUserAgent(xUserAgent *string) *GetCharactersCharacterIDRolesParams {
	o.SetXUserAgent(xUserAgent)
	return o
}

// SetXUserAgent adds the xUserAgent to the get characters character id roles params
func (o *GetCharactersCharacterIDRolesParams) SetXUserAgent(xUserAgent *string) {
	o.XUserAgent = xUserAgent
}

// WithCharacterID adds the characterID to the get characters character id roles params
func (o *GetCharactersCharacterIDRolesParams) WithCharacterID(characterID int32) *GetCharactersCharacterIDRolesParams {
	o.SetCharacterID(characterID)
	return o
}

// SetCharacterID adds the characterId to the get characters character id roles params
func (o *GetCharactersCharacterIDRolesParams) SetCharacterID(characterID int32) {
	o.CharacterID = characterID
}

// WithDatasource adds the datasource to the get characters character id roles params
func (o *GetCharactersCharacterIDRolesParams) WithDatasource(datasource *string) *GetCharactersCharacterIDRolesParams {
	o.SetDatasource(datasource)
	return o
}

// SetDatasource adds the datasource to the get characters character id roles params
func (o *GetCharactersCharacterIDRolesParams) SetDatasource(datasource *string) {
	o.Datasource = datasource
}

// WithToken adds the token to the get characters character id roles params
func (o *GetCharactersCharacterIDRolesParams) WithToken(token *string) *GetCharactersCharacterIDRolesParams {
	o.SetToken(token)
	return o
}

// SetToken adds the token to the get characters character id roles params
func (o *GetCharactersCharacterIDRolesParams) SetToken(token *string) {
	o.Token = token
}

// WithUserAgent adds the userAgent to the get characters character id roles params
func (o *GetCharactersCharacterIDRolesParams) WithUserAgent(userAgent *string) *GetCharactersCharacterIDRolesParams {
	o.SetUserAgent(userAgent)
	return o
}

// SetUserAgent adds the userAgent to the get characters character id roles params
func (o *GetCharactersCharacterIDRolesParams) SetUserAgent(userAgent *string) {
	o.UserAgent = userAgent
}

// WriteToRequest writes these params to a swagger request
func (o *GetCharactersCharacterIDRolesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param character_id
	if err := r.SetPathParam("character_id", swag.FormatInt32(o.CharacterID)); err != nil {
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
