# InlineResponse200249DeviceTypePolicies

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceType** | Pointer to **string** | The device type. Can be one of &#39;Android&#39;, &#39;BlackBerry&#39;, &#39;Chrome OS&#39;, &#39;iPad&#39;, &#39;iPhone&#39;, &#39;iPod&#39;, &#39;Mac OS X&#39;, &#39;Windows&#39;, &#39;Windows Phone&#39;, &#39;B&amp;N Nook&#39; or &#39;Other OS&#39; | [optional] 
**DevicePolicy** | Pointer to **string** | The device policy. Can be one of &#39;Allowed&#39;, &#39;Blocked&#39; or &#39;Group policy&#39; | [optional] 
**GroupPolicyId** | Pointer to **int32** | ID of the group policy object. Only present when the device policy is \&quot;Group policy\&quot;. | [optional] 

## Methods

### NewInlineResponse200249DeviceTypePolicies

`func NewInlineResponse200249DeviceTypePolicies() *InlineResponse200249DeviceTypePolicies`

NewInlineResponse200249DeviceTypePolicies instantiates a new InlineResponse200249DeviceTypePolicies object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200249DeviceTypePoliciesWithDefaults

`func NewInlineResponse200249DeviceTypePoliciesWithDefaults() *InlineResponse200249DeviceTypePolicies`

NewInlineResponse200249DeviceTypePoliciesWithDefaults instantiates a new InlineResponse200249DeviceTypePolicies object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceType

`func (o *InlineResponse200249DeviceTypePolicies) GetDeviceType() string`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *InlineResponse200249DeviceTypePolicies) GetDeviceTypeOk() (*string, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *InlineResponse200249DeviceTypePolicies) SetDeviceType(v string)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *InlineResponse200249DeviceTypePolicies) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### GetDevicePolicy

`func (o *InlineResponse200249DeviceTypePolicies) GetDevicePolicy() string`

GetDevicePolicy returns the DevicePolicy field if non-nil, zero value otherwise.

### GetDevicePolicyOk

`func (o *InlineResponse200249DeviceTypePolicies) GetDevicePolicyOk() (*string, bool)`

GetDevicePolicyOk returns a tuple with the DevicePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevicePolicy

`func (o *InlineResponse200249DeviceTypePolicies) SetDevicePolicy(v string)`

SetDevicePolicy sets DevicePolicy field to given value.

### HasDevicePolicy

`func (o *InlineResponse200249DeviceTypePolicies) HasDevicePolicy() bool`

HasDevicePolicy returns a boolean if a field has been set.

### GetGroupPolicyId

`func (o *InlineResponse200249DeviceTypePolicies) GetGroupPolicyId() int32`

GetGroupPolicyId returns the GroupPolicyId field if non-nil, zero value otherwise.

### GetGroupPolicyIdOk

`func (o *InlineResponse200249DeviceTypePolicies) GetGroupPolicyIdOk() (*int32, bool)`

GetGroupPolicyIdOk returns a tuple with the GroupPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupPolicyId

`func (o *InlineResponse200249DeviceTypePolicies) SetGroupPolicyId(v int32)`

SetGroupPolicyId sets GroupPolicyId field to given value.

### HasGroupPolicyId

`func (o *InlineResponse200249DeviceTypePolicies) HasGroupPolicyId() bool`

HasGroupPolicyId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


