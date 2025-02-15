# InlineResponse20029Ports

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Lldp** | Pointer to [**InlineResponse20029Lldp**](InlineResponse20029Lldp.md) |  | [optional] 
**Cdp** | Pointer to [**InlineResponse20029Cdp**](InlineResponse20029Cdp.md) |  | [optional] 
**DeviceMac** | Pointer to **string** | MAC address for the device | [optional] 
**Device** | Pointer to [**InlineResponse20029Device**](InlineResponse20029Device.md) |  | [optional] 

## Methods

### NewInlineResponse20029Ports

`func NewInlineResponse20029Ports() *InlineResponse20029Ports`

NewInlineResponse20029Ports instantiates a new InlineResponse20029Ports object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20029PortsWithDefaults

`func NewInlineResponse20029PortsWithDefaults() *InlineResponse20029Ports`

NewInlineResponse20029PortsWithDefaults instantiates a new InlineResponse20029Ports object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLldp

`func (o *InlineResponse20029Ports) GetLldp() InlineResponse20029Lldp`

GetLldp returns the Lldp field if non-nil, zero value otherwise.

### GetLldpOk

`func (o *InlineResponse20029Ports) GetLldpOk() (*InlineResponse20029Lldp, bool)`

GetLldpOk returns a tuple with the Lldp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLldp

`func (o *InlineResponse20029Ports) SetLldp(v InlineResponse20029Lldp)`

SetLldp sets Lldp field to given value.

### HasLldp

`func (o *InlineResponse20029Ports) HasLldp() bool`

HasLldp returns a boolean if a field has been set.

### GetCdp

`func (o *InlineResponse20029Ports) GetCdp() InlineResponse20029Cdp`

GetCdp returns the Cdp field if non-nil, zero value otherwise.

### GetCdpOk

`func (o *InlineResponse20029Ports) GetCdpOk() (*InlineResponse20029Cdp, bool)`

GetCdpOk returns a tuple with the Cdp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdp

`func (o *InlineResponse20029Ports) SetCdp(v InlineResponse20029Cdp)`

SetCdp sets Cdp field to given value.

### HasCdp

`func (o *InlineResponse20029Ports) HasCdp() bool`

HasCdp returns a boolean if a field has been set.

### GetDeviceMac

`func (o *InlineResponse20029Ports) GetDeviceMac() string`

GetDeviceMac returns the DeviceMac field if non-nil, zero value otherwise.

### GetDeviceMacOk

`func (o *InlineResponse20029Ports) GetDeviceMacOk() (*string, bool)`

GetDeviceMacOk returns a tuple with the DeviceMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceMac

`func (o *InlineResponse20029Ports) SetDeviceMac(v string)`

SetDeviceMac sets DeviceMac field to given value.

### HasDeviceMac

`func (o *InlineResponse20029Ports) HasDeviceMac() bool`

HasDeviceMac returns a boolean if a field has been set.

### GetDevice

`func (o *InlineResponse20029Ports) GetDevice() InlineResponse20029Device`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *InlineResponse20029Ports) GetDeviceOk() (*InlineResponse20029Device, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *InlineResponse20029Ports) SetDevice(v InlineResponse20029Device)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *InlineResponse20029Ports) HasDevice() bool`

HasDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


