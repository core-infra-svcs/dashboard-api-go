# InlineResponse200372Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Serial** | Pointer to **string** | The cloud ID of the wireless LAN controller | [optional] 
**Intervals** | Pointer to [**[]InlineResponse200372Intervals**](InlineResponse200372Intervals.md) | Time interval snapshots of CPU usage data of the wireless LAN controller | [optional] 

## Methods

### NewInlineResponse200372Items

`func NewInlineResponse200372Items() *InlineResponse200372Items`

NewInlineResponse200372Items instantiates a new InlineResponse200372Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200372ItemsWithDefaults

`func NewInlineResponse200372ItemsWithDefaults() *InlineResponse200372Items`

NewInlineResponse200372ItemsWithDefaults instantiates a new InlineResponse200372Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSerial

`func (o *InlineResponse200372Items) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *InlineResponse200372Items) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *InlineResponse200372Items) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *InlineResponse200372Items) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetIntervals

`func (o *InlineResponse200372Items) GetIntervals() []InlineResponse200372Intervals`

GetIntervals returns the Intervals field if non-nil, zero value otherwise.

### GetIntervalsOk

`func (o *InlineResponse200372Items) GetIntervalsOk() (*[]InlineResponse200372Intervals, bool)`

GetIntervalsOk returns a tuple with the Intervals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervals

`func (o *InlineResponse200372Items) SetIntervals(v []InlineResponse200372Intervals)`

SetIntervals sets Intervals field to given value.

### HasIntervals

`func (o *InlineResponse200372Items) HasIntervals() bool`

HasIntervals returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


