# InlineResponse200257Usage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Overall** | Pointer to [**InlineResponse200257UsageOverall**](InlineResponse200257UsageOverall.md) |  | [optional] 
**Average** | Pointer to **float32** | Average data usage (in kb) of each client in organization | [optional] 

## Methods

### NewInlineResponse200257Usage

`func NewInlineResponse200257Usage() *InlineResponse200257Usage`

NewInlineResponse200257Usage instantiates a new InlineResponse200257Usage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200257UsageWithDefaults

`func NewInlineResponse200257UsageWithDefaults() *InlineResponse200257Usage`

NewInlineResponse200257UsageWithDefaults instantiates a new InlineResponse200257Usage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOverall

`func (o *InlineResponse200257Usage) GetOverall() InlineResponse200257UsageOverall`

GetOverall returns the Overall field if non-nil, zero value otherwise.

### GetOverallOk

`func (o *InlineResponse200257Usage) GetOverallOk() (*InlineResponse200257UsageOverall, bool)`

GetOverallOk returns a tuple with the Overall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverall

`func (o *InlineResponse200257Usage) SetOverall(v InlineResponse200257UsageOverall)`

SetOverall sets Overall field to given value.

### HasOverall

`func (o *InlineResponse200257Usage) HasOverall() bool`

HasOverall returns a boolean if a field has been set.

### GetAverage

`func (o *InlineResponse200257Usage) GetAverage() float32`

GetAverage returns the Average field if non-nil, zero value otherwise.

### GetAverageOk

`func (o *InlineResponse200257Usage) GetAverageOk() (*float32, bool)`

GetAverageOk returns a tuple with the Average field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverage

`func (o *InlineResponse200257Usage) SetAverage(v float32)`

SetAverage sets Average field to given value.

### HasAverage

`func (o *InlineResponse200257Usage) HasAverage() bool`

HasAverage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


