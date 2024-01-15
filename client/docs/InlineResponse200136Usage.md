# InlineResponse200136Usage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Overall** | Pointer to [**InlineResponse200136UsageOverall**](InlineResponse200136UsageOverall.md) |  | [optional] 
**Average** | Pointer to **float32** | Average data usage (in kb) of each client in organization | [optional] 

## Methods

### NewInlineResponse200136Usage

`func NewInlineResponse200136Usage() *InlineResponse200136Usage`

NewInlineResponse200136Usage instantiates a new InlineResponse200136Usage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200136UsageWithDefaults

`func NewInlineResponse200136UsageWithDefaults() *InlineResponse200136Usage`

NewInlineResponse200136UsageWithDefaults instantiates a new InlineResponse200136Usage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOverall

`func (o *InlineResponse200136Usage) GetOverall() InlineResponse200136UsageOverall`

GetOverall returns the Overall field if non-nil, zero value otherwise.

### GetOverallOk

`func (o *InlineResponse200136Usage) GetOverallOk() (*InlineResponse200136UsageOverall, bool)`

GetOverallOk returns a tuple with the Overall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverall

`func (o *InlineResponse200136Usage) SetOverall(v InlineResponse200136UsageOverall)`

SetOverall sets Overall field to given value.

### HasOverall

`func (o *InlineResponse200136Usage) HasOverall() bool`

HasOverall returns a boolean if a field has been set.

### GetAverage

`func (o *InlineResponse200136Usage) GetAverage() float32`

GetAverage returns the Average field if non-nil, zero value otherwise.

### GetAverageOk

`func (o *InlineResponse200136Usage) GetAverageOk() (*float32, bool)`

GetAverageOk returns a tuple with the Average field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverage

`func (o *InlineResponse200136Usage) SetAverage(v float32)`

SetAverage sets Average field to given value.

### HasAverage

`func (o *InlineResponse200136Usage) HasAverage() bool`

HasAverage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


