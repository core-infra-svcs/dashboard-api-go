# InlineResponse20031Ports

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Lldp** | Pointer to [**InlineResponse20031Lldp**](InlineResponse20031Lldp.md) |  | [optional] 
**Cdp** | Pointer to [**InlineResponse20031Cdp**](InlineResponse20031Cdp.md) |  | [optional] 
**DeviceMac** | Pointer to **string** | MAC address for the device | [optional] 
**Device** | Pointer to [**InlineResponse20031Device**](InlineResponse20031Device.md) |  | [optional] 

## Methods

### NewInlineResponse20031Ports

`func NewInlineResponse20031Ports() *InlineResponse20031Ports`

NewInlineResponse20031Ports instantiates a new InlineResponse20031Ports object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20031PortsWithDefaults

`func NewInlineResponse20031PortsWithDefaults() *InlineResponse20031Ports`

NewInlineResponse20031PortsWithDefaults instantiates a new InlineResponse20031Ports object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLldp

`func (o *InlineResponse20031Ports) GetLldp() InlineResponse20031Lldp`

GetLldp returns the Lldp field if non-nil, zero value otherwise.

### GetLldpOk

`func (o *InlineResponse20031Ports) GetLldpOk() (*InlineResponse20031Lldp, bool)`

GetLldpOk returns a tuple with the Lldp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLldp

`func (o *InlineResponse20031Ports) SetLldp(v InlineResponse20031Lldp)`

SetLldp sets Lldp field to given value.

### HasLldp

`func (o *InlineResponse20031Ports) HasLldp() bool`

HasLldp returns a boolean if a field has been set.

### GetCdp

`func (o *InlineResponse20031Ports) GetCdp() InlineResponse20031Cdp`

GetCdp returns the Cdp field if non-nil, zero value otherwise.

### GetCdpOk

`func (o *InlineResponse20031Ports) GetCdpOk() (*InlineResponse20031Cdp, bool)`

GetCdpOk returns a tuple with the Cdp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdp

`func (o *InlineResponse20031Ports) SetCdp(v InlineResponse20031Cdp)`

SetCdp sets Cdp field to given value.

### HasCdp

`func (o *InlineResponse20031Ports) HasCdp() bool`

HasCdp returns a boolean if a field has been set.

### GetDeviceMac

`func (o *InlineResponse20031Ports) GetDeviceMac() string`

GetDeviceMac returns the DeviceMac field if non-nil, zero value otherwise.

### GetDeviceMacOk

`func (o *InlineResponse20031Ports) GetDeviceMacOk() (*string, bool)`

GetDeviceMacOk returns a tuple with the DeviceMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceMac

`func (o *InlineResponse20031Ports) SetDeviceMac(v string)`

SetDeviceMac sets DeviceMac field to given value.

### HasDeviceMac

`func (o *InlineResponse20031Ports) HasDeviceMac() bool`

HasDeviceMac returns a boolean if a field has been set.

### GetDevice

`func (o *InlineResponse20031Ports) GetDevice() InlineResponse20031Device`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *InlineResponse20031Ports) GetDeviceOk() (*InlineResponse20031Device, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *InlineResponse20031Ports) SetDevice(v InlineResponse20031Device)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *InlineResponse20031Ports) HasDevice() bool`

HasDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


