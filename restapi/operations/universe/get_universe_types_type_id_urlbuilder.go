package universe

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"strings"

	"github.com/go-openapi/swag"
)

// GetUniverseTypesTypeIDURL generates an URL for the get universe types type id operation
type GetUniverseTypesTypeIDURL struct {
	TypeID int32

	Datasource *string
	Language   *string
	UserAgent  *string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetUniverseTypesTypeIDURL) WithBasePath(bp string) *GetUniverseTypesTypeIDURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetUniverseTypesTypeIDURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *GetUniverseTypesTypeIDURL) Build() (*url.URL, error) {
	var result url.URL

	var _path = "/universe/types/{type_id}/"

	typeID := swag.FormatInt32(o.TypeID)
	if typeID != "" {
		_path = strings.Replace(_path, "{type_id}", typeID, -1)
	} else {
		return nil, errors.New("TypeID is required on GetUniverseTypesTypeIDURL")
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

	var language string
	if o.Language != nil {
		language = *o.Language
	}
	if language != "" {
		qs.Set("language", language)
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
func (o *GetUniverseTypesTypeIDURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *GetUniverseTypesTypeIDURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *GetUniverseTypesTypeIDURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on GetUniverseTypesTypeIDURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on GetUniverseTypesTypeIDURL")
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
func (o *GetUniverseTypesTypeIDURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}