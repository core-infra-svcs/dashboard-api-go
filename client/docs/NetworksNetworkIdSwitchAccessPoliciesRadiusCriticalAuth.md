# NetworksNetworkIdSwitchAccessPoliciesRadiusCriticalAuth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataVlanId** | Pointer to **NullableInt32** | VLAN that clients who use data will be placed on when RADIUS authentication fails. Will be null if hostMode is Multi-Auth | [optional] 
**VoiceVlanId** | Pointer to **NullableInt32** | VLAN that clients who use voice will be placed on when RADIUS authentication fails. Will be null if hostMode is Multi-Auth | [optional] 
**SuspendPortBounce** | Pointer to **bool** | Disable port bounce when RADIUS servers are unreachable | [optional] 
**DataGroupPolicyId** | Pointer to **string** | Group policy Number for data VLAN (Requires MS 18 or higher) | [optional] 
**VoiceGroupPolicyId** | Pointer to **string** | Group policy Number for voice VLAN (Requires MS 18 or higher) | [optional] 
**DataSgtId** | Pointer to **int32** | Security Group Tag ID for data VLAN (Requires MS 18 or higher) | [optional] 
**VoiceSgtId** | Pointer to **int32** | Security Group Tag ID for voice VLAN (Requires MS 18 or higher) | [optional] 

## Methods

### NewNetworksNetworkIdSwitchAccessPoliciesRadiusCriticalAuth

`func NewNetworksNetworkIdSwitchAccessPoliciesRadiusCriticalAuth() *NetworksNetworkIdSwitchAccessPoliciesRadiusCriticalAuth`

NewNetworksNetworkIdSwitchAccessPoliciesRadiusCriticalAuth instantiates a new NetworksNetworkIdSwitchAccessPoliciesRadiusCriticalAuth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworksNetworkIdSwitchAccessPoliciesRadiusCriticalAuthWithDefaults

`func NewNetworksNetworkIdSwitchAccessPoliciesRadiusCriticalAuthWithDefaults() *NetworksNetworkIdSwitchAccessPoliciesRadiusCriticalAuth`

NewNetworksNetworkIdSwitchAccessPoliciesRadiusCriticalAuthWithDefaults instantiates a new NetworksNetworkIdSwitchAccessPoliciesRadiusCriticalAuth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataVlanId

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusCriticalAuth) GetDataVlanId() int32`

GetDataVlanId returns the DataVlanId field if non-nil, zero value otherwise.

### GetDataVlanIdOk

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusCriticalAuth) GetDataVlanIdOk() (*int32, bool)`

GetDataVlanIdOk returns a tuple with the DataVlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataVlanId

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusCriticalAuth) SetDataVlanId(v int32)`

SetDataVlanId sets DataVlanId field to given value.

### HasDataVlanId

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusCriticalAuth) HasDataVlanId() bool`

HasDataVlanId returns a boolean if a field has been set.

### SetDataVlanIdNil

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusCriticalAuth) SetDataVlanIdNil(b bool)`

 SetDataVlanIdNil sets the value for DataVlanId to be an explicit nil

### UnsetDataVlanId
`func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusCriticalAuth) UnsetDataVlanId()`

UnsetDataVlanId ensures that no value is present for DataVlanId, not even an explicit nil
### GetVoiceVlanId

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusCriticalAuth) GetVoiceVlanId() int32`

GetVoiceVlanId returns the VoiceVlanId field if non-nil, zero value otherwise.

### GetVoiceVlanIdOk

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusCriticalAuth) GetVoiceVlanIdOk() (*int32, bool)`

GetVoiceVlanIdOk returns a tuple with the VoiceVlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceVlanId

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusCriticalAuth) SetVoiceVlanId(v int32)`

SetVoiceVlanId sets VoiceVlanId field to given value.

### HasVoiceVlanId

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusCriticalAuth) HasVoiceVlanId() bool`

HasVoiceVlanId returns a boolean if a field has been set.

### SetVoiceVlanIdNil

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusCriticalAuth) SetVoiceVlanIdNil(b bool)`

 SetVoiceVlanIdNil sets the value for VoiceVlanId to be an explicit nil

### UnsetVoiceVlanId
`func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusCriticalAuth) UnsetVoiceVlanId()`

UnsetVoiceVlanId ensures that no value is present for VoiceVlanId, not even an explicit nil
### GetSuspendPortBounce

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusCriticalAuth) GetSuspendPortBounce() bool`

GetSuspendPortBounce returns the SuspendPortBounce field if non-nil, zero value otherwise.

### GetSuspendPortBounceOk

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusCriticalAuth) GetSuspendPortBounceOk() (*bool, bool)`

GetSuspendPortBounceOk returns a tuple with the SuspendPortBounce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspendPortBounce

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusCriticalAuth) SetSuspendPortBounce(v bool)`

SetSuspendPortBounce sets SuspendPortBounce field to given value.

### HasSuspendPortBounce

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusCriticalAuth) HasSuspendPortBounce() bool`

HasSuspendPortBounce returns a boolean if a field has been set.

### GetDataGroupPolicyId

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusCriticalAuth) GetDataGroupPolicyId() string`

GetDataGroupPolicyId returns the DataGroupPolicyId field if non-nil, zero value otherwise.

### GetDataGroupPolicyIdOk

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusCriticalAuth) GetDataGroupPolicyIdOk() (*string, bool)`

GetDataGroupPolicyIdOk returns a tuple with the DataGroupPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataGroupPolicyId

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusCriticalAuth) SetDataGroupPolicyId(v string)`

SetDataGroupPolicyId sets DataGroupPolicyId field to given value.

### HasDataGroupPolicyId

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusCriticalAuth) HasDataGroupPolicyId() bool`

HasDataGroupPolicyId returns a boolean if a field has been set.

### GetVoiceGroupPolicyId

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusCriticalAuth) GetVoiceGroupPolicyId() string`

GetVoiceGroupPolicyId returns the VoiceGroupPolicyId field if non-nil, zero value otherwise.

### GetVoiceGroupPolicyIdOk

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusCriticalAuth) GetVoiceGroupPolicyIdOk() (*string, bool)`

GetVoiceGroupPolicyIdOk returns a tuple with the VoiceGroupPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceGroupPolicyId

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusCriticalAuth) SetVoiceGroupPolicyId(v string)`

SetVoiceGroupPolicyId sets VoiceGroupPolicyId field to given value.

### HasVoiceGroupPolicyId

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusCriticalAuth) HasVoiceGroupPolicyId() bool`

HasVoiceGroupPolicyId returns a boolean if a field has been set.

### GetDataSgtId

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusCriticalAuth) GetDataSgtId() int32`

GetDataSgtId returns the DataSgtId field if non-nil, zero value otherwise.

### GetDataSgtIdOk

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusCriticalAuth) GetDataSgtIdOk() (*int32, bool)`

GetDataSgtIdOk returns a tuple with the DataSgtId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSgtId

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusCriticalAuth) SetDataSgtId(v int32)`

SetDataSgtId sets DataSgtId field to given value.

### HasDataSgtId

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusCriticalAuth) HasDataSgtId() bool`

HasDataSgtId returns a boolean if a field has been set.

### GetVoiceSgtId

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusCriticalAuth) GetVoiceSgtId() int32`

GetVoiceSgtId returns the VoiceSgtId field if non-nil, zero value otherwise.

### GetVoiceSgtIdOk

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusCriticalAuth) GetVoiceSgtIdOk() (*int32, bool)`

GetVoiceSgtIdOk returns a tuple with the VoiceSgtId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceSgtId

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusCriticalAuth) SetVoiceSgtId(v int32)`

SetVoiceSgtId sets VoiceSgtId field to given value.

### HasVoiceSgtId

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusCriticalAuth) HasVoiceSgtId() bool`

HasVoiceSgtId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


