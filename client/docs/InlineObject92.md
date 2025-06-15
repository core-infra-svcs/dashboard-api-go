# InlineObject92

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DevicePolicy** | **string** | The policy to assign. Can be &#39;Whitelisted&#39;, &#39;Blocked&#39;, &#39;Normal&#39; or &#39;Group policy&#39;. Required. | 
**GroupPolicyId** | Pointer to **string** | [Optional] If &#39;devicePolicy&#39; is set to &#39;Group policy&#39; this param is used to specify the group policy ID. | [optional] 

## Methods

### NewInlineObject92

`func NewInlineObject92(devicePolicy string, ) *InlineObject92`

NewInlineObject92 instantiates a new InlineObject92 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject92WithDefaults

`func NewInlineObject92WithDefaults() *InlineObject92`

NewInlineObject92WithDefaults instantiates a new InlineObject92 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDevicePolicy

`func (o *InlineObject92) GetDevicePolicy() string`

GetDevicePolicy returns the DevicePolicy field if non-nil, zero value otherwise.

### GetDevicePolicyOk

`func (o *InlineObject92) GetDevicePolicyOk() (*string, bool)`

GetDevicePolicyOk returns a tuple with the DevicePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevicePolicy

`func (o *InlineObject92) SetDevicePolicy(v string)`

SetDevicePolicy sets DevicePolicy field to given value.


### GetGroupPolicyId

`func (o *InlineObject92) GetGroupPolicyId() string`

GetGroupPolicyId returns the GroupPolicyId field if non-nil, zero value otherwise.

### GetGroupPolicyIdOk

`func (o *InlineObject92) GetGroupPolicyIdOk() (*string, bool)`

GetGroupPolicyIdOk returns a tuple with the GroupPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupPolicyId

`func (o *InlineObject92) SetGroupPolicyId(v string)`

SetGroupPolicyId sets GroupPolicyId field to given value.

### HasGroupPolicyId

`func (o *InlineObject92) HasGroupPolicyId() bool`

HasGroupPolicyId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


