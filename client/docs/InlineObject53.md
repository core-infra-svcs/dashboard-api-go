# InlineObject53

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | The status of the port | [optional] 
**DropUntaggedTraffic** | Pointer to **bool** | Trunk port can Drop all Untagged traffic. When true, no VLAN is required. Access ports cannot have dropUntaggedTraffic set to true. | [optional] 
**Type** | Pointer to **string** | The type of the port: &#39;access&#39; or &#39;trunk&#39;. | [optional] 
**Vlan** | Pointer to **int32** | Native VLAN when the port is in Trunk mode. Access VLAN when the port is in Access mode. | [optional] 
**AllowedVlans** | Pointer to **string** | Comma-delimited list of the VLAN ID&#39;s allowed on the port, or &#39;all&#39; to permit all VLAN&#39;s on the port. | [optional] 
**AccessPolicy** | Pointer to **string** | The name of the policy. Only applicable to Access ports. Valid values are: &#39;open&#39;, &#39;8021x-radius&#39;, &#39;mac-radius&#39;, &#39;hybris-radius&#39; for MX64 or Z3 or any MX supporting the per port authentication feature. Otherwise, &#39;open&#39; is the only valid value and &#39;open&#39; is the default value if the field is missing. | [optional] 

## Methods

### NewInlineObject53

`func NewInlineObject53() *InlineObject53`

NewInlineObject53 instantiates a new InlineObject53 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject53WithDefaults

`func NewInlineObject53WithDefaults() *InlineObject53`

NewInlineObject53WithDefaults instantiates a new InlineObject53 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *InlineObject53) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InlineObject53) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InlineObject53) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InlineObject53) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetDropUntaggedTraffic

`func (o *InlineObject53) GetDropUntaggedTraffic() bool`

GetDropUntaggedTraffic returns the DropUntaggedTraffic field if non-nil, zero value otherwise.

### GetDropUntaggedTrafficOk

`func (o *InlineObject53) GetDropUntaggedTrafficOk() (*bool, bool)`

GetDropUntaggedTrafficOk returns a tuple with the DropUntaggedTraffic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDropUntaggedTraffic

`func (o *InlineObject53) SetDropUntaggedTraffic(v bool)`

SetDropUntaggedTraffic sets DropUntaggedTraffic field to given value.

### HasDropUntaggedTraffic

`func (o *InlineObject53) HasDropUntaggedTraffic() bool`

HasDropUntaggedTraffic returns a boolean if a field has been set.

### GetType

`func (o *InlineObject53) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InlineObject53) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InlineObject53) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *InlineObject53) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVlan

`func (o *InlineObject53) GetVlan() int32`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *InlineObject53) GetVlanOk() (*int32, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *InlineObject53) SetVlan(v int32)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *InlineObject53) HasVlan() bool`

HasVlan returns a boolean if a field has been set.

### GetAllowedVlans

`func (o *InlineObject53) GetAllowedVlans() string`

GetAllowedVlans returns the AllowedVlans field if non-nil, zero value otherwise.

### GetAllowedVlansOk

`func (o *InlineObject53) GetAllowedVlansOk() (*string, bool)`

GetAllowedVlansOk returns a tuple with the AllowedVlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedVlans

`func (o *InlineObject53) SetAllowedVlans(v string)`

SetAllowedVlans sets AllowedVlans field to given value.

### HasAllowedVlans

`func (o *InlineObject53) HasAllowedVlans() bool`

HasAllowedVlans returns a boolean if a field has been set.

### GetAccessPolicy

`func (o *InlineObject53) GetAccessPolicy() string`

GetAccessPolicy returns the AccessPolicy field if non-nil, zero value otherwise.

### GetAccessPolicyOk

`func (o *InlineObject53) GetAccessPolicyOk() (*string, bool)`

GetAccessPolicyOk returns a tuple with the AccessPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessPolicy

`func (o *InlineObject53) SetAccessPolicy(v string)`

SetAccessPolicy sets AccessPolicy field to given value.

### HasAccessPolicy

`func (o *InlineObject53) HasAccessPolicy() bool`

HasAccessPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


