# InlineResponse20030Ports

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Lldp** | Pointer to [**InlineResponse20030Lldp**](InlineResponse20030Lldp.md) |  | [optional] 
**Cdp** | Pointer to [**InlineResponse20030Cdp**](InlineResponse20030Cdp.md) |  | [optional] 
**DeviceMac** | Pointer to **string** | MAC address for the device | [optional] 
**Device** | Pointer to [**InlineResponse20030Device**](InlineResponse20030Device.md) |  | [optional] 

## Methods

### NewInlineResponse20030Ports

`func NewInlineResponse20030Ports() *InlineResponse20030Ports`

NewInlineResponse20030Ports instantiates a new InlineResponse20030Ports object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20030PortsWithDefaults

`func NewInlineResponse20030PortsWithDefaults() *InlineResponse20030Ports`

NewInlineResponse20030PortsWithDefaults instantiates a new InlineResponse20030Ports object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLldp

`func (o *InlineResponse20030Ports) GetLldp() InlineResponse20030Lldp`

GetLldp returns the Lldp field if non-nil, zero value otherwise.

### GetLldpOk

`func (o *InlineResponse20030Ports) GetLldpOk() (*InlineResponse20030Lldp, bool)`

GetLldpOk returns a tuple with the Lldp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLldp

`func (o *InlineResponse20030Ports) SetLldp(v InlineResponse20030Lldp)`

SetLldp sets Lldp field to given value.

### HasLldp

`func (o *InlineResponse20030Ports) HasLldp() bool`

HasLldp returns a boolean if a field has been set.

### GetCdp

`func (o *InlineResponse20030Ports) GetCdp() InlineResponse20030Cdp`

GetCdp returns the Cdp field if non-nil, zero value otherwise.

### GetCdpOk

`func (o *InlineResponse20030Ports) GetCdpOk() (*InlineResponse20030Cdp, bool)`

GetCdpOk returns a tuple with the Cdp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdp

`func (o *InlineResponse20030Ports) SetCdp(v InlineResponse20030Cdp)`

SetCdp sets Cdp field to given value.

### HasCdp

`func (o *InlineResponse20030Ports) HasCdp() bool`

HasCdp returns a boolean if a field has been set.

### GetDeviceMac

`func (o *InlineResponse20030Ports) GetDeviceMac() string`

GetDeviceMac returns the DeviceMac field if non-nil, zero value otherwise.

### GetDeviceMacOk

`func (o *InlineResponse20030Ports) GetDeviceMacOk() (*string, bool)`

GetDeviceMacOk returns a tuple with the DeviceMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceMac

`func (o *InlineResponse20030Ports) SetDeviceMac(v string)`

SetDeviceMac sets DeviceMac field to given value.

### HasDeviceMac

`func (o *InlineResponse20030Ports) HasDeviceMac() bool`

HasDeviceMac returns a boolean if a field has been set.

### GetDevice

`func (o *InlineResponse20030Ports) GetDevice() InlineResponse20030Device`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *InlineResponse20030Ports) GetDeviceOk() (*InlineResponse20030Device, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *InlineResponse20030Ports) SetDevice(v InlineResponse20030Device)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *InlineResponse20030Ports) HasDevice() bool`

HasDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


