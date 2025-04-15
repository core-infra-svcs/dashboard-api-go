# InlineResponse200328Intervals

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartTs** | Pointer to **time.Time** | The starting timestamp of the given interval. | [optional] 
**EndTs** | Pointer to **time.Time** | The end timestamp of the given interval. | [optional] 
**Data** | Pointer to [**InlineResponse200328Data**](InlineResponse200328Data.md) |  | [optional] 
**Bandwidth** | Pointer to [**InlineResponse200328Bandwidth**](InlineResponse200328Bandwidth.md) |  | [optional] 
**Energy** | Pointer to [**InlineResponse200328Energy**](InlineResponse200328Energy.md) |  | [optional] 

## Methods

### NewInlineResponse200328Intervals

`func NewInlineResponse200328Intervals() *InlineResponse200328Intervals`

NewInlineResponse200328Intervals instantiates a new InlineResponse200328Intervals object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200328IntervalsWithDefaults

`func NewInlineResponse200328IntervalsWithDefaults() *InlineResponse200328Intervals`

NewInlineResponse200328IntervalsWithDefaults instantiates a new InlineResponse200328Intervals object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartTs

`func (o *InlineResponse200328Intervals) GetStartTs() time.Time`

GetStartTs returns the StartTs field if non-nil, zero value otherwise.

### GetStartTsOk

`func (o *InlineResponse200328Intervals) GetStartTsOk() (*time.Time, bool)`

GetStartTsOk returns a tuple with the StartTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTs

`func (o *InlineResponse200328Intervals) SetStartTs(v time.Time)`

SetStartTs sets StartTs field to given value.

### HasStartTs

`func (o *InlineResponse200328Intervals) HasStartTs() bool`

HasStartTs returns a boolean if a field has been set.

### GetEndTs

`func (o *InlineResponse200328Intervals) GetEndTs() time.Time`

GetEndTs returns the EndTs field if non-nil, zero value otherwise.

### GetEndTsOk

`func (o *InlineResponse200328Intervals) GetEndTsOk() (*time.Time, bool)`

GetEndTsOk returns a tuple with the EndTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTs

`func (o *InlineResponse200328Intervals) SetEndTs(v time.Time)`

SetEndTs sets EndTs field to given value.

### HasEndTs

`func (o *InlineResponse200328Intervals) HasEndTs() bool`

HasEndTs returns a boolean if a field has been set.

### GetData

`func (o *InlineResponse200328Intervals) GetData() InlineResponse200328Data`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *InlineResponse200328Intervals) GetDataOk() (*InlineResponse200328Data, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *InlineResponse200328Intervals) SetData(v InlineResponse200328Data)`

SetData sets Data field to given value.

### HasData

`func (o *InlineResponse200328Intervals) HasData() bool`

HasData returns a boolean if a field has been set.

### GetBandwidth

`func (o *InlineResponse200328Intervals) GetBandwidth() InlineResponse200328Bandwidth`

GetBandwidth returns the Bandwidth field if non-nil, zero value otherwise.

### GetBandwidthOk

`func (o *InlineResponse200328Intervals) GetBandwidthOk() (*InlineResponse200328Bandwidth, bool)`

GetBandwidthOk returns a tuple with the Bandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidth

`func (o *InlineResponse200328Intervals) SetBandwidth(v InlineResponse200328Bandwidth)`

SetBandwidth sets Bandwidth field to given value.

### HasBandwidth

`func (o *InlineResponse200328Intervals) HasBandwidth() bool`

HasBandwidth returns a boolean if a field has been set.

### GetEnergy

`func (o *InlineResponse200328Intervals) GetEnergy() InlineResponse200328Energy`

GetEnergy returns the Energy field if non-nil, zero value otherwise.

### GetEnergyOk

`func (o *InlineResponse200328Intervals) GetEnergyOk() (*InlineResponse200328Energy, bool)`

GetEnergyOk returns a tuple with the Energy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnergy

`func (o *InlineResponse200328Intervals) SetEnergy(v InlineResponse200328Energy)`

SetEnergy sets Energy field to given value.

### HasEnergy

`func (o *InlineResponse200328Intervals) HasEnergy() bool`

HasEnergy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


