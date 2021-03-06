package planetary_interaction

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"strings"

	"github.com/go-openapi/swag"
)

// GetUniverseSchematicsSchematicIDURL generates an URL for the get universe schematics schematic id operation
type GetUniverseSchematicsSchematicIDURL struct {
	SchematicID int32

	Datasource *string
	UserAgent  *string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetUniverseSchematicsSchematicIDURL) WithBasePath(bp string) *GetUniverseSchematicsSchematicIDURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetUniverseSchematicsSchematicIDURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *GetUniverseSchematicsSchematicIDURL) Build() (*url.URL, error) {
	var result url.URL

	var _path = "/universe/schematics/{schematic_id}/"

	schematicID := swag.FormatInt32(o.SchematicID)
	if schematicID != "" {
		_path = strings.Replace(_path, "{schematic_id}", schematicID, -1)
	} else {
		return nil, errors.New("SchematicID is required on GetUniverseSchematicsSchematicIDURL")
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
func (o *GetUniverseSchematicsSchematicIDURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *GetUniverseSchematicsSchematicIDURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *GetUniverseSchematicsSchematicIDURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on GetUniverseSchematicsSchematicIDURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on GetUniverseSchematicsSchematicIDURL")
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
func (o *GetUniverseSchematicsSchematicIDURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
