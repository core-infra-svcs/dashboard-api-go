# InlineResponse20026Ports

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Lldp** | Pointer to [**InlineResponse20026Lldp**](InlineResponse20026Lldp.md) |  | [optional] 
**Cdp** | Pointer to [**InlineResponse20026Cdp**](InlineResponse20026Cdp.md) |  | [optional] 
**DeviceMac** | Pointer to **string** | MAC address for the device | [optional] 
**Device** | Pointer to [**InlineResponse20026Device**](InlineResponse20026Device.md) |  | [optional] 

## Methods

### NewInlineResponse20026Ports

`func NewInlineResponse20026Ports() *InlineResponse20026Ports`

NewInlineResponse20026Ports instantiates a new InlineResponse20026Ports object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20026PortsWithDefaults

`func NewInlineResponse20026PortsWithDefaults() *InlineResponse20026Ports`

NewInlineResponse20026PortsWithDefaults instantiates a new InlineResponse20026Ports object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLldp

`func (o *InlineResponse20026Ports) GetLldp() InlineResponse20026Lldp`

GetLldp returns the Lldp field if non-nil, zero value otherwise.

### GetLldpOk

`func (o *InlineResponse20026Ports) GetLldpOk() (*InlineResponse20026Lldp, bool)`

GetLldpOk returns a tuple with the Lldp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLldp

`func (o *InlineResponse20026Ports) SetLldp(v InlineResponse20026Lldp)`

SetLldp sets Lldp field to given value.

### HasLldp

`func (o *InlineResponse20026Ports) HasLldp() bool`

HasLldp returns a boolean if a field has been set.

### GetCdp

`func (o *InlineResponse20026Ports) GetCdp() InlineResponse20026Cdp`

GetCdp returns the Cdp field if non-nil, zero value otherwise.

### GetCdpOk

`func (o *InlineResponse20026Ports) GetCdpOk() (*InlineResponse20026Cdp, bool)`

GetCdpOk returns a tuple with the Cdp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdp

`func (o *InlineResponse20026Ports) SetCdp(v InlineResponse20026Cdp)`

SetCdp sets Cdp field to given value.

### HasCdp

`func (o *InlineResponse20026Ports) HasCdp() bool`

HasCdp returns a boolean if a field has been set.

### GetDeviceMac

`func (o *InlineResponse20026Ports) GetDeviceMac() string`

GetDeviceMac returns the DeviceMac field if non-nil, zero value otherwise.

### GetDeviceMacOk

`func (o *InlineResponse20026Ports) GetDeviceMacOk() (*string, bool)`

GetDeviceMacOk returns a tuple with the DeviceMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceMac

`func (o *InlineResponse20026Ports) SetDeviceMac(v string)`

SetDeviceMac sets DeviceMac field to given value.

### HasDeviceMac

`func (o *InlineResponse20026Ports) HasDeviceMac() bool`

HasDeviceMac returns a boolean if a field has been set.

### GetDevice

`func (o *InlineResponse20026Ports) GetDevice() InlineResponse20026Device`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *InlineResponse20026Ports) GetDeviceOk() (*InlineResponse20026Device, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *InlineResponse20026Ports) SetDevice(v InlineResponse20026Device)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *InlineResponse20026Ports) HasDevice() bool`

HasDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


