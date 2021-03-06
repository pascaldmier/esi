package search

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new search API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for search API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetCharactersCharacterIDSearch searches on a string

Search for entities that match a given sub-string.

---

Alternate route: `/v2/characters/{character_id}/search/`


---

This route is cached for up to 3600 seconds
*/
func (a *Client) GetCharactersCharacterIDSearch(params *GetCharactersCharacterIDSearchParams, authInfo runtime.ClientAuthInfoWriter) (*GetCharactersCharacterIDSearchOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCharactersCharacterIDSearchParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "get_characters_character_id_search",
		Method:             "GET",
		PathPattern:        "/characters/{character_id}/search/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCharactersCharacterIDSearchReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetCharactersCharacterIDSearchOK), nil

}

/*
GetSearch searches on a string

Search for entities that match a given sub-string.

---

Alternate route: `/v1/search/`

Alternate route: `/legacy/search/`


---

This route is cached for up to 3600 seconds
*/
func (a *Client) GetSearch(params *GetSearchParams) (*GetSearchOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSearchParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "get_search",
		Method:             "GET",
		PathPattern:        "/search/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetSearchReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetSearchOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
