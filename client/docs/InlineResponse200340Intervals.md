# InlineResponse200340Intervals

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartTs** | Pointer to **time.Time** | The starting timestamp of the given interval. | [optional] 
**EndTs** | Pointer to **time.Time** | The end timestamp of the given interval. | [optional] 
**Data** | Pointer to [**InlineResponse200340Data**](InlineResponse200340Data.md) |  | [optional] 
**Bandwidth** | Pointer to [**InlineResponse200340Bandwidth**](InlineResponse200340Bandwidth.md) |  | [optional] 
**Energy** | Pointer to [**InlineResponse200340Energy**](InlineResponse200340Energy.md) |  | [optional] 

## Methods

### NewInlineResponse200340Intervals

`func NewInlineResponse200340Intervals() *InlineResponse200340Intervals`

NewInlineResponse200340Intervals instantiates a new InlineResponse200340Intervals object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200340IntervalsWithDefaults

`func NewInlineResponse200340IntervalsWithDefaults() *InlineResponse200340Intervals`

NewInlineResponse200340IntervalsWithDefaults instantiates a new InlineResponse200340Intervals object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartTs

`func (o *InlineResponse200340Intervals) GetStartTs() time.Time`

GetStartTs returns the StartTs field if non-nil, zero value otherwise.

### GetStartTsOk

`func (o *InlineResponse200340Intervals) GetStartTsOk() (*time.Time, bool)`

GetStartTsOk returns a tuple with the StartTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTs

`func (o *InlineResponse200340Intervals) SetStartTs(v time.Time)`

SetStartTs sets StartTs field to given value.

### HasStartTs

`func (o *InlineResponse200340Intervals) HasStartTs() bool`

HasStartTs returns a boolean if a field has been set.

### GetEndTs

`func (o *InlineResponse200340Intervals) GetEndTs() time.Time`

GetEndTs returns the EndTs field if non-nil, zero value otherwise.

### GetEndTsOk

`func (o *InlineResponse200340Intervals) GetEndTsOk() (*time.Time, bool)`

GetEndTsOk returns a tuple with the EndTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTs

`func (o *InlineResponse200340Intervals) SetEndTs(v time.Time)`

SetEndTs sets EndTs field to given value.

### HasEndTs

`func (o *InlineResponse200340Intervals) HasEndTs() bool`

HasEndTs returns a boolean if a field has been set.

### GetData

`func (o *InlineResponse200340Intervals) GetData() InlineResponse200340Data`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *InlineResponse200340Intervals) GetDataOk() (*InlineResponse200340Data, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *InlineResponse200340Intervals) SetData(v InlineResponse200340Data)`

SetData sets Data field to given value.

### HasData

`func (o *InlineResponse200340Intervals) HasData() bool`

HasData returns a boolean if a field has been set.

### GetBandwidth

`func (o *InlineResponse200340Intervals) GetBandwidth() InlineResponse200340Bandwidth`

GetBandwidth returns the Bandwidth field if non-nil, zero value otherwise.

### GetBandwidthOk

`func (o *InlineResponse200340Intervals) GetBandwidthOk() (*InlineResponse200340Bandwidth, bool)`

GetBandwidthOk returns a tuple with the Bandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidth

`func (o *InlineResponse200340Intervals) SetBandwidth(v InlineResponse200340Bandwidth)`

SetBandwidth sets Bandwidth field to given value.

### HasBandwidth

`func (o *InlineResponse200340Intervals) HasBandwidth() bool`

HasBandwidth returns a boolean if a field has been set.

### GetEnergy

`func (o *InlineResponse200340Intervals) GetEnergy() InlineResponse200340Energy`

GetEnergy returns the Energy field if non-nil, zero value otherwise.

### GetEnergyOk

`func (o *InlineResponse200340Intervals) GetEnergyOk() (*InlineResponse200340Energy, bool)`

GetEnergyOk returns a tuple with the Energy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnergy

`func (o *InlineResponse200340Intervals) SetEnergy(v InlineResponse200340Energy)`

SetEnergy sets Energy field to given value.

### HasEnergy

`func (o *InlineResponse200340Intervals) HasEnergy() bool`

HasEnergy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


