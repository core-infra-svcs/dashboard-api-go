# InlineResponse200249

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | If true, the SSID device type group policies are enabled. | [optional] 
**DeviceTypePolicies** | Pointer to [**[]InlineResponse200249DeviceTypePolicies**](InlineResponse200249DeviceTypePolicies.md) | List of device type policies. | [optional] 

## Methods

### NewInlineResponse200249

`func NewInlineResponse200249() *InlineResponse200249`

NewInlineResponse200249 instantiates a new InlineResponse200249 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200249WithDefaults

`func NewInlineResponse200249WithDefaults() *InlineResponse200249`

NewInlineResponse200249WithDefaults instantiates a new InlineResponse200249 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *InlineResponse200249) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InlineResponse200249) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InlineResponse200249) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InlineResponse200249) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetDeviceTypePolicies

`func (o *InlineResponse200249) GetDeviceTypePolicies() []InlineResponse200249DeviceTypePolicies`

GetDeviceTypePolicies returns the DeviceTypePolicies field if non-nil, zero value otherwise.

### GetDeviceTypePoliciesOk

`func (o *InlineResponse200249) GetDeviceTypePoliciesOk() (*[]InlineResponse200249DeviceTypePolicies, bool)`

GetDeviceTypePoliciesOk returns a tuple with the DeviceTypePolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceTypePolicies

`func (o *InlineResponse200249) SetDeviceTypePolicies(v []InlineResponse200249DeviceTypePolicies)`

SetDeviceTypePolicies sets DeviceTypePolicies field to given value.

### HasDeviceTypePolicies

`func (o *InlineResponse200249) HasDeviceTypePolicies() bool`

HasDeviceTypePolicies returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


