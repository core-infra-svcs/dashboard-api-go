# InlineResponse20074Usage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Overall** | Pointer to [**InlineResponse20074UsageOverall**](InlineResponse20074UsageOverall.md) |  | [optional] 
**Average** | Pointer to **float32** | Average data usage (in kb) of each client in organization | [optional] 

## Methods

### NewInlineResponse20074Usage

`func NewInlineResponse20074Usage() *InlineResponse20074Usage`

NewInlineResponse20074Usage instantiates a new InlineResponse20074Usage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20074UsageWithDefaults

`func NewInlineResponse20074UsageWithDefaults() *InlineResponse20074Usage`

NewInlineResponse20074UsageWithDefaults instantiates a new InlineResponse20074Usage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOverall

`func (o *InlineResponse20074Usage) GetOverall() InlineResponse20074UsageOverall`

GetOverall returns the Overall field if non-nil, zero value otherwise.

### GetOverallOk

`func (o *InlineResponse20074Usage) GetOverallOk() (*InlineResponse20074UsageOverall, bool)`

GetOverallOk returns a tuple with the Overall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverall

`func (o *InlineResponse20074Usage) SetOverall(v InlineResponse20074UsageOverall)`

SetOverall sets Overall field to given value.

### HasOverall

`func (o *InlineResponse20074Usage) HasOverall() bool`

HasOverall returns a boolean if a field has been set.

### GetAverage

`func (o *InlineResponse20074Usage) GetAverage() float32`

GetAverage returns the Average field if non-nil, zero value otherwise.

### GetAverageOk

`func (o *InlineResponse20074Usage) GetAverageOk() (*float32, bool)`

GetAverageOk returns a tuple with the Average field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverage

`func (o *InlineResponse20074Usage) SetAverage(v float32)`

SetAverage sets Average field to given value.

### HasAverage

`func (o *InlineResponse20074Usage) HasAverage() bool`

HasAverage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


