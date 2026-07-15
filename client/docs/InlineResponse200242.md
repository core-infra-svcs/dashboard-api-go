# InlineResponse200242

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MeshingEnabled** | Pointer to **bool** | Toggle for enabling or disabling meshing in a network | [optional] 
**Ipv6BridgeEnabled** | Pointer to **bool** | Toggle for enabling or disabling IPv6 bridging in a network (Note: if enabled, SSIDs must also be configured to use bridge mode) | [optional] 
**LocationAnalyticsEnabled** | Pointer to **bool** | Toggle for enabling or disabling location analytics for your network | [optional] 
**UpgradeStrategy** | Pointer to **string** | The default strategy that network devices will use to perform an upgrade. Requires firmware version MR 26.8 or higher. | [optional] 
**Upgrade** | Pointer to [**InlineResponse200242Upgrade**](InlineResponse200242Upgrade.md) |  | [optional] 
**LedLightsOn** | Pointer to **bool** | Toggle for enabling or disabling LED lights on all APs in the network (making them run dark) | [optional] 
**MulticastToUnicastConversion** | Pointer to [**InlineResponse200242MulticastToUnicastConversion**](InlineResponse200242MulticastToUnicastConversion.md) |  | [optional] 
**NamedVlans** | Pointer to [**InlineResponse200242NamedVlans**](InlineResponse200242NamedVlans.md) |  | [optional] 
**RegulatoryDomain** | Pointer to [**InlineResponse200242RegulatoryDomain**](InlineResponse200242RegulatoryDomain.md) |  | [optional] 

## Methods

### NewInlineResponse200242

`func NewInlineResponse200242() *InlineResponse200242`

NewInlineResponse200242 instantiates a new InlineResponse200242 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200242WithDefaults

`func NewInlineResponse200242WithDefaults() *InlineResponse200242`

NewInlineResponse200242WithDefaults instantiates a new InlineResponse200242 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeshingEnabled

`func (o *InlineResponse200242) GetMeshingEnabled() bool`

GetMeshingEnabled returns the MeshingEnabled field if non-nil, zero value otherwise.

### GetMeshingEnabledOk

`func (o *InlineResponse200242) GetMeshingEnabledOk() (*bool, bool)`

GetMeshingEnabledOk returns a tuple with the MeshingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeshingEnabled

`func (o *InlineResponse200242) SetMeshingEnabled(v bool)`

SetMeshingEnabled sets MeshingEnabled field to given value.

### HasMeshingEnabled

`func (o *InlineResponse200242) HasMeshingEnabled() bool`

HasMeshingEnabled returns a boolean if a field has been set.

### GetIpv6BridgeEnabled

`func (o *InlineResponse200242) GetIpv6BridgeEnabled() bool`

GetIpv6BridgeEnabled returns the Ipv6BridgeEnabled field if non-nil, zero value otherwise.

### GetIpv6BridgeEnabledOk

`func (o *InlineResponse200242) GetIpv6BridgeEnabledOk() (*bool, bool)`

GetIpv6BridgeEnabledOk returns a tuple with the Ipv6BridgeEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6BridgeEnabled

`func (o *InlineResponse200242) SetIpv6BridgeEnabled(v bool)`

SetIpv6BridgeEnabled sets Ipv6BridgeEnabled field to given value.

### HasIpv6BridgeEnabled

`func (o *InlineResponse200242) HasIpv6BridgeEnabled() bool`

HasIpv6BridgeEnabled returns a boolean if a field has been set.

### GetLocationAnalyticsEnabled

`func (o *InlineResponse200242) GetLocationAnalyticsEnabled() bool`

GetLocationAnalyticsEnabled returns the LocationAnalyticsEnabled field if non-nil, zero value otherwise.

### GetLocationAnalyticsEnabledOk

`func (o *InlineResponse200242) GetLocationAnalyticsEnabledOk() (*bool, bool)`

GetLocationAnalyticsEnabledOk returns a tuple with the LocationAnalyticsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationAnalyticsEnabled

`func (o *InlineResponse200242) SetLocationAnalyticsEnabled(v bool)`

SetLocationAnalyticsEnabled sets LocationAnalyticsEnabled field to given value.

### HasLocationAnalyticsEnabled

`func (o *InlineResponse200242) HasLocationAnalyticsEnabled() bool`

HasLocationAnalyticsEnabled returns a boolean if a field has been set.

### GetUpgradeStrategy

`func (o *InlineResponse200242) GetUpgradeStrategy() string`

GetUpgradeStrategy returns the UpgradeStrategy field if non-nil, zero value otherwise.

### GetUpgradeStrategyOk

`func (o *InlineResponse200242) GetUpgradeStrategyOk() (*string, bool)`

GetUpgradeStrategyOk returns a tuple with the UpgradeStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeStrategy

`func (o *InlineResponse200242) SetUpgradeStrategy(v string)`

SetUpgradeStrategy sets UpgradeStrategy field to given value.

### HasUpgradeStrategy

`func (o *InlineResponse200242) HasUpgradeStrategy() bool`

HasUpgradeStrategy returns a boolean if a field has been set.

### GetUpgrade

`func (o *InlineResponse200242) GetUpgrade() InlineResponse200242Upgrade`

GetUpgrade returns the Upgrade field if non-nil, zero value otherwise.

### GetUpgradeOk

`func (o *InlineResponse200242) GetUpgradeOk() (*InlineResponse200242Upgrade, bool)`

GetUpgradeOk returns a tuple with the Upgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgrade

`func (o *InlineResponse200242) SetUpgrade(v InlineResponse200242Upgrade)`

SetUpgrade sets Upgrade field to given value.

### HasUpgrade

`func (o *InlineResponse200242) HasUpgrade() bool`

HasUpgrade returns a boolean if a field has been set.

### GetLedLightsOn

`func (o *InlineResponse200242) GetLedLightsOn() bool`

GetLedLightsOn returns the LedLightsOn field if non-nil, zero value otherwise.

### GetLedLightsOnOk

`func (o *InlineResponse200242) GetLedLightsOnOk() (*bool, bool)`

GetLedLightsOnOk returns a tuple with the LedLightsOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLedLightsOn

`func (o *InlineResponse200242) SetLedLightsOn(v bool)`

SetLedLightsOn sets LedLightsOn field to given value.

### HasLedLightsOn

`func (o *InlineResponse200242) HasLedLightsOn() bool`

HasLedLightsOn returns a boolean if a field has been set.

### GetMulticastToUnicastConversion

`func (o *InlineResponse200242) GetMulticastToUnicastConversion() InlineResponse200242MulticastToUnicastConversion`

GetMulticastToUnicastConversion returns the MulticastToUnicastConversion field if non-nil, zero value otherwise.

### GetMulticastToUnicastConversionOk

`func (o *InlineResponse200242) GetMulticastToUnicastConversionOk() (*InlineResponse200242MulticastToUnicastConversion, bool)`

GetMulticastToUnicastConversionOk returns a tuple with the MulticastToUnicastConversion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMulticastToUnicastConversion

`func (o *InlineResponse200242) SetMulticastToUnicastConversion(v InlineResponse200242MulticastToUnicastConversion)`

SetMulticastToUnicastConversion sets MulticastToUnicastConversion field to given value.

### HasMulticastToUnicastConversion

`func (o *InlineResponse200242) HasMulticastToUnicastConversion() bool`

HasMulticastToUnicastConversion returns a boolean if a field has been set.

### GetNamedVlans

`func (o *InlineResponse200242) GetNamedVlans() InlineResponse200242NamedVlans`

GetNamedVlans returns the NamedVlans field if non-nil, zero value otherwise.

### GetNamedVlansOk

`func (o *InlineResponse200242) GetNamedVlansOk() (*InlineResponse200242NamedVlans, bool)`

GetNamedVlansOk returns a tuple with the NamedVlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamedVlans

`func (o *InlineResponse200242) SetNamedVlans(v InlineResponse200242NamedVlans)`

SetNamedVlans sets NamedVlans field to given value.

### HasNamedVlans

`func (o *InlineResponse200242) HasNamedVlans() bool`

HasNamedVlans returns a boolean if a field has been set.

### GetRegulatoryDomain

`func (o *InlineResponse200242) GetRegulatoryDomain() InlineResponse200242RegulatoryDomain`

GetRegulatoryDomain returns the RegulatoryDomain field if non-nil, zero value otherwise.

### GetRegulatoryDomainOk

`func (o *InlineResponse200242) GetRegulatoryDomainOk() (*InlineResponse200242RegulatoryDomain, bool)`

GetRegulatoryDomainOk returns a tuple with the RegulatoryDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegulatoryDomain

`func (o *InlineResponse200242) SetRegulatoryDomain(v InlineResponse200242RegulatoryDomain)`

SetRegulatoryDomain sets RegulatoryDomain field to given value.

### HasRegulatoryDomain

`func (o *InlineResponse200242) HasRegulatoryDomain() bool`

HasRegulatoryDomain returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


