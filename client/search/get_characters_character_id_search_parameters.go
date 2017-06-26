package search

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

// NewGetCharactersCharacterIDSearchParams creates a new GetCharactersCharacterIDSearchParams object
// with the default values initialized.
func NewGetCharactersCharacterIDSearchParams() *GetCharactersCharacterIDSearchParams {
	var (
		datasourceDefault = string("tranquility")
		languageDefault   = string("en-us")
		strictDefault     = bool(false)
	)
	return &GetCharactersCharacterIDSearchParams{
		Datasource: &datasourceDefault,
		Language:   &languageDefault,
		Strict:     &strictDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCharactersCharacterIDSearchParamsWithTimeout creates a new GetCharactersCharacterIDSearchParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCharactersCharacterIDSearchParamsWithTimeout(timeout time.Duration) *GetCharactersCharacterIDSearchParams {
	var (
		datasourceDefault = string("tranquility")
		languageDefault   = string("en-us")
		strictDefault     = bool(false)
	)
	return &GetCharactersCharacterIDSearchParams{
		Datasource: &datasourceDefault,
		Language:   &languageDefault,
		Strict:     &strictDefault,

		timeout: timeout,
	}
}

// NewGetCharactersCharacterIDSearchParamsWithContext creates a new GetCharactersCharacterIDSearchParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCharactersCharacterIDSearchParamsWithContext(ctx context.Context) *GetCharactersCharacterIDSearchParams {
	var (
		datasourceDefault = string("tranquility")
		languageDefault   = string("en-us")
		strictDefault     = bool(false)
	)
	return &GetCharactersCharacterIDSearchParams{
		Datasource: &datasourceDefault,
		Language:   &languageDefault,
		Strict:     &strictDefault,

		Context: ctx,
	}
}

// NewGetCharactersCharacterIDSearchParamsWithHTTPClient creates a new GetCharactersCharacterIDSearchParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCharactersCharacterIDSearchParamsWithHTTPClient(client *http.Client) *GetCharactersCharacterIDSearchParams {
	var (
		datasourceDefault = string("tranquility")
		languageDefault   = string("en-us")
		strictDefault     = bool(false)
	)
	return &GetCharactersCharacterIDSearchParams{
		Datasource: &datasourceDefault,
		Language:   &languageDefault,
		Strict:     &strictDefault,
		HTTPClient: client,
	}
}

/*GetCharactersCharacterIDSearchParams contains all the parameters to send to the API endpoint
for the get characters character id search operation typically these are written to a http.Request
*/
type GetCharactersCharacterIDSearchParams struct {

	/*XUserAgent
	  Client identifier, takes precedence over User-Agent

	*/
	XUserAgent *string
	/*Categories
	  Type of entities to search for

	*/
	Categories []string
	/*CharacterID
	  An EVE character ID

	*/
	CharacterID int32
	/*Datasource
	  The server name you would like data from

	*/
	Datasource *string
	/*Language
	  Search locale

	*/
	Language *string
	/*Search
	  The string to search on

	*/
	Search string
	/*Strict
	  Whether the search should be a strict match

	*/
	Strict *bool
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

// WithTimeout adds the timeout to the get characters character id search params
func (o *GetCharactersCharacterIDSearchParams) WithTimeout(timeout time.Duration) *GetCharactersCharacterIDSearchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get characters character id search params
func (o *GetCharactersCharacterIDSearchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get characters character id search params
func (o *GetCharactersCharacterIDSearchParams) WithContext(ctx context.Context) *GetCharactersCharacterIDSearchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get characters character id search params
func (o *GetCharactersCharacterIDSearchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get characters character id search params
func (o *GetCharactersCharacterIDSearchParams) WithHTTPClient(client *http.Client) *GetCharactersCharacterIDSearchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get characters character id search params
func (o *GetCharactersCharacterIDSearchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXUserAgent adds the xUserAgent to the get characters character id search params
func (o *GetCharactersCharacterIDSearchParams) WithXUserAgent(xUserAgent *string) *GetCharactersCharacterIDSearchParams {
	o.SetXUserAgent(xUserAgent)
	return o
}

// SetXUserAgent adds the xUserAgent to the get characters character id search params
func (o *GetCharactersCharacterIDSearchParams) SetXUserAgent(xUserAgent *string) {
	o.XUserAgent = xUserAgent
}

// WithCategories adds the categories to the get characters character id search params
func (o *GetCharactersCharacterIDSearchParams) WithCategories(categories []string) *GetCharactersCharacterIDSearchParams {
	o.SetCategories(categories)
	return o
}

// SetCategories adds the categories to the get characters character id search params
func (o *GetCharactersCharacterIDSearchParams) SetCategories(categories []string) {
	o.Categories = categories
}

// WithCharacterID adds the characterID to the get characters character id search params
func (o *GetCharactersCharacterIDSearchParams) WithCharacterID(characterID int32) *GetCharactersCharacterIDSearchParams {
	o.SetCharacterID(characterID)
	return o
}

// SetCharacterID adds the characterId to the get characters character id search params
func (o *GetCharactersCharacterIDSearchParams) SetCharacterID(characterID int32) {
	o.CharacterID = characterID
}

// WithDatasource adds the datasource to the get characters character id search params
func (o *GetCharactersCharacterIDSearchParams) WithDatasource(datasource *string) *GetCharactersCharacterIDSearchParams {
	o.SetDatasource(datasource)
	return o
}

// SetDatasource adds the datasource to the get characters character id search params
func (o *GetCharactersCharacterIDSearchParams) SetDatasource(datasource *string) {
	o.Datasource = datasource
}

// WithLanguage adds the language to the get characters character id search params
func (o *GetCharactersCharacterIDSearchParams) WithLanguage(language *string) *GetCharactersCharacterIDSearchParams {
	o.SetLanguage(language)
	return o
}

// SetLanguage adds the language to the get characters character id search params
func (o *GetCharactersCharacterIDSearchParams) SetLanguage(language *string) {
	o.Language = language
}

// WithSearch adds the search to the get characters character id search params
func (o *GetCharactersCharacterIDSearchParams) WithSearch(search string) *GetCharactersCharacterIDSearchParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the get characters character id search params
func (o *GetCharactersCharacterIDSearchParams) SetSearch(search string) {
	o.Search = search
}

// WithStrict adds the strict to the get characters character id search params
func (o *GetCharactersCharacterIDSearchParams) WithStrict(strict *bool) *GetCharactersCharacterIDSearchParams {
	o.SetStrict(strict)
	return o
}

// SetStrict adds the strict to the get characters character id search params
func (o *GetCharactersCharacterIDSearchParams) SetStrict(strict *bool) {
	o.Strict = strict
}

// WithToken adds the token to the get characters character id search params
func (o *GetCharactersCharacterIDSearchParams) WithToken(token *string) *GetCharactersCharacterIDSearchParams {
	o.SetToken(token)
	return o
}

// SetToken adds the token to the get characters character id search params
func (o *GetCharactersCharacterIDSearchParams) SetToken(token *string) {
	o.Token = token
}

// WithUserAgent adds the userAgent to the get characters character id search params
func (o *GetCharactersCharacterIDSearchParams) WithUserAgent(userAgent *string) *GetCharactersCharacterIDSearchParams {
	o.SetUserAgent(userAgent)
	return o
}

// SetUserAgent adds the userAgent to the get characters character id search params
func (o *GetCharactersCharacterIDSearchParams) SetUserAgent(userAgent *string) {
	o.UserAgent = userAgent
}

// WriteToRequest writes these params to a swagger request
func (o *GetCharactersCharacterIDSearchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	valuesCategories := o.Categories

	joinedCategories := swag.JoinByFormat(valuesCategories, "")
	// query array param categories
	if err := r.SetQueryParam("categories", joinedCategories...); err != nil {
		return err
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

	if o.Language != nil {

		// query param language
		var qrLanguage string
		if o.Language != nil {
			qrLanguage = *o.Language
		}
		qLanguage := qrLanguage
		if qLanguage != "" {
			if err := r.SetQueryParam("language", qLanguage); err != nil {
				return err
			}
		}

	}

	// query param search
	qrSearch := o.Search
	qSearch := qrSearch
	if qSearch != "" {
		if err := r.SetQueryParam("search", qSearch); err != nil {
			return err
		}
	}

	if o.Strict != nil {

		// query param strict
		var qrStrict bool
		if o.Strict != nil {
			qrStrict = *o.Strict
		}
		qStrict := swag.FormatBool(qrStrict)
		if qStrict != "" {
			if err := r.SetQueryParam("strict", qStrict); err != nil {
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