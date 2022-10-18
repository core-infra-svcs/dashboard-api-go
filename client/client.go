/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 October, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.26.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"bytes"
	"context"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"golang.org/x/oauth2"
)

var (
	jsonCheck = regexp.MustCompile(`(?i:(?:application|text)/(?:vnd\.[^;]+\+)?json)`)
	xmlCheck  = regexp.MustCompile(`(?i:(?:application|text)/xml)`)
)

// APIClient manages communication with the Meraki Dashboard API API v1.26.0
// In most cases there should be only one, shared, APIClient.
type APIClient struct {
	cfg    *Configuration
	common service // Reuse a single struct instead of allocating one for each service on the heap.

	// API Services

	AccessControlListsApi *AccessControlListsApiService

	AccessPoliciesApi *AccessPoliciesApiService

	AclsApi *AclsApiService

	ActionBatchesApi *ActionBatchesApiService

	AdaptivePolicyApi *AdaptivePolicyApiService

	AddressesApi *AddressesApiService

	AdministeredApi *AdministeredApiService

	AdminsApi *AdminsApiService

	AirMarshalApi *AirMarshalApiService

	AlertTypesApi *AlertTypesApiService

	AlertsApi *AlertsApiService

	AlternateManagementInterfaceApi *AlternateManagementInterfaceApiService

	AnalyticsApi *AnalyticsApiService

	ApiRequestsApi *ApiRequestsApiService

	ApnsCertApi *ApnsCertApiService

	ApplianceApi *ApplianceApiService

	AppliancesApi *AppliancesApiService

	ApplicationCategoriesApi *ApplicationCategoriesApiService

	ApplicationUsageApi *ApplicationUsageApiService

	ApplicationsApi *ApplicationsApiService

	ArpInspectionApi *ArpInspectionApiService

	ArtifactsApi *ArtifactsApiService

	AuthenticationTokenApi *AuthenticationTokenApiService

	AvailabilitiesApi *AvailabilitiesApiService

	BandwidthUsageHistoryApi *BandwidthUsageHistoryApiService

	BgpApi *BgpApiService

	BillingApi *BillingApiService

	BluetoothApi *BluetoothApiService

	BluetoothClientsApi *BluetoothClientsApiService

	BonjourForwardingApi *BonjourForwardingApiService

	BrandingPoliciesApi *BrandingPoliciesApiService

	ByClientApi *ByClientApiService

	ByDeviceApi *ByDeviceApiService

	ByEnergyUsageApi *ByEnergyUsageApiService

	ByMetricApi *ByMetricApiService

	BySwitchApi *BySwitchApiService

	ByUsageApi *ByUsageApiService

	ByUtilizationApi *ByUtilizationApiService

	BypassActivationLockAttemptsApi *BypassActivationLockAttemptsApiService

	CameraApi *CameraApiService

	CategoriesApi *CategoriesApiService

	CellularFirewallRulesApi *CellularFirewallRulesApiService

	CellularGatewayApi *CellularGatewayApiService

	CellularUsageHistoryApi *CellularUsageHistoryApiService

	CertsApi *CertsApiService

	ChannelUtilizationApi *ChannelUtilizationApiService

	ChannelUtilizationHistoryApi *ChannelUtilizationHistoryApiService

	ClaimApi *ClaimApiService

	ClientCountHistoryApi *ClientCountHistoryApiService

	ClientsApi *ClientsApiService

	ConfigTemplatesApi *ConfigTemplatesApiService

	ConfigurationChangesApi *ConfigurationChangesApiService

	ConfigureApi *ConfigureApiService

	ConnectionStatsApi *ConnectionStatsApiService

	ConnectivityApi *ConnectivityApiService

	ConnectivityEventsApi *ConnectivityEventsApiService

	ConnectivityMonitoringDestinationsApi *ConnectivityMonitoringDestinationsApiService

	ContentFilteringApi *ContentFilteringApiService

	CurrentApi *CurrentApiService

	CustomAnalyticsApi *CustomAnalyticsApiService

	CustomPerformanceClassesApi *CustomPerformanceClassesApiService

	DataRateHistoryApi *DataRateHistoryApiService

	DelegatedApi *DelegatedApiService

	DesktopLogsApi *DesktopLogsApiService

	DeviceCommandLogsApi *DeviceCommandLogsApiService

	DeviceProfilesApi *DeviceProfilesApiService

	DeviceTypeGroupPoliciesApi *DeviceTypeGroupPoliciesApiService

	DevicesApi *DevicesApiService

	DhcpApi *DhcpApiService

	DhcpServerPolicyApi *DhcpServerPolicyApiService

	DscpTaggingOptionsApi *DscpTaggingOptionsApiService

	DscpToCosMappingsApi *DscpToCosMappingsApiService

	EapOverrideApi *EapOverrideApiService

	EarlyAccessApi *EarlyAccessApiService

	EventTypesApi *EventTypesApiService

	EventsApi *EventsApiService

	FailedConnectionsApi *FailedConnectionsApiService

	FeaturesApi *FeaturesApiService

	FieldsApi *FieldsApiService

	FirewallApi *FirewallApiService

	FirewalledServicesApi *FirewalledServicesApiService

	FirmwareApi *FirmwareApiService

	FirmwareUpgradesApi *FirmwareUpgradesApiService

	FloorPlansApi *FloorPlansApiService

	GroupPoliciesApi *GroupPoliciesApiService

	GroupsApi *GroupsApiService

	HealthApi *HealthApiService

	HealthByTimeApi *HealthByTimeApiService

	HistoryApi *HistoryApiService

	Hotspot20Api *Hotspot20ApiService

	HttpServersApi *HttpServersApiService

	IdentitiesApi *IdentitiesApiService

	IdentityPsksApi *IdentityPsksApiService

	IdpsApi *IdpsApiService

	InboundCellularFirewallRulesApi *InboundCellularFirewallRulesApiService

	InboundFirewallRulesApi *InboundFirewallRulesApiService

	InsightApi *InsightApiService

	InterfacesApi *InterfacesApiService

	IntrusionApi *IntrusionApiService

	InventoryApi *InventoryApiService

	L3FirewallRulesApi *L3FirewallRulesApiService

	L7FirewallRulesApi *L7FirewallRulesApiService

	LanApi *LanApiService

	LatencyHistoryApi *LatencyHistoryApiService

	LatencyStatsApi *LatencyStatsApiService

	LatestApi *LatestApiService

	LicensesApi *LicensesApiService

	LinkAggregationsApi *LinkAggregationsApiService

	LinkLayerApi *LinkLayerApiService

	LiveApi *LiveApiService

	LiveToolsApi *LiveToolsApiService

	LldpCdpApi *LldpCdpApiService

	LoginSecurityApi *LoginSecurityApiService

	LogsApi *LogsApiService

	LossAndLatencyHistoryApi *LossAndLatencyHistoryApiService

	MalwareApi *MalwareApiService

	ManagementInterfaceApi *ManagementInterfaceApiService

	ManufacturersApi *ManufacturersApiService

	MeApi *MeApiService

	MerakiAuthUsersApi *MerakiAuthUsersApiService

	MeshStatusesApi *MeshStatusesApiService

	ModelsApi *ModelsApiService

	MonitorApi *MonitorApiService

	MonitoredMediaServersApi *MonitoredMediaServersApiService

	MqttBrokersApi *MqttBrokersApiService

	MtuApi *MtuApiService

	MulticastApi *MulticastApiService

	NetflowApi *NetflowApiService

	NetworkAdaptersApi *NetworkAdaptersApiService

	NetworkHealthApi *NetworkHealthApiService

	NetworksApi *NetworksApiService

	ObjectDetectionModelsApi *ObjectDetectionModelsApiService

	OnboardingApi *OnboardingApiService

	OneToManyNatRulesApi *OneToManyNatRulesApiService

	OneToOneNatRulesApi *OneToOneNatRulesApiService

	OpenapiSpecApi *OpenapiSpecApiService

	OptInsApi *OptInsApiService

	OrderApi *OrderApiService

	OrganizationsApi *OrganizationsApiService

	OspfApi *OspfApiService

	OverviewApi *OverviewApiService

	PacketsApi *PacketsApiService

	PayloadTemplatesApi *PayloadTemplatesApiService

	PerformanceApi *PerformanceApiService

	PerformanceHistoryApi *PerformanceHistoryApiService

	PiiApi *PiiApiService

	PiiKeysApi *PiiKeysApiService

	PingApi *PingApiService

	PingDeviceApi *PingDeviceApiService

	PoliciesApi *PoliciesApiService

	PolicyApi *PolicyApiService

	PortForwardingRulesApi *PortForwardingRulesApiService

	PortSchedulesApi *PortSchedulesApiService

	PortsApi *PortsApiService

	PowerModulesApi *PowerModulesApiService

	PrefixesApi *PrefixesApiService

	PrioritiesApi *PrioritiesApiService

	ProfilesApi *ProfilesApiService

	QosRulesApi *QosRulesApiService

	QualityAndRetentionApi *QualityAndRetentionApiService

	QualityRetentionProfilesApi *QualityRetentionProfilesApiService

	RadioApi *RadioApiService

	ReadingsApi *ReadingsApiService

	RecentApi *RecentApiService

	RendezvousPointsApi *RendezvousPointsApiService

	RequestsApi *RequestsApiService

	RestrictionsApi *RestrictionsApiService

	RfProfilesApi *RfProfilesApiService

	RollbacksApi *RollbacksApiService

	RoutingApi *RoutingApiService

	RulesApi *RulesApiService

	SamlApi *SamlApiService

	SamlRolesApi *SamlRolesApiService

	SchedulesApi *SchedulesApiService

	SearchApi *SearchApiService

	SecurityApi *SecurityApiService

	SecurityCentersApi *SecurityCentersApiService

	SeenApi *SeenApiService

	SenseApi *SenseApiService

	SensorApi *SensorApiService

	ServersApi *ServersApiService

	SettingsApi *SettingsApiService

	SignalQualityHistoryApi *SignalQualityHistoryApiService

	SingleLanApi *SingleLanApiService

	SiteToSiteVpnApi *SiteToSiteVpnApiService

	SmApi *SmApiService

	SmDevicesForKeyApi *SmDevicesForKeyApiService

	SmOwnersForKeyApi *SmOwnersForKeyApiService

	SnmpApi *SnmpApiService

	SoftwaresApi *SoftwaresApiService

	SplashApi *SplashApiService

	SplashAuthorizationStatusApi *SplashAuthorizationStatusApiService

	SplashLoginAttemptsApi *SplashLoginAttemptsApiService

	SsidsApi *SsidsApiService

	StacksApi *StacksApiService

	StagedApi *StagedApiService

	StagesApi *StagesApiService

	StaticRoutesApi *StaticRoutesApiService

	StaticsApi *StaticsApiService

	StatsApi *StatsApiService

	StatusApi *StatusApiService

	StatusesApi *StatusesApiService

	StormControlApi *StormControlApiService

	StpApi *StpApiService

	SubnetPoolApi *SubnetPoolApiService

	SubnetsApi *SubnetsApiService

	SummaryApi *SummaryApiService

	SwitchApi *SwitchApiService

	SwitchesApi *SwitchesApiService

	SyslogServersApi *SyslogServersApiService

	TargetGroupsApi *TargetGroupsApiService

	ThirdPartyVPNPeersApi *ThirdPartyVPNPeersApiService

	TopApi *TopApiService

	TopologyApi *TopologyApiService

	TrafficApi *TrafficApiService

	TrafficAnalysisApi *TrafficAnalysisApiService

	TrafficHistoryApi *TrafficHistoryApiService

	TrafficShapingApi *TrafficShapingApiService

	TrustedAccessConfigsApi *TrustedAccessConfigsApiService

	TrustedServersApi *TrustedServersApiService

	UpgradesApi *UpgradesApiService

	UplinkApi *UplinkApiService

	UplinkBandwidthApi *UplinkBandwidthApiService

	UplinkSelectionApi *UplinkSelectionApiService

	UplinksApi *UplinksApiService

	UplinksLossAndLatencyApi *UplinksLossAndLatencyApiService

	UsageHistoriesApi *UsageHistoriesApiService

	UsageHistoryApi *UsageHistoryApiService

	UserAccessDevicesApi *UserAccessDevicesApiService

	UsersApi *UsersApiService

	V4Api *V4ApiService

	VideoApi *VideoApiService

	VideoLinkApi *VideoLinkApiService

	VlanAssignmentsApi *VlanAssignmentsApiService

	VlansApi *VlansApiService

	VmxApi *VmxApiService

	VpnApi *VpnApiService

	VpnFirewallRulesApi *VpnFirewallRulesApiService

	VppAccountsApi *VppAccountsApiService

	WarmSpareApi *WarmSpareApiService

	WarningsApi *WarningsApiService

	WebhookLogsApi *WebhookLogsApiService

	WebhookTestsApi *WebhookTestsApiService

	WebhooksApi *WebhooksApiService

	WirelessApi *WirelessApiService

	WirelessProfilesApi *WirelessProfilesApiService

	WlanListsApi *WlanListsApiService

	ZonesApi *ZonesApiService
}

type service struct {
	client *APIClient
}

// NewAPIClient creates a new API client. Requires a userAgent string describing your application.
// optionally a custom http.Client to allow for advanced features such as caching.
func NewAPIClient(cfg *Configuration) *APIClient {
	if cfg.HTTPClient == nil {
		cfg.HTTPClient = http.DefaultClient
	}

	c := &APIClient{}
	c.cfg = cfg
	c.common.client = c

	// API Services
	c.AccessControlListsApi = (*AccessControlListsApiService)(&c.common)
	c.AccessPoliciesApi = (*AccessPoliciesApiService)(&c.common)
	c.AclsApi = (*AclsApiService)(&c.common)
	c.ActionBatchesApi = (*ActionBatchesApiService)(&c.common)
	c.AdaptivePolicyApi = (*AdaptivePolicyApiService)(&c.common)
	c.AddressesApi = (*AddressesApiService)(&c.common)
	c.AdministeredApi = (*AdministeredApiService)(&c.common)
	c.AdminsApi = (*AdminsApiService)(&c.common)
	c.AirMarshalApi = (*AirMarshalApiService)(&c.common)
	c.AlertTypesApi = (*AlertTypesApiService)(&c.common)
	c.AlertsApi = (*AlertsApiService)(&c.common)
	c.AlternateManagementInterfaceApi = (*AlternateManagementInterfaceApiService)(&c.common)
	c.AnalyticsApi = (*AnalyticsApiService)(&c.common)
	c.ApiRequestsApi = (*ApiRequestsApiService)(&c.common)
	c.ApnsCertApi = (*ApnsCertApiService)(&c.common)
	c.ApplianceApi = (*ApplianceApiService)(&c.common)
	c.AppliancesApi = (*AppliancesApiService)(&c.common)
	c.ApplicationCategoriesApi = (*ApplicationCategoriesApiService)(&c.common)
	c.ApplicationUsageApi = (*ApplicationUsageApiService)(&c.common)
	c.ApplicationsApi = (*ApplicationsApiService)(&c.common)
	c.ArpInspectionApi = (*ArpInspectionApiService)(&c.common)
	c.ArtifactsApi = (*ArtifactsApiService)(&c.common)
	c.AuthenticationTokenApi = (*AuthenticationTokenApiService)(&c.common)
	c.AvailabilitiesApi = (*AvailabilitiesApiService)(&c.common)
	c.BandwidthUsageHistoryApi = (*BandwidthUsageHistoryApiService)(&c.common)
	c.BgpApi = (*BgpApiService)(&c.common)
	c.BillingApi = (*BillingApiService)(&c.common)
	c.BluetoothApi = (*BluetoothApiService)(&c.common)
	c.BluetoothClientsApi = (*BluetoothClientsApiService)(&c.common)
	c.BonjourForwardingApi = (*BonjourForwardingApiService)(&c.common)
	c.BrandingPoliciesApi = (*BrandingPoliciesApiService)(&c.common)
	c.ByClientApi = (*ByClientApiService)(&c.common)
	c.ByDeviceApi = (*ByDeviceApiService)(&c.common)
	c.ByEnergyUsageApi = (*ByEnergyUsageApiService)(&c.common)
	c.ByMetricApi = (*ByMetricApiService)(&c.common)
	c.BySwitchApi = (*BySwitchApiService)(&c.common)
	c.ByUsageApi = (*ByUsageApiService)(&c.common)
	c.ByUtilizationApi = (*ByUtilizationApiService)(&c.common)
	c.BypassActivationLockAttemptsApi = (*BypassActivationLockAttemptsApiService)(&c.common)
	c.CameraApi = (*CameraApiService)(&c.common)
	c.CategoriesApi = (*CategoriesApiService)(&c.common)
	c.CellularFirewallRulesApi = (*CellularFirewallRulesApiService)(&c.common)
	c.CellularGatewayApi = (*CellularGatewayApiService)(&c.common)
	c.CellularUsageHistoryApi = (*CellularUsageHistoryApiService)(&c.common)
	c.CertsApi = (*CertsApiService)(&c.common)
	c.ChannelUtilizationApi = (*ChannelUtilizationApiService)(&c.common)
	c.ChannelUtilizationHistoryApi = (*ChannelUtilizationHistoryApiService)(&c.common)
	c.ClaimApi = (*ClaimApiService)(&c.common)
	c.ClientCountHistoryApi = (*ClientCountHistoryApiService)(&c.common)
	c.ClientsApi = (*ClientsApiService)(&c.common)
	c.ConfigTemplatesApi = (*ConfigTemplatesApiService)(&c.common)
	c.ConfigurationChangesApi = (*ConfigurationChangesApiService)(&c.common)
	c.ConfigureApi = (*ConfigureApiService)(&c.common)
	c.ConnectionStatsApi = (*ConnectionStatsApiService)(&c.common)
	c.ConnectivityApi = (*ConnectivityApiService)(&c.common)
	c.ConnectivityEventsApi = (*ConnectivityEventsApiService)(&c.common)
	c.ConnectivityMonitoringDestinationsApi = (*ConnectivityMonitoringDestinationsApiService)(&c.common)
	c.ContentFilteringApi = (*ContentFilteringApiService)(&c.common)
	c.CurrentApi = (*CurrentApiService)(&c.common)
	c.CustomAnalyticsApi = (*CustomAnalyticsApiService)(&c.common)
	c.CustomPerformanceClassesApi = (*CustomPerformanceClassesApiService)(&c.common)
	c.DataRateHistoryApi = (*DataRateHistoryApiService)(&c.common)
	c.DelegatedApi = (*DelegatedApiService)(&c.common)
	c.DesktopLogsApi = (*DesktopLogsApiService)(&c.common)
	c.DeviceCommandLogsApi = (*DeviceCommandLogsApiService)(&c.common)
	c.DeviceProfilesApi = (*DeviceProfilesApiService)(&c.common)
	c.DeviceTypeGroupPoliciesApi = (*DeviceTypeGroupPoliciesApiService)(&c.common)
	c.DevicesApi = (*DevicesApiService)(&c.common)
	c.DhcpApi = (*DhcpApiService)(&c.common)
	c.DhcpServerPolicyApi = (*DhcpServerPolicyApiService)(&c.common)
	c.DscpTaggingOptionsApi = (*DscpTaggingOptionsApiService)(&c.common)
	c.DscpToCosMappingsApi = (*DscpToCosMappingsApiService)(&c.common)
	c.EapOverrideApi = (*EapOverrideApiService)(&c.common)
	c.EarlyAccessApi = (*EarlyAccessApiService)(&c.common)
	c.EventTypesApi = (*EventTypesApiService)(&c.common)
	c.EventsApi = (*EventsApiService)(&c.common)
	c.FailedConnectionsApi = (*FailedConnectionsApiService)(&c.common)
	c.FeaturesApi = (*FeaturesApiService)(&c.common)
	c.FieldsApi = (*FieldsApiService)(&c.common)
	c.FirewallApi = (*FirewallApiService)(&c.common)
	c.FirewalledServicesApi = (*FirewalledServicesApiService)(&c.common)
	c.FirmwareApi = (*FirmwareApiService)(&c.common)
	c.FirmwareUpgradesApi = (*FirmwareUpgradesApiService)(&c.common)
	c.FloorPlansApi = (*FloorPlansApiService)(&c.common)
	c.GroupPoliciesApi = (*GroupPoliciesApiService)(&c.common)
	c.GroupsApi = (*GroupsApiService)(&c.common)
	c.HealthApi = (*HealthApiService)(&c.common)
	c.HealthByTimeApi = (*HealthByTimeApiService)(&c.common)
	c.HistoryApi = (*HistoryApiService)(&c.common)
	c.Hotspot20Api = (*Hotspot20ApiService)(&c.common)
	c.HttpServersApi = (*HttpServersApiService)(&c.common)
	c.IdentitiesApi = (*IdentitiesApiService)(&c.common)
	c.IdentityPsksApi = (*IdentityPsksApiService)(&c.common)
	c.IdpsApi = (*IdpsApiService)(&c.common)
	c.InboundCellularFirewallRulesApi = (*InboundCellularFirewallRulesApiService)(&c.common)
	c.InboundFirewallRulesApi = (*InboundFirewallRulesApiService)(&c.common)
	c.InsightApi = (*InsightApiService)(&c.common)
	c.InterfacesApi = (*InterfacesApiService)(&c.common)
	c.IntrusionApi = (*IntrusionApiService)(&c.common)
	c.InventoryApi = (*InventoryApiService)(&c.common)
	c.L3FirewallRulesApi = (*L3FirewallRulesApiService)(&c.common)
	c.L7FirewallRulesApi = (*L7FirewallRulesApiService)(&c.common)
	c.LanApi = (*LanApiService)(&c.common)
	c.LatencyHistoryApi = (*LatencyHistoryApiService)(&c.common)
	c.LatencyStatsApi = (*LatencyStatsApiService)(&c.common)
	c.LatestApi = (*LatestApiService)(&c.common)
	c.LicensesApi = (*LicensesApiService)(&c.common)
	c.LinkAggregationsApi = (*LinkAggregationsApiService)(&c.common)
	c.LinkLayerApi = (*LinkLayerApiService)(&c.common)
	c.LiveApi = (*LiveApiService)(&c.common)
	c.LiveToolsApi = (*LiveToolsApiService)(&c.common)
	c.LldpCdpApi = (*LldpCdpApiService)(&c.common)
	c.LoginSecurityApi = (*LoginSecurityApiService)(&c.common)
	c.LogsApi = (*LogsApiService)(&c.common)
	c.LossAndLatencyHistoryApi = (*LossAndLatencyHistoryApiService)(&c.common)
	c.MalwareApi = (*MalwareApiService)(&c.common)
	c.ManagementInterfaceApi = (*ManagementInterfaceApiService)(&c.common)
	c.ManufacturersApi = (*ManufacturersApiService)(&c.common)
	c.MeApi = (*MeApiService)(&c.common)
	c.MerakiAuthUsersApi = (*MerakiAuthUsersApiService)(&c.common)
	c.MeshStatusesApi = (*MeshStatusesApiService)(&c.common)
	c.ModelsApi = (*ModelsApiService)(&c.common)
	c.MonitorApi = (*MonitorApiService)(&c.common)
	c.MonitoredMediaServersApi = (*MonitoredMediaServersApiService)(&c.common)
	c.MqttBrokersApi = (*MqttBrokersApiService)(&c.common)
	c.MtuApi = (*MtuApiService)(&c.common)
	c.MulticastApi = (*MulticastApiService)(&c.common)
	c.NetflowApi = (*NetflowApiService)(&c.common)
	c.NetworkAdaptersApi = (*NetworkAdaptersApiService)(&c.common)
	c.NetworkHealthApi = (*NetworkHealthApiService)(&c.common)
	c.NetworksApi = (*NetworksApiService)(&c.common)
	c.ObjectDetectionModelsApi = (*ObjectDetectionModelsApiService)(&c.common)
	c.OnboardingApi = (*OnboardingApiService)(&c.common)
	c.OneToManyNatRulesApi = (*OneToManyNatRulesApiService)(&c.common)
	c.OneToOneNatRulesApi = (*OneToOneNatRulesApiService)(&c.common)
	c.OpenapiSpecApi = (*OpenapiSpecApiService)(&c.common)
	c.OptInsApi = (*OptInsApiService)(&c.common)
	c.OrderApi = (*OrderApiService)(&c.common)
	c.OrganizationsApi = (*OrganizationsApiService)(&c.common)
	c.OspfApi = (*OspfApiService)(&c.common)
	c.OverviewApi = (*OverviewApiService)(&c.common)
	c.PacketsApi = (*PacketsApiService)(&c.common)
	c.PayloadTemplatesApi = (*PayloadTemplatesApiService)(&c.common)
	c.PerformanceApi = (*PerformanceApiService)(&c.common)
	c.PerformanceHistoryApi = (*PerformanceHistoryApiService)(&c.common)
	c.PiiApi = (*PiiApiService)(&c.common)
	c.PiiKeysApi = (*PiiKeysApiService)(&c.common)
	c.PingApi = (*PingApiService)(&c.common)
	c.PingDeviceApi = (*PingDeviceApiService)(&c.common)
	c.PoliciesApi = (*PoliciesApiService)(&c.common)
	c.PolicyApi = (*PolicyApiService)(&c.common)
	c.PortForwardingRulesApi = (*PortForwardingRulesApiService)(&c.common)
	c.PortSchedulesApi = (*PortSchedulesApiService)(&c.common)
	c.PortsApi = (*PortsApiService)(&c.common)
	c.PowerModulesApi = (*PowerModulesApiService)(&c.common)
	c.PrefixesApi = (*PrefixesApiService)(&c.common)
	c.PrioritiesApi = (*PrioritiesApiService)(&c.common)
	c.ProfilesApi = (*ProfilesApiService)(&c.common)
	c.QosRulesApi = (*QosRulesApiService)(&c.common)
	c.QualityAndRetentionApi = (*QualityAndRetentionApiService)(&c.common)
	c.QualityRetentionProfilesApi = (*QualityRetentionProfilesApiService)(&c.common)
	c.RadioApi = (*RadioApiService)(&c.common)
	c.ReadingsApi = (*ReadingsApiService)(&c.common)
	c.RecentApi = (*RecentApiService)(&c.common)
	c.RendezvousPointsApi = (*RendezvousPointsApiService)(&c.common)
	c.RequestsApi = (*RequestsApiService)(&c.common)
	c.RestrictionsApi = (*RestrictionsApiService)(&c.common)
	c.RfProfilesApi = (*RfProfilesApiService)(&c.common)
	c.RollbacksApi = (*RollbacksApiService)(&c.common)
	c.RoutingApi = (*RoutingApiService)(&c.common)
	c.RulesApi = (*RulesApiService)(&c.common)
	c.SamlApi = (*SamlApiService)(&c.common)
	c.SamlRolesApi = (*SamlRolesApiService)(&c.common)
	c.SchedulesApi = (*SchedulesApiService)(&c.common)
	c.SearchApi = (*SearchApiService)(&c.common)
	c.SecurityApi = (*SecurityApiService)(&c.common)
	c.SecurityCentersApi = (*SecurityCentersApiService)(&c.common)
	c.SeenApi = (*SeenApiService)(&c.common)
	c.SenseApi = (*SenseApiService)(&c.common)
	c.SensorApi = (*SensorApiService)(&c.common)
	c.ServersApi = (*ServersApiService)(&c.common)
	c.SettingsApi = (*SettingsApiService)(&c.common)
	c.SignalQualityHistoryApi = (*SignalQualityHistoryApiService)(&c.common)
	c.SingleLanApi = (*SingleLanApiService)(&c.common)
	c.SiteToSiteVpnApi = (*SiteToSiteVpnApiService)(&c.common)
	c.SmApi = (*SmApiService)(&c.common)
	c.SmDevicesForKeyApi = (*SmDevicesForKeyApiService)(&c.common)
	c.SmOwnersForKeyApi = (*SmOwnersForKeyApiService)(&c.common)
	c.SnmpApi = (*SnmpApiService)(&c.common)
	c.SoftwaresApi = (*SoftwaresApiService)(&c.common)
	c.SplashApi = (*SplashApiService)(&c.common)
	c.SplashAuthorizationStatusApi = (*SplashAuthorizationStatusApiService)(&c.common)
	c.SplashLoginAttemptsApi = (*SplashLoginAttemptsApiService)(&c.common)
	c.SsidsApi = (*SsidsApiService)(&c.common)
	c.StacksApi = (*StacksApiService)(&c.common)
	c.StagedApi = (*StagedApiService)(&c.common)
	c.StagesApi = (*StagesApiService)(&c.common)
	c.StaticRoutesApi = (*StaticRoutesApiService)(&c.common)
	c.StaticsApi = (*StaticsApiService)(&c.common)
	c.StatsApi = (*StatsApiService)(&c.common)
	c.StatusApi = (*StatusApiService)(&c.common)
	c.StatusesApi = (*StatusesApiService)(&c.common)
	c.StormControlApi = (*StormControlApiService)(&c.common)
	c.StpApi = (*StpApiService)(&c.common)
	c.SubnetPoolApi = (*SubnetPoolApiService)(&c.common)
	c.SubnetsApi = (*SubnetsApiService)(&c.common)
	c.SummaryApi = (*SummaryApiService)(&c.common)
	c.SwitchApi = (*SwitchApiService)(&c.common)
	c.SwitchesApi = (*SwitchesApiService)(&c.common)
	c.SyslogServersApi = (*SyslogServersApiService)(&c.common)
	c.TargetGroupsApi = (*TargetGroupsApiService)(&c.common)
	c.ThirdPartyVPNPeersApi = (*ThirdPartyVPNPeersApiService)(&c.common)
	c.TopApi = (*TopApiService)(&c.common)
	c.TopologyApi = (*TopologyApiService)(&c.common)
	c.TrafficApi = (*TrafficApiService)(&c.common)
	c.TrafficAnalysisApi = (*TrafficAnalysisApiService)(&c.common)
	c.TrafficHistoryApi = (*TrafficHistoryApiService)(&c.common)
	c.TrafficShapingApi = (*TrafficShapingApiService)(&c.common)
	c.TrustedAccessConfigsApi = (*TrustedAccessConfigsApiService)(&c.common)
	c.TrustedServersApi = (*TrustedServersApiService)(&c.common)
	c.UpgradesApi = (*UpgradesApiService)(&c.common)
	c.UplinkApi = (*UplinkApiService)(&c.common)
	c.UplinkBandwidthApi = (*UplinkBandwidthApiService)(&c.common)
	c.UplinkSelectionApi = (*UplinkSelectionApiService)(&c.common)
	c.UplinksApi = (*UplinksApiService)(&c.common)
	c.UplinksLossAndLatencyApi = (*UplinksLossAndLatencyApiService)(&c.common)
	c.UsageHistoriesApi = (*UsageHistoriesApiService)(&c.common)
	c.UsageHistoryApi = (*UsageHistoryApiService)(&c.common)
	c.UserAccessDevicesApi = (*UserAccessDevicesApiService)(&c.common)
	c.UsersApi = (*UsersApiService)(&c.common)
	c.V4Api = (*V4ApiService)(&c.common)
	c.VideoApi = (*VideoApiService)(&c.common)
	c.VideoLinkApi = (*VideoLinkApiService)(&c.common)
	c.VlanAssignmentsApi = (*VlanAssignmentsApiService)(&c.common)
	c.VlansApi = (*VlansApiService)(&c.common)
	c.VmxApi = (*VmxApiService)(&c.common)
	c.VpnApi = (*VpnApiService)(&c.common)
	c.VpnFirewallRulesApi = (*VpnFirewallRulesApiService)(&c.common)
	c.VppAccountsApi = (*VppAccountsApiService)(&c.common)
	c.WarmSpareApi = (*WarmSpareApiService)(&c.common)
	c.WarningsApi = (*WarningsApiService)(&c.common)
	c.WebhookLogsApi = (*WebhookLogsApiService)(&c.common)
	c.WebhookTestsApi = (*WebhookTestsApiService)(&c.common)
	c.WebhooksApi = (*WebhooksApiService)(&c.common)
	c.WirelessApi = (*WirelessApiService)(&c.common)
	c.WirelessProfilesApi = (*WirelessProfilesApiService)(&c.common)
	c.WlanListsApi = (*WlanListsApiService)(&c.common)
	c.ZonesApi = (*ZonesApiService)(&c.common)

	return c
}

func atoi(in string) (int, error) {
	return strconv.Atoi(in)
}

// selectHeaderContentType select a content type from the available list.
func selectHeaderContentType(contentTypes []string) string {
	if len(contentTypes) == 0 {
		return ""
	}
	if contains(contentTypes, "application/json") {
		return "application/json"
	}
	return contentTypes[0] // use the first content type specified in 'consumes'
}

// selectHeaderAccept join all accept types and return
func selectHeaderAccept(accepts []string) string {
	if len(accepts) == 0 {
		return ""
	}

	if contains(accepts, "application/json") {
		return "application/json"
	}

	return strings.Join(accepts, ",")
}

// contains is a case insensitive match, finding needle in a haystack
func contains(haystack []string, needle string) bool {
	for _, a := range haystack {
		if strings.ToLower(a) == strings.ToLower(needle) {
			return true
		}
	}
	return false
}

// Verify optional parameters are of the correct type.
func typeCheckParameter(obj interface{}, expected string, name string) error {
	// Make sure there is an object.
	if obj == nil {
		return nil
	}

	// Check the type is as expected.
	if reflect.TypeOf(obj).String() != expected {
		return fmt.Errorf("Expected %s to be of type %s but received %s.", name, expected, reflect.TypeOf(obj).String())
	}
	return nil
}

// parameterToString convert interface{} parameters to string, using a delimiter if format is provided.
func parameterToString(obj interface{}, collectionFormat string) string {
	var delimiter string

	switch collectionFormat {
	case "pipes":
		delimiter = "|"
	case "ssv":
		delimiter = " "
	case "tsv":
		delimiter = "\t"
	case "csv":
		delimiter = ","
	}

	if reflect.TypeOf(obj).Kind() == reflect.Slice {
		return strings.Trim(strings.Replace(fmt.Sprint(obj), " ", delimiter, -1), "[]")
	} else if t, ok := obj.(time.Time); ok {
		return t.Format(time.RFC3339)
	}

	return fmt.Sprintf("%v", obj)
}

// helper for converting interface{} parameters to json strings
func parameterToJson(obj interface{}) (string, error) {
	jsonBuf, err := json.Marshal(obj)
	if err != nil {
		return "", err
	}
	return string(jsonBuf), err
}

// callAPI do the request.
func (c *APIClient) callAPI(request *http.Request) (*http.Response, error) {
	if c.cfg.Debug {
		dump, err := httputil.DumpRequestOut(request, true)
		if err != nil {
			return nil, err
		}
		log.Printf("\n%s\n", string(dump))
	}

	resp, err := c.cfg.HTTPClient.Do(request)
	if err != nil {
		return resp, err
	}

	if c.cfg.Debug {
		dump, err := httputil.DumpResponse(resp, true)
		if err != nil {
			return resp, err
		}
		log.Printf("\n%s\n", string(dump))
	}
	return resp, err
}

// Allow modification of underlying config for alternate implementations and testing
// Caution: modifying the configuration while live can cause data races and potentially unwanted behavior
func (c *APIClient) GetConfig() *Configuration {
	return c.cfg
}

// prepareRequest build the request
func (c *APIClient) prepareRequest(
	ctx context.Context,
	path string, method string,
	postBody interface{},
	headerParams map[string]string,
	queryParams url.Values,
	formParams url.Values,
	formFileName string,
	fileName string,
	fileBytes []byte) (localVarRequest *http.Request, err error) {

	var body *bytes.Buffer

	// Detect postBody type and post.
	if postBody != nil {
		contentType := headerParams["Content-Type"]
		if contentType == "" {
			contentType = detectContentType(postBody)
			headerParams["Content-Type"] = contentType
		}

		body, err = setBody(postBody, contentType)
		if err != nil {
			return nil, err
		}
	}

	// add form parameters and file if available.
	if strings.HasPrefix(headerParams["Content-Type"], "multipart/form-data") && len(formParams) > 0 || (len(fileBytes) > 0 && fileName != "") {
		if body != nil {
			return nil, errors.New("Cannot specify postBody and multipart form at the same time.")
		}
		body = &bytes.Buffer{}
		w := multipart.NewWriter(body)

		for k, v := range formParams {
			for _, iv := range v {
				if strings.HasPrefix(k, "@") { // file
					err = addFile(w, k[1:], iv)
					if err != nil {
						return nil, err
					}
				} else { // form value
					w.WriteField(k, iv)
				}
			}
		}
		if len(fileBytes) > 0 && fileName != "" {
			w.Boundary()
			//_, fileNm := filepath.Split(fileName)
			part, err := w.CreateFormFile(formFileName, filepath.Base(fileName))
			if err != nil {
				return nil, err
			}
			_, err = part.Write(fileBytes)
			if err != nil {
				return nil, err
			}
		}

		// Set the Boundary in the Content-Type
		headerParams["Content-Type"] = w.FormDataContentType()

		// Set Content-Length
		headerParams["Content-Length"] = fmt.Sprintf("%d", body.Len())
		w.Close()
	}

	if strings.HasPrefix(headerParams["Content-Type"], "application/x-www-form-urlencoded") && len(formParams) > 0 {
		if body != nil {
			return nil, errors.New("Cannot specify postBody and x-www-form-urlencoded form at the same time.")
		}
		body = &bytes.Buffer{}
		body.WriteString(formParams.Encode())
		// Set Content-Length
		headerParams["Content-Length"] = fmt.Sprintf("%d", body.Len())
	}

	// Setup path and query parameters
	url, err := url.Parse(path)
	if err != nil {
		return nil, err
	}

	// Override request host, if applicable
	if c.cfg.Host != "" {
		url.Host = c.cfg.Host
	}

	// Override request scheme, if applicable
	if c.cfg.Scheme != "" {
		url.Scheme = c.cfg.Scheme
	}

	// Adding Query Param
	query := url.Query()
	for k, v := range queryParams {
		for _, iv := range v {
			query.Add(k, iv)
		}
	}

	// Encode the parameters.
	url.RawQuery = query.Encode()

	// Generate a new request
	if body != nil {
		localVarRequest, err = http.NewRequest(method, url.String(), body)
	} else {
		localVarRequest, err = http.NewRequest(method, url.String(), nil)
	}
	if err != nil {
		return nil, err
	}

	// add header parameters, if any
	if len(headerParams) > 0 {
		headers := http.Header{}
		for h, v := range headerParams {
			headers.Set(h, v)
		}
		localVarRequest.Header = headers
	}

	// Add the user agent to the request.
	localVarRequest.Header.Add("User-Agent", c.cfg.UserAgent)

	if ctx != nil {
		// add context to the request
		localVarRequest = localVarRequest.WithContext(ctx)

		// Walk through any authentication.

		// OAuth2 authentication
		if tok, ok := ctx.Value(ContextOAuth2).(oauth2.TokenSource); ok {
			// We were able to grab an oauth2 token from the context
			var latestToken *oauth2.Token
			if latestToken, err = tok.Token(); err != nil {
				return nil, err
			}

			latestToken.SetAuthHeader(localVarRequest)
		}

		// Basic HTTP Authentication
		if auth, ok := ctx.Value(ContextBasicAuth).(BasicAuth); ok {
			localVarRequest.SetBasicAuth(auth.UserName, auth.Password)
		}

		// AccessToken Authentication
		if auth, ok := ctx.Value(ContextAccessToken).(string); ok {
			localVarRequest.Header.Add("Authorization", "Bearer "+auth)
		}

	}

	for header, value := range c.cfg.DefaultHeader {
		localVarRequest.Header.Add(header, value)
	}
	return localVarRequest, nil
}

func (c *APIClient) decode(v interface{}, b []byte, contentType string) (err error) {
	if len(b) == 0 {
		return nil
	}
	if s, ok := v.(*string); ok {
		*s = string(b)
		return nil
	}
	if f, ok := v.(**os.File); ok {
		*f, err = ioutil.TempFile("", "HttpClientFile")
		if err != nil {
			return
		}
		_, err = (*f).Write(b)
		if err != nil {
			return
		}
		_, err = (*f).Seek(0, io.SeekStart)
		return
	}
	if xmlCheck.MatchString(contentType) {
		if err = xml.Unmarshal(b, v); err != nil {
			return err
		}
		return nil
	}
	if jsonCheck.MatchString(contentType) {
		if actualObj, ok := v.(interface{ GetActualInstance() interface{} }); ok { // oneOf, anyOf schemas
			if unmarshalObj, ok := actualObj.(interface{ UnmarshalJSON([]byte) error }); ok { // make sure it has UnmarshalJSON defined
				if err = unmarshalObj.UnmarshalJSON(b); err != nil {
					return err
				}
			} else {
				return errors.New("Unknown type with GetActualInstance but no unmarshalObj.UnmarshalJSON defined")
			}
		} else if err = json.Unmarshal(b, v); err != nil { // simple model
			return err
		}
		return nil
	}
	return errors.New("undefined response type")
}

// Add a file to the multipart request
func addFile(w *multipart.Writer, fieldName, path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	part, err := w.CreateFormFile(fieldName, filepath.Base(path))
	if err != nil {
		return err
	}
	_, err = io.Copy(part, file)

	return err
}

// Prevent trying to import "fmt"
func reportError(format string, a ...interface{}) error {
	return fmt.Errorf(format, a...)
}

// Set request body from an interface{}
func setBody(body interface{}, contentType string) (bodyBuf *bytes.Buffer, err error) {
	if bodyBuf == nil {
		bodyBuf = &bytes.Buffer{}
	}

	if reader, ok := body.(io.Reader); ok {
		_, err = bodyBuf.ReadFrom(reader)
	} else if fp, ok := body.(**os.File); ok {
		_, err = bodyBuf.ReadFrom(*fp)
	} else if b, ok := body.([]byte); ok {
		_, err = bodyBuf.Write(b)
	} else if s, ok := body.(string); ok {
		_, err = bodyBuf.WriteString(s)
	} else if s, ok := body.(*string); ok {
		_, err = bodyBuf.WriteString(*s)
	} else if jsonCheck.MatchString(contentType) {
		err = json.NewEncoder(bodyBuf).Encode(body)
	} else if xmlCheck.MatchString(contentType) {
		err = xml.NewEncoder(bodyBuf).Encode(body)
	}

	if err != nil {
		return nil, err
	}

	if bodyBuf.Len() == 0 {
		err = fmt.Errorf("Invalid body type %s\n", contentType)
		return nil, err
	}
	return bodyBuf, nil
}

// detectContentType method is used to figure out `Request.Body` content type for request header
func detectContentType(body interface{}) string {
	contentType := "text/plain; charset=utf-8"
	kind := reflect.TypeOf(body).Kind()

	switch kind {
	case reflect.Struct, reflect.Map, reflect.Ptr:
		contentType = "application/json; charset=utf-8"
	case reflect.String:
		contentType = "text/plain; charset=utf-8"
	default:
		if b, ok := body.([]byte); ok {
			contentType = http.DetectContentType(b)
		} else if kind == reflect.Slice {
			contentType = "application/json; charset=utf-8"
		}
	}

	return contentType
}

// Ripped from https://github.com/gregjones/httpcache/blob/master/httpcache.go
type cacheControl map[string]string

func parseCacheControl(headers http.Header) cacheControl {
	cc := cacheControl{}
	ccHeader := headers.Get("Cache-Control")
	for _, part := range strings.Split(ccHeader, ",") {
		part = strings.Trim(part, " ")
		if part == "" {
			continue
		}
		if strings.ContainsRune(part, '=') {
			keyval := strings.Split(part, "=")
			cc[strings.Trim(keyval[0], " ")] = strings.Trim(keyval[1], ",")
		} else {
			cc[part] = ""
		}
	}
	return cc
}

// CacheExpires helper function to determine remaining time before repeating a request.
func CacheExpires(r *http.Response) time.Time {
	// Figure out when the cache expires.
	var expires time.Time
	now, err := time.Parse(time.RFC1123, r.Header.Get("date"))
	if err != nil {
		return time.Now()
	}
	respCacheControl := parseCacheControl(r.Header)

	if maxAge, ok := respCacheControl["max-age"]; ok {
		lifetime, err := time.ParseDuration(maxAge + "s")
		if err != nil {
			expires = now
		} else {
			expires = now.Add(lifetime)
		}
	} else {
		expiresHeader := r.Header.Get("Expires")
		if expiresHeader != "" {
			expires, err = time.Parse(time.RFC1123, expiresHeader)
			if err != nil {
				expires = now
			}
		}
	}
	return expires
}

func strlen(s string) int {
	return utf8.RuneCountInString(s)
}

// GenericOpenAPIError Provides access to the body, error and model on returned errors.
type GenericOpenAPIError struct {
	body  []byte
	error string
	model interface{}
}

// Error returns non-empty string if there was an error.
func (e GenericOpenAPIError) Error() string {
	return e.error
}

// Body returns the raw bytes of the response
func (e GenericOpenAPIError) Body() []byte {
	return e.body
}

// Model returns the unpacked model of the error
func (e GenericOpenAPIError) Model() interface{} {
	return e.model
}
