# NetworksNetworkIdSwitchAccessPoliciesRadiusCache

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Enable to cache authorization and authentication responses on the RADIUS server | [optional] 
**Timeout** | Pointer to **int32** | If RADIUS caching is enabled, this value dictates how long the cache will remain in the RADIUS server, in hours, to allow network access without authentication | [optional] 

## Methods

### NewNetworksNetworkIdSwitchAccessPoliciesRadiusCache

`func NewNetworksNetworkIdSwitchAccessPoliciesRadiusCache() *NetworksNetworkIdSwitchAccessPoliciesRadiusCache`

NewNetworksNetworkIdSwitchAccessPoliciesRadiusCache instantiates a new NetworksNetworkIdSwitchAccessPoliciesRadiusCache object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworksNetworkIdSwitchAccessPoliciesRadiusCacheWithDefaults

`func NewNetworksNetworkIdSwitchAccessPoliciesRadiusCacheWithDefaults() *NetworksNetworkIdSwitchAccessPoliciesRadiusCache`

NewNetworksNetworkIdSwitchAccessPoliciesRadiusCacheWithDefaults instantiates a new NetworksNetworkIdSwitchAccessPoliciesRadiusCache object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusCache) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusCache) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusCache) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusCache) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetTimeout

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusCache) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusCache) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusCache) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusCache) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


