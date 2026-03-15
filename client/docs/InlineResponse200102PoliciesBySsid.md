# InlineResponse200102PoliciesBySsid

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SsidNumber** | Pointer to **int32** | The SSID number for the policy rule | [optional] 
**DevicePolicy** | Pointer to **string** | The device policy applied to the client for this SSID | [optional] 
**GroupPolicyId** | Pointer to **string** | The group policy identifier for this SSID | [optional] 

## Methods

### NewInlineResponse200102PoliciesBySsid

`func NewInlineResponse200102PoliciesBySsid() *InlineResponse200102PoliciesBySsid`

NewInlineResponse200102PoliciesBySsid instantiates a new InlineResponse200102PoliciesBySsid object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200102PoliciesBySsidWithDefaults

`func NewInlineResponse200102PoliciesBySsidWithDefaults() *InlineResponse200102PoliciesBySsid`

NewInlineResponse200102PoliciesBySsidWithDefaults instantiates a new InlineResponse200102PoliciesBySsid object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSsidNumber

`func (o *InlineResponse200102PoliciesBySsid) GetSsidNumber() int32`

GetSsidNumber returns the SsidNumber field if non-nil, zero value otherwise.

### GetSsidNumberOk

`func (o *InlineResponse200102PoliciesBySsid) GetSsidNumberOk() (*int32, bool)`

GetSsidNumberOk returns a tuple with the SsidNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsidNumber

`func (o *InlineResponse200102PoliciesBySsid) SetSsidNumber(v int32)`

SetSsidNumber sets SsidNumber field to given value.

### HasSsidNumber

`func (o *InlineResponse200102PoliciesBySsid) HasSsidNumber() bool`

HasSsidNumber returns a boolean if a field has been set.

### GetDevicePolicy

`func (o *InlineResponse200102PoliciesBySsid) GetDevicePolicy() string`

GetDevicePolicy returns the DevicePolicy field if non-nil, zero value otherwise.

### GetDevicePolicyOk

`func (o *InlineResponse200102PoliciesBySsid) GetDevicePolicyOk() (*string, bool)`

GetDevicePolicyOk returns a tuple with the DevicePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevicePolicy

`func (o *InlineResponse200102PoliciesBySsid) SetDevicePolicy(v string)`

SetDevicePolicy sets DevicePolicy field to given value.

### HasDevicePolicy

`func (o *InlineResponse200102PoliciesBySsid) HasDevicePolicy() bool`

HasDevicePolicy returns a boolean if a field has been set.

### GetGroupPolicyId

`func (o *InlineResponse200102PoliciesBySsid) GetGroupPolicyId() string`

GetGroupPolicyId returns the GroupPolicyId field if non-nil, zero value otherwise.

### GetGroupPolicyIdOk

`func (o *InlineResponse200102PoliciesBySsid) GetGroupPolicyIdOk() (*string, bool)`

GetGroupPolicyIdOk returns a tuple with the GroupPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupPolicyId

`func (o *InlineResponse200102PoliciesBySsid) SetGroupPolicyId(v string)`

SetGroupPolicyId sets GroupPolicyId field to given value.

### HasGroupPolicyId

`func (o *InlineResponse200102PoliciesBySsid) HasGroupPolicyId() bool`

HasGroupPolicyId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


