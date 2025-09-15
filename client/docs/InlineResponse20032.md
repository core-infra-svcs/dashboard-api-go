# InlineResponse20032

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartTime** | Pointer to **time.Time** | Start time of the sample | [optional] 
**EndTime** | Pointer to **time.Time** | End time of the sample | [optional] 
**LossPercent** | Pointer to **float32** | Percentage of packets lost | [optional] 
**LatencyMs** | Pointer to **float32** | Latency in milliseconds | [optional] 
**Goodput** | Pointer to **int32** | Number of useful information bits delivered | [optional] 
**Jitter** | Pointer to **float32** | Jitter, in milliseconds | [optional] 

## Methods

### NewInlineResponse20032

`func NewInlineResponse20032() *InlineResponse20032`

NewInlineResponse20032 instantiates a new InlineResponse20032 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20032WithDefaults

`func NewInlineResponse20032WithDefaults() *InlineResponse20032`

NewInlineResponse20032WithDefaults instantiates a new InlineResponse20032 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartTime

`func (o *InlineResponse20032) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *InlineResponse20032) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *InlineResponse20032) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *InlineResponse20032) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetEndTime

`func (o *InlineResponse20032) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *InlineResponse20032) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *InlineResponse20032) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *InlineResponse20032) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetLossPercent

`func (o *InlineResponse20032) GetLossPercent() float32`

GetLossPercent returns the LossPercent field if non-nil, zero value otherwise.

### GetLossPercentOk

`func (o *InlineResponse20032) GetLossPercentOk() (*float32, bool)`

GetLossPercentOk returns a tuple with the LossPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLossPercent

`func (o *InlineResponse20032) SetLossPercent(v float32)`

SetLossPercent sets LossPercent field to given value.

### HasLossPercent

`func (o *InlineResponse20032) HasLossPercent() bool`

HasLossPercent returns a boolean if a field has been set.

### GetLatencyMs

`func (o *InlineResponse20032) GetLatencyMs() float32`

GetLatencyMs returns the LatencyMs field if non-nil, zero value otherwise.

### GetLatencyMsOk

`func (o *InlineResponse20032) GetLatencyMsOk() (*float32, bool)`

GetLatencyMsOk returns a tuple with the LatencyMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatencyMs

`func (o *InlineResponse20032) SetLatencyMs(v float32)`

SetLatencyMs sets LatencyMs field to given value.

### HasLatencyMs

`func (o *InlineResponse20032) HasLatencyMs() bool`

HasLatencyMs returns a boolean if a field has been set.

### GetGoodput

`func (o *InlineResponse20032) GetGoodput() int32`

GetGoodput returns the Goodput field if non-nil, zero value otherwise.

### GetGoodputOk

`func (o *InlineResponse20032) GetGoodputOk() (*int32, bool)`

GetGoodputOk returns a tuple with the Goodput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoodput

`func (o *InlineResponse20032) SetGoodput(v int32)`

SetGoodput sets Goodput field to given value.

### HasGoodput

`func (o *InlineResponse20032) HasGoodput() bool`

HasGoodput returns a boolean if a field has been set.

### GetJitter

`func (o *InlineResponse20032) GetJitter() float32`

GetJitter returns the Jitter field if non-nil, zero value otherwise.

### GetJitterOk

`func (o *InlineResponse20032) GetJitterOk() (*float32, bool)`

GetJitterOk returns a tuple with the Jitter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJitter

`func (o *InlineResponse20032) SetJitter(v float32)`

SetJitter sets Jitter field to given value.

### HasJitter

`func (o *InlineResponse20032) HasJitter() bool`

HasJitter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


