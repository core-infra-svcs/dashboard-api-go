# InlineObject217

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MeshingEnabled** | Pointer to **bool** | Toggle for enabling or disabling meshing in a network | [optional] 
**Ipv6BridgeEnabled** | Pointer to **bool** | Toggle for enabling or disabling IPv6 bridging in a network (Note: if enabled, SSIDs must also be configured to use bridge mode) | [optional] 
**LocationAnalyticsEnabled** | Pointer to **bool** | Toggle for enabling or disabling location analytics for your network | [optional] 
**UpgradeStrategy** | Pointer to **string** | The default strategy that network devices will use to perform an upgrade. Requires firmware version MR 26.8 or higher. | [optional] 
**Upgrade** | Pointer to [**InlineResponse200245Upgrade**](InlineResponse200245Upgrade.md) |  | [optional] 
**LedLightsOn** | Pointer to **bool** | Toggle for enabling or disabling LED lights on all APs in the network (making them run dark) | [optional] 
**MulticastToUnicastConversion** | Pointer to [**InlineResponse200245MulticastToUnicastConversion**](InlineResponse200245MulticastToUnicastConversion.md) |  | [optional] 
**NamedVlans** | Pointer to [**NetworksNetworkIdWirelessSettingsNamedVlans**](NetworksNetworkIdWirelessSettingsNamedVlans.md) |  | [optional] 

## Methods

### NewInlineObject217

`func NewInlineObject217() *InlineObject217`

NewInlineObject217 instantiates a new InlineObject217 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject217WithDefaults

`func NewInlineObject217WithDefaults() *InlineObject217`

NewInlineObject217WithDefaults instantiates a new InlineObject217 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeshingEnabled

`func (o *InlineObject217) GetMeshingEnabled() bool`

GetMeshingEnabled returns the MeshingEnabled field if non-nil, zero value otherwise.

### GetMeshingEnabledOk

`func (o *InlineObject217) GetMeshingEnabledOk() (*bool, bool)`

GetMeshingEnabledOk returns a tuple with the MeshingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeshingEnabled

`func (o *InlineObject217) SetMeshingEnabled(v bool)`

SetMeshingEnabled sets MeshingEnabled field to given value.

### HasMeshingEnabled

`func (o *InlineObject217) HasMeshingEnabled() bool`

HasMeshingEnabled returns a boolean if a field has been set.

### GetIpv6BridgeEnabled

`func (o *InlineObject217) GetIpv6BridgeEnabled() bool`

GetIpv6BridgeEnabled returns the Ipv6BridgeEnabled field if non-nil, zero value otherwise.

### GetIpv6BridgeEnabledOk

`func (o *InlineObject217) GetIpv6BridgeEnabledOk() (*bool, bool)`

GetIpv6BridgeEnabledOk returns a tuple with the Ipv6BridgeEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6BridgeEnabled

`func (o *InlineObject217) SetIpv6BridgeEnabled(v bool)`

SetIpv6BridgeEnabled sets Ipv6BridgeEnabled field to given value.

### HasIpv6BridgeEnabled

`func (o *InlineObject217) HasIpv6BridgeEnabled() bool`

HasIpv6BridgeEnabled returns a boolean if a field has been set.

### GetLocationAnalyticsEnabled

`func (o *InlineObject217) GetLocationAnalyticsEnabled() bool`

GetLocationAnalyticsEnabled returns the LocationAnalyticsEnabled field if non-nil, zero value otherwise.

### GetLocationAnalyticsEnabledOk

`func (o *InlineObject217) GetLocationAnalyticsEnabledOk() (*bool, bool)`

GetLocationAnalyticsEnabledOk returns a tuple with the LocationAnalyticsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationAnalyticsEnabled

`func (o *InlineObject217) SetLocationAnalyticsEnabled(v bool)`

SetLocationAnalyticsEnabled sets LocationAnalyticsEnabled field to given value.

### HasLocationAnalyticsEnabled

`func (o *InlineObject217) HasLocationAnalyticsEnabled() bool`

HasLocationAnalyticsEnabled returns a boolean if a field has been set.

### GetUpgradeStrategy

`func (o *InlineObject217) GetUpgradeStrategy() string`

GetUpgradeStrategy returns the UpgradeStrategy field if non-nil, zero value otherwise.

### GetUpgradeStrategyOk

`func (o *InlineObject217) GetUpgradeStrategyOk() (*string, bool)`

GetUpgradeStrategyOk returns a tuple with the UpgradeStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeStrategy

`func (o *InlineObject217) SetUpgradeStrategy(v string)`

SetUpgradeStrategy sets UpgradeStrategy field to given value.

### HasUpgradeStrategy

`func (o *InlineObject217) HasUpgradeStrategy() bool`

HasUpgradeStrategy returns a boolean if a field has been set.

### GetUpgrade

`func (o *InlineObject217) GetUpgrade() InlineResponse200245Upgrade`

GetUpgrade returns the Upgrade field if non-nil, zero value otherwise.

### GetUpgradeOk

`func (o *InlineObject217) GetUpgradeOk() (*InlineResponse200245Upgrade, bool)`

GetUpgradeOk returns a tuple with the Upgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgrade

`func (o *InlineObject217) SetUpgrade(v InlineResponse200245Upgrade)`

SetUpgrade sets Upgrade field to given value.

### HasUpgrade

`func (o *InlineObject217) HasUpgrade() bool`

HasUpgrade returns a boolean if a field has been set.

### GetLedLightsOn

`func (o *InlineObject217) GetLedLightsOn() bool`

GetLedLightsOn returns the LedLightsOn field if non-nil, zero value otherwise.

### GetLedLightsOnOk

`func (o *InlineObject217) GetLedLightsOnOk() (*bool, bool)`

GetLedLightsOnOk returns a tuple with the LedLightsOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLedLightsOn

`func (o *InlineObject217) SetLedLightsOn(v bool)`

SetLedLightsOn sets LedLightsOn field to given value.

### HasLedLightsOn

`func (o *InlineObject217) HasLedLightsOn() bool`

HasLedLightsOn returns a boolean if a field has been set.

### GetMulticastToUnicastConversion

`func (o *InlineObject217) GetMulticastToUnicastConversion() InlineResponse200245MulticastToUnicastConversion`

GetMulticastToUnicastConversion returns the MulticastToUnicastConversion field if non-nil, zero value otherwise.

### GetMulticastToUnicastConversionOk

`func (o *InlineObject217) GetMulticastToUnicastConversionOk() (*InlineResponse200245MulticastToUnicastConversion, bool)`

GetMulticastToUnicastConversionOk returns a tuple with the MulticastToUnicastConversion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMulticastToUnicastConversion

`func (o *InlineObject217) SetMulticastToUnicastConversion(v InlineResponse200245MulticastToUnicastConversion)`

SetMulticastToUnicastConversion sets MulticastToUnicastConversion field to given value.

### HasMulticastToUnicastConversion

`func (o *InlineObject217) HasMulticastToUnicastConversion() bool`

HasMulticastToUnicastConversion returns a boolean if a field has been set.

### GetNamedVlans

`func (o *InlineObject217) GetNamedVlans() NetworksNetworkIdWirelessSettingsNamedVlans`

GetNamedVlans returns the NamedVlans field if non-nil, zero value otherwise.

### GetNamedVlansOk

`func (o *InlineObject217) GetNamedVlansOk() (*NetworksNetworkIdWirelessSettingsNamedVlans, bool)`

GetNamedVlansOk returns a tuple with the NamedVlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamedVlans

`func (o *InlineObject217) SetNamedVlans(v NetworksNetworkIdWirelessSettingsNamedVlans)`

SetNamedVlans sets NamedVlans field to given value.

### HasNamedVlans

`func (o *InlineObject217) HasNamedVlans() bool`

HasNamedVlans returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


