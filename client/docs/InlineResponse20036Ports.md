# InlineResponse20036Ports

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Lldp** | Pointer to [**InlineResponse20036Lldp**](InlineResponse20036Lldp.md) |  | [optional] 
**Cdp** | Pointer to [**InlineResponse20036Cdp**](InlineResponse20036Cdp.md) |  | [optional] 
**DeviceMac** | Pointer to **string** | MAC address for the device | [optional] 
**Device** | Pointer to [**InlineResponse20036Device**](InlineResponse20036Device.md) |  | [optional] 

## Methods

### NewInlineResponse20036Ports

`func NewInlineResponse20036Ports() *InlineResponse20036Ports`

NewInlineResponse20036Ports instantiates a new InlineResponse20036Ports object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20036PortsWithDefaults

`func NewInlineResponse20036PortsWithDefaults() *InlineResponse20036Ports`

NewInlineResponse20036PortsWithDefaults instantiates a new InlineResponse20036Ports object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLldp

`func (o *InlineResponse20036Ports) GetLldp() InlineResponse20036Lldp`

GetLldp returns the Lldp field if non-nil, zero value otherwise.

### GetLldpOk

`func (o *InlineResponse20036Ports) GetLldpOk() (*InlineResponse20036Lldp, bool)`

GetLldpOk returns a tuple with the Lldp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLldp

`func (o *InlineResponse20036Ports) SetLldp(v InlineResponse20036Lldp)`

SetLldp sets Lldp field to given value.

### HasLldp

`func (o *InlineResponse20036Ports) HasLldp() bool`

HasLldp returns a boolean if a field has been set.

### GetCdp

`func (o *InlineResponse20036Ports) GetCdp() InlineResponse20036Cdp`

GetCdp returns the Cdp field if non-nil, zero value otherwise.

### GetCdpOk

`func (o *InlineResponse20036Ports) GetCdpOk() (*InlineResponse20036Cdp, bool)`

GetCdpOk returns a tuple with the Cdp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdp

`func (o *InlineResponse20036Ports) SetCdp(v InlineResponse20036Cdp)`

SetCdp sets Cdp field to given value.

### HasCdp

`func (o *InlineResponse20036Ports) HasCdp() bool`

HasCdp returns a boolean if a field has been set.

### GetDeviceMac

`func (o *InlineResponse20036Ports) GetDeviceMac() string`

GetDeviceMac returns the DeviceMac field if non-nil, zero value otherwise.

### GetDeviceMacOk

`func (o *InlineResponse20036Ports) GetDeviceMacOk() (*string, bool)`

GetDeviceMacOk returns a tuple with the DeviceMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceMac

`func (o *InlineResponse20036Ports) SetDeviceMac(v string)`

SetDeviceMac sets DeviceMac field to given value.

### HasDeviceMac

`func (o *InlineResponse20036Ports) HasDeviceMac() bool`

HasDeviceMac returns a boolean if a field has been set.

### GetDevice

`func (o *InlineResponse20036Ports) GetDevice() InlineResponse20036Device`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *InlineResponse20036Ports) GetDeviceOk() (*InlineResponse20036Device, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *InlineResponse20036Ports) SetDevice(v InlineResponse20036Device)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *InlineResponse20036Ports) HasDevice() bool`

HasDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


