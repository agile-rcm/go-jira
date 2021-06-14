package jira

import (
	"context"
	"fmt"
)

// PluginService handles Plugins for the Jira instance / API.
//
// Jira API docs: https://docs.atlassian.com/jira/REST/latest/#api/2/plugins
type PluginService struct {
	client *Client
}

type PluginStruct struct {
	Plugins []Plugin `json:"plugins"`
	Links   LinksCommon     `json:"links"`
}
type Links struct {
	Self          string `json:"self"`
	PluginSummary string `json:"plugin-summary"`
	Modify        string `json:"modify"`
	PluginIcon    string `json:"plugin-icon"`
	PluginLogo    string `json:"plugin-logo"`
	Manage        string `json:"manage"`
}
type Vendor struct {
	Name            string `json:"name"`
	MarketplaceLink string `json:"marketplaceLink"`
	Link            string `json:"link"`
}
type Plugin struct {
	Enabled       bool   `json:"enabled"`
	Links         Links  `json:"links"`
	Name          string `json:"name"`
	Version       string `json:"version"`
	UserInstalled bool   `json:"userInstalled"`
	Optional      bool   `json:"optional"`
	Static        bool   `json:"static"`
	Unloadable    bool   `json:"unloadable"`
	Description   string `json:"description"`
	Key           string `json:"key"`
	UsesLicensing bool   `json:"usesLicensing"`
	Remotable     bool   `json:"remotable"`
	Vendor        Vendor `json:"vendor"`
}
type LinksCommon struct {
	Self          string `json:"self"`
	Marketplace   string `json:"marketplace"`
	Notifications string `json:"notifications"`
	Requests      string `json:"requests"`
	Categories    string `json:"categories"`
	Banners       string `json:"banners"`
	Featured      string `json:"featured"`
	HighestRated  string `json:"highest-rated"`
	TopGrossing   string `json:"top-grossing"`
	Popular       string `json:"popular"`
	Trending      string `json:"trending"`
	Atlassian     string `json:"atlassian"`
	Available     string `json:"available"`
	TopVendor     string `json:"top-vendor"`
	EnterSafeMode string `json:"enter-safe-mode"`
	OsgiBundles   string `json:"osgi-bundles"`
}

//###############
type PluginDetails struct {
	Links struct {
		Self          string `json:"self"`
		PluginSummary string `json:"plugin-summary"`
		Modify        string `json:"modify"`
		PluginIcon    string `json:"plugin-icon"`
		PluginLogo    string `json:"plugin-logo"`
		Manage        string `json:"manage"`
	} `json:"links"`
	Key              string `json:"key"`
	Enabled          bool   `json:"enabled"`
	EnabledByDefault bool   `json:"enabledByDefault"`
	Version          string `json:"version"`
	Description      string `json:"description"`
	Name             string `json:"name"`
	Modules          []struct {
		Key         string `json:"key"`
		CompleteKey string `json:"completeKey"`
		Links       struct {
			Self   string `json:"self"`
			Modify string `json:"modify"`
			Plugin string `json:"plugin"`
		} `json:"links"`
		Enabled          bool   `json:"enabled"`
		Optional         bool   `json:"optional"`
		Name             string `json:"name"`
		RecognisableType bool   `json:"recognisableType"`
		Broken           bool   `json:"broken"`
	} `json:"modules"`
	UserInstalled           bool `json:"userInstalled"`
	Optional                bool `json:"optional"`
	UnrecognisedModuleTypes bool `json:"unrecognisedModuleTypes"`
	Unloadable              bool `json:"unloadable"`
	Static                  bool `json:"static"`
	UsesLicensing           bool `json:"usesLicensing"`
	Remotable               bool `json:"remotable"`
	Vendor                  struct {
		Name            string `json:"name"`
		MarketplaceLink string `json:"marketplaceLink"`
		Link            string `json:"link"`
	} `json:"vendor"`
}

type PluginLicenseDetails struct {
	Links struct {
		Self              string `json:"self"`
		Alternate         string `json:"alternate"`
		License           string `json:"license"`
		UpdateLicense     string `json:"update-license"`
		ValidateDowngrade string `json:"validate-downgrade"`
		LicenseCallback   string `json:"license-callback"`
	} `json:"links"`
	PluginKey                    string `json:"pluginKey"`
	Valid                        bool   `json:"valid"`
	Evaluation                   bool   `json:"evaluation"`
	NearlyExpired                bool   `json:"nearlyExpired"`
	MaximumNumberOfUsers         int    `json:"maximumNumberOfUsers"`
	MaintenanceExpiryDate        int64  `json:"maintenanceExpiryDate"`
	MaintenanceExpired           bool   `json:"maintenanceExpired"`
	LicenseType                  string `json:"licenseType"`
	LicenseTypeDescriptionKey    string `json:"licenseTypeDescriptionKey"`
	CreationDateString           string `json:"creationDateString"`
	RawLicense                   string `json:"rawLicense"`
	Renewable                    bool   `json:"renewable"`
	MaintenanceExpiryDateString  string `json:"maintenanceExpiryDateString"`
	OrganizationName             string `json:"organizationName"`
	ContactEmail                 string `json:"contactEmail"`
	Enterprise                   bool   `json:"enterprise"`
	DataCenter                   bool   `json:"dataCenter"`
	Subscription                 bool   `json:"subscription"`
	Active                       bool   `json:"active"`
	AutoRenewal                  bool   `json:"autoRenewal"`
	Upgradable                   bool   `json:"upgradable"`
	Crossgradeable               bool   `json:"crossgradeable"`
	PurchasePastServerCutoffDate bool   `json:"purchasePastServerCutoffDate"`
	TypeI18NSingular             string `json:"typeI18nSingular"`
	TypeI18NPlural               string `json:"typeI18nPlural"`
	SupportEntitlementNumber     string `json:"supportEntitlementNumber"`
}

func (s *PluginService) GetPluginsWithContext(ctx context.Context, plugins *PluginStruct) ( *Response, error) {
	apiEndpoint := "rest/plugins/latest/"
	req, err := s.client.NewRequestWithContext(ctx, "GET", apiEndpoint, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req, plugins)

	if err != nil {
		jerr := NewJiraError(resp, err)
		return resp, jerr
	}

	return resp, nil
}

// Get wraps GetWithContext using the background context.
func (s *PluginService) GetPlugins(plugins *PluginStruct) ( *Response, error) {
	return s.GetPluginsWithContext(context.Background(), plugins )
}

func (s *PluginService) GetPluginDetailsWithContext(ctx context.Context, pluginName string, plugin *PluginDetails) ( *Response, error) {
	apiEndpoint := fmt.Sprintf("rest/plugins/latest/%s-key", pluginName)
	req, err := s.client.NewRequestWithContext(ctx, "GET", apiEndpoint, nil)
	if err != nil {
		return nil, err
	}

	// pluginStruct := new(PluginStruct)
	resp, err := s.client.Do(req, plugin)

	if err != nil {
		jerr := NewJiraError(resp, err)
		return resp, jerr
	}

	return resp, nil
}

// Get wraps GetWithContext using the background context.
func (s *PluginService) GetPluginDetails(pluginName string,plugin *PluginDetails) ( *Response, error) {
	return s.GetPluginDetailsWithContext(context.Background(),pluginName,plugin)
}

func (s *PluginService) GetPluginLicenseDetailsWithContext(ctx context.Context, pluginName string, plugin *PluginLicenseDetails) ( *Response, error) {
	apiEndpoint := fmt.Sprintf("rest/plugins/latest/%s-key/license", pluginName)
	req, err := s.client.NewRequestWithContext(ctx, "GET", apiEndpoint, nil)
	if err != nil {
		return nil, err
	}

	// pluginStruct := new(PluginStruct)
	resp, err := s.client.Do(req, plugin)

	if err != nil {
		jerr := NewJiraError(resp, err)
		return resp, jerr
	}

	return resp, nil
}

// Get wraps GetWithContext using the background context.
func (s *PluginService) GetPluginLicenseDetails(pluginName string,plugin *PluginLicenseDetails) ( *Response, error) {
	return s.GetPluginLicenseDetailsWithContext(context.Background(),pluginName,plugin)
}