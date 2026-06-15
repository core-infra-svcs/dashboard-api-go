# InlineResponse200126PoliciesBySsid

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SsidNumber** | Pointer to **int32** | The SSID number for the policy rule | [optional] 
**DevicePolicy** | Pointer to **string** | The device policy applied to the client for this SSID | [optional] 
**GroupPolicyId** | Pointer to **string** | The group policy identifier for this SSID | [optional] 

## Methods

### NewInlineResponse200126PoliciesBySsid

`func NewInlineResponse200126PoliciesBySsid() *InlineResponse200126PoliciesBySsid`

NewInlineResponse200126PoliciesBySsid instantiates a new InlineResponse200126PoliciesBySsid object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200126PoliciesBySsidWithDefaults

`func NewInlineResponse200126PoliciesBySsidWithDefaults() *InlineResponse200126PoliciesBySsid`

NewInlineResponse200126PoliciesBySsidWithDefaults instantiates a new InlineResponse200126PoliciesBySsid object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSsidNumber

`func (o *InlineResponse200126PoliciesBySsid) GetSsidNumber() int32`

GetSsidNumber returns the SsidNumber field if non-nil, zero value otherwise.

### GetSsidNumberOk

`func (o *InlineResponse200126PoliciesBySsid) GetSsidNumberOk() (*int32, bool)`

GetSsidNumberOk returns a tuple with the SsidNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsidNumber

`func (o *InlineResponse200126PoliciesBySsid) SetSsidNumber(v int32)`

SetSsidNumber sets SsidNumber field to given value.

### HasSsidNumber

`func (o *InlineResponse200126PoliciesBySsid) HasSsidNumber() bool`

HasSsidNumber returns a boolean if a field has been set.

### GetDevicePolicy

`func (o *InlineResponse200126PoliciesBySsid) GetDevicePolicy() string`

GetDevicePolicy returns the DevicePolicy field if non-nil, zero value otherwise.

### GetDevicePolicyOk

`func (o *InlineResponse200126PoliciesBySsid) GetDevicePolicyOk() (*string, bool)`

GetDevicePolicyOk returns a tuple with the DevicePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevicePolicy

`func (o *InlineResponse200126PoliciesBySsid) SetDevicePolicy(v string)`

SetDevicePolicy sets DevicePolicy field to given value.

### HasDevicePolicy

`func (o *InlineResponse200126PoliciesBySsid) HasDevicePolicy() bool`

HasDevicePolicy returns a boolean if a field has been set.

### GetGroupPolicyId

`func (o *InlineResponse200126PoliciesBySsid) GetGroupPolicyId() string`

GetGroupPolicyId returns the GroupPolicyId field if non-nil, zero value otherwise.

### GetGroupPolicyIdOk

`func (o *InlineResponse200126PoliciesBySsid) GetGroupPolicyIdOk() (*string, bool)`

GetGroupPolicyIdOk returns a tuple with the GroupPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupPolicyId

`func (o *InlineResponse200126PoliciesBySsid) SetGroupPolicyId(v string)`

SetGroupPolicyId sets GroupPolicyId field to given value.

### HasGroupPolicyId

`func (o *InlineResponse200126PoliciesBySsid) HasGroupPolicyId() bool`

HasGroupPolicyId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


