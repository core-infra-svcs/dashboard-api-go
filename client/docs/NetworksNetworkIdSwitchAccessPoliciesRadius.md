# NetworksNetworkIdSwitchAccessPoliciesRadius

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CriticalAuth** | Pointer to [**NetworksNetworkIdSwitchAccessPoliciesRadiusCriticalAuth**](NetworksNetworkIdSwitchAccessPoliciesRadiusCriticalAuth.md) |  | [optional] 
**FailedAuthVlanId** | Pointer to **NullableInt32** | VLAN that clients will be placed on when RADIUS authentication fails. Will be null if hostMode is Multi-Auth | [optional] 
**FailedAuthGroupPolicyId** | Pointer to **string** | Group policy Number for failed authentication group policy (Requires MS 18 or higher) | [optional] 
**FailedAuthSgtId** | Pointer to **int32** | Security Group Tag ID for failed authentication group policy (Requires MS 18 or higher) | [optional] 
**ReAuthenticationInterval** | Pointer to **NullableInt32** | Re-authentication period in seconds. Will be null if hostMode is Multi-Auth | [optional] 
**Cache** | Pointer to [**NetworksNetworkIdSwitchAccessPoliciesRadiusCache**](NetworksNetworkIdSwitchAccessPoliciesRadiusCache.md) |  | [optional] 
**Authentication** | Pointer to [**NetworksNetworkIdSwitchAccessPoliciesRadiusAuthentication**](NetworksNetworkIdSwitchAccessPoliciesRadiusAuthentication.md) |  | [optional] 
**PreAuthenticationGroupPolicyId** | Pointer to **string** | Group policy Number for pre-authentication group policy (Requires MS 18 or higher) | [optional] 

## Methods

### NewNetworksNetworkIdSwitchAccessPoliciesRadius

`func NewNetworksNetworkIdSwitchAccessPoliciesRadius() *NetworksNetworkIdSwitchAccessPoliciesRadius`

NewNetworksNetworkIdSwitchAccessPoliciesRadius instantiates a new NetworksNetworkIdSwitchAccessPoliciesRadius object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworksNetworkIdSwitchAccessPoliciesRadiusWithDefaults

`func NewNetworksNetworkIdSwitchAccessPoliciesRadiusWithDefaults() *NetworksNetworkIdSwitchAccessPoliciesRadius`

NewNetworksNetworkIdSwitchAccessPoliciesRadiusWithDefaults instantiates a new NetworksNetworkIdSwitchAccessPoliciesRadius object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCriticalAuth

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadius) GetCriticalAuth() NetworksNetworkIdSwitchAccessPoliciesRadiusCriticalAuth`

GetCriticalAuth returns the CriticalAuth field if non-nil, zero value otherwise.

### GetCriticalAuthOk

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadius) GetCriticalAuthOk() (*NetworksNetworkIdSwitchAccessPoliciesRadiusCriticalAuth, bool)`

GetCriticalAuthOk returns a tuple with the CriticalAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriticalAuth

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadius) SetCriticalAuth(v NetworksNetworkIdSwitchAccessPoliciesRadiusCriticalAuth)`

SetCriticalAuth sets CriticalAuth field to given value.

### HasCriticalAuth

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadius) HasCriticalAuth() bool`

HasCriticalAuth returns a boolean if a field has been set.

### GetFailedAuthVlanId

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadius) GetFailedAuthVlanId() int32`

GetFailedAuthVlanId returns the FailedAuthVlanId field if non-nil, zero value otherwise.

### GetFailedAuthVlanIdOk

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadius) GetFailedAuthVlanIdOk() (*int32, bool)`

GetFailedAuthVlanIdOk returns a tuple with the FailedAuthVlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedAuthVlanId

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadius) SetFailedAuthVlanId(v int32)`

SetFailedAuthVlanId sets FailedAuthVlanId field to given value.

### HasFailedAuthVlanId

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadius) HasFailedAuthVlanId() bool`

HasFailedAuthVlanId returns a boolean if a field has been set.

### SetFailedAuthVlanIdNil

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadius) SetFailedAuthVlanIdNil(b bool)`

 SetFailedAuthVlanIdNil sets the value for FailedAuthVlanId to be an explicit nil

### UnsetFailedAuthVlanId
`func (o *NetworksNetworkIdSwitchAccessPoliciesRadius) UnsetFailedAuthVlanId()`

UnsetFailedAuthVlanId ensures that no value is present for FailedAuthVlanId, not even an explicit nil
### GetFailedAuthGroupPolicyId

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadius) GetFailedAuthGroupPolicyId() string`

GetFailedAuthGroupPolicyId returns the FailedAuthGroupPolicyId field if non-nil, zero value otherwise.

### GetFailedAuthGroupPolicyIdOk

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadius) GetFailedAuthGroupPolicyIdOk() (*string, bool)`

GetFailedAuthGroupPolicyIdOk returns a tuple with the FailedAuthGroupPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedAuthGroupPolicyId

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadius) SetFailedAuthGroupPolicyId(v string)`

SetFailedAuthGroupPolicyId sets FailedAuthGroupPolicyId field to given value.

### HasFailedAuthGroupPolicyId

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadius) HasFailedAuthGroupPolicyId() bool`

HasFailedAuthGroupPolicyId returns a boolean if a field has been set.

### GetFailedAuthSgtId

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadius) GetFailedAuthSgtId() int32`

GetFailedAuthSgtId returns the FailedAuthSgtId field if non-nil, zero value otherwise.

### GetFailedAuthSgtIdOk

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadius) GetFailedAuthSgtIdOk() (*int32, bool)`

GetFailedAuthSgtIdOk returns a tuple with the FailedAuthSgtId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedAuthSgtId

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadius) SetFailedAuthSgtId(v int32)`

SetFailedAuthSgtId sets FailedAuthSgtId field to given value.

### HasFailedAuthSgtId

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadius) HasFailedAuthSgtId() bool`

HasFailedAuthSgtId returns a boolean if a field has been set.

### GetReAuthenticationInterval

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadius) GetReAuthenticationInterval() int32`

GetReAuthenticationInterval returns the ReAuthenticationInterval field if non-nil, zero value otherwise.

### GetReAuthenticationIntervalOk

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadius) GetReAuthenticationIntervalOk() (*int32, bool)`

GetReAuthenticationIntervalOk returns a tuple with the ReAuthenticationInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReAuthenticationInterval

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadius) SetReAuthenticationInterval(v int32)`

SetReAuthenticationInterval sets ReAuthenticationInterval field to given value.

### HasReAuthenticationInterval

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadius) HasReAuthenticationInterval() bool`

HasReAuthenticationInterval returns a boolean if a field has been set.

### SetReAuthenticationIntervalNil

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadius) SetReAuthenticationIntervalNil(b bool)`

 SetReAuthenticationIntervalNil sets the value for ReAuthenticationInterval to be an explicit nil

### UnsetReAuthenticationInterval
`func (o *NetworksNetworkIdSwitchAccessPoliciesRadius) UnsetReAuthenticationInterval()`

UnsetReAuthenticationInterval ensures that no value is present for ReAuthenticationInterval, not even an explicit nil
### GetCache

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadius) GetCache() NetworksNetworkIdSwitchAccessPoliciesRadiusCache`

GetCache returns the Cache field if non-nil, zero value otherwise.

### GetCacheOk

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadius) GetCacheOk() (*NetworksNetworkIdSwitchAccessPoliciesRadiusCache, bool)`

GetCacheOk returns a tuple with the Cache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCache

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadius) SetCache(v NetworksNetworkIdSwitchAccessPoliciesRadiusCache)`

SetCache sets Cache field to given value.

### HasCache

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadius) HasCache() bool`

HasCache returns a boolean if a field has been set.

### GetAuthentication

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadius) GetAuthentication() NetworksNetworkIdSwitchAccessPoliciesRadiusAuthentication`

GetAuthentication returns the Authentication field if non-nil, zero value otherwise.

### GetAuthenticationOk

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadius) GetAuthenticationOk() (*NetworksNetworkIdSwitchAccessPoliciesRadiusAuthentication, bool)`

GetAuthenticationOk returns a tuple with the Authentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthentication

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadius) SetAuthentication(v NetworksNetworkIdSwitchAccessPoliciesRadiusAuthentication)`

SetAuthentication sets Authentication field to given value.

### HasAuthentication

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadius) HasAuthentication() bool`

HasAuthentication returns a boolean if a field has been set.

### GetPreAuthenticationGroupPolicyId

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadius) GetPreAuthenticationGroupPolicyId() string`

GetPreAuthenticationGroupPolicyId returns the PreAuthenticationGroupPolicyId field if non-nil, zero value otherwise.

### GetPreAuthenticationGroupPolicyIdOk

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadius) GetPreAuthenticationGroupPolicyIdOk() (*string, bool)`

GetPreAuthenticationGroupPolicyIdOk returns a tuple with the PreAuthenticationGroupPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreAuthenticationGroupPolicyId

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadius) SetPreAuthenticationGroupPolicyId(v string)`

SetPreAuthenticationGroupPolicyId sets PreAuthenticationGroupPolicyId field to given value.

### HasPreAuthenticationGroupPolicyId

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadius) HasPreAuthenticationGroupPolicyId() bool`

HasPreAuthenticationGroupPolicyId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


