package skills

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

// NewGetCharactersCharacterIDSkillsParams creates a new GetCharactersCharacterIDSkillsParams object
// with the default values initialized.
func NewGetCharactersCharacterIDSkillsParams() *GetCharactersCharacterIDSkillsParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetCharactersCharacterIDSkillsParams{
		Datasource: &datasourceDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCharactersCharacterIDSkillsParamsWithTimeout creates a new GetCharactersCharacterIDSkillsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCharactersCharacterIDSkillsParamsWithTimeout(timeout time.Duration) *GetCharactersCharacterIDSkillsParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetCharactersCharacterIDSkillsParams{
		Datasource: &datasourceDefault,

		timeout: timeout,
	}
}

// NewGetCharactersCharacterIDSkillsParamsWithContext creates a new GetCharactersCharacterIDSkillsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCharactersCharacterIDSkillsParamsWithContext(ctx context.Context) *GetCharactersCharacterIDSkillsParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetCharactersCharacterIDSkillsParams{
		Datasource: &datasourceDefault,

		Context: ctx,
	}
}

// NewGetCharactersCharacterIDSkillsParamsWithHTTPClient creates a new GetCharactersCharacterIDSkillsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCharactersCharacterIDSkillsParamsWithHTTPClient(client *http.Client) *GetCharactersCharacterIDSkillsParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetCharactersCharacterIDSkillsParams{
		Datasource: &datasourceDefault,
		HTTPClient: client,
	}
}

/*GetCharactersCharacterIDSkillsParams contains all the parameters to send to the API endpoint
for the get characters character id skills operation typically these are written to a http.Request
*/
type GetCharactersCharacterIDSkillsParams struct {

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

// WithTimeout adds the timeout to the get characters character id skills params
func (o *GetCharactersCharacterIDSkillsParams) WithTimeout(timeout time.Duration) *GetCharactersCharacterIDSkillsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get characters character id skills params
func (o *GetCharactersCharacterIDSkillsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get characters character id skills params
func (o *GetCharactersCharacterIDSkillsParams) WithContext(ctx context.Context) *GetCharactersCharacterIDSkillsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get characters character id skills params
func (o *GetCharactersCharacterIDSkillsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get characters character id skills params
func (o *GetCharactersCharacterIDSkillsParams) WithHTTPClient(client *http.Client) *GetCharactersCharacterIDSkillsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get characters character id skills params
func (o *GetCharactersCharacterIDSkillsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXUserAgent adds the xUserAgent to the get characters character id skills params
func (o *GetCharactersCharacterIDSkillsParams) WithXUserAgent(xUserAgent *string) *GetCharactersCharacterIDSkillsParams {
	o.SetXUserAgent(xUserAgent)
	return o
}

// SetXUserAgent adds the xUserAgent to the get characters character id skills params
func (o *GetCharactersCharacterIDSkillsParams) SetXUserAgent(xUserAgent *string) {
	o.XUserAgent = xUserAgent
}

// WithCharacterID adds the characterID to the get characters character id skills params
func (o *GetCharactersCharacterIDSkillsParams) WithCharacterID(characterID int32) *GetCharactersCharacterIDSkillsParams {
	o.SetCharacterID(characterID)
	return o
}

// SetCharacterID adds the characterId to the get characters character id skills params
func (o *GetCharactersCharacterIDSkillsParams) SetCharacterID(characterID int32) {
	o.CharacterID = characterID
}

// WithDatasource adds the datasource to the get characters character id skills params
func (o *GetCharactersCharacterIDSkillsParams) WithDatasource(datasource *string) *GetCharactersCharacterIDSkillsParams {
	o.SetDatasource(datasource)
	return o
}

// SetDatasource adds the datasource to the get characters character id skills params
func (o *GetCharactersCharacterIDSkillsParams) SetDatasource(datasource *string) {
	o.Datasource = datasource
}

// WithToken adds the token to the get characters character id skills params
func (o *GetCharactersCharacterIDSkillsParams) WithToken(token *string) *GetCharactersCharacterIDSkillsParams {
	o.SetToken(token)
	return o
}

// SetToken adds the token to the get characters character id skills params
func (o *GetCharactersCharacterIDSkillsParams) SetToken(token *string) {
	o.Token = token
}

// WithUserAgent adds the userAgent to the get characters character id skills params
func (o *GetCharactersCharacterIDSkillsParams) WithUserAgent(userAgent *string) *GetCharactersCharacterIDSkillsParams {
	o.SetUserAgent(userAgent)
	return o
}

// SetUserAgent adds the userAgent to the get characters character id skills params
func (o *GetCharactersCharacterIDSkillsParams) SetUserAgent(userAgent *string) {
	o.UserAgent = userAgent
}

// WriteToRequest writes these params to a swagger request
func (o *GetCharactersCharacterIDSkillsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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