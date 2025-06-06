# InlineResponse20038

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PortId** | Pointer to **string** | The string identifier of this port on the switch. This is commonly just the port number but may contain additional identifying information such as the slot and module-type if the port is located on a port module. | [optional] 
**Packets** | Pointer to [**[]DevicesSerialSwitchPortsStatusesPacketsPackets**](DevicesSerialSwitchPortsStatusesPacketsPackets.md) | The packet counts on the switch. | [optional] 

## Methods

### NewInlineResponse20038

`func NewInlineResponse20038() *InlineResponse20038`

NewInlineResponse20038 instantiates a new InlineResponse20038 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20038WithDefaults

`func NewInlineResponse20038WithDefaults() *InlineResponse20038`

NewInlineResponse20038WithDefaults instantiates a new InlineResponse20038 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPortId

`func (o *InlineResponse20038) GetPortId() string`

GetPortId returns the PortId field if non-nil, zero value otherwise.

### GetPortIdOk

`func (o *InlineResponse20038) GetPortIdOk() (*string, bool)`

GetPortIdOk returns a tuple with the PortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortId

`func (o *InlineResponse20038) SetPortId(v string)`

SetPortId sets PortId field to given value.

### HasPortId

`func (o *InlineResponse20038) HasPortId() bool`

HasPortId returns a boolean if a field has been set.

### GetPackets

`func (o *InlineResponse20038) GetPackets() []DevicesSerialSwitchPortsStatusesPacketsPackets`

GetPackets returns the Packets field if non-nil, zero value otherwise.

### GetPacketsOk

`func (o *InlineResponse20038) GetPacketsOk() (*[]DevicesSerialSwitchPortsStatusesPacketsPackets, bool)`

GetPacketsOk returns a tuple with the Packets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackets

`func (o *InlineResponse20038) SetPackets(v []DevicesSerialSwitchPortsStatusesPacketsPackets)`

SetPackets sets Packets field to given value.

### HasPackets

`func (o *InlineResponse20038) HasPackets() bool`

HasPackets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


