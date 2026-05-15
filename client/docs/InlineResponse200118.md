# InlineResponse200118

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mac** | Pointer to **string** | The MAC address of the client | [optional] 
**DevicePolicy** | Pointer to **string** | The name of the client&#39;s policy | [optional] 
**GroupPolicyId** | Pointer to **string** | The group policy identifier of the client | [optional] 
**PoliciesBySsid** | Pointer to [**[]InlineResponse200118PoliciesBySsid**](InlineResponse200118PoliciesBySsid.md) | The policies applied to the client by SSID | [optional] 

## Methods

### NewInlineResponse200118

`func NewInlineResponse200118() *InlineResponse200118`

NewInlineResponse200118 instantiates a new InlineResponse200118 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200118WithDefaults

`func NewInlineResponse200118WithDefaults() *InlineResponse200118`

NewInlineResponse200118WithDefaults instantiates a new InlineResponse200118 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMac

`func (o *InlineResponse200118) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *InlineResponse200118) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *InlineResponse200118) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *InlineResponse200118) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetDevicePolicy

`func (o *InlineResponse200118) GetDevicePolicy() string`

GetDevicePolicy returns the DevicePolicy field if non-nil, zero value otherwise.

### GetDevicePolicyOk

`func (o *InlineResponse200118) GetDevicePolicyOk() (*string, bool)`

GetDevicePolicyOk returns a tuple with the DevicePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevicePolicy

`func (o *InlineResponse200118) SetDevicePolicy(v string)`

SetDevicePolicy sets DevicePolicy field to given value.

### HasDevicePolicy

`func (o *InlineResponse200118) HasDevicePolicy() bool`

HasDevicePolicy returns a boolean if a field has been set.

### GetGroupPolicyId

`func (o *InlineResponse200118) GetGroupPolicyId() string`

GetGroupPolicyId returns the GroupPolicyId field if non-nil, zero value otherwise.

### GetGroupPolicyIdOk

`func (o *InlineResponse200118) GetGroupPolicyIdOk() (*string, bool)`

GetGroupPolicyIdOk returns a tuple with the GroupPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupPolicyId

`func (o *InlineResponse200118) SetGroupPolicyId(v string)`

SetGroupPolicyId sets GroupPolicyId field to given value.

### HasGroupPolicyId

`func (o *InlineResponse200118) HasGroupPolicyId() bool`

HasGroupPolicyId returns a boolean if a field has been set.

### GetPoliciesBySsid

`func (o *InlineResponse200118) GetPoliciesBySsid() []InlineResponse200118PoliciesBySsid`

GetPoliciesBySsid returns the PoliciesBySsid field if non-nil, zero value otherwise.

### GetPoliciesBySsidOk

`func (o *InlineResponse200118) GetPoliciesBySsidOk() (*[]InlineResponse200118PoliciesBySsid, bool)`

GetPoliciesBySsidOk returns a tuple with the PoliciesBySsid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoliciesBySsid

`func (o *InlineResponse200118) SetPoliciesBySsid(v []InlineResponse200118PoliciesBySsid)`

SetPoliciesBySsid sets PoliciesBySsid field to given value.

### HasPoliciesBySsid

`func (o *InlineResponse200118) HasPoliciesBySsid() bool`

HasPoliciesBySsid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


