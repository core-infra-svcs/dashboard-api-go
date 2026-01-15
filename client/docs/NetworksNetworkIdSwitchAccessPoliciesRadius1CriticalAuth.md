# NetworksNetworkIdSwitchAccessPoliciesRadius1CriticalAuth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataVlanId** | Pointer to **NullableInt32** | VLAN that clients who use data will be placed on when RADIUS authentication fails. Will be null if hostMode is Multi-Auth | [optional] 
**VoiceVlanId** | Pointer to **NullableInt32** | VLAN that clients who use voice will be placed on when RADIUS authentication fails. Will be null if hostMode is Multi-Auth | [optional] 
**SuspendPortBounce** | Pointer to **bool** | Disable port bounce when RADIUS servers are unreachable | [optional] 
**DataGroupPolicyId** | Pointer to **NullableString** | Group policy Number for data VLAN | [optional] 
**VoiceGroupPolicyId** | Pointer to **NullableString** | Group policy Number for voice VLAN | [optional] 
**DataSgtId** | Pointer to **NullableInt32** | Security Group Tag ID for data VLAN | [optional] 
**VoiceSgtId** | Pointer to **NullableInt32** | Security Group Tag ID for voice VLAN | [optional] 

## Methods

### NewNetworksNetworkIdSwitchAccessPoliciesRadius1CriticalAuth

`func NewNetworksNetworkIdSwitchAccessPoliciesRadius1CriticalAuth() *NetworksNetworkIdSwitchAccessPoliciesRadius1CriticalAuth`

NewNetworksNetworkIdSwitchAccessPoliciesRadius1CriticalAuth instantiates a new NetworksNetworkIdSwitchAccessPoliciesRadius1CriticalAuth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworksNetworkIdSwitchAccessPoliciesRadius1CriticalAuthWithDefaults

`func NewNetworksNetworkIdSwitchAccessPoliciesRadius1CriticalAuthWithDefaults() *NetworksNetworkIdSwitchAccessPoliciesRadius1CriticalAuth`

NewNetworksNetworkIdSwitchAccessPoliciesRadius1CriticalAuthWithDefaults instantiates a new NetworksNetworkIdSwitchAccessPoliciesRadius1CriticalAuth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataVlanId

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadius1CriticalAuth) GetDataVlanId() int32`

GetDataVlanId returns the DataVlanId field if non-nil, zero value otherwise.

### GetDataVlanIdOk

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadius1CriticalAuth) GetDataVlanIdOk() (*int32, bool)`

GetDataVlanIdOk returns a tuple with the DataVlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataVlanId

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadius1CriticalAuth) SetDataVlanId(v int32)`

SetDataVlanId sets DataVlanId field to given value.

### HasDataVlanId

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadius1CriticalAuth) HasDataVlanId() bool`

HasDataVlanId returns a boolean if a field has been set.

### SetDataVlanIdNil

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadius1CriticalAuth) SetDataVlanIdNil(b bool)`

 SetDataVlanIdNil sets the value for DataVlanId to be an explicit nil

### UnsetDataVlanId
`func (o *NetworksNetworkIdSwitchAccessPoliciesRadius1CriticalAuth) UnsetDataVlanId()`

UnsetDataVlanId ensures that no value is present for DataVlanId, not even an explicit nil
### GetVoiceVlanId

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadius1CriticalAuth) GetVoiceVlanId() int32`

GetVoiceVlanId returns the VoiceVlanId field if non-nil, zero value otherwise.

### GetVoiceVlanIdOk

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadius1CriticalAuth) GetVoiceVlanIdOk() (*int32, bool)`

GetVoiceVlanIdOk returns a tuple with the VoiceVlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceVlanId

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadius1CriticalAuth) SetVoiceVlanId(v int32)`

SetVoiceVlanId sets VoiceVlanId field to given value.

### HasVoiceVlanId

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadius1CriticalAuth) HasVoiceVlanId() bool`

HasVoiceVlanId returns a boolean if a field has been set.

### SetVoiceVlanIdNil

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadius1CriticalAuth) SetVoiceVlanIdNil(b bool)`

 SetVoiceVlanIdNil sets the value for VoiceVlanId to be an explicit nil

### UnsetVoiceVlanId
`func (o *NetworksNetworkIdSwitchAccessPoliciesRadius1CriticalAuth) UnsetVoiceVlanId()`

UnsetVoiceVlanId ensures that no value is present for VoiceVlanId, not even an explicit nil
### GetSuspendPortBounce

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadius1CriticalAuth) GetSuspendPortBounce() bool`

GetSuspendPortBounce returns the SuspendPortBounce field if non-nil, zero value otherwise.

### GetSuspendPortBounceOk

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadius1CriticalAuth) GetSuspendPortBounceOk() (*bool, bool)`

GetSuspendPortBounceOk returns a tuple with the SuspendPortBounce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspendPortBounce

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadius1CriticalAuth) SetSuspendPortBounce(v bool)`

SetSuspendPortBounce sets SuspendPortBounce field to given value.

### HasSuspendPortBounce

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadius1CriticalAuth) HasSuspendPortBounce() bool`

HasSuspendPortBounce returns a boolean if a field has been set.

### GetDataGroupPolicyId

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadius1CriticalAuth) GetDataGroupPolicyId() string`

GetDataGroupPolicyId returns the DataGroupPolicyId field if non-nil, zero value otherwise.

### GetDataGroupPolicyIdOk

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadius1CriticalAuth) GetDataGroupPolicyIdOk() (*string, bool)`

GetDataGroupPolicyIdOk returns a tuple with the DataGroupPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataGroupPolicyId

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadius1CriticalAuth) SetDataGroupPolicyId(v string)`

SetDataGroupPolicyId sets DataGroupPolicyId field to given value.

### HasDataGroupPolicyId

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadius1CriticalAuth) HasDataGroupPolicyId() bool`

HasDataGroupPolicyId returns a boolean if a field has been set.

### SetDataGroupPolicyIdNil

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadius1CriticalAuth) SetDataGroupPolicyIdNil(b bool)`

 SetDataGroupPolicyIdNil sets the value for DataGroupPolicyId to be an explicit nil

### UnsetDataGroupPolicyId
`func (o *NetworksNetworkIdSwitchAccessPoliciesRadius1CriticalAuth) UnsetDataGroupPolicyId()`

UnsetDataGroupPolicyId ensures that no value is present for DataGroupPolicyId, not even an explicit nil
### GetVoiceGroupPolicyId

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadius1CriticalAuth) GetVoiceGroupPolicyId() string`

GetVoiceGroupPolicyId returns the VoiceGroupPolicyId field if non-nil, zero value otherwise.

### GetVoiceGroupPolicyIdOk

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadius1CriticalAuth) GetVoiceGroupPolicyIdOk() (*string, bool)`

GetVoiceGroupPolicyIdOk returns a tuple with the VoiceGroupPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceGroupPolicyId

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadius1CriticalAuth) SetVoiceGroupPolicyId(v string)`

SetVoiceGroupPolicyId sets VoiceGroupPolicyId field to given value.

### HasVoiceGroupPolicyId

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadius1CriticalAuth) HasVoiceGroupPolicyId() bool`

HasVoiceGroupPolicyId returns a boolean if a field has been set.

### SetVoiceGroupPolicyIdNil

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadius1CriticalAuth) SetVoiceGroupPolicyIdNil(b bool)`

 SetVoiceGroupPolicyIdNil sets the value for VoiceGroupPolicyId to be an explicit nil

### UnsetVoiceGroupPolicyId
`func (o *NetworksNetworkIdSwitchAccessPoliciesRadius1CriticalAuth) UnsetVoiceGroupPolicyId()`

UnsetVoiceGroupPolicyId ensures that no value is present for VoiceGroupPolicyId, not even an explicit nil
### GetDataSgtId

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadius1CriticalAuth) GetDataSgtId() int32`

GetDataSgtId returns the DataSgtId field if non-nil, zero value otherwise.

### GetDataSgtIdOk

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadius1CriticalAuth) GetDataSgtIdOk() (*int32, bool)`

GetDataSgtIdOk returns a tuple with the DataSgtId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSgtId

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadius1CriticalAuth) SetDataSgtId(v int32)`

SetDataSgtId sets DataSgtId field to given value.

### HasDataSgtId

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadius1CriticalAuth) HasDataSgtId() bool`

HasDataSgtId returns a boolean if a field has been set.

### SetDataSgtIdNil

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadius1CriticalAuth) SetDataSgtIdNil(b bool)`

 SetDataSgtIdNil sets the value for DataSgtId to be an explicit nil

### UnsetDataSgtId
`func (o *NetworksNetworkIdSwitchAccessPoliciesRadius1CriticalAuth) UnsetDataSgtId()`

UnsetDataSgtId ensures that no value is present for DataSgtId, not even an explicit nil
### GetVoiceSgtId

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadius1CriticalAuth) GetVoiceSgtId() int32`

GetVoiceSgtId returns the VoiceSgtId field if non-nil, zero value otherwise.

### GetVoiceSgtIdOk

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadius1CriticalAuth) GetVoiceSgtIdOk() (*int32, bool)`

GetVoiceSgtIdOk returns a tuple with the VoiceSgtId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceSgtId

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadius1CriticalAuth) SetVoiceSgtId(v int32)`

SetVoiceSgtId sets VoiceSgtId field to given value.

### HasVoiceSgtId

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadius1CriticalAuth) HasVoiceSgtId() bool`

HasVoiceSgtId returns a boolean if a field has been set.

### SetVoiceSgtIdNil

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadius1CriticalAuth) SetVoiceSgtIdNil(b bool)`

 SetVoiceSgtIdNil sets the value for VoiceSgtId to be an explicit nil

### UnsetVoiceSgtId
`func (o *NetworksNetworkIdSwitchAccessPoliciesRadius1CriticalAuth) UnsetVoiceSgtId()`

UnsetVoiceSgtId ensures that no value is present for VoiceSgtId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


