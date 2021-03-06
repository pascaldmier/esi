package location

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

// NewGetCharactersCharacterIDShipParams creates a new GetCharactersCharacterIDShipParams object
// with the default values initialized.
func NewGetCharactersCharacterIDShipParams() *GetCharactersCharacterIDShipParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetCharactersCharacterIDShipParams{
		Datasource: &datasourceDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCharactersCharacterIDShipParamsWithTimeout creates a new GetCharactersCharacterIDShipParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCharactersCharacterIDShipParamsWithTimeout(timeout time.Duration) *GetCharactersCharacterIDShipParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetCharactersCharacterIDShipParams{
		Datasource: &datasourceDefault,

		timeout: timeout,
	}
}

// NewGetCharactersCharacterIDShipParamsWithContext creates a new GetCharactersCharacterIDShipParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCharactersCharacterIDShipParamsWithContext(ctx context.Context) *GetCharactersCharacterIDShipParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetCharactersCharacterIDShipParams{
		Datasource: &datasourceDefault,

		Context: ctx,
	}
}

// NewGetCharactersCharacterIDShipParamsWithHTTPClient creates a new GetCharactersCharacterIDShipParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCharactersCharacterIDShipParamsWithHTTPClient(client *http.Client) *GetCharactersCharacterIDShipParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetCharactersCharacterIDShipParams{
		Datasource: &datasourceDefault,
		HTTPClient: client,
	}
}

/*GetCharactersCharacterIDShipParams contains all the parameters to send to the API endpoint
for the get characters character id ship operation typically these are written to a http.Request
*/
type GetCharactersCharacterIDShipParams struct {

	/*XUserAgent
	  Client identifier, takes precedence over User-Agent

	*/
	XUserAgent *string
	/*CharacterID
	  An EVE character ID

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

// WithTimeout adds the timeout to the get characters character id ship params
func (o *GetCharactersCharacterIDShipParams) WithTimeout(timeout time.Duration) *GetCharactersCharacterIDShipParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get characters character id ship params
func (o *GetCharactersCharacterIDShipParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get characters character id ship params
func (o *GetCharactersCharacterIDShipParams) WithContext(ctx context.Context) *GetCharactersCharacterIDShipParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get characters character id ship params
func (o *GetCharactersCharacterIDShipParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get characters character id ship params
func (o *GetCharactersCharacterIDShipParams) WithHTTPClient(client *http.Client) *GetCharactersCharacterIDShipParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get characters character id ship params
func (o *GetCharactersCharacterIDShipParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXUserAgent adds the xUserAgent to the get characters character id ship params
func (o *GetCharactersCharacterIDShipParams) WithXUserAgent(xUserAgent *string) *GetCharactersCharacterIDShipParams {
	o.SetXUserAgent(xUserAgent)
	return o
}

// SetXUserAgent adds the xUserAgent to the get characters character id ship params
func (o *GetCharactersCharacterIDShipParams) SetXUserAgent(xUserAgent *string) {
	o.XUserAgent = xUserAgent
}

// WithCharacterID adds the characterID to the get characters character id ship params
func (o *GetCharactersCharacterIDShipParams) WithCharacterID(characterID int32) *GetCharactersCharacterIDShipParams {
	o.SetCharacterID(characterID)
	return o
}

// SetCharacterID adds the characterId to the get characters character id ship params
func (o *GetCharactersCharacterIDShipParams) SetCharacterID(characterID int32) {
	o.CharacterID = characterID
}

// WithDatasource adds the datasource to the get characters character id ship params
func (o *GetCharactersCharacterIDShipParams) WithDatasource(datasource *string) *GetCharactersCharacterIDShipParams {
	o.SetDatasource(datasource)
	return o
}

// SetDatasource adds the datasource to the get characters character id ship params
func (o *GetCharactersCharacterIDShipParams) SetDatasource(datasource *string) {
	o.Datasource = datasource
}

// WithToken adds the token to the get characters character id ship params
func (o *GetCharactersCharacterIDShipParams) WithToken(token *string) *GetCharactersCharacterIDShipParams {
	o.SetToken(token)
	return o
}

// SetToken adds the token to the get characters character id ship params
func (o *GetCharactersCharacterIDShipParams) SetToken(token *string) {
	o.Token = token
}

// WithUserAgent adds the userAgent to the get characters character id ship params
func (o *GetCharactersCharacterIDShipParams) WithUserAgent(userAgent *string) *GetCharactersCharacterIDShipParams {
	o.SetUserAgent(userAgent)
	return o
}

// SetUserAgent adds the userAgent to the get characters character id ship params
func (o *GetCharactersCharacterIDShipParams) SetUserAgent(userAgent *string) {
	o.UserAgent = userAgent
}

// WriteToRequest writes these params to a swagger request
func (o *GetCharactersCharacterIDShipParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
