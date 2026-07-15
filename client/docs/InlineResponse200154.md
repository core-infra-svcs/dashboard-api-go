# InlineResponse200154

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SupportedMetrics** | Pointer to **[]string** | List of metrics that are supported for alerts, based on available sensor devices in the network | [optional] 
**Counts** | Pointer to [**InlineResponse200154Counts**](InlineResponse200154Counts.md) |  | [optional] 

## Methods

### NewInlineResponse200154

`func NewInlineResponse200154() *InlineResponse200154`

NewInlineResponse200154 instantiates a new InlineResponse200154 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200154WithDefaults

`func NewInlineResponse200154WithDefaults() *InlineResponse200154`

NewInlineResponse200154WithDefaults instantiates a new InlineResponse200154 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSupportedMetrics

`func (o *InlineResponse200154) GetSupportedMetrics() []string`

GetSupportedMetrics returns the SupportedMetrics field if non-nil, zero value otherwise.

### GetSupportedMetricsOk

`func (o *InlineResponse200154) GetSupportedMetricsOk() (*[]string, bool)`

GetSupportedMetricsOk returns a tuple with the SupportedMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedMetrics

`func (o *InlineResponse200154) SetSupportedMetrics(v []string)`

SetSupportedMetrics sets SupportedMetrics field to given value.

### HasSupportedMetrics

`func (o *InlineResponse200154) HasSupportedMetrics() bool`

HasSupportedMetrics returns a boolean if a field has been set.

### GetCounts

`func (o *InlineResponse200154) GetCounts() InlineResponse200154Counts`

GetCounts returns the Counts field if non-nil, zero value otherwise.

### GetCountsOk

`func (o *InlineResponse200154) GetCountsOk() (*InlineResponse200154Counts, bool)`

GetCountsOk returns a tuple with the Counts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounts

`func (o *InlineResponse200154) SetCounts(v InlineResponse200154Counts)`

SetCounts sets Counts field to given value.

### HasCounts

`func (o *InlineResponse200154) HasCounts() bool`

HasCounts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


