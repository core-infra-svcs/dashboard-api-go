# InlineResponse200353Usage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Overall** | Pointer to [**InlineResponse200353UsageOverall**](InlineResponse200353UsageOverall.md) |  | [optional] 
**Average** | Pointer to **float32** | Average data usage in MiB of each client in organization | [optional] 

## Methods

### NewInlineResponse200353Usage

`func NewInlineResponse200353Usage() *InlineResponse200353Usage`

NewInlineResponse200353Usage instantiates a new InlineResponse200353Usage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200353UsageWithDefaults

`func NewInlineResponse200353UsageWithDefaults() *InlineResponse200353Usage`

NewInlineResponse200353UsageWithDefaults instantiates a new InlineResponse200353Usage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOverall

`func (o *InlineResponse200353Usage) GetOverall() InlineResponse200353UsageOverall`

GetOverall returns the Overall field if non-nil, zero value otherwise.

### GetOverallOk

`func (o *InlineResponse200353Usage) GetOverallOk() (*InlineResponse200353UsageOverall, bool)`

GetOverallOk returns a tuple with the Overall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverall

`func (o *InlineResponse200353Usage) SetOverall(v InlineResponse200353UsageOverall)`

SetOverall sets Overall field to given value.

### HasOverall

`func (o *InlineResponse200353Usage) HasOverall() bool`

HasOverall returns a boolean if a field has been set.

### GetAverage

`func (o *InlineResponse200353Usage) GetAverage() float32`

GetAverage returns the Average field if non-nil, zero value otherwise.

### GetAverageOk

`func (o *InlineResponse200353Usage) GetAverageOk() (*float32, bool)`

GetAverageOk returns a tuple with the Average field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverage

`func (o *InlineResponse200353Usage) SetAverage(v float32)`

SetAverage sets Average field to given value.

### HasAverage

`func (o *InlineResponse200353Usage) HasAverage() bool`

HasAverage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


