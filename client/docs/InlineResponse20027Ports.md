# InlineResponse20027Ports

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Lldp** | Pointer to [**InlineResponse20027Lldp**](InlineResponse20027Lldp.md) |  | [optional] 
**Cdp** | Pointer to [**InlineResponse20027Cdp**](InlineResponse20027Cdp.md) |  | [optional] 
**DeviceMac** | Pointer to **string** | MAC address for the device | [optional] 
**Device** | Pointer to [**InlineResponse20027Device**](InlineResponse20027Device.md) |  | [optional] 

## Methods

### NewInlineResponse20027Ports

`func NewInlineResponse20027Ports() *InlineResponse20027Ports`

NewInlineResponse20027Ports instantiates a new InlineResponse20027Ports object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20027PortsWithDefaults

`func NewInlineResponse20027PortsWithDefaults() *InlineResponse20027Ports`

NewInlineResponse20027PortsWithDefaults instantiates a new InlineResponse20027Ports object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLldp

`func (o *InlineResponse20027Ports) GetLldp() InlineResponse20027Lldp`

GetLldp returns the Lldp field if non-nil, zero value otherwise.

### GetLldpOk

`func (o *InlineResponse20027Ports) GetLldpOk() (*InlineResponse20027Lldp, bool)`

GetLldpOk returns a tuple with the Lldp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLldp

`func (o *InlineResponse20027Ports) SetLldp(v InlineResponse20027Lldp)`

SetLldp sets Lldp field to given value.

### HasLldp

`func (o *InlineResponse20027Ports) HasLldp() bool`

HasLldp returns a boolean if a field has been set.

### GetCdp

`func (o *InlineResponse20027Ports) GetCdp() InlineResponse20027Cdp`

GetCdp returns the Cdp field if non-nil, zero value otherwise.

### GetCdpOk

`func (o *InlineResponse20027Ports) GetCdpOk() (*InlineResponse20027Cdp, bool)`

GetCdpOk returns a tuple with the Cdp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdp

`func (o *InlineResponse20027Ports) SetCdp(v InlineResponse20027Cdp)`

SetCdp sets Cdp field to given value.

### HasCdp

`func (o *InlineResponse20027Ports) HasCdp() bool`

HasCdp returns a boolean if a field has been set.

### GetDeviceMac

`func (o *InlineResponse20027Ports) GetDeviceMac() string`

GetDeviceMac returns the DeviceMac field if non-nil, zero value otherwise.

### GetDeviceMacOk

`func (o *InlineResponse20027Ports) GetDeviceMacOk() (*string, bool)`

GetDeviceMacOk returns a tuple with the DeviceMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceMac

`func (o *InlineResponse20027Ports) SetDeviceMac(v string)`

SetDeviceMac sets DeviceMac field to given value.

### HasDeviceMac

`func (o *InlineResponse20027Ports) HasDeviceMac() bool`

HasDeviceMac returns a boolean if a field has been set.

### GetDevice

`func (o *InlineResponse20027Ports) GetDevice() InlineResponse20027Device`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *InlineResponse20027Ports) GetDeviceOk() (*InlineResponse20027Device, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *InlineResponse20027Ports) SetDevice(v InlineResponse20027Device)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *InlineResponse20027Ports) HasDevice() bool`

HasDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


