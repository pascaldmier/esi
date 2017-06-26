package market

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"strings"

	"github.com/go-openapi/swag"
)

// GetMarketsStructuresStructureIDURL generates an URL for the get markets structures structure id operation
type GetMarketsStructuresStructureIDURL struct {
	StructureID int64

	Datasource *string
	Page       *int32
	Token      *string
	UserAgent  *string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetMarketsStructuresStructureIDURL) WithBasePath(bp string) *GetMarketsStructuresStructureIDURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetMarketsStructuresStructureIDURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *GetMarketsStructuresStructureIDURL) Build() (*url.URL, error) {
	var result url.URL

	var _path = "/markets/structures/{structure_id}/"

	structureID := swag.FormatInt64(o.StructureID)
	if structureID != "" {
		_path = strings.Replace(_path, "{structure_id}", structureID, -1)
	} else {
		return nil, errors.New("StructureID is required on GetMarketsStructuresStructureIDURL")
	}
	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/latest"
	}
	result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	var datasource string
	if o.Datasource != nil {
		datasource = *o.Datasource
	}
	if datasource != "" {
		qs.Set("datasource", datasource)
	}

	var page string
	if o.Page != nil {
		page = swag.FormatInt32(*o.Page)
	}
	if page != "" {
		qs.Set("page", page)
	}

	var token string
	if o.Token != nil {
		token = *o.Token
	}
	if token != "" {
		qs.Set("token", token)
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
func (o *GetMarketsStructuresStructureIDURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *GetMarketsStructuresStructureIDURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *GetMarketsStructuresStructureIDURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on GetMarketsStructuresStructureIDURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on GetMarketsStructuresStructureIDURL")
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
func (o *GetMarketsStructuresStructureIDURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
