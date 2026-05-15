# InlineResponse200233

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MeshingEnabled** | Pointer to **bool** | Toggle for enabling or disabling meshing in a network | [optional] 
**Ipv6BridgeEnabled** | Pointer to **bool** | Toggle for enabling or disabling IPv6 bridging in a network (Note: if enabled, SSIDs must also be configured to use bridge mode) | [optional] 
**LocationAnalyticsEnabled** | Pointer to **bool** | Toggle for enabling or disabling location analytics for your network | [optional] 
**UpgradeStrategy** | Pointer to **string** | The default strategy that network devices will use to perform an upgrade. Requires firmware version MR 26.8 or higher. | [optional] 
**Upgrade** | Pointer to [**InlineResponse200233Upgrade**](InlineResponse200233Upgrade.md) |  | [optional] 
**LedLightsOn** | Pointer to **bool** | Toggle for enabling or disabling LED lights on all APs in the network (making them run dark) | [optional] 
**MulticastToUnicastConversion** | Pointer to [**InlineResponse200233MulticastToUnicastConversion**](InlineResponse200233MulticastToUnicastConversion.md) |  | [optional] 
**NamedVlans** | Pointer to [**InlineResponse200233NamedVlans**](InlineResponse200233NamedVlans.md) |  | [optional] 
**RegulatoryDomain** | Pointer to [**InlineResponse200233RegulatoryDomain**](InlineResponse200233RegulatoryDomain.md) |  | [optional] 

## Methods

### NewInlineResponse200233

`func NewInlineResponse200233() *InlineResponse200233`

NewInlineResponse200233 instantiates a new InlineResponse200233 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200233WithDefaults

`func NewInlineResponse200233WithDefaults() *InlineResponse200233`

NewInlineResponse200233WithDefaults instantiates a new InlineResponse200233 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeshingEnabled

`func (o *InlineResponse200233) GetMeshingEnabled() bool`

GetMeshingEnabled returns the MeshingEnabled field if non-nil, zero value otherwise.

### GetMeshingEnabledOk

`func (o *InlineResponse200233) GetMeshingEnabledOk() (*bool, bool)`

GetMeshingEnabledOk returns a tuple with the MeshingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeshingEnabled

`func (o *InlineResponse200233) SetMeshingEnabled(v bool)`

SetMeshingEnabled sets MeshingEnabled field to given value.

### HasMeshingEnabled

`func (o *InlineResponse200233) HasMeshingEnabled() bool`

HasMeshingEnabled returns a boolean if a field has been set.

### GetIpv6BridgeEnabled

`func (o *InlineResponse200233) GetIpv6BridgeEnabled() bool`

GetIpv6BridgeEnabled returns the Ipv6BridgeEnabled field if non-nil, zero value otherwise.

### GetIpv6BridgeEnabledOk

`func (o *InlineResponse200233) GetIpv6BridgeEnabledOk() (*bool, bool)`

GetIpv6BridgeEnabledOk returns a tuple with the Ipv6BridgeEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6BridgeEnabled

`func (o *InlineResponse200233) SetIpv6BridgeEnabled(v bool)`

SetIpv6BridgeEnabled sets Ipv6BridgeEnabled field to given value.

### HasIpv6BridgeEnabled

`func (o *InlineResponse200233) HasIpv6BridgeEnabled() bool`

HasIpv6BridgeEnabled returns a boolean if a field has been set.

### GetLocationAnalyticsEnabled

`func (o *InlineResponse200233) GetLocationAnalyticsEnabled() bool`

GetLocationAnalyticsEnabled returns the LocationAnalyticsEnabled field if non-nil, zero value otherwise.

### GetLocationAnalyticsEnabledOk

`func (o *InlineResponse200233) GetLocationAnalyticsEnabledOk() (*bool, bool)`

GetLocationAnalyticsEnabledOk returns a tuple with the LocationAnalyticsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationAnalyticsEnabled

`func (o *InlineResponse200233) SetLocationAnalyticsEnabled(v bool)`

SetLocationAnalyticsEnabled sets LocationAnalyticsEnabled field to given value.

### HasLocationAnalyticsEnabled

`func (o *InlineResponse200233) HasLocationAnalyticsEnabled() bool`

HasLocationAnalyticsEnabled returns a boolean if a field has been set.

### GetUpgradeStrategy

`func (o *InlineResponse200233) GetUpgradeStrategy() string`

GetUpgradeStrategy returns the UpgradeStrategy field if non-nil, zero value otherwise.

### GetUpgradeStrategyOk

`func (o *InlineResponse200233) GetUpgradeStrategyOk() (*string, bool)`

GetUpgradeStrategyOk returns a tuple with the UpgradeStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeStrategy

`func (o *InlineResponse200233) SetUpgradeStrategy(v string)`

SetUpgradeStrategy sets UpgradeStrategy field to given value.

### HasUpgradeStrategy

`func (o *InlineResponse200233) HasUpgradeStrategy() bool`

HasUpgradeStrategy returns a boolean if a field has been set.

### GetUpgrade

`func (o *InlineResponse200233) GetUpgrade() InlineResponse200233Upgrade`

GetUpgrade returns the Upgrade field if non-nil, zero value otherwise.

### GetUpgradeOk

`func (o *InlineResponse200233) GetUpgradeOk() (*InlineResponse200233Upgrade, bool)`

GetUpgradeOk returns a tuple with the Upgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgrade

`func (o *InlineResponse200233) SetUpgrade(v InlineResponse200233Upgrade)`

SetUpgrade sets Upgrade field to given value.

### HasUpgrade

`func (o *InlineResponse200233) HasUpgrade() bool`

HasUpgrade returns a boolean if a field has been set.

### GetLedLightsOn

`func (o *InlineResponse200233) GetLedLightsOn() bool`

GetLedLightsOn returns the LedLightsOn field if non-nil, zero value otherwise.

### GetLedLightsOnOk

`func (o *InlineResponse200233) GetLedLightsOnOk() (*bool, bool)`

GetLedLightsOnOk returns a tuple with the LedLightsOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLedLightsOn

`func (o *InlineResponse200233) SetLedLightsOn(v bool)`

SetLedLightsOn sets LedLightsOn field to given value.

### HasLedLightsOn

`func (o *InlineResponse200233) HasLedLightsOn() bool`

HasLedLightsOn returns a boolean if a field has been set.

### GetMulticastToUnicastConversion

`func (o *InlineResponse200233) GetMulticastToUnicastConversion() InlineResponse200233MulticastToUnicastConversion`

GetMulticastToUnicastConversion returns the MulticastToUnicastConversion field if non-nil, zero value otherwise.

### GetMulticastToUnicastConversionOk

`func (o *InlineResponse200233) GetMulticastToUnicastConversionOk() (*InlineResponse200233MulticastToUnicastConversion, bool)`

GetMulticastToUnicastConversionOk returns a tuple with the MulticastToUnicastConversion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMulticastToUnicastConversion

`func (o *InlineResponse200233) SetMulticastToUnicastConversion(v InlineResponse200233MulticastToUnicastConversion)`

SetMulticastToUnicastConversion sets MulticastToUnicastConversion field to given value.

### HasMulticastToUnicastConversion

`func (o *InlineResponse200233) HasMulticastToUnicastConversion() bool`

HasMulticastToUnicastConversion returns a boolean if a field has been set.

### GetNamedVlans

`func (o *InlineResponse200233) GetNamedVlans() InlineResponse200233NamedVlans`

GetNamedVlans returns the NamedVlans field if non-nil, zero value otherwise.

### GetNamedVlansOk

`func (o *InlineResponse200233) GetNamedVlansOk() (*InlineResponse200233NamedVlans, bool)`

GetNamedVlansOk returns a tuple with the NamedVlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamedVlans

`func (o *InlineResponse200233) SetNamedVlans(v InlineResponse200233NamedVlans)`

SetNamedVlans sets NamedVlans field to given value.

### HasNamedVlans

`func (o *InlineResponse200233) HasNamedVlans() bool`

HasNamedVlans returns a boolean if a field has been set.

### GetRegulatoryDomain

`func (o *InlineResponse200233) GetRegulatoryDomain() InlineResponse200233RegulatoryDomain`

GetRegulatoryDomain returns the RegulatoryDomain field if non-nil, zero value otherwise.

### GetRegulatoryDomainOk

`func (o *InlineResponse200233) GetRegulatoryDomainOk() (*InlineResponse200233RegulatoryDomain, bool)`

GetRegulatoryDomainOk returns a tuple with the RegulatoryDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegulatoryDomain

`func (o *InlineResponse200233) SetRegulatoryDomain(v InlineResponse200233RegulatoryDomain)`

SetRegulatoryDomain sets RegulatoryDomain field to given value.

### HasRegulatoryDomain

`func (o *InlineResponse200233) HasRegulatoryDomain() bool`

HasRegulatoryDomain returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


