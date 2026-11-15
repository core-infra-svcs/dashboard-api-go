# InlineResponse20056

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PortId** | Pointer to **string** | The string identifier of this port on the switch. This is commonly just the port number but may contain additional identifying information such as the slot and module-type if the port is located on a port module. | [optional] 
**Packets** | Pointer to [**[]DevicesSerialSwitchPortsStatusesPacketsPackets**](DevicesSerialSwitchPortsStatusesPacketsPackets.md) | The packet counts on the switch. Note that this data is collected periodically from the switch. Any data from the past five minutes may be incomplete. | [optional] 

## Methods

### NewInlineResponse20056

`func NewInlineResponse20056() *InlineResponse20056`

NewInlineResponse20056 instantiates a new InlineResponse20056 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20056WithDefaults

`func NewInlineResponse20056WithDefaults() *InlineResponse20056`

NewInlineResponse20056WithDefaults instantiates a new InlineResponse20056 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPortId

`func (o *InlineResponse20056) GetPortId() string`

GetPortId returns the PortId field if non-nil, zero value otherwise.

### GetPortIdOk

`func (o *InlineResponse20056) GetPortIdOk() (*string, bool)`

GetPortIdOk returns a tuple with the PortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortId

`func (o *InlineResponse20056) SetPortId(v string)`

SetPortId sets PortId field to given value.

### HasPortId

`func (o *InlineResponse20056) HasPortId() bool`

HasPortId returns a boolean if a field has been set.

### GetPackets

`func (o *InlineResponse20056) GetPackets() []DevicesSerialSwitchPortsStatusesPacketsPackets`

GetPackets returns the Packets field if non-nil, zero value otherwise.

### GetPacketsOk

`func (o *InlineResponse20056) GetPacketsOk() (*[]DevicesSerialSwitchPortsStatusesPacketsPackets, bool)`

GetPacketsOk returns a tuple with the Packets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackets

`func (o *InlineResponse20056) SetPackets(v []DevicesSerialSwitchPortsStatusesPacketsPackets)`

SetPackets sets Packets field to given value.

### HasPackets

`func (o *InlineResponse20056) HasPackets() bool`

HasPackets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


