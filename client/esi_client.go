package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/pascaldmier/esi/client/alliance"
	"github.com/pascaldmier/esi/client/assets"
	"github.com/pascaldmier/esi/client/bookmarks"
	"github.com/pascaldmier/esi/client/calendar"
	"github.com/pascaldmier/esi/client/character"
	"github.com/pascaldmier/esi/client/clones"
	"github.com/pascaldmier/esi/client/contacts"
	"github.com/pascaldmier/esi/client/corporation"
	"github.com/pascaldmier/esi/client/dogma"
	"github.com/pascaldmier/esi/client/fittings"
	"github.com/pascaldmier/esi/client/fleets"
	"github.com/pascaldmier/esi/client/incursions"
	"github.com/pascaldmier/esi/client/industry"
	"github.com/pascaldmier/esi/client/insurance"
	"github.com/pascaldmier/esi/client/killmails"
	"github.com/pascaldmier/esi/client/location"
	"github.com/pascaldmier/esi/client/loyalty"
	"github.com/pascaldmier/esi/client/mail"
	"github.com/pascaldmier/esi/client/market"
	"github.com/pascaldmier/esi/client/opportunities"
	"github.com/pascaldmier/esi/client/planetary_interaction"
	"github.com/pascaldmier/esi/client/search"
	"github.com/pascaldmier/esi/client/skills"
	"github.com/pascaldmier/esi/client/sovereignty"
	"github.com/pascaldmier/esi/client/status"
	"github.com/pascaldmier/esi/client/universe"
	"github.com/pascaldmier/esi/client/user_interface"
	"github.com/pascaldmier/esi/client/wallet"
	"github.com/pascaldmier/esi/client/wars"
)

// Default esi HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "esi.tech.ccp.is"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/latest"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"https"}

// NewHTTPClient creates a new esi HTTP client.
func NewHTTPClient(formats strfmt.Registry) *Esi {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new esi HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *Esi {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new esi client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Esi {
	cli := new(Esi)
	cli.Transport = transport

	cli.Alliance = alliance.New(transport, formats)

	cli.Assets = assets.New(transport, formats)

	cli.Bookmarks = bookmarks.New(transport, formats)

	cli.Calendar = calendar.New(transport, formats)

	cli.Character = character.New(transport, formats)

	cli.Clones = clones.New(transport, formats)

	cli.Contacts = contacts.New(transport, formats)

	cli.Corporation = corporation.New(transport, formats)

	cli.Dogma = dogma.New(transport, formats)

	cli.Fittings = fittings.New(transport, formats)

	cli.Fleets = fleets.New(transport, formats)

	cli.Incursions = incursions.New(transport, formats)

	cli.Industry = industry.New(transport, formats)

	cli.Insurance = insurance.New(transport, formats)

	cli.Killmails = killmails.New(transport, formats)

	cli.Location = location.New(transport, formats)

	cli.Loyalty = loyalty.New(transport, formats)

	cli.Mail = mail.New(transport, formats)

	cli.Market = market.New(transport, formats)

	cli.Opportunities = opportunities.New(transport, formats)

	cli.PlanetaryInteraction = planetary_interaction.New(transport, formats)

	cli.Search = search.New(transport, formats)

	cli.Skills = skills.New(transport, formats)

	cli.Sovereignty = sovereignty.New(transport, formats)

	cli.Status = status.New(transport, formats)

	cli.Universe = universe.New(transport, formats)

	cli.UserInterface = user_interface.New(transport, formats)

	cli.Wallet = wallet.New(transport, formats)

	cli.Wars = wars.New(transport, formats)

	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// Esi is a client for esi
type Esi struct {
	Alliance *alliance.Client

	Assets *assets.Client

	Bookmarks *bookmarks.Client

	Calendar *calendar.Client

	Character *character.Client

	Clones *clones.Client

	Contacts *contacts.Client

	Corporation *corporation.Client

	Dogma *dogma.Client

	Fittings *fittings.Client

	Fleets *fleets.Client

	Incursions *incursions.Client

	Industry *industry.Client

	Insurance *insurance.Client

	Killmails *killmails.Client

	Location *location.Client

	Loyalty *loyalty.Client

	Mail *mail.Client

	Market *market.Client

	Opportunities *opportunities.Client

	PlanetaryInteraction *planetary_interaction.Client

	Search *search.Client

	Skills *skills.Client

	Sovereignty *sovereignty.Client

	Status *status.Client

	Universe *universe.Client

	UserInterface *user_interface.Client

	Wallet *wallet.Client

	Wars *wars.Client

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *Esi) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport

	c.Alliance.SetTransport(transport)

	c.Assets.SetTransport(transport)

	c.Bookmarks.SetTransport(transport)

	c.Calendar.SetTransport(transport)

	c.Character.SetTransport(transport)

	c.Clones.SetTransport(transport)

	c.Contacts.SetTransport(transport)

	c.Corporation.SetTransport(transport)

	c.Dogma.SetTransport(transport)

	c.Fittings.SetTransport(transport)

	c.Fleets.SetTransport(transport)

	c.Incursions.SetTransport(transport)

	c.Industry.SetTransport(transport)

	c.Insurance.SetTransport(transport)

	c.Killmails.SetTransport(transport)

	c.Location.SetTransport(transport)

	c.Loyalty.SetTransport(transport)

	c.Mail.SetTransport(transport)

	c.Market.SetTransport(transport)

	c.Opportunities.SetTransport(transport)

	c.PlanetaryInteraction.SetTransport(transport)

	c.Search.SetTransport(transport)

	c.Skills.SetTransport(transport)

	c.Sovereignty.SetTransport(transport)

	c.Status.SetTransport(transport)

	c.Universe.SetTransport(transport)

	c.UserInterface.SetTransport(transport)

	c.Wallet.SetTransport(transport)

	c.Wars.SetTransport(transport)

}
