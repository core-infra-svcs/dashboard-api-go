# InlineResponse20033

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PortId** | Pointer to **string** | The string identifier of this port on the switch. This is commonly just the port number but may contain additional identifying information such as the slot and module-type if the port is located on a port module. | [optional] 
**Packets** | Pointer to [**[]DevicesSerialSwitchPortsStatusesPacketsPackets**](DevicesSerialSwitchPortsStatusesPacketsPackets.md) | The packet counts on the switch. | [optional] 

## Methods

### NewInlineResponse20033

`func NewInlineResponse20033() *InlineResponse20033`

NewInlineResponse20033 instantiates a new InlineResponse20033 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20033WithDefaults

`func NewInlineResponse20033WithDefaults() *InlineResponse20033`

NewInlineResponse20033WithDefaults instantiates a new InlineResponse20033 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPortId

`func (o *InlineResponse20033) GetPortId() string`

GetPortId returns the PortId field if non-nil, zero value otherwise.

### GetPortIdOk

`func (o *InlineResponse20033) GetPortIdOk() (*string, bool)`

GetPortIdOk returns a tuple with the PortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortId

`func (o *InlineResponse20033) SetPortId(v string)`

SetPortId sets PortId field to given value.

### HasPortId

`func (o *InlineResponse20033) HasPortId() bool`

HasPortId returns a boolean if a field has been set.

### GetPackets

`func (o *InlineResponse20033) GetPackets() []DevicesSerialSwitchPortsStatusesPacketsPackets`

GetPackets returns the Packets field if non-nil, zero value otherwise.

### GetPacketsOk

`func (o *InlineResponse20033) GetPacketsOk() (*[]DevicesSerialSwitchPortsStatusesPacketsPackets, bool)`

GetPacketsOk returns a tuple with the Packets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackets

`func (o *InlineResponse20033) SetPackets(v []DevicesSerialSwitchPortsStatusesPacketsPackets)`

SetPackets sets Packets field to given value.

### HasPackets

`func (o *InlineResponse20033) HasPackets() bool`

HasPackets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


