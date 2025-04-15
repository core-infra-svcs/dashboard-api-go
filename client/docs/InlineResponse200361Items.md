# InlineResponse200361Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Serial** | Pointer to **string** | The cloud ID of the wireless LAN controller | [optional] 
**Intervals** | Pointer to [**[]InlineResponse200361Intervals**](InlineResponse200361Intervals.md) | Time interval snapshots of interfaces usage data of the wireless LAN controller | [optional] 

## Methods

### NewInlineResponse200361Items

`func NewInlineResponse200361Items() *InlineResponse200361Items`

NewInlineResponse200361Items instantiates a new InlineResponse200361Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200361ItemsWithDefaults

`func NewInlineResponse200361ItemsWithDefaults() *InlineResponse200361Items`

NewInlineResponse200361ItemsWithDefaults instantiates a new InlineResponse200361Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSerial

`func (o *InlineResponse200361Items) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *InlineResponse200361Items) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *InlineResponse200361Items) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *InlineResponse200361Items) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetIntervals

`func (o *InlineResponse200361Items) GetIntervals() []InlineResponse200361Intervals`

GetIntervals returns the Intervals field if non-nil, zero value otherwise.

### GetIntervalsOk

`func (o *InlineResponse200361Items) GetIntervalsOk() (*[]InlineResponse200361Intervals, bool)`

GetIntervalsOk returns a tuple with the Intervals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervals

`func (o *InlineResponse200361Items) SetIntervals(v []InlineResponse200361Intervals)`

SetIntervals sets Intervals field to given value.

### HasIntervals

`func (o *InlineResponse200361Items) HasIntervals() bool`

HasIntervals returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


