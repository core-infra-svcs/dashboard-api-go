# InlineResponse20048Ports

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Lldp** | Pointer to [**InlineResponse20048Lldp**](InlineResponse20048Lldp.md) |  | [optional] 
**Cdp** | Pointer to [**InlineResponse20048Cdp**](InlineResponse20048Cdp.md) |  | [optional] 
**DeviceMac** | Pointer to **string** | MAC address for the device | [optional] 
**Device** | Pointer to [**InlineResponse20048Device**](InlineResponse20048Device.md) |  | [optional] 

## Methods

### NewInlineResponse20048Ports

`func NewInlineResponse20048Ports() *InlineResponse20048Ports`

NewInlineResponse20048Ports instantiates a new InlineResponse20048Ports object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20048PortsWithDefaults

`func NewInlineResponse20048PortsWithDefaults() *InlineResponse20048Ports`

NewInlineResponse20048PortsWithDefaults instantiates a new InlineResponse20048Ports object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLldp

`func (o *InlineResponse20048Ports) GetLldp() InlineResponse20048Lldp`

GetLldp returns the Lldp field if non-nil, zero value otherwise.

### GetLldpOk

`func (o *InlineResponse20048Ports) GetLldpOk() (*InlineResponse20048Lldp, bool)`

GetLldpOk returns a tuple with the Lldp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLldp

`func (o *InlineResponse20048Ports) SetLldp(v InlineResponse20048Lldp)`

SetLldp sets Lldp field to given value.

### HasLldp

`func (o *InlineResponse20048Ports) HasLldp() bool`

HasLldp returns a boolean if a field has been set.

### GetCdp

`func (o *InlineResponse20048Ports) GetCdp() InlineResponse20048Cdp`

GetCdp returns the Cdp field if non-nil, zero value otherwise.

### GetCdpOk

`func (o *InlineResponse20048Ports) GetCdpOk() (*InlineResponse20048Cdp, bool)`

GetCdpOk returns a tuple with the Cdp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdp

`func (o *InlineResponse20048Ports) SetCdp(v InlineResponse20048Cdp)`

SetCdp sets Cdp field to given value.

### HasCdp

`func (o *InlineResponse20048Ports) HasCdp() bool`

HasCdp returns a boolean if a field has been set.

### GetDeviceMac

`func (o *InlineResponse20048Ports) GetDeviceMac() string`

GetDeviceMac returns the DeviceMac field if non-nil, zero value otherwise.

### GetDeviceMacOk

`func (o *InlineResponse20048Ports) GetDeviceMacOk() (*string, bool)`

GetDeviceMacOk returns a tuple with the DeviceMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceMac

`func (o *InlineResponse20048Ports) SetDeviceMac(v string)`

SetDeviceMac sets DeviceMac field to given value.

### HasDeviceMac

`func (o *InlineResponse20048Ports) HasDeviceMac() bool`

HasDeviceMac returns a boolean if a field has been set.

### GetDevice

`func (o *InlineResponse20048Ports) GetDevice() InlineResponse20048Device`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *InlineResponse20048Ports) GetDeviceOk() (*InlineResponse20048Device, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *InlineResponse20048Ports) SetDevice(v InlineResponse20048Device)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *InlineResponse20048Ports) HasDevice() bool`

HasDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


