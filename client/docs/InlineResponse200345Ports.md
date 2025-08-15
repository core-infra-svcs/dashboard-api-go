# InlineResponse200345Ports

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PortId** | Pointer to **string** | The string identifier of this port on the switch. This is commonly just the port number but may contain additional identifying information such as the slot and module-type if the port is located on a port module. | [optional] 
**Intervals** | Pointer to [**[]InlineResponse200345Intervals**](InlineResponse200345Intervals.md) | An array of intervals for a port with bandwidth, traffic, and power usage data. | [optional] 

## Methods

### NewInlineResponse200345Ports

`func NewInlineResponse200345Ports() *InlineResponse200345Ports`

NewInlineResponse200345Ports instantiates a new InlineResponse200345Ports object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200345PortsWithDefaults

`func NewInlineResponse200345PortsWithDefaults() *InlineResponse200345Ports`

NewInlineResponse200345PortsWithDefaults instantiates a new InlineResponse200345Ports object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPortId

`func (o *InlineResponse200345Ports) GetPortId() string`

GetPortId returns the PortId field if non-nil, zero value otherwise.

### GetPortIdOk

`func (o *InlineResponse200345Ports) GetPortIdOk() (*string, bool)`

GetPortIdOk returns a tuple with the PortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortId

`func (o *InlineResponse200345Ports) SetPortId(v string)`

SetPortId sets PortId field to given value.

### HasPortId

`func (o *InlineResponse200345Ports) HasPortId() bool`

HasPortId returns a boolean if a field has been set.

### GetIntervals

`func (o *InlineResponse200345Ports) GetIntervals() []InlineResponse200345Intervals`

GetIntervals returns the Intervals field if non-nil, zero value otherwise.

### GetIntervalsOk

`func (o *InlineResponse200345Ports) GetIntervalsOk() (*[]InlineResponse200345Intervals, bool)`

GetIntervalsOk returns a tuple with the Intervals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervals

`func (o *InlineResponse200345Ports) SetIntervals(v []InlineResponse200345Intervals)`

SetIntervals sets Intervals field to given value.

### HasIntervals

`func (o *InlineResponse200345Ports) HasIntervals() bool`

HasIntervals returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


