# NetworksNetworkIdGroupPoliciesBandwidth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Settings** | Pointer to **string** | How bandwidth limits are enforced. Can be &#39;network default&#39;, &#39;ignore&#39; or &#39;custom&#39;. | [optional] 
**BandwidthLimits** | Pointer to [**NetworksNetworkIdGroupPoliciesBandwidthBandwidthLimits**](NetworksNetworkIdGroupPoliciesBandwidthBandwidthLimits.md) |  | [optional] 

## Methods

### NewNetworksNetworkIdGroupPoliciesBandwidth

`func NewNetworksNetworkIdGroupPoliciesBandwidth() *NetworksNetworkIdGroupPoliciesBandwidth`

NewNetworksNetworkIdGroupPoliciesBandwidth instantiates a new NetworksNetworkIdGroupPoliciesBandwidth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworksNetworkIdGroupPoliciesBandwidthWithDefaults

`func NewNetworksNetworkIdGroupPoliciesBandwidthWithDefaults() *NetworksNetworkIdGroupPoliciesBandwidth`

NewNetworksNetworkIdGroupPoliciesBandwidthWithDefaults instantiates a new NetworksNetworkIdGroupPoliciesBandwidth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSettings

`func (o *NetworksNetworkIdGroupPoliciesBandwidth) GetSettings() string`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *NetworksNetworkIdGroupPoliciesBandwidth) GetSettingsOk() (*string, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *NetworksNetworkIdGroupPoliciesBandwidth) SetSettings(v string)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *NetworksNetworkIdGroupPoliciesBandwidth) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetBandwidthLimits

`func (o *NetworksNetworkIdGroupPoliciesBandwidth) GetBandwidthLimits() NetworksNetworkIdGroupPoliciesBandwidthBandwidthLimits`

GetBandwidthLimits returns the BandwidthLimits field if non-nil, zero value otherwise.

### GetBandwidthLimitsOk

`func (o *NetworksNetworkIdGroupPoliciesBandwidth) GetBandwidthLimitsOk() (*NetworksNetworkIdGroupPoliciesBandwidthBandwidthLimits, bool)`

GetBandwidthLimitsOk returns a tuple with the BandwidthLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidthLimits

`func (o *NetworksNetworkIdGroupPoliciesBandwidth) SetBandwidthLimits(v NetworksNetworkIdGroupPoliciesBandwidthBandwidthLimits)`

SetBandwidthLimits sets BandwidthLimits field to given value.

### HasBandwidthLimits

`func (o *NetworksNetworkIdGroupPoliciesBandwidth) HasBandwidthLimits() bool`

HasBandwidthLimits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


