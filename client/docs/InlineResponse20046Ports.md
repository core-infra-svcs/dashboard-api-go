# InlineResponse20046Ports

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Lldp** | Pointer to [**InlineResponse20046Lldp**](InlineResponse20046Lldp.md) |  | [optional] 
**Cdp** | Pointer to [**InlineResponse20046Cdp**](InlineResponse20046Cdp.md) |  | [optional] 
**DeviceMac** | Pointer to **string** | MAC address for the device | [optional] 
**Device** | Pointer to [**InlineResponse20046Device**](InlineResponse20046Device.md) |  | [optional] 

## Methods

### NewInlineResponse20046Ports

`func NewInlineResponse20046Ports() *InlineResponse20046Ports`

NewInlineResponse20046Ports instantiates a new InlineResponse20046Ports object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20046PortsWithDefaults

`func NewInlineResponse20046PortsWithDefaults() *InlineResponse20046Ports`

NewInlineResponse20046PortsWithDefaults instantiates a new InlineResponse20046Ports object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLldp

`func (o *InlineResponse20046Ports) GetLldp() InlineResponse20046Lldp`

GetLldp returns the Lldp field if non-nil, zero value otherwise.

### GetLldpOk

`func (o *InlineResponse20046Ports) GetLldpOk() (*InlineResponse20046Lldp, bool)`

GetLldpOk returns a tuple with the Lldp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLldp

`func (o *InlineResponse20046Ports) SetLldp(v InlineResponse20046Lldp)`

SetLldp sets Lldp field to given value.

### HasLldp

`func (o *InlineResponse20046Ports) HasLldp() bool`

HasLldp returns a boolean if a field has been set.

### GetCdp

`func (o *InlineResponse20046Ports) GetCdp() InlineResponse20046Cdp`

GetCdp returns the Cdp field if non-nil, zero value otherwise.

### GetCdpOk

`func (o *InlineResponse20046Ports) GetCdpOk() (*InlineResponse20046Cdp, bool)`

GetCdpOk returns a tuple with the Cdp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdp

`func (o *InlineResponse20046Ports) SetCdp(v InlineResponse20046Cdp)`

SetCdp sets Cdp field to given value.

### HasCdp

`func (o *InlineResponse20046Ports) HasCdp() bool`

HasCdp returns a boolean if a field has been set.

### GetDeviceMac

`func (o *InlineResponse20046Ports) GetDeviceMac() string`

GetDeviceMac returns the DeviceMac field if non-nil, zero value otherwise.

### GetDeviceMacOk

`func (o *InlineResponse20046Ports) GetDeviceMacOk() (*string, bool)`

GetDeviceMacOk returns a tuple with the DeviceMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceMac

`func (o *InlineResponse20046Ports) SetDeviceMac(v string)`

SetDeviceMac sets DeviceMac field to given value.

### HasDeviceMac

`func (o *InlineResponse20046Ports) HasDeviceMac() bool`

HasDeviceMac returns a boolean if a field has been set.

### GetDevice

`func (o *InlineResponse20046Ports) GetDevice() InlineResponse20046Device`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *InlineResponse20046Ports) GetDeviceOk() (*InlineResponse20046Device, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *InlineResponse20046Ports) SetDevice(v InlineResponse20046Device)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *InlineResponse20046Ports) HasDevice() bool`

HasDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


