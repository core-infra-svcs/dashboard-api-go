# InlineResponse20041Ports

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Lldp** | Pointer to [**InlineResponse20041Lldp**](InlineResponse20041Lldp.md) |  | [optional] 
**Cdp** | Pointer to [**InlineResponse20041Cdp**](InlineResponse20041Cdp.md) |  | [optional] 
**DeviceMac** | Pointer to **string** | MAC address for the device | [optional] 
**Device** | Pointer to [**InlineResponse20041Device**](InlineResponse20041Device.md) |  | [optional] 

## Methods

### NewInlineResponse20041Ports

`func NewInlineResponse20041Ports() *InlineResponse20041Ports`

NewInlineResponse20041Ports instantiates a new InlineResponse20041Ports object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20041PortsWithDefaults

`func NewInlineResponse20041PortsWithDefaults() *InlineResponse20041Ports`

NewInlineResponse20041PortsWithDefaults instantiates a new InlineResponse20041Ports object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLldp

`func (o *InlineResponse20041Ports) GetLldp() InlineResponse20041Lldp`

GetLldp returns the Lldp field if non-nil, zero value otherwise.

### GetLldpOk

`func (o *InlineResponse20041Ports) GetLldpOk() (*InlineResponse20041Lldp, bool)`

GetLldpOk returns a tuple with the Lldp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLldp

`func (o *InlineResponse20041Ports) SetLldp(v InlineResponse20041Lldp)`

SetLldp sets Lldp field to given value.

### HasLldp

`func (o *InlineResponse20041Ports) HasLldp() bool`

HasLldp returns a boolean if a field has been set.

### GetCdp

`func (o *InlineResponse20041Ports) GetCdp() InlineResponse20041Cdp`

GetCdp returns the Cdp field if non-nil, zero value otherwise.

### GetCdpOk

`func (o *InlineResponse20041Ports) GetCdpOk() (*InlineResponse20041Cdp, bool)`

GetCdpOk returns a tuple with the Cdp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdp

`func (o *InlineResponse20041Ports) SetCdp(v InlineResponse20041Cdp)`

SetCdp sets Cdp field to given value.

### HasCdp

`func (o *InlineResponse20041Ports) HasCdp() bool`

HasCdp returns a boolean if a field has been set.

### GetDeviceMac

`func (o *InlineResponse20041Ports) GetDeviceMac() string`

GetDeviceMac returns the DeviceMac field if non-nil, zero value otherwise.

### GetDeviceMacOk

`func (o *InlineResponse20041Ports) GetDeviceMacOk() (*string, bool)`

GetDeviceMacOk returns a tuple with the DeviceMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceMac

`func (o *InlineResponse20041Ports) SetDeviceMac(v string)`

SetDeviceMac sets DeviceMac field to given value.

### HasDeviceMac

`func (o *InlineResponse20041Ports) HasDeviceMac() bool`

HasDeviceMac returns a boolean if a field has been set.

### GetDevice

`func (o *InlineResponse20041Ports) GetDevice() InlineResponse20041Device`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *InlineResponse20041Ports) GetDeviceOk() (*InlineResponse20041Device, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *InlineResponse20041Ports) SetDevice(v InlineResponse20041Device)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *InlineResponse20041Ports) HasDevice() bool`

HasDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


