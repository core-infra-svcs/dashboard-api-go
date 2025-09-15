# InlineResponse200280Device

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Serial** | Pointer to **string** | The serial of the device | [optional] 
**Switchports** | Pointer to **string** | The switchports on which to take the packet capture | [optional] 
**Interface** | Pointer to **string** | The interfaces on which to take the packet capture (applicable for Catalyst devices) | [optional] 

## Methods

### NewInlineResponse200280Device

`func NewInlineResponse200280Device() *InlineResponse200280Device`

NewInlineResponse200280Device instantiates a new InlineResponse200280Device object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200280DeviceWithDefaults

`func NewInlineResponse200280DeviceWithDefaults() *InlineResponse200280Device`

NewInlineResponse200280DeviceWithDefaults instantiates a new InlineResponse200280Device object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSerial

`func (o *InlineResponse200280Device) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *InlineResponse200280Device) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *InlineResponse200280Device) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *InlineResponse200280Device) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetSwitchports

`func (o *InlineResponse200280Device) GetSwitchports() string`

GetSwitchports returns the Switchports field if non-nil, zero value otherwise.

### GetSwitchportsOk

`func (o *InlineResponse200280Device) GetSwitchportsOk() (*string, bool)`

GetSwitchportsOk returns a tuple with the Switchports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchports

`func (o *InlineResponse200280Device) SetSwitchports(v string)`

SetSwitchports sets Switchports field to given value.

### HasSwitchports

`func (o *InlineResponse200280Device) HasSwitchports() bool`

HasSwitchports returns a boolean if a field has been set.

### GetInterface

`func (o *InlineResponse200280Device) GetInterface() string`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *InlineResponse200280Device) GetInterfaceOk() (*string, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *InlineResponse200280Device) SetInterface(v string)`

SetInterface sets Interface field to given value.

### HasInterface

`func (o *InlineResponse200280Device) HasInterface() bool`

HasInterface returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


