package restapi

import (
	"crypto/tls"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	graceful "github.com/tylerb/graceful"

	"github.com/pascaldmier/esi/restapi/operations"
	"github.com/pascaldmier/esi/restapi/operations/alliance"
	"github.com/pascaldmier/esi/restapi/operations/assets"
	"github.com/pascaldmier/esi/restapi/operations/bookmarks"
	"github.com/pascaldmier/esi/restapi/operations/calendar"
	"github.com/pascaldmier/esi/restapi/operations/character"
	"github.com/pascaldmier/esi/restapi/operations/clones"
	"github.com/pascaldmier/esi/restapi/operations/contacts"
	"github.com/pascaldmier/esi/restapi/operations/corporation"
	"github.com/pascaldmier/esi/restapi/operations/dogma"
	"github.com/pascaldmier/esi/restapi/operations/fittings"
	"github.com/pascaldmier/esi/restapi/operations/fleets"
	"github.com/pascaldmier/esi/restapi/operations/incursions"
	"github.com/pascaldmier/esi/restapi/operations/industry"
	"github.com/pascaldmier/esi/restapi/operations/insurance"
	"github.com/pascaldmier/esi/restapi/operations/killmails"
	"github.com/pascaldmier/esi/restapi/operations/location"
	"github.com/pascaldmier/esi/restapi/operations/loyalty"
	"github.com/pascaldmier/esi/restapi/operations/mail"
	"github.com/pascaldmier/esi/restapi/operations/market"
	"github.com/pascaldmier/esi/restapi/operations/opportunities"
	"github.com/pascaldmier/esi/restapi/operations/planetary_interaction"
	"github.com/pascaldmier/esi/restapi/operations/search"
	"github.com/pascaldmier/esi/restapi/operations/skills"
	"github.com/pascaldmier/esi/restapi/operations/sovereignty"
	"github.com/pascaldmier/esi/restapi/operations/status"
	"github.com/pascaldmier/esi/restapi/operations/universe"
	"github.com/pascaldmier/esi/restapi/operations/user_interface"
	"github.com/pascaldmier/esi/restapi/operations/wallet"
	"github.com/pascaldmier/esi/restapi/operations/wars"
)

// This file is safe to edit. Once it exists it will not be overwritten

//go:generate swagger generate server --target .. --name esi --spec ../swagger/swagger.flattened.json

func configureFlags(api *operations.EsiAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.EsiAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// s.api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.EvessoAuth = func(token string, scopes []string) (interface{}, error) {
		return nil, errors.NotImplemented("oauth2 bearer auth (evesso) has not yet been implemented")
	}

	api.ContactsDeleteCharactersCharacterIDContactsHandler = contacts.DeleteCharactersCharacterIDContactsHandlerFunc(func(params contacts.DeleteCharactersCharacterIDContactsParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation contacts.DeleteCharactersCharacterIDContacts has not yet been implemented")
	})
	api.FittingsDeleteCharactersCharacterIDFittingsFittingIDHandler = fittings.DeleteCharactersCharacterIDFittingsFittingIDHandlerFunc(func(params fittings.DeleteCharactersCharacterIDFittingsFittingIDParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation fittings.DeleteCharactersCharacterIDFittingsFittingID has not yet been implemented")
	})
	api.MailDeleteCharactersCharacterIDMailLabelsLabelIDHandler = mail.DeleteCharactersCharacterIDMailLabelsLabelIDHandlerFunc(func(params mail.DeleteCharactersCharacterIDMailLabelsLabelIDParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation mail.DeleteCharactersCharacterIDMailLabelsLabelID has not yet been implemented")
	})
	api.MailDeleteCharactersCharacterIDMailMailIDHandler = mail.DeleteCharactersCharacterIDMailMailIDHandlerFunc(func(params mail.DeleteCharactersCharacterIDMailMailIDParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation mail.DeleteCharactersCharacterIDMailMailID has not yet been implemented")
	})
	api.FleetsDeleteFleetsFleetIDMembersMemberIDHandler = fleets.DeleteFleetsFleetIDMembersMemberIDHandlerFunc(func(params fleets.DeleteFleetsFleetIDMembersMemberIDParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation fleets.DeleteFleetsFleetIDMembersMemberID has not yet been implemented")
	})
	api.FleetsDeleteFleetsFleetIDSquadsSquadIDHandler = fleets.DeleteFleetsFleetIDSquadsSquadIDHandlerFunc(func(params fleets.DeleteFleetsFleetIDSquadsSquadIDParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation fleets.DeleteFleetsFleetIDSquadsSquadID has not yet been implemented")
	})
	api.FleetsDeleteFleetsFleetIDWingsWingIDHandler = fleets.DeleteFleetsFleetIDWingsWingIDHandlerFunc(func(params fleets.DeleteFleetsFleetIDWingsWingIDParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation fleets.DeleteFleetsFleetIDWingsWingID has not yet been implemented")
	})
	api.AllianceGetAlliancesHandler = alliance.GetAlliancesHandlerFunc(func(params alliance.GetAlliancesParams) middleware.Responder {
		return middleware.NotImplemented("operation alliance.GetAlliances has not yet been implemented")
	})
	api.AllianceGetAlliancesAllianceIDHandler = alliance.GetAlliancesAllianceIDHandlerFunc(func(params alliance.GetAlliancesAllianceIDParams) middleware.Responder {
		return middleware.NotImplemented("operation alliance.GetAlliancesAllianceID has not yet been implemented")
	})
	api.AllianceGetAlliancesAllianceIDCorporationsHandler = alliance.GetAlliancesAllianceIDCorporationsHandlerFunc(func(params alliance.GetAlliancesAllianceIDCorporationsParams) middleware.Responder {
		return middleware.NotImplemented("operation alliance.GetAlliancesAllianceIDCorporations has not yet been implemented")
	})
	api.AllianceGetAlliancesAllianceIDIconsHandler = alliance.GetAlliancesAllianceIDIconsHandlerFunc(func(params alliance.GetAlliancesAllianceIDIconsParams) middleware.Responder {
		return middleware.NotImplemented("operation alliance.GetAlliancesAllianceIDIcons has not yet been implemented")
	})
	api.AllianceGetAlliancesNamesHandler = alliance.GetAlliancesNamesHandlerFunc(func(params alliance.GetAlliancesNamesParams) middleware.Responder {
		return middleware.NotImplemented("operation alliance.GetAlliancesNames has not yet been implemented")
	})
	api.CharacterGetCharactersCharacterIDHandler = character.GetCharactersCharacterIDHandlerFunc(func(params character.GetCharactersCharacterIDParams) middleware.Responder {
		return middleware.NotImplemented("operation character.GetCharactersCharacterID has not yet been implemented")
	})
	api.CharacterGetCharactersCharacterIDAgentsResearchHandler = character.GetCharactersCharacterIDAgentsResearchHandlerFunc(func(params character.GetCharactersCharacterIDAgentsResearchParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation character.GetCharactersCharacterIDAgentsResearch has not yet been implemented")
	})
	api.AssetsGetCharactersCharacterIDAssetsHandler = assets.GetCharactersCharacterIDAssetsHandlerFunc(func(params assets.GetCharactersCharacterIDAssetsParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation assets.GetCharactersCharacterIDAssets has not yet been implemented")
	})
	api.CharacterGetCharactersCharacterIDBlueprintsHandler = character.GetCharactersCharacterIDBlueprintsHandlerFunc(func(params character.GetCharactersCharacterIDBlueprintsParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation character.GetCharactersCharacterIDBlueprints has not yet been implemented")
	})
	api.BookmarksGetCharactersCharacterIDBookmarksHandler = bookmarks.GetCharactersCharacterIDBookmarksHandlerFunc(func(params bookmarks.GetCharactersCharacterIDBookmarksParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation bookmarks.GetCharactersCharacterIDBookmarks has not yet been implemented")
	})
	api.BookmarksGetCharactersCharacterIDBookmarksFoldersHandler = bookmarks.GetCharactersCharacterIDBookmarksFoldersHandlerFunc(func(params bookmarks.GetCharactersCharacterIDBookmarksFoldersParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation bookmarks.GetCharactersCharacterIDBookmarksFolders has not yet been implemented")
	})
	api.CalendarGetCharactersCharacterIDCalendarHandler = calendar.GetCharactersCharacterIDCalendarHandlerFunc(func(params calendar.GetCharactersCharacterIDCalendarParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation calendar.GetCharactersCharacterIDCalendar has not yet been implemented")
	})
	api.CalendarGetCharactersCharacterIDCalendarEventIDHandler = calendar.GetCharactersCharacterIDCalendarEventIDHandlerFunc(func(params calendar.GetCharactersCharacterIDCalendarEventIDParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation calendar.GetCharactersCharacterIDCalendarEventID has not yet been implemented")
	})
	api.CharacterGetCharactersCharacterIDChatChannelsHandler = character.GetCharactersCharacterIDChatChannelsHandlerFunc(func(params character.GetCharactersCharacterIDChatChannelsParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation character.GetCharactersCharacterIDChatChannels has not yet been implemented")
	})
	api.ClonesGetCharactersCharacterIDClonesHandler = clones.GetCharactersCharacterIDClonesHandlerFunc(func(params clones.GetCharactersCharacterIDClonesParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation clones.GetCharactersCharacterIDClones has not yet been implemented")
	})
	api.ContactsGetCharactersCharacterIDContactsHandler = contacts.GetCharactersCharacterIDContactsHandlerFunc(func(params contacts.GetCharactersCharacterIDContactsParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation contacts.GetCharactersCharacterIDContacts has not yet been implemented")
	})
	api.ContactsGetCharactersCharacterIDContactsLabelsHandler = contacts.GetCharactersCharacterIDContactsLabelsHandlerFunc(func(params contacts.GetCharactersCharacterIDContactsLabelsParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation contacts.GetCharactersCharacterIDContactsLabels has not yet been implemented")
	})
	api.CharacterGetCharactersCharacterIDCorporationhistoryHandler = character.GetCharactersCharacterIDCorporationhistoryHandlerFunc(func(params character.GetCharactersCharacterIDCorporationhistoryParams) middleware.Responder {
		return middleware.NotImplemented("operation character.GetCharactersCharacterIDCorporationhistory has not yet been implemented")
	})
	api.FittingsGetCharactersCharacterIDFittingsHandler = fittings.GetCharactersCharacterIDFittingsHandlerFunc(func(params fittings.GetCharactersCharacterIDFittingsParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation fittings.GetCharactersCharacterIDFittings has not yet been implemented")
	})
	api.IndustryGetCharactersCharacterIDIndustryJobsHandler = industry.GetCharactersCharacterIDIndustryJobsHandlerFunc(func(params industry.GetCharactersCharacterIDIndustryJobsParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation industry.GetCharactersCharacterIDIndustryJobs has not yet been implemented")
	})
	api.KillmailsGetCharactersCharacterIDKillmailsRecentHandler = killmails.GetCharactersCharacterIDKillmailsRecentHandlerFunc(func(params killmails.GetCharactersCharacterIDKillmailsRecentParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation killmails.GetCharactersCharacterIDKillmailsRecent has not yet been implemented")
	})
	api.LocationGetCharactersCharacterIDLocationHandler = location.GetCharactersCharacterIDLocationHandlerFunc(func(params location.GetCharactersCharacterIDLocationParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation location.GetCharactersCharacterIDLocation has not yet been implemented")
	})
	api.LoyaltyGetCharactersCharacterIDLoyaltyPointsHandler = loyalty.GetCharactersCharacterIDLoyaltyPointsHandlerFunc(func(params loyalty.GetCharactersCharacterIDLoyaltyPointsParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation loyalty.GetCharactersCharacterIDLoyaltyPoints has not yet been implemented")
	})
	api.MailGetCharactersCharacterIDMailHandler = mail.GetCharactersCharacterIDMailHandlerFunc(func(params mail.GetCharactersCharacterIDMailParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation mail.GetCharactersCharacterIDMail has not yet been implemented")
	})
	api.MailGetCharactersCharacterIDMailLabelsHandler = mail.GetCharactersCharacterIDMailLabelsHandlerFunc(func(params mail.GetCharactersCharacterIDMailLabelsParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation mail.GetCharactersCharacterIDMailLabels has not yet been implemented")
	})
	api.MailGetCharactersCharacterIDMailListsHandler = mail.GetCharactersCharacterIDMailListsHandlerFunc(func(params mail.GetCharactersCharacterIDMailListsParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation mail.GetCharactersCharacterIDMailLists has not yet been implemented")
	})
	api.MailGetCharactersCharacterIDMailMailIDHandler = mail.GetCharactersCharacterIDMailMailIDHandlerFunc(func(params mail.GetCharactersCharacterIDMailMailIDParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation mail.GetCharactersCharacterIDMailMailID has not yet been implemented")
	})
	api.CharacterGetCharactersCharacterIDMedalsHandler = character.GetCharactersCharacterIDMedalsHandlerFunc(func(params character.GetCharactersCharacterIDMedalsParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation character.GetCharactersCharacterIDMedals has not yet been implemented")
	})
	api.LocationGetCharactersCharacterIDOnlineHandler = location.GetCharactersCharacterIDOnlineHandlerFunc(func(params location.GetCharactersCharacterIDOnlineParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation location.GetCharactersCharacterIDOnline has not yet been implemented")
	})
	api.OpportunitiesGetCharactersCharacterIDOpportunitiesHandler = opportunities.GetCharactersCharacterIDOpportunitiesHandlerFunc(func(params opportunities.GetCharactersCharacterIDOpportunitiesParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation opportunities.GetCharactersCharacterIDOpportunities has not yet been implemented")
	})
	api.MarketGetCharactersCharacterIDOrdersHandler = market.GetCharactersCharacterIDOrdersHandlerFunc(func(params market.GetCharactersCharacterIDOrdersParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation market.GetCharactersCharacterIDOrders has not yet been implemented")
	})
	api.PlanetaryInteractionGetCharactersCharacterIDPlanetsHandler = planetary_interaction.GetCharactersCharacterIDPlanetsHandlerFunc(func(params planetary_interaction.GetCharactersCharacterIDPlanetsParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation planetary_interaction.GetCharactersCharacterIDPlanets has not yet been implemented")
	})
	api.PlanetaryInteractionGetCharactersCharacterIDPlanetsPlanetIDHandler = planetary_interaction.GetCharactersCharacterIDPlanetsPlanetIDHandlerFunc(func(params planetary_interaction.GetCharactersCharacterIDPlanetsPlanetIDParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation planetary_interaction.GetCharactersCharacterIDPlanetsPlanetID has not yet been implemented")
	})
	api.CharacterGetCharactersCharacterIDPortraitHandler = character.GetCharactersCharacterIDPortraitHandlerFunc(func(params character.GetCharactersCharacterIDPortraitParams) middleware.Responder {
		return middleware.NotImplemented("operation character.GetCharactersCharacterIDPortrait has not yet been implemented")
	})
	api.CharacterGetCharactersCharacterIDRolesHandler = character.GetCharactersCharacterIDRolesHandlerFunc(func(params character.GetCharactersCharacterIDRolesParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation character.GetCharactersCharacterIDRoles has not yet been implemented")
	})
	api.SearchGetCharactersCharacterIDSearchHandler = search.GetCharactersCharacterIDSearchHandlerFunc(func(params search.GetCharactersCharacterIDSearchParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation search.GetCharactersCharacterIDSearch has not yet been implemented")
	})
	api.LocationGetCharactersCharacterIDShipHandler = location.GetCharactersCharacterIDShipHandlerFunc(func(params location.GetCharactersCharacterIDShipParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation location.GetCharactersCharacterIDShip has not yet been implemented")
	})
	api.SkillsGetCharactersCharacterIDSkillqueueHandler = skills.GetCharactersCharacterIDSkillqueueHandlerFunc(func(params skills.GetCharactersCharacterIDSkillqueueParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation skills.GetCharactersCharacterIDSkillqueue has not yet been implemented")
	})
	api.SkillsGetCharactersCharacterIDSkillsHandler = skills.GetCharactersCharacterIDSkillsHandlerFunc(func(params skills.GetCharactersCharacterIDSkillsParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation skills.GetCharactersCharacterIDSkills has not yet been implemented")
	})
	api.CharacterGetCharactersCharacterIDStandingsHandler = character.GetCharactersCharacterIDStandingsHandlerFunc(func(params character.GetCharactersCharacterIDStandingsParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation character.GetCharactersCharacterIDStandings has not yet been implemented")
	})
	api.WalletGetCharactersCharacterIDWalletsHandler = wallet.GetCharactersCharacterIDWalletsHandlerFunc(func(params wallet.GetCharactersCharacterIDWalletsParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation wallet.GetCharactersCharacterIDWallets has not yet been implemented")
	})
	api.CharacterGetCharactersNamesHandler = character.GetCharactersNamesHandlerFunc(func(params character.GetCharactersNamesParams) middleware.Responder {
		return middleware.NotImplemented("operation character.GetCharactersNames has not yet been implemented")
	})
	api.CorporationGetCorporationsCorporationIDHandler = corporation.GetCorporationsCorporationIDHandlerFunc(func(params corporation.GetCorporationsCorporationIDParams) middleware.Responder {
		return middleware.NotImplemented("operation corporation.GetCorporationsCorporationID has not yet been implemented")
	})
	api.CorporationGetCorporationsCorporationIDAlliancehistoryHandler = corporation.GetCorporationsCorporationIDAlliancehistoryHandlerFunc(func(params corporation.GetCorporationsCorporationIDAlliancehistoryParams) middleware.Responder {
		return middleware.NotImplemented("operation corporation.GetCorporationsCorporationIDAlliancehistory has not yet been implemented")
	})
	api.CorporationGetCorporationsCorporationIDIconsHandler = corporation.GetCorporationsCorporationIDIconsHandlerFunc(func(params corporation.GetCorporationsCorporationIDIconsParams) middleware.Responder {
		return middleware.NotImplemented("operation corporation.GetCorporationsCorporationIDIcons has not yet been implemented")
	})
	api.CorporationGetCorporationsCorporationIDMembersHandler = corporation.GetCorporationsCorporationIDMembersHandlerFunc(func(params corporation.GetCorporationsCorporationIDMembersParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation corporation.GetCorporationsCorporationIDMembers has not yet been implemented")
	})
	api.CorporationGetCorporationsCorporationIDRolesHandler = corporation.GetCorporationsCorporationIDRolesHandlerFunc(func(params corporation.GetCorporationsCorporationIDRolesParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation corporation.GetCorporationsCorporationIDRoles has not yet been implemented")
	})
	api.CorporationGetCorporationsCorporationIDStructuresHandler = corporation.GetCorporationsCorporationIDStructuresHandlerFunc(func(params corporation.GetCorporationsCorporationIDStructuresParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation corporation.GetCorporationsCorporationIDStructures has not yet been implemented")
	})
	api.CorporationGetCorporationsNamesHandler = corporation.GetCorporationsNamesHandlerFunc(func(params corporation.GetCorporationsNamesParams) middleware.Responder {
		return middleware.NotImplemented("operation corporation.GetCorporationsNames has not yet been implemented")
	})
	api.CorporationGetCorporationsNpccorpsHandler = corporation.GetCorporationsNpccorpsHandlerFunc(func(params corporation.GetCorporationsNpccorpsParams) middleware.Responder {
		return middleware.NotImplemented("operation corporation.GetCorporationsNpccorps has not yet been implemented")
	})
	api.DogmaGetDogmaAttributesHandler = dogma.GetDogmaAttributesHandlerFunc(func(params dogma.GetDogmaAttributesParams) middleware.Responder {
		return middleware.NotImplemented("operation dogma.GetDogmaAttributes has not yet been implemented")
	})
	api.DogmaGetDogmaAttributesAttributeIDHandler = dogma.GetDogmaAttributesAttributeIDHandlerFunc(func(params dogma.GetDogmaAttributesAttributeIDParams) middleware.Responder {
		return middleware.NotImplemented("operation dogma.GetDogmaAttributesAttributeID has not yet been implemented")
	})
	api.DogmaGetDogmaEffectsHandler = dogma.GetDogmaEffectsHandlerFunc(func(params dogma.GetDogmaEffectsParams) middleware.Responder {
		return middleware.NotImplemented("operation dogma.GetDogmaEffects has not yet been implemented")
	})
	api.DogmaGetDogmaEffectsEffectIDHandler = dogma.GetDogmaEffectsEffectIDHandlerFunc(func(params dogma.GetDogmaEffectsEffectIDParams) middleware.Responder {
		return middleware.NotImplemented("operation dogma.GetDogmaEffectsEffectID has not yet been implemented")
	})
	api.FleetsGetFleetsFleetIDHandler = fleets.GetFleetsFleetIDHandlerFunc(func(params fleets.GetFleetsFleetIDParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation fleets.GetFleetsFleetID has not yet been implemented")
	})
	api.FleetsGetFleetsFleetIDMembersHandler = fleets.GetFleetsFleetIDMembersHandlerFunc(func(params fleets.GetFleetsFleetIDMembersParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation fleets.GetFleetsFleetIDMembers has not yet been implemented")
	})
	api.FleetsGetFleetsFleetIDWingsHandler = fleets.GetFleetsFleetIDWingsHandlerFunc(func(params fleets.GetFleetsFleetIDWingsParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation fleets.GetFleetsFleetIDWings has not yet been implemented")
	})
	api.IncursionsGetIncursionsHandler = incursions.GetIncursionsHandlerFunc(func(params incursions.GetIncursionsParams) middleware.Responder {
		return middleware.NotImplemented("operation incursions.GetIncursions has not yet been implemented")
	})
	api.IndustryGetIndustryFacilitiesHandler = industry.GetIndustryFacilitiesHandlerFunc(func(params industry.GetIndustryFacilitiesParams) middleware.Responder {
		return middleware.NotImplemented("operation industry.GetIndustryFacilities has not yet been implemented")
	})
	api.IndustryGetIndustrySystemsHandler = industry.GetIndustrySystemsHandlerFunc(func(params industry.GetIndustrySystemsParams) middleware.Responder {
		return middleware.NotImplemented("operation industry.GetIndustrySystems has not yet been implemented")
	})
	api.InsuranceGetInsurancePricesHandler = insurance.GetInsurancePricesHandlerFunc(func(params insurance.GetInsurancePricesParams) middleware.Responder {
		return middleware.NotImplemented("operation insurance.GetInsurancePrices has not yet been implemented")
	})
	api.KillmailsGetKillmailsKillmailIDKillmailHashHandler = killmails.GetKillmailsKillmailIDKillmailHashHandlerFunc(func(params killmails.GetKillmailsKillmailIDKillmailHashParams) middleware.Responder {
		return middleware.NotImplemented("operation killmails.GetKillmailsKillmailIDKillmailHash has not yet been implemented")
	})
	api.LoyaltyGetLoyaltyStoresCorporationIDOffersHandler = loyalty.GetLoyaltyStoresCorporationIDOffersHandlerFunc(func(params loyalty.GetLoyaltyStoresCorporationIDOffersParams) middleware.Responder {
		return middleware.NotImplemented("operation loyalty.GetLoyaltyStoresCorporationIDOffers has not yet been implemented")
	})
	api.MarketGetMarketsGroupsHandler = market.GetMarketsGroupsHandlerFunc(func(params market.GetMarketsGroupsParams) middleware.Responder {
		return middleware.NotImplemented("operation market.GetMarketsGroups has not yet been implemented")
	})
	api.MarketGetMarketsGroupsMarketGroupIDHandler = market.GetMarketsGroupsMarketGroupIDHandlerFunc(func(params market.GetMarketsGroupsMarketGroupIDParams) middleware.Responder {
		return middleware.NotImplemented("operation market.GetMarketsGroupsMarketGroupID has not yet been implemented")
	})
	api.MarketGetMarketsPricesHandler = market.GetMarketsPricesHandlerFunc(func(params market.GetMarketsPricesParams) middleware.Responder {
		return middleware.NotImplemented("operation market.GetMarketsPrices has not yet been implemented")
	})
	api.MarketGetMarketsRegionIDHistoryHandler = market.GetMarketsRegionIDHistoryHandlerFunc(func(params market.GetMarketsRegionIDHistoryParams) middleware.Responder {
		return middleware.NotImplemented("operation market.GetMarketsRegionIDHistory has not yet been implemented")
	})
	api.MarketGetMarketsRegionIDOrdersHandler = market.GetMarketsRegionIDOrdersHandlerFunc(func(params market.GetMarketsRegionIDOrdersParams) middleware.Responder {
		return middleware.NotImplemented("operation market.GetMarketsRegionIDOrders has not yet been implemented")
	})
	api.MarketGetMarketsStructuresStructureIDHandler = market.GetMarketsStructuresStructureIDHandlerFunc(func(params market.GetMarketsStructuresStructureIDParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation market.GetMarketsStructuresStructureID has not yet been implemented")
	})
	api.OpportunitiesGetOpportunitiesGroupsHandler = opportunities.GetOpportunitiesGroupsHandlerFunc(func(params opportunities.GetOpportunitiesGroupsParams) middleware.Responder {
		return middleware.NotImplemented("operation opportunities.GetOpportunitiesGroups has not yet been implemented")
	})
	api.OpportunitiesGetOpportunitiesGroupsGroupIDHandler = opportunities.GetOpportunitiesGroupsGroupIDHandlerFunc(func(params opportunities.GetOpportunitiesGroupsGroupIDParams) middleware.Responder {
		return middleware.NotImplemented("operation opportunities.GetOpportunitiesGroupsGroupID has not yet been implemented")
	})
	api.OpportunitiesGetOpportunitiesTasksHandler = opportunities.GetOpportunitiesTasksHandlerFunc(func(params opportunities.GetOpportunitiesTasksParams) middleware.Responder {
		return middleware.NotImplemented("operation opportunities.GetOpportunitiesTasks has not yet been implemented")
	})
	api.OpportunitiesGetOpportunitiesTasksTaskIDHandler = opportunities.GetOpportunitiesTasksTaskIDHandlerFunc(func(params opportunities.GetOpportunitiesTasksTaskIDParams) middleware.Responder {
		return middleware.NotImplemented("operation opportunities.GetOpportunitiesTasksTaskID has not yet been implemented")
	})
	api.SearchGetSearchHandler = search.GetSearchHandlerFunc(func(params search.GetSearchParams) middleware.Responder {
		return middleware.NotImplemented("operation search.GetSearch has not yet been implemented")
	})
	api.SovereigntyGetSovereigntyCampaignsHandler = sovereignty.GetSovereigntyCampaignsHandlerFunc(func(params sovereignty.GetSovereigntyCampaignsParams) middleware.Responder {
		return middleware.NotImplemented("operation sovereignty.GetSovereigntyCampaigns has not yet been implemented")
	})
	api.SovereigntyGetSovereigntyMapHandler = sovereignty.GetSovereigntyMapHandlerFunc(func(params sovereignty.GetSovereigntyMapParams) middleware.Responder {
		return middleware.NotImplemented("operation sovereignty.GetSovereigntyMap has not yet been implemented")
	})
	api.SovereigntyGetSovereigntyStructuresHandler = sovereignty.GetSovereigntyStructuresHandlerFunc(func(params sovereignty.GetSovereigntyStructuresParams) middleware.Responder {
		return middleware.NotImplemented("operation sovereignty.GetSovereigntyStructures has not yet been implemented")
	})
	api.StatusGetStatusHandler = status.GetStatusHandlerFunc(func(params status.GetStatusParams) middleware.Responder {
		return middleware.NotImplemented("operation status.GetStatus has not yet been implemented")
	})
	api.UniverseGetUniverseBloodlinesHandler = universe.GetUniverseBloodlinesHandlerFunc(func(params universe.GetUniverseBloodlinesParams) middleware.Responder {
		return middleware.NotImplemented("operation universe.GetUniverseBloodlines has not yet been implemented")
	})
	api.UniverseGetUniverseCategoriesHandler = universe.GetUniverseCategoriesHandlerFunc(func(params universe.GetUniverseCategoriesParams) middleware.Responder {
		return middleware.NotImplemented("operation universe.GetUniverseCategories has not yet been implemented")
	})
	api.UniverseGetUniverseCategoriesCategoryIDHandler = universe.GetUniverseCategoriesCategoryIDHandlerFunc(func(params universe.GetUniverseCategoriesCategoryIDParams) middleware.Responder {
		return middleware.NotImplemented("operation universe.GetUniverseCategoriesCategoryID has not yet been implemented")
	})
	api.UniverseGetUniverseConstellationsHandler = universe.GetUniverseConstellationsHandlerFunc(func(params universe.GetUniverseConstellationsParams) middleware.Responder {
		return middleware.NotImplemented("operation universe.GetUniverseConstellations has not yet been implemented")
	})
	api.UniverseGetUniverseConstellationsConstellationIDHandler = universe.GetUniverseConstellationsConstellationIDHandlerFunc(func(params universe.GetUniverseConstellationsConstellationIDParams) middleware.Responder {
		return middleware.NotImplemented("operation universe.GetUniverseConstellationsConstellationID has not yet been implemented")
	})
	api.UniverseGetUniverseFactionsHandler = universe.GetUniverseFactionsHandlerFunc(func(params universe.GetUniverseFactionsParams) middleware.Responder {
		return middleware.NotImplemented("operation universe.GetUniverseFactions has not yet been implemented")
	})
	api.UniverseGetUniverseGraphicsHandler = universe.GetUniverseGraphicsHandlerFunc(func(params universe.GetUniverseGraphicsParams) middleware.Responder {
		return middleware.NotImplemented("operation universe.GetUniverseGraphics has not yet been implemented")
	})
	api.UniverseGetUniverseGraphicsGraphicIDHandler = universe.GetUniverseGraphicsGraphicIDHandlerFunc(func(params universe.GetUniverseGraphicsGraphicIDParams) middleware.Responder {
		return middleware.NotImplemented("operation universe.GetUniverseGraphicsGraphicID has not yet been implemented")
	})
	api.UniverseGetUniverseGroupsHandler = universe.GetUniverseGroupsHandlerFunc(func(params universe.GetUniverseGroupsParams) middleware.Responder {
		return middleware.NotImplemented("operation universe.GetUniverseGroups has not yet been implemented")
	})
	api.UniverseGetUniverseGroupsGroupIDHandler = universe.GetUniverseGroupsGroupIDHandlerFunc(func(params universe.GetUniverseGroupsGroupIDParams) middleware.Responder {
		return middleware.NotImplemented("operation universe.GetUniverseGroupsGroupID has not yet been implemented")
	})
	api.UniverseGetUniverseMoonsMoonIDHandler = universe.GetUniverseMoonsMoonIDHandlerFunc(func(params universe.GetUniverseMoonsMoonIDParams) middleware.Responder {
		return middleware.NotImplemented("operation universe.GetUniverseMoonsMoonID has not yet been implemented")
	})
	api.UniverseGetUniversePlanetsPlanetIDHandler = universe.GetUniversePlanetsPlanetIDHandlerFunc(func(params universe.GetUniversePlanetsPlanetIDParams) middleware.Responder {
		return middleware.NotImplemented("operation universe.GetUniversePlanetsPlanetID has not yet been implemented")
	})
	api.UniverseGetUniverseRacesHandler = universe.GetUniverseRacesHandlerFunc(func(params universe.GetUniverseRacesParams) middleware.Responder {
		return middleware.NotImplemented("operation universe.GetUniverseRaces has not yet been implemented")
	})
	api.UniverseGetUniverseRegionsHandler = universe.GetUniverseRegionsHandlerFunc(func(params universe.GetUniverseRegionsParams) middleware.Responder {
		return middleware.NotImplemented("operation universe.GetUniverseRegions has not yet been implemented")
	})
	api.UniverseGetUniverseRegionsRegionIDHandler = universe.GetUniverseRegionsRegionIDHandlerFunc(func(params universe.GetUniverseRegionsRegionIDParams) middleware.Responder {
		return middleware.NotImplemented("operation universe.GetUniverseRegionsRegionID has not yet been implemented")
	})
	api.PlanetaryInteractionGetUniverseSchematicsSchematicIDHandler = planetary_interaction.GetUniverseSchematicsSchematicIDHandlerFunc(func(params planetary_interaction.GetUniverseSchematicsSchematicIDParams) middleware.Responder {
		return middleware.NotImplemented("operation planetary_interaction.GetUniverseSchematicsSchematicID has not yet been implemented")
	})
	api.UniverseGetUniverseStargatesStargateIDHandler = universe.GetUniverseStargatesStargateIDHandlerFunc(func(params universe.GetUniverseStargatesStargateIDParams) middleware.Responder {
		return middleware.NotImplemented("operation universe.GetUniverseStargatesStargateID has not yet been implemented")
	})
	api.UniverseGetUniverseStationsStationIDHandler = universe.GetUniverseStationsStationIDHandlerFunc(func(params universe.GetUniverseStationsStationIDParams) middleware.Responder {
		return middleware.NotImplemented("operation universe.GetUniverseStationsStationID has not yet been implemented")
	})
	api.UniverseGetUniverseStructuresHandler = universe.GetUniverseStructuresHandlerFunc(func(params universe.GetUniverseStructuresParams) middleware.Responder {
		return middleware.NotImplemented("operation universe.GetUniverseStructures has not yet been implemented")
	})
	api.UniverseGetUniverseStructuresStructureIDHandler = universe.GetUniverseStructuresStructureIDHandlerFunc(func(params universe.GetUniverseStructuresStructureIDParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation universe.GetUniverseStructuresStructureID has not yet been implemented")
	})
	api.UniverseGetUniverseSystemJumpsHandler = universe.GetUniverseSystemJumpsHandlerFunc(func(params universe.GetUniverseSystemJumpsParams) middleware.Responder {
		return middleware.NotImplemented("operation universe.GetUniverseSystemJumps has not yet been implemented")
	})
	api.UniverseGetUniverseSystemKillsHandler = universe.GetUniverseSystemKillsHandlerFunc(func(params universe.GetUniverseSystemKillsParams) middleware.Responder {
		return middleware.NotImplemented("operation universe.GetUniverseSystemKills has not yet been implemented")
	})
	api.UniverseGetUniverseSystemsHandler = universe.GetUniverseSystemsHandlerFunc(func(params universe.GetUniverseSystemsParams) middleware.Responder {
		return middleware.NotImplemented("operation universe.GetUniverseSystems has not yet been implemented")
	})
	api.UniverseGetUniverseSystemsSystemIDHandler = universe.GetUniverseSystemsSystemIDHandlerFunc(func(params universe.GetUniverseSystemsSystemIDParams) middleware.Responder {
		return middleware.NotImplemented("operation universe.GetUniverseSystemsSystemID has not yet been implemented")
	})
	api.UniverseGetUniverseTypesHandler = universe.GetUniverseTypesHandlerFunc(func(params universe.GetUniverseTypesParams) middleware.Responder {
		return middleware.NotImplemented("operation universe.GetUniverseTypes has not yet been implemented")
	})
	api.UniverseGetUniverseTypesTypeIDHandler = universe.GetUniverseTypesTypeIDHandlerFunc(func(params universe.GetUniverseTypesTypeIDParams) middleware.Responder {
		return middleware.NotImplemented("operation universe.GetUniverseTypesTypeID has not yet been implemented")
	})
	api.WarsGetWarsHandler = wars.GetWarsHandlerFunc(func(params wars.GetWarsParams) middleware.Responder {
		return middleware.NotImplemented("operation wars.GetWars has not yet been implemented")
	})
	api.WarsGetWarsWarIDHandler = wars.GetWarsWarIDHandlerFunc(func(params wars.GetWarsWarIDParams) middleware.Responder {
		return middleware.NotImplemented("operation wars.GetWarsWarID has not yet been implemented")
	})
	api.WarsGetWarsWarIDKillmailsHandler = wars.GetWarsWarIDKillmailsHandlerFunc(func(params wars.GetWarsWarIDKillmailsParams) middleware.Responder {
		return middleware.NotImplemented("operation wars.GetWarsWarIDKillmails has not yet been implemented")
	})
	api.CharacterPostCharactersAffiliationHandler = character.PostCharactersAffiliationHandlerFunc(func(params character.PostCharactersAffiliationParams) middleware.Responder {
		return middleware.NotImplemented("operation character.PostCharactersAffiliation has not yet been implemented")
	})
	api.ContactsPostCharactersCharacterIDContactsHandler = contacts.PostCharactersCharacterIDContactsHandlerFunc(func(params contacts.PostCharactersCharacterIDContactsParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation contacts.PostCharactersCharacterIDContacts has not yet been implemented")
	})
	api.CharacterPostCharactersCharacterIDCspaHandler = character.PostCharactersCharacterIDCspaHandlerFunc(func(params character.PostCharactersCharacterIDCspaParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation character.PostCharactersCharacterIDCspa has not yet been implemented")
	})
	api.FittingsPostCharactersCharacterIDFittingsHandler = fittings.PostCharactersCharacterIDFittingsHandlerFunc(func(params fittings.PostCharactersCharacterIDFittingsParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation fittings.PostCharactersCharacterIDFittings has not yet been implemented")
	})
	api.MailPostCharactersCharacterIDMailHandler = mail.PostCharactersCharacterIDMailHandlerFunc(func(params mail.PostCharactersCharacterIDMailParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation mail.PostCharactersCharacterIDMail has not yet been implemented")
	})
	api.MailPostCharactersCharacterIDMailLabelsHandler = mail.PostCharactersCharacterIDMailLabelsHandlerFunc(func(params mail.PostCharactersCharacterIDMailLabelsParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation mail.PostCharactersCharacterIDMailLabels has not yet been implemented")
	})
	api.FleetsPostFleetsFleetIDMembersHandler = fleets.PostFleetsFleetIDMembersHandlerFunc(func(params fleets.PostFleetsFleetIDMembersParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation fleets.PostFleetsFleetIDMembers has not yet been implemented")
	})
	api.FleetsPostFleetsFleetIDWingsHandler = fleets.PostFleetsFleetIDWingsHandlerFunc(func(params fleets.PostFleetsFleetIDWingsParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation fleets.PostFleetsFleetIDWings has not yet been implemented")
	})
	api.FleetsPostFleetsFleetIDWingsWingIDSquadsHandler = fleets.PostFleetsFleetIDWingsWingIDSquadsHandlerFunc(func(params fleets.PostFleetsFleetIDWingsWingIDSquadsParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation fleets.PostFleetsFleetIDWingsWingIDSquads has not yet been implemented")
	})
	api.UserInterfacePostUIAutopilotWaypointHandler = user_interface.PostUIAutopilotWaypointHandlerFunc(func(params user_interface.PostUIAutopilotWaypointParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation user_interface.PostUIAutopilotWaypoint has not yet been implemented")
	})
	api.UserInterfacePostUIOpenwindowContractHandler = user_interface.PostUIOpenwindowContractHandlerFunc(func(params user_interface.PostUIOpenwindowContractParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation user_interface.PostUIOpenwindowContract has not yet been implemented")
	})
	api.UserInterfacePostUIOpenwindowInformationHandler = user_interface.PostUIOpenwindowInformationHandlerFunc(func(params user_interface.PostUIOpenwindowInformationParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation user_interface.PostUIOpenwindowInformation has not yet been implemented")
	})
	api.UserInterfacePostUIOpenwindowMarketdetailsHandler = user_interface.PostUIOpenwindowMarketdetailsHandlerFunc(func(params user_interface.PostUIOpenwindowMarketdetailsParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation user_interface.PostUIOpenwindowMarketdetails has not yet been implemented")
	})
	api.UserInterfacePostUIOpenwindowNewmailHandler = user_interface.PostUIOpenwindowNewmailHandlerFunc(func(params user_interface.PostUIOpenwindowNewmailParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation user_interface.PostUIOpenwindowNewmail has not yet been implemented")
	})
	api.UniversePostUniverseNamesHandler = universe.PostUniverseNamesHandlerFunc(func(params universe.PostUniverseNamesParams) middleware.Responder {
		return middleware.NotImplemented("operation universe.PostUniverseNames has not yet been implemented")
	})
	api.CalendarPutCharactersCharacterIDCalendarEventIDHandler = calendar.PutCharactersCharacterIDCalendarEventIDHandlerFunc(func(params calendar.PutCharactersCharacterIDCalendarEventIDParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation calendar.PutCharactersCharacterIDCalendarEventID has not yet been implemented")
	})
	api.ContactsPutCharactersCharacterIDContactsHandler = contacts.PutCharactersCharacterIDContactsHandlerFunc(func(params contacts.PutCharactersCharacterIDContactsParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation contacts.PutCharactersCharacterIDContacts has not yet been implemented")
	})
	api.MailPutCharactersCharacterIDMailMailIDHandler = mail.PutCharactersCharacterIDMailMailIDHandlerFunc(func(params mail.PutCharactersCharacterIDMailMailIDParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation mail.PutCharactersCharacterIDMailMailID has not yet been implemented")
	})
	api.CorporationPutCorporationsCorporationIDStructuresStructureIDHandler = corporation.PutCorporationsCorporationIDStructuresStructureIDHandlerFunc(func(params corporation.PutCorporationsCorporationIDStructuresStructureIDParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation corporation.PutCorporationsCorporationIDStructuresStructureID has not yet been implemented")
	})
	api.FleetsPutFleetsFleetIDHandler = fleets.PutFleetsFleetIDHandlerFunc(func(params fleets.PutFleetsFleetIDParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation fleets.PutFleetsFleetID has not yet been implemented")
	})
	api.FleetsPutFleetsFleetIDMembersMemberIDHandler = fleets.PutFleetsFleetIDMembersMemberIDHandlerFunc(func(params fleets.PutFleetsFleetIDMembersMemberIDParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation fleets.PutFleetsFleetIDMembersMemberID has not yet been implemented")
	})
	api.FleetsPutFleetsFleetIDSquadsSquadIDHandler = fleets.PutFleetsFleetIDSquadsSquadIDHandlerFunc(func(params fleets.PutFleetsFleetIDSquadsSquadIDParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation fleets.PutFleetsFleetIDSquadsSquadID has not yet been implemented")
	})
	api.FleetsPutFleetsFleetIDWingsWingIDHandler = fleets.PutFleetsFleetIDWingsWingIDHandlerFunc(func(params fleets.PutFleetsFleetIDWingsWingIDParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation fleets.PutFleetsFleetIDWingsWingID has not yet been implemented")
	})

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *graceful.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
