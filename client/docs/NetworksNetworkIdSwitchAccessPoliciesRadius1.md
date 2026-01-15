# NetworksNetworkIdSwitchAccessPoliciesRadius1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CriticalAuth** | Pointer to [**NetworksNetworkIdSwitchAccessPoliciesRadius1CriticalAuth**](NetworksNetworkIdSwitchAccessPoliciesRadius1CriticalAuth.md) |  | [optional] 
**FailedAuthVlanId** | Pointer to **NullableInt32** | VLAN that clients will be placed on when RADIUS authentication fails. Will be null if hostMode is Multi-Auth | [optional] 
**FailedAuthGroupPolicyId** | Pointer to **NullableString** | Group policy Number for failed authentication group policy | [optional] 
**FailedAuthSgtId** | Pointer to **NullableInt32** | Security Group Tag ID for failed authentication group policy | [optional] 
**ReAuthenticationInterval** | Pointer to **NullableInt32** | Re-authentication period in seconds. Will be null if hostMode is Multi-Auth | [optional] 
**Cache** | Pointer to [**NetworksNetworkIdSwitchAccessPoliciesRadiusCache**](NetworksNetworkIdSwitchAccessPoliciesRadiusCache.md) |  | [optional] 
**Authentication** | Pointer to [**NetworksNetworkIdSwitchAccessPoliciesRadius1Authentication**](NetworksNetworkIdSwitchAccessPoliciesRadius1Authentication.md) |  | [optional] 
**PreAuthenticationGroupPolicyId** | Pointer to **NullableString** | Group policy Number for pre-authentication group policy | [optional] 

## Methods

### NewNetworksNetworkIdSwitchAccessPoliciesRadius1

`func NewNetworksNetworkIdSwitchAccessPoliciesRadius1() *NetworksNetworkIdSwitchAccessPoliciesRadius1`

NewNetworksNetworkIdSwitchAccessPoliciesRadius1 instantiates a new NetworksNetworkIdSwitchAccessPoliciesRadius1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworksNetworkIdSwitchAccessPoliciesRadius1WithDefaults

`func NewNetworksNetworkIdSwitchAccessPoliciesRadius1WithDefaults() *NetworksNetworkIdSwitchAccessPoliciesRadius1`

NewNetworksNetworkIdSwitchAccessPoliciesRadius1WithDefaults instantiates a new NetworksNetworkIdSwitchAccessPoliciesRadius1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCriticalAuth

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadius1) GetCriticalAuth() NetworksNetworkIdSwitchAccessPoliciesRadius1CriticalAuth`

GetCriticalAuth returns the CriticalAuth field if non-nil, zero value otherwise.

### GetCriticalAuthOk

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadius1) GetCriticalAuthOk() (*NetworksNetworkIdSwitchAccessPoliciesRadius1CriticalAuth, bool)`

GetCriticalAuthOk returns a tuple with the CriticalAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriticalAuth

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadius1) SetCriticalAuth(v NetworksNetworkIdSwitchAccessPoliciesRadius1CriticalAuth)`

SetCriticalAuth sets CriticalAuth field to given value.

### HasCriticalAuth

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadius1) HasCriticalAuth() bool`

HasCriticalAuth returns a boolean if a field has been set.

### GetFailedAuthVlanId

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadius1) GetFailedAuthVlanId() int32`

GetFailedAuthVlanId returns the FailedAuthVlanId field if non-nil, zero value otherwise.

### GetFailedAuthVlanIdOk

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadius1) GetFailedAuthVlanIdOk() (*int32, bool)`

GetFailedAuthVlanIdOk returns a tuple with the FailedAuthVlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedAuthVlanId

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadius1) SetFailedAuthVlanId(v int32)`

SetFailedAuthVlanId sets FailedAuthVlanId field to given value.

### HasFailedAuthVlanId

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadius1) HasFailedAuthVlanId() bool`

HasFailedAuthVlanId returns a boolean if a field has been set.

### SetFailedAuthVlanIdNil

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadius1) SetFailedAuthVlanIdNil(b bool)`

 SetFailedAuthVlanIdNil sets the value for FailedAuthVlanId to be an explicit nil

### UnsetFailedAuthVlanId
`func (o *NetworksNetworkIdSwitchAccessPoliciesRadius1) UnsetFailedAuthVlanId()`

UnsetFailedAuthVlanId ensures that no value is present for FailedAuthVlanId, not even an explicit nil
### GetFailedAuthGroupPolicyId

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadius1) GetFailedAuthGroupPolicyId() string`

GetFailedAuthGroupPolicyId returns the FailedAuthGroupPolicyId field if non-nil, zero value otherwise.

### GetFailedAuthGroupPolicyIdOk

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadius1) GetFailedAuthGroupPolicyIdOk() (*string, bool)`

GetFailedAuthGroupPolicyIdOk returns a tuple with the FailedAuthGroupPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedAuthGroupPolicyId

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadius1) SetFailedAuthGroupPolicyId(v string)`

SetFailedAuthGroupPolicyId sets FailedAuthGroupPolicyId field to given value.

### HasFailedAuthGroupPolicyId

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadius1) HasFailedAuthGroupPolicyId() bool`

HasFailedAuthGroupPolicyId returns a boolean if a field has been set.

### SetFailedAuthGroupPolicyIdNil

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadius1) SetFailedAuthGroupPolicyIdNil(b bool)`

 SetFailedAuthGroupPolicyIdNil sets the value for FailedAuthGroupPolicyId to be an explicit nil

### UnsetFailedAuthGroupPolicyId
`func (o *NetworksNetworkIdSwitchAccessPoliciesRadius1) UnsetFailedAuthGroupPolicyId()`

UnsetFailedAuthGroupPolicyId ensures that no value is present for FailedAuthGroupPolicyId, not even an explicit nil
### GetFailedAuthSgtId

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadius1) GetFailedAuthSgtId() int32`

GetFailedAuthSgtId returns the FailedAuthSgtId field if non-nil, zero value otherwise.

### GetFailedAuthSgtIdOk

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadius1) GetFailedAuthSgtIdOk() (*int32, bool)`

GetFailedAuthSgtIdOk returns a tuple with the FailedAuthSgtId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedAuthSgtId

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadius1) SetFailedAuthSgtId(v int32)`

SetFailedAuthSgtId sets FailedAuthSgtId field to given value.

### HasFailedAuthSgtId

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadius1) HasFailedAuthSgtId() bool`

HasFailedAuthSgtId returns a boolean if a field has been set.

### SetFailedAuthSgtIdNil

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadius1) SetFailedAuthSgtIdNil(b bool)`

 SetFailedAuthSgtIdNil sets the value for FailedAuthSgtId to be an explicit nil

### UnsetFailedAuthSgtId
`func (o *NetworksNetworkIdSwitchAccessPoliciesRadius1) UnsetFailedAuthSgtId()`

UnsetFailedAuthSgtId ensures that no value is present for FailedAuthSgtId, not even an explicit nil
### GetReAuthenticationInterval

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadius1) GetReAuthenticationInterval() int32`

GetReAuthenticationInterval returns the ReAuthenticationInterval field if non-nil, zero value otherwise.

### GetReAuthenticationIntervalOk

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadius1) GetReAuthenticationIntervalOk() (*int32, bool)`

GetReAuthenticationIntervalOk returns a tuple with the ReAuthenticationInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReAuthenticationInterval

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadius1) SetReAuthenticationInterval(v int32)`

SetReAuthenticationInterval sets ReAuthenticationInterval field to given value.

### HasReAuthenticationInterval

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadius1) HasReAuthenticationInterval() bool`

HasReAuthenticationInterval returns a boolean if a field has been set.

### SetReAuthenticationIntervalNil

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadius1) SetReAuthenticationIntervalNil(b bool)`

 SetReAuthenticationIntervalNil sets the value for ReAuthenticationInterval to be an explicit nil

### UnsetReAuthenticationInterval
`func (o *NetworksNetworkIdSwitchAccessPoliciesRadius1) UnsetReAuthenticationInterval()`

UnsetReAuthenticationInterval ensures that no value is present for ReAuthenticationInterval, not even an explicit nil
### GetCache

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadius1) GetCache() NetworksNetworkIdSwitchAccessPoliciesRadiusCache`

GetCache returns the Cache field if non-nil, zero value otherwise.

### GetCacheOk

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadius1) GetCacheOk() (*NetworksNetworkIdSwitchAccessPoliciesRadiusCache, bool)`

GetCacheOk returns a tuple with the Cache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCache

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadius1) SetCache(v NetworksNetworkIdSwitchAccessPoliciesRadiusCache)`

SetCache sets Cache field to given value.

### HasCache

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadius1) HasCache() bool`

HasCache returns a boolean if a field has been set.

### GetAuthentication

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadius1) GetAuthentication() NetworksNetworkIdSwitchAccessPoliciesRadius1Authentication`

GetAuthentication returns the Authentication field if non-nil, zero value otherwise.

### GetAuthenticationOk

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadius1) GetAuthenticationOk() (*NetworksNetworkIdSwitchAccessPoliciesRadius1Authentication, bool)`

GetAuthenticationOk returns a tuple with the Authentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthentication

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadius1) SetAuthentication(v NetworksNetworkIdSwitchAccessPoliciesRadius1Authentication)`

SetAuthentication sets Authentication field to given value.

### HasAuthentication

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadius1) HasAuthentication() bool`

HasAuthentication returns a boolean if a field has been set.

### GetPreAuthenticationGroupPolicyId

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadius1) GetPreAuthenticationGroupPolicyId() string`

GetPreAuthenticationGroupPolicyId returns the PreAuthenticationGroupPolicyId field if non-nil, zero value otherwise.

### GetPreAuthenticationGroupPolicyIdOk

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadius1) GetPreAuthenticationGroupPolicyIdOk() (*string, bool)`

GetPreAuthenticationGroupPolicyIdOk returns a tuple with the PreAuthenticationGroupPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreAuthenticationGroupPolicyId

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadius1) SetPreAuthenticationGroupPolicyId(v string)`

SetPreAuthenticationGroupPolicyId sets PreAuthenticationGroupPolicyId field to given value.

### HasPreAuthenticationGroupPolicyId

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadius1) HasPreAuthenticationGroupPolicyId() bool`

HasPreAuthenticationGroupPolicyId returns a boolean if a field has been set.

### SetPreAuthenticationGroupPolicyIdNil

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadius1) SetPreAuthenticationGroupPolicyIdNil(b bool)`

 SetPreAuthenticationGroupPolicyIdNil sets the value for PreAuthenticationGroupPolicyId to be an explicit nil

### UnsetPreAuthenticationGroupPolicyId
`func (o *NetworksNetworkIdSwitchAccessPoliciesRadius1) UnsetPreAuthenticationGroupPolicyId()`

UnsetPreAuthenticationGroupPolicyId ensures that no value is present for PreAuthenticationGroupPolicyId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


