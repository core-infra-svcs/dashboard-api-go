# InlineResponse200273Readings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartTs** | Pointer to **time.Time** | Timestamp for the beginning of the bandwidth usage snapshot, exclusive. | [optional] 
**EndTs** | Pointer to **time.Time** | Timestamp for the end of the bandwidth usage snapshot, inclusive. | [optional] 
**SfpProductId** | Pointer to **string** | The product ID for the Small Form Factor Pluggable (SFP) currently inserted into the given port | [optional] 
**ByMetric** | Pointer to [**InlineResponse200273ByMetric**](InlineResponse200273ByMetric.md) |  | [optional] 

## Methods

### NewInlineResponse200273Readings

`func NewInlineResponse200273Readings() *InlineResponse200273Readings`

NewInlineResponse200273Readings instantiates a new InlineResponse200273Readings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200273ReadingsWithDefaults

`func NewInlineResponse200273ReadingsWithDefaults() *InlineResponse200273Readings`

NewInlineResponse200273ReadingsWithDefaults instantiates a new InlineResponse200273Readings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartTs

`func (o *InlineResponse200273Readings) GetStartTs() time.Time`

GetStartTs returns the StartTs field if non-nil, zero value otherwise.

### GetStartTsOk

`func (o *InlineResponse200273Readings) GetStartTsOk() (*time.Time, bool)`

GetStartTsOk returns a tuple with the StartTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTs

`func (o *InlineResponse200273Readings) SetStartTs(v time.Time)`

SetStartTs sets StartTs field to given value.

### HasStartTs

`func (o *InlineResponse200273Readings) HasStartTs() bool`

HasStartTs returns a boolean if a field has been set.

### GetEndTs

`func (o *InlineResponse200273Readings) GetEndTs() time.Time`

GetEndTs returns the EndTs field if non-nil, zero value otherwise.

### GetEndTsOk

`func (o *InlineResponse200273Readings) GetEndTsOk() (*time.Time, bool)`

GetEndTsOk returns a tuple with the EndTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTs

`func (o *InlineResponse200273Readings) SetEndTs(v time.Time)`

SetEndTs sets EndTs field to given value.

### HasEndTs

`func (o *InlineResponse200273Readings) HasEndTs() bool`

HasEndTs returns a boolean if a field has been set.

### GetSfpProductId

`func (o *InlineResponse200273Readings) GetSfpProductId() string`

GetSfpProductId returns the SfpProductId field if non-nil, zero value otherwise.

### GetSfpProductIdOk

`func (o *InlineResponse200273Readings) GetSfpProductIdOk() (*string, bool)`

GetSfpProductIdOk returns a tuple with the SfpProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSfpProductId

`func (o *InlineResponse200273Readings) SetSfpProductId(v string)`

SetSfpProductId sets SfpProductId field to given value.

### HasSfpProductId

`func (o *InlineResponse200273Readings) HasSfpProductId() bool`

HasSfpProductId returns a boolean if a field has been set.

### GetByMetric

`func (o *InlineResponse200273Readings) GetByMetric() InlineResponse200273ByMetric`

GetByMetric returns the ByMetric field if non-nil, zero value otherwise.

### GetByMetricOk

`func (o *InlineResponse200273Readings) GetByMetricOk() (*InlineResponse200273ByMetric, bool)`

GetByMetricOk returns a tuple with the ByMetric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByMetric

`func (o *InlineResponse200273Readings) SetByMetric(v InlineResponse200273ByMetric)`

SetByMetric sets ByMetric field to given value.

### HasByMetric

`func (o *InlineResponse200273Readings) HasByMetric() bool`

HasByMetric returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


