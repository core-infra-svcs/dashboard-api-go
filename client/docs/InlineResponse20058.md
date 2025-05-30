# InlineResponse20058

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Number** | Pointer to **int32** | Number of the port | [optional] 
**Enabled** | Pointer to **bool** | The status of the port | [optional] 
**Type** | Pointer to **string** | The type of the port: &#39;access&#39; or &#39;trunk&#39;. | [optional] 
**DropUntaggedTraffic** | Pointer to **bool** | Whether the trunk port can drop all untagged traffic. | [optional] 
**Vlan** | Pointer to **int32** | Native VLAN when the port is in Trunk mode. Access VLAN when the port is in Access mode. | [optional] 
**AllowedVlans** | Pointer to **string** | Comma-delimited list of the VLAN ID&#39;s allowed on the port, or &#39;all&#39; to permit all VLAN&#39;s on the port. | [optional] 
**AccessPolicy** | Pointer to **string** | The name of the policy. Only applicable to Access ports. | [optional] 

## Methods

### NewInlineResponse20058

`func NewInlineResponse20058() *InlineResponse20058`

NewInlineResponse20058 instantiates a new InlineResponse20058 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20058WithDefaults

`func NewInlineResponse20058WithDefaults() *InlineResponse20058`

NewInlineResponse20058WithDefaults instantiates a new InlineResponse20058 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumber

`func (o *InlineResponse20058) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *InlineResponse20058) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *InlineResponse20058) SetNumber(v int32)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *InlineResponse20058) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetEnabled

`func (o *InlineResponse20058) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InlineResponse20058) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InlineResponse20058) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InlineResponse20058) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetType

`func (o *InlineResponse20058) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InlineResponse20058) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InlineResponse20058) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *InlineResponse20058) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDropUntaggedTraffic

`func (o *InlineResponse20058) GetDropUntaggedTraffic() bool`

GetDropUntaggedTraffic returns the DropUntaggedTraffic field if non-nil, zero value otherwise.

### GetDropUntaggedTrafficOk

`func (o *InlineResponse20058) GetDropUntaggedTrafficOk() (*bool, bool)`

GetDropUntaggedTrafficOk returns a tuple with the DropUntaggedTraffic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDropUntaggedTraffic

`func (o *InlineResponse20058) SetDropUntaggedTraffic(v bool)`

SetDropUntaggedTraffic sets DropUntaggedTraffic field to given value.

### HasDropUntaggedTraffic

`func (o *InlineResponse20058) HasDropUntaggedTraffic() bool`

HasDropUntaggedTraffic returns a boolean if a field has been set.

### GetVlan

`func (o *InlineResponse20058) GetVlan() int32`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *InlineResponse20058) GetVlanOk() (*int32, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *InlineResponse20058) SetVlan(v int32)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *InlineResponse20058) HasVlan() bool`

HasVlan returns a boolean if a field has been set.

### GetAllowedVlans

`func (o *InlineResponse20058) GetAllowedVlans() string`

GetAllowedVlans returns the AllowedVlans field if non-nil, zero value otherwise.

### GetAllowedVlansOk

`func (o *InlineResponse20058) GetAllowedVlansOk() (*string, bool)`

GetAllowedVlansOk returns a tuple with the AllowedVlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedVlans

`func (o *InlineResponse20058) SetAllowedVlans(v string)`

SetAllowedVlans sets AllowedVlans field to given value.

### HasAllowedVlans

`func (o *InlineResponse20058) HasAllowedVlans() bool`

HasAllowedVlans returns a boolean if a field has been set.

### GetAccessPolicy

`func (o *InlineResponse20058) GetAccessPolicy() string`

GetAccessPolicy returns the AccessPolicy field if non-nil, zero value otherwise.

### GetAccessPolicyOk

`func (o *InlineResponse20058) GetAccessPolicyOk() (*string, bool)`

GetAccessPolicyOk returns a tuple with the AccessPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessPolicy

`func (o *InlineResponse20058) SetAccessPolicy(v string)`

SetAccessPolicy sets AccessPolicy field to given value.

### HasAccessPolicy

`func (o *InlineResponse20058) HasAccessPolicy() bool`

HasAccessPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


