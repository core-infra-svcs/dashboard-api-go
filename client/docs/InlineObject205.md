# InlineObject205

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceSerial** | Pointer to **string** | The serial number of the device to assign this license to. Set this to  null to unassign the license. If a different license is already active on the device, this parameter will control queueing/dequeuing this license. | [optional] 

## Methods

### NewInlineObject205

`func NewInlineObject205() *InlineObject205`

NewInlineObject205 instantiates a new InlineObject205 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject205WithDefaults

`func NewInlineObject205WithDefaults() *InlineObject205`

NewInlineObject205WithDefaults instantiates a new InlineObject205 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceSerial

`func (o *InlineObject205) GetDeviceSerial() string`

GetDeviceSerial returns the DeviceSerial field if non-nil, zero value otherwise.

### GetDeviceSerialOk

`func (o *InlineObject205) GetDeviceSerialOk() (*string, bool)`

GetDeviceSerialOk returns a tuple with the DeviceSerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceSerial

`func (o *InlineObject205) SetDeviceSerial(v string)`

SetDeviceSerial sets DeviceSerial field to given value.

### HasDeviceSerial

`func (o *InlineObject205) HasDeviceSerial() bool`

HasDeviceSerial returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


