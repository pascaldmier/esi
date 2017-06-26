package dogma

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new dogma API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for dogma API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetDogmaAttributes gets attributes

Get a list of dogma attribute ids

---

Alternate route: `/v1/dogma/attributes/`

Alternate route: `/legacy/dogma/attributes/`

Alternate route: `/dev/dogma/attributes/`


---

This route expires daily at 11:05
*/
func (a *Client) GetDogmaAttributes(params *GetDogmaAttributesParams) (*GetDogmaAttributesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDogmaAttributesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "get_dogma_attributes",
		Method:             "GET",
		PathPattern:        "/dogma/attributes/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDogmaAttributesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetDogmaAttributesOK), nil

}

/*
GetDogmaAttributesAttributeID gets attribute information

Get information on a dogma attribute

---

Alternate route: `/v1/dogma/attributes/{attribute_id}/`

Alternate route: `/legacy/dogma/attributes/{attribute_id}/`

Alternate route: `/dev/dogma/attributes/{attribute_id}/`


---

This route expires daily at 11:05
*/
func (a *Client) GetDogmaAttributesAttributeID(params *GetDogmaAttributesAttributeIDParams) (*GetDogmaAttributesAttributeIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDogmaAttributesAttributeIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "get_dogma_attributes_attribute_id",
		Method:             "GET",
		PathPattern:        "/dogma/attributes/{attribute_id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDogmaAttributesAttributeIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetDogmaAttributesAttributeIDOK), nil

}

/*
GetDogmaEffects gets effects

Get a list of dogma effect ids

---

Alternate route: `/v1/dogma/effects/`

Alternate route: `/legacy/dogma/effects/`

Alternate route: `/dev/dogma/effects/`


---

This route expires daily at 11:05
*/
func (a *Client) GetDogmaEffects(params *GetDogmaEffectsParams) (*GetDogmaEffectsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDogmaEffectsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "get_dogma_effects",
		Method:             "GET",
		PathPattern:        "/dogma/effects/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDogmaEffectsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetDogmaEffectsOK), nil

}

/*
GetDogmaEffectsEffectID gets effect information

Get information on a dogma effect

---

Alternate route: `/v1/dogma/effects/{effect_id}/`

Alternate route: `/legacy/dogma/effects/{effect_id}/`


---

This route expires daily at 11:05
*/
func (a *Client) GetDogmaEffectsEffectID(params *GetDogmaEffectsEffectIDParams) (*GetDogmaEffectsEffectIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDogmaEffectsEffectIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "get_dogma_effects_effect_id",
		Method:             "GET",
		PathPattern:        "/dogma/effects/{effect_id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDogmaEffectsEffectIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetDogmaEffectsEffectIDOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
