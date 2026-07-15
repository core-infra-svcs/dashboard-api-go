# InlineResponse200473Readings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartTs** | Pointer to **time.Time** | The start time of the query range | [optional] 
**EndTs** | Pointer to **time.Time** | The end time of the query range | [optional] 
**Counts** | Pointer to [**InlineResponse200473Counts**](InlineResponse200473Counts.md) |  | [optional] 

## Methods

### NewInlineResponse200473Readings

`func NewInlineResponse200473Readings() *InlineResponse200473Readings`

NewInlineResponse200473Readings instantiates a new InlineResponse200473Readings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200473ReadingsWithDefaults

`func NewInlineResponse200473ReadingsWithDefaults() *InlineResponse200473Readings`

NewInlineResponse200473ReadingsWithDefaults instantiates a new InlineResponse200473Readings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartTs

`func (o *InlineResponse200473Readings) GetStartTs() time.Time`

GetStartTs returns the StartTs field if non-nil, zero value otherwise.

### GetStartTsOk

`func (o *InlineResponse200473Readings) GetStartTsOk() (*time.Time, bool)`

GetStartTsOk returns a tuple with the StartTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTs

`func (o *InlineResponse200473Readings) SetStartTs(v time.Time)`

SetStartTs sets StartTs field to given value.

### HasStartTs

`func (o *InlineResponse200473Readings) HasStartTs() bool`

HasStartTs returns a boolean if a field has been set.

### GetEndTs

`func (o *InlineResponse200473Readings) GetEndTs() time.Time`

GetEndTs returns the EndTs field if non-nil, zero value otherwise.

### GetEndTsOk

`func (o *InlineResponse200473Readings) GetEndTsOk() (*time.Time, bool)`

GetEndTsOk returns a tuple with the EndTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTs

`func (o *InlineResponse200473Readings) SetEndTs(v time.Time)`

SetEndTs sets EndTs field to given value.

### HasEndTs

`func (o *InlineResponse200473Readings) HasEndTs() bool`

HasEndTs returns a boolean if a field has been set.

### GetCounts

`func (o *InlineResponse200473Readings) GetCounts() InlineResponse200473Counts`

GetCounts returns the Counts field if non-nil, zero value otherwise.

### GetCountsOk

`func (o *InlineResponse200473Readings) GetCountsOk() (*InlineResponse200473Counts, bool)`

GetCountsOk returns a tuple with the Counts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounts

`func (o *InlineResponse200473Readings) SetCounts(v InlineResponse200473Counts)`

SetCounts sets Counts field to given value.

### HasCounts

`func (o *InlineResponse200473Readings) HasCounts() bool`

HasCounts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


