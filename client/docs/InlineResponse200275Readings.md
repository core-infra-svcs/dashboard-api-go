# InlineResponse200275Readings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartTs** | Pointer to **time.Time** | Timestamp for the beginning of the bandwidth usage snapshot, exclusive. | [optional] 
**EndTs** | Pointer to **time.Time** | Timestamp for the end of the bandwidth usage snapshot, inclusive. | [optional] 
**SfpProductId** | Pointer to **string** | The product ID for the Small Form Factor Pluggable (SFP) currently inserted into the given port | [optional] 
**ByMetric** | Pointer to [**InlineResponse200275ByMetric**](InlineResponse200275ByMetric.md) |  | [optional] 

## Methods

### NewInlineResponse200275Readings

`func NewInlineResponse200275Readings() *InlineResponse200275Readings`

NewInlineResponse200275Readings instantiates a new InlineResponse200275Readings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200275ReadingsWithDefaults

`func NewInlineResponse200275ReadingsWithDefaults() *InlineResponse200275Readings`

NewInlineResponse200275ReadingsWithDefaults instantiates a new InlineResponse200275Readings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartTs

`func (o *InlineResponse200275Readings) GetStartTs() time.Time`

GetStartTs returns the StartTs field if non-nil, zero value otherwise.

### GetStartTsOk

`func (o *InlineResponse200275Readings) GetStartTsOk() (*time.Time, bool)`

GetStartTsOk returns a tuple with the StartTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTs

`func (o *InlineResponse200275Readings) SetStartTs(v time.Time)`

SetStartTs sets StartTs field to given value.

### HasStartTs

`func (o *InlineResponse200275Readings) HasStartTs() bool`

HasStartTs returns a boolean if a field has been set.

### GetEndTs

`func (o *InlineResponse200275Readings) GetEndTs() time.Time`

GetEndTs returns the EndTs field if non-nil, zero value otherwise.

### GetEndTsOk

`func (o *InlineResponse200275Readings) GetEndTsOk() (*time.Time, bool)`

GetEndTsOk returns a tuple with the EndTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTs

`func (o *InlineResponse200275Readings) SetEndTs(v time.Time)`

SetEndTs sets EndTs field to given value.

### HasEndTs

`func (o *InlineResponse200275Readings) HasEndTs() bool`

HasEndTs returns a boolean if a field has been set.

### GetSfpProductId

`func (o *InlineResponse200275Readings) GetSfpProductId() string`

GetSfpProductId returns the SfpProductId field if non-nil, zero value otherwise.

### GetSfpProductIdOk

`func (o *InlineResponse200275Readings) GetSfpProductIdOk() (*string, bool)`

GetSfpProductIdOk returns a tuple with the SfpProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSfpProductId

`func (o *InlineResponse200275Readings) SetSfpProductId(v string)`

SetSfpProductId sets SfpProductId field to given value.

### HasSfpProductId

`func (o *InlineResponse200275Readings) HasSfpProductId() bool`

HasSfpProductId returns a boolean if a field has been set.

### GetByMetric

`func (o *InlineResponse200275Readings) GetByMetric() InlineResponse200275ByMetric`

GetByMetric returns the ByMetric field if non-nil, zero value otherwise.

### GetByMetricOk

`func (o *InlineResponse200275Readings) GetByMetricOk() (*InlineResponse200275ByMetric, bool)`

GetByMetricOk returns a tuple with the ByMetric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByMetric

`func (o *InlineResponse200275Readings) SetByMetric(v InlineResponse200275ByMetric)`

SetByMetric sets ByMetric field to given value.

### HasByMetric

`func (o *InlineResponse200275Readings) HasByMetric() bool`

HasByMetric returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


