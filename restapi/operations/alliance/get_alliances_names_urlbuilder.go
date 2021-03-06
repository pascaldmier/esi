package alliance

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"

	"github.com/go-openapi/swag"
)

// GetAlliancesNamesURL generates an URL for the get alliances names operation
type GetAlliancesNamesURL struct {
	AllianceIds []int64
	Datasource  *string
	UserAgent   *string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetAlliancesNamesURL) WithBasePath(bp string) *GetAlliancesNamesURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetAlliancesNamesURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *GetAlliancesNamesURL) Build() (*url.URL, error) {
	var result url.URL

	var _path = "/alliances/names/"

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/latest"
	}
	result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	var allianceIdsIR []string
	for _, allianceIdsI := range o.AllianceIds {
		allianceIdsIS := swag.FormatInt64(allianceIdsI)
		if allianceIdsIS != "" {
			allianceIdsIR = append(allianceIdsIR, allianceIdsIS)
		}
	}

	allianceIds := swag.JoinByFormat(allianceIdsIR, "")

	if len(allianceIds) > 0 {
		qsv := allianceIds[0]
		if qsv != "" {
			qs.Set("alliance_ids", qsv)
		}
	}

	var datasource string
	if o.Datasource != nil {
		datasource = *o.Datasource
	}
	if datasource != "" {
		qs.Set("datasource", datasource)
	}

	var userAgent string
	if o.UserAgent != nil {
		userAgent = *o.UserAgent
	}
	if userAgent != "" {
		qs.Set("user_agent", userAgent)
	}

	result.RawQuery = qs.Encode()

	return &result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *GetAlliancesNamesURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *GetAlliancesNamesURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *GetAlliancesNamesURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on GetAlliancesNamesURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on GetAlliancesNamesURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *GetAlliancesNamesURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
