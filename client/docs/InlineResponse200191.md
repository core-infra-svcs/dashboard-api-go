# InlineResponse200191

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartTs** | Pointer to **time.Time** | The start time of the query range | [optional] 
**EndTs** | Pointer to **time.Time** | The end time of the query range | [optional] 
**AvgLatencyMs** | Pointer to **int32** | Average latency in milliseconds | [optional] 

## Methods

### NewInlineResponse200191

`func NewInlineResponse200191() *InlineResponse200191`

NewInlineResponse200191 instantiates a new InlineResponse200191 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200191WithDefaults

`func NewInlineResponse200191WithDefaults() *InlineResponse200191`

NewInlineResponse200191WithDefaults instantiates a new InlineResponse200191 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartTs

`func (o *InlineResponse200191) GetStartTs() time.Time`

GetStartTs returns the StartTs field if non-nil, zero value otherwise.

### GetStartTsOk

`func (o *InlineResponse200191) GetStartTsOk() (*time.Time, bool)`

GetStartTsOk returns a tuple with the StartTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTs

`func (o *InlineResponse200191) SetStartTs(v time.Time)`

SetStartTs sets StartTs field to given value.

### HasStartTs

`func (o *InlineResponse200191) HasStartTs() bool`

HasStartTs returns a boolean if a field has been set.

### GetEndTs

`func (o *InlineResponse200191) GetEndTs() time.Time`

GetEndTs returns the EndTs field if non-nil, zero value otherwise.

### GetEndTsOk

`func (o *InlineResponse200191) GetEndTsOk() (*time.Time, bool)`

GetEndTsOk returns a tuple with the EndTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTs

`func (o *InlineResponse200191) SetEndTs(v time.Time)`

SetEndTs sets EndTs field to given value.

### HasEndTs

`func (o *InlineResponse200191) HasEndTs() bool`

HasEndTs returns a boolean if a field has been set.

### GetAvgLatencyMs

`func (o *InlineResponse200191) GetAvgLatencyMs() int32`

GetAvgLatencyMs returns the AvgLatencyMs field if non-nil, zero value otherwise.

### GetAvgLatencyMsOk

`func (o *InlineResponse200191) GetAvgLatencyMsOk() (*int32, bool)`

GetAvgLatencyMsOk returns a tuple with the AvgLatencyMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgLatencyMs

`func (o *InlineResponse200191) SetAvgLatencyMs(v int32)`

SetAvgLatencyMs sets AvgLatencyMs field to given value.

### HasAvgLatencyMs

`func (o *InlineResponse200191) HasAvgLatencyMs() bool`

HasAvgLatencyMs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


