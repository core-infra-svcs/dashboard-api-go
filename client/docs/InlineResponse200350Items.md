# InlineResponse200350Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Serial** | Pointer to **string** | Serial of device | [optional] 
**Intervals** | Pointer to [**[]InlineResponse200350Intervals**](InlineResponse200350Intervals.md) | Historical cellular data usage buckets for the device. Maximum 8928 intervals per device based on the 31-day timespan and 5-minute interval limits. | [optional] 

## Methods

### NewInlineResponse200350Items

`func NewInlineResponse200350Items() *InlineResponse200350Items`

NewInlineResponse200350Items instantiates a new InlineResponse200350Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200350ItemsWithDefaults

`func NewInlineResponse200350ItemsWithDefaults() *InlineResponse200350Items`

NewInlineResponse200350ItemsWithDefaults instantiates a new InlineResponse200350Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSerial

`func (o *InlineResponse200350Items) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *InlineResponse200350Items) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *InlineResponse200350Items) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *InlineResponse200350Items) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetIntervals

`func (o *InlineResponse200350Items) GetIntervals() []InlineResponse200350Intervals`

GetIntervals returns the Intervals field if non-nil, zero value otherwise.

### GetIntervalsOk

`func (o *InlineResponse200350Items) GetIntervalsOk() (*[]InlineResponse200350Intervals, bool)`

GetIntervalsOk returns a tuple with the Intervals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervals

`func (o *InlineResponse200350Items) SetIntervals(v []InlineResponse200350Intervals)`

SetIntervals sets Intervals field to given value.

### HasIntervals

`func (o *InlineResponse200350Items) HasIntervals() bool`

HasIntervals returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


