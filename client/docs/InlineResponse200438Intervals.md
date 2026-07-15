# InlineResponse200438Intervals

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartTs** | Pointer to **time.Time** | The starting timestamp of the given interval. | [optional] 
**EndTs** | Pointer to **time.Time** | The end timestamp of the given interval. | [optional] 
**Data** | Pointer to [**InlineResponse200438Data**](InlineResponse200438Data.md) |  | [optional] 
**Bandwidth** | Pointer to [**InlineResponse200438Bandwidth**](InlineResponse200438Bandwidth.md) |  | [optional] 
**Energy** | Pointer to [**InlineResponse200438Energy**](InlineResponse200438Energy.md) |  | [optional] 

## Methods

### NewInlineResponse200438Intervals

`func NewInlineResponse200438Intervals() *InlineResponse200438Intervals`

NewInlineResponse200438Intervals instantiates a new InlineResponse200438Intervals object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200438IntervalsWithDefaults

`func NewInlineResponse200438IntervalsWithDefaults() *InlineResponse200438Intervals`

NewInlineResponse200438IntervalsWithDefaults instantiates a new InlineResponse200438Intervals object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartTs

`func (o *InlineResponse200438Intervals) GetStartTs() time.Time`

GetStartTs returns the StartTs field if non-nil, zero value otherwise.

### GetStartTsOk

`func (o *InlineResponse200438Intervals) GetStartTsOk() (*time.Time, bool)`

GetStartTsOk returns a tuple with the StartTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTs

`func (o *InlineResponse200438Intervals) SetStartTs(v time.Time)`

SetStartTs sets StartTs field to given value.

### HasStartTs

`func (o *InlineResponse200438Intervals) HasStartTs() bool`

HasStartTs returns a boolean if a field has been set.

### GetEndTs

`func (o *InlineResponse200438Intervals) GetEndTs() time.Time`

GetEndTs returns the EndTs field if non-nil, zero value otherwise.

### GetEndTsOk

`func (o *InlineResponse200438Intervals) GetEndTsOk() (*time.Time, bool)`

GetEndTsOk returns a tuple with the EndTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTs

`func (o *InlineResponse200438Intervals) SetEndTs(v time.Time)`

SetEndTs sets EndTs field to given value.

### HasEndTs

`func (o *InlineResponse200438Intervals) HasEndTs() bool`

HasEndTs returns a boolean if a field has been set.

### GetData

`func (o *InlineResponse200438Intervals) GetData() InlineResponse200438Data`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *InlineResponse200438Intervals) GetDataOk() (*InlineResponse200438Data, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *InlineResponse200438Intervals) SetData(v InlineResponse200438Data)`

SetData sets Data field to given value.

### HasData

`func (o *InlineResponse200438Intervals) HasData() bool`

HasData returns a boolean if a field has been set.

### GetBandwidth

`func (o *InlineResponse200438Intervals) GetBandwidth() InlineResponse200438Bandwidth`

GetBandwidth returns the Bandwidth field if non-nil, zero value otherwise.

### GetBandwidthOk

`func (o *InlineResponse200438Intervals) GetBandwidthOk() (*InlineResponse200438Bandwidth, bool)`

GetBandwidthOk returns a tuple with the Bandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidth

`func (o *InlineResponse200438Intervals) SetBandwidth(v InlineResponse200438Bandwidth)`

SetBandwidth sets Bandwidth field to given value.

### HasBandwidth

`func (o *InlineResponse200438Intervals) HasBandwidth() bool`

HasBandwidth returns a boolean if a field has been set.

### GetEnergy

`func (o *InlineResponse200438Intervals) GetEnergy() InlineResponse200438Energy`

GetEnergy returns the Energy field if non-nil, zero value otherwise.

### GetEnergyOk

`func (o *InlineResponse200438Intervals) GetEnergyOk() (*InlineResponse200438Energy, bool)`

GetEnergyOk returns a tuple with the Energy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnergy

`func (o *InlineResponse200438Intervals) SetEnergy(v InlineResponse200438Energy)`

SetEnergy sets Energy field to given value.

### HasEnergy

`func (o *InlineResponse200438Intervals) HasEnergy() bool`

HasEnergy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


