# InlineResponse20025

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SupportedMetrics** | Pointer to **[]string** | List of metrics that are supported for alerts, based on available sensor devices in the network | [optional] 
**Counts** | Pointer to [**InlineResponse20025Counts**](InlineResponse20025Counts.md) |  | [optional] 

## Methods

### NewInlineResponse20025

`func NewInlineResponse20025() *InlineResponse20025`

NewInlineResponse20025 instantiates a new InlineResponse20025 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20025WithDefaults

`func NewInlineResponse20025WithDefaults() *InlineResponse20025`

NewInlineResponse20025WithDefaults instantiates a new InlineResponse20025 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSupportedMetrics

`func (o *InlineResponse20025) GetSupportedMetrics() []string`

GetSupportedMetrics returns the SupportedMetrics field if non-nil, zero value otherwise.

### GetSupportedMetricsOk

`func (o *InlineResponse20025) GetSupportedMetricsOk() (*[]string, bool)`

GetSupportedMetricsOk returns a tuple with the SupportedMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedMetrics

`func (o *InlineResponse20025) SetSupportedMetrics(v []string)`

SetSupportedMetrics sets SupportedMetrics field to given value.

### HasSupportedMetrics

`func (o *InlineResponse20025) HasSupportedMetrics() bool`

HasSupportedMetrics returns a boolean if a field has been set.

### GetCounts

`func (o *InlineResponse20025) GetCounts() InlineResponse20025Counts`

GetCounts returns the Counts field if non-nil, zero value otherwise.

### GetCountsOk

`func (o *InlineResponse20025) GetCountsOk() (*InlineResponse20025Counts, bool)`

GetCountsOk returns a tuple with the Counts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounts

`func (o *InlineResponse20025) SetCounts(v InlineResponse20025Counts)`

SetCounts sets Counts field to given value.

### HasCounts

`func (o *InlineResponse20025) HasCounts() bool`

HasCounts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


