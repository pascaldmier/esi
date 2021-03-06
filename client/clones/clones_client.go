package clones

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new clones API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for clones API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetCharactersCharacterIDClones gets clones

A list of the character's clones

---

Alternate route: `/v2/characters/{character_id}/clones/`

Alternate route: `/dev/characters/{character_id}/clones/`


---

This route is cached for up to 120 seconds
*/
func (a *Client) GetCharactersCharacterIDClones(params *GetCharactersCharacterIDClonesParams, authInfo runtime.ClientAuthInfoWriter) (*GetCharactersCharacterIDClonesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCharactersCharacterIDClonesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "get_characters_character_id_clones",
		Method:             "GET",
		PathPattern:        "/characters/{character_id}/clones/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCharactersCharacterIDClonesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetCharactersCharacterIDClonesOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
