# InlineObject193

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MeshingEnabled** | Pointer to **bool** | Toggle for enabling or disabling meshing in a network | [optional] 
**Ipv6BridgeEnabled** | Pointer to **bool** | Toggle for enabling or disabling IPv6 bridging in a network (Note: if enabled, SSIDs must also be configured to use bridge mode) | [optional] 
**LocationAnalyticsEnabled** | Pointer to **bool** | Toggle for enabling or disabling location analytics for your network | [optional] 
**UpgradeStrategy** | Pointer to **string** | The default strategy that network devices will use to perform an upgrade. Requires firmware version MR 26.8 or higher. | [optional] 
**LedLightsOn** | Pointer to **bool** | Toggle for enabling or disabling LED lights on all APs in the network (making them run dark) | [optional] 
**NamedVlans** | Pointer to [**NetworksNetworkIdWirelessSettingsNamedVlans**](NetworksNetworkIdWirelessSettingsNamedVlans.md) |  | [optional] 

## Methods

### NewInlineObject193

`func NewInlineObject193() *InlineObject193`

NewInlineObject193 instantiates a new InlineObject193 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject193WithDefaults

`func NewInlineObject193WithDefaults() *InlineObject193`

NewInlineObject193WithDefaults instantiates a new InlineObject193 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeshingEnabled

`func (o *InlineObject193) GetMeshingEnabled() bool`

GetMeshingEnabled returns the MeshingEnabled field if non-nil, zero value otherwise.

### GetMeshingEnabledOk

`func (o *InlineObject193) GetMeshingEnabledOk() (*bool, bool)`

GetMeshingEnabledOk returns a tuple with the MeshingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeshingEnabled

`func (o *InlineObject193) SetMeshingEnabled(v bool)`

SetMeshingEnabled sets MeshingEnabled field to given value.

### HasMeshingEnabled

`func (o *InlineObject193) HasMeshingEnabled() bool`

HasMeshingEnabled returns a boolean if a field has been set.

### GetIpv6BridgeEnabled

`func (o *InlineObject193) GetIpv6BridgeEnabled() bool`

GetIpv6BridgeEnabled returns the Ipv6BridgeEnabled field if non-nil, zero value otherwise.

### GetIpv6BridgeEnabledOk

`func (o *InlineObject193) GetIpv6BridgeEnabledOk() (*bool, bool)`

GetIpv6BridgeEnabledOk returns a tuple with the Ipv6BridgeEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6BridgeEnabled

`func (o *InlineObject193) SetIpv6BridgeEnabled(v bool)`

SetIpv6BridgeEnabled sets Ipv6BridgeEnabled field to given value.

### HasIpv6BridgeEnabled

`func (o *InlineObject193) HasIpv6BridgeEnabled() bool`

HasIpv6BridgeEnabled returns a boolean if a field has been set.

### GetLocationAnalyticsEnabled

`func (o *InlineObject193) GetLocationAnalyticsEnabled() bool`

GetLocationAnalyticsEnabled returns the LocationAnalyticsEnabled field if non-nil, zero value otherwise.

### GetLocationAnalyticsEnabledOk

`func (o *InlineObject193) GetLocationAnalyticsEnabledOk() (*bool, bool)`

GetLocationAnalyticsEnabledOk returns a tuple with the LocationAnalyticsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationAnalyticsEnabled

`func (o *InlineObject193) SetLocationAnalyticsEnabled(v bool)`

SetLocationAnalyticsEnabled sets LocationAnalyticsEnabled field to given value.

### HasLocationAnalyticsEnabled

`func (o *InlineObject193) HasLocationAnalyticsEnabled() bool`

HasLocationAnalyticsEnabled returns a boolean if a field has been set.

### GetUpgradeStrategy

`func (o *InlineObject193) GetUpgradeStrategy() string`

GetUpgradeStrategy returns the UpgradeStrategy field if non-nil, zero value otherwise.

### GetUpgradeStrategyOk

`func (o *InlineObject193) GetUpgradeStrategyOk() (*string, bool)`

GetUpgradeStrategyOk returns a tuple with the UpgradeStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeStrategy

`func (o *InlineObject193) SetUpgradeStrategy(v string)`

SetUpgradeStrategy sets UpgradeStrategy field to given value.

### HasUpgradeStrategy

`func (o *InlineObject193) HasUpgradeStrategy() bool`

HasUpgradeStrategy returns a boolean if a field has been set.

### GetLedLightsOn

`func (o *InlineObject193) GetLedLightsOn() bool`

GetLedLightsOn returns the LedLightsOn field if non-nil, zero value otherwise.

### GetLedLightsOnOk

`func (o *InlineObject193) GetLedLightsOnOk() (*bool, bool)`

GetLedLightsOnOk returns a tuple with the LedLightsOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLedLightsOn

`func (o *InlineObject193) SetLedLightsOn(v bool)`

SetLedLightsOn sets LedLightsOn field to given value.

### HasLedLightsOn

`func (o *InlineObject193) HasLedLightsOn() bool`

HasLedLightsOn returns a boolean if a field has been set.

### GetNamedVlans

`func (o *InlineObject193) GetNamedVlans() NetworksNetworkIdWirelessSettingsNamedVlans`

GetNamedVlans returns the NamedVlans field if non-nil, zero value otherwise.

### GetNamedVlansOk

`func (o *InlineObject193) GetNamedVlansOk() (*NetworksNetworkIdWirelessSettingsNamedVlans, bool)`

GetNamedVlansOk returns a tuple with the NamedVlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamedVlans

`func (o *InlineObject193) SetNamedVlans(v NetworksNetworkIdWirelessSettingsNamedVlans)`

SetNamedVlans sets NamedVlans field to given value.

### HasNamedVlans

`func (o *InlineObject193) HasNamedVlans() bool`

HasNamedVlans returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


