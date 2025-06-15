# InlineResponse200340Ports

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PortId** | Pointer to **string** | The string identifier of this port on the switch. This is commonly just the port number but may contain additional identifying information such as the slot and module-type if the port is located on a port module. | [optional] 
**Intervals** | Pointer to [**[]InlineResponse200340Intervals**](InlineResponse200340Intervals.md) | An array of intervals for a port with bandwidth, traffic, and power usage data. | [optional] 

## Methods

### NewInlineResponse200340Ports

`func NewInlineResponse200340Ports() *InlineResponse200340Ports`

NewInlineResponse200340Ports instantiates a new InlineResponse200340Ports object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200340PortsWithDefaults

`func NewInlineResponse200340PortsWithDefaults() *InlineResponse200340Ports`

NewInlineResponse200340PortsWithDefaults instantiates a new InlineResponse200340Ports object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPortId

`func (o *InlineResponse200340Ports) GetPortId() string`

GetPortId returns the PortId field if non-nil, zero value otherwise.

### GetPortIdOk

`func (o *InlineResponse200340Ports) GetPortIdOk() (*string, bool)`

GetPortIdOk returns a tuple with the PortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortId

`func (o *InlineResponse200340Ports) SetPortId(v string)`

SetPortId sets PortId field to given value.

### HasPortId

`func (o *InlineResponse200340Ports) HasPortId() bool`

HasPortId returns a boolean if a field has been set.

### GetIntervals

`func (o *InlineResponse200340Ports) GetIntervals() []InlineResponse200340Intervals`

GetIntervals returns the Intervals field if non-nil, zero value otherwise.

### GetIntervalsOk

`func (o *InlineResponse200340Ports) GetIntervalsOk() (*[]InlineResponse200340Intervals, bool)`

GetIntervalsOk returns a tuple with the Intervals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervals

`func (o *InlineResponse200340Ports) SetIntervals(v []InlineResponse200340Intervals)`

SetIntervals sets Intervals field to given value.

### HasIntervals

`func (o *InlineResponse200340Ports) HasIntervals() bool`

HasIntervals returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


