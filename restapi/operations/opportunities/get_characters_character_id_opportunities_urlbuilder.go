package opportunities

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"strings"

	"github.com/go-openapi/swag"
)

// GetCharactersCharacterIDOpportunitiesURL generates an URL for the get characters character id opportunities operation
type GetCharactersCharacterIDOpportunitiesURL struct {
	CharacterID int32

	Datasource *string
	Token      *string
	UserAgent  *string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetCharactersCharacterIDOpportunitiesURL) WithBasePath(bp string) *GetCharactersCharacterIDOpportunitiesURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetCharactersCharacterIDOpportunitiesURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *GetCharactersCharacterIDOpportunitiesURL) Build() (*url.URL, error) {
	var result url.URL

	var _path = "/characters/{character_id}/opportunities/"

	characterID := swag.FormatInt32(o.CharacterID)
	if characterID != "" {
		_path = strings.Replace(_path, "{character_id}", characterID, -1)
	} else {
		return nil, errors.New("CharacterID is required on GetCharactersCharacterIDOpportunitiesURL")
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
func (o *GetCharactersCharacterIDOpportunitiesURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *GetCharactersCharacterIDOpportunitiesURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *GetCharactersCharacterIDOpportunitiesURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on GetCharactersCharacterIDOpportunitiesURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on GetCharactersCharacterIDOpportunitiesURL")
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
func (o *GetCharactersCharacterIDOpportunitiesURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
