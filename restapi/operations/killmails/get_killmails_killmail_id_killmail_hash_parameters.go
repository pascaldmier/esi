package killmails

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetKillmailsKillmailIDKillmailHashParams creates a new GetKillmailsKillmailIDKillmailHashParams object
// with the default values initialized.
func NewGetKillmailsKillmailIDKillmailHashParams() GetKillmailsKillmailIDKillmailHashParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return GetKillmailsKillmailIDKillmailHashParams{
		Datasource: &datasourceDefault,
	}
}

// GetKillmailsKillmailIDKillmailHashParams contains all the bound params for the get killmails killmail id killmail hash operation
// typically these are obtained from a http.Request
//
// swagger:parameters get_killmails_killmail_id_killmail_hash
type GetKillmailsKillmailIDKillmailHashParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request

	/*Client identifier, takes precedence over User-Agent
	  In: header
	*/
	XUserAgent *string
	/*The server name you would like data from
	  In: query
	  Default: "tranquility"
	*/
	Datasource *string
	/*The killmail hash for verification
	  Required: true
	  In: path
	*/
	KillmailHash string
	/*The killmail ID to be queried
	  Required: true
	  In: path
	*/
	KillmailID int32
	/*Client identifier, takes precedence over headers
	  In: query
	*/
	UserAgent *string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *GetKillmailsKillmailIDKillmailHashParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	if err := o.bindXUserAgent(r.Header[http.CanonicalHeaderKey("X-User-Agent")], true, route.Formats); err != nil {
		res = append(res, err)
	}

	qDatasource, qhkDatasource, _ := qs.GetOK("datasource")
	if err := o.bindDatasource(qDatasource, qhkDatasource, route.Formats); err != nil {
		res = append(res, err)
	}

	rKillmailHash, rhkKillmailHash, _ := route.Params.GetOK("killmail_hash")
	if err := o.bindKillmailHash(rKillmailHash, rhkKillmailHash, route.Formats); err != nil {
		res = append(res, err)
	}

	rKillmailID, rhkKillmailID, _ := route.Params.GetOK("killmail_id")
	if err := o.bindKillmailID(rKillmailID, rhkKillmailID, route.Formats); err != nil {
		res = append(res, err)
	}

	qUserAgent, qhkUserAgent, _ := qs.GetOK("user_agent")
	if err := o.bindUserAgent(qUserAgent, qhkUserAgent, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetKillmailsKillmailIDKillmailHashParams) bindXUserAgent(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.XUserAgent = &raw

	return nil
}

func (o *GetKillmailsKillmailIDKillmailHashParams) bindDatasource(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		var datasourceDefault string = string("tranquility")
		o.Datasource = &datasourceDefault
		return nil
	}

	o.Datasource = &raw

	if err := o.validateDatasource(formats); err != nil {
		return err
	}

	return nil
}

func (o *GetKillmailsKillmailIDKillmailHashParams) validateDatasource(formats strfmt.Registry) error {

	if err := validate.Enum("datasource", "query", *o.Datasource, []interface{}{"tranquility", "singularity"}); err != nil {
		return err
	}

	return nil
}

func (o *GetKillmailsKillmailIDKillmailHashParams) bindKillmailHash(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.KillmailHash = raw

	return nil
}

func (o *GetKillmailsKillmailIDKillmailHashParams) bindKillmailID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	value, err := swag.ConvertInt32(raw)
	if err != nil {
		return errors.InvalidType("killmail_id", "path", "int32", raw)
	}
	o.KillmailID = value

	return nil
}

func (o *GetKillmailsKillmailIDKillmailHashParams) bindUserAgent(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.UserAgent = &raw

	return nil
}
