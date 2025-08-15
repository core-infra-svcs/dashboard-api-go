# InlineResponse200117

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SupportedMetrics** | Pointer to **[]string** | List of metrics that are supported for alerts, based on available sensor devices in the network | [optional] 
**Counts** | Pointer to [**InlineResponse200117Counts**](InlineResponse200117Counts.md) |  | [optional] 

## Methods

### NewInlineResponse200117

`func NewInlineResponse200117() *InlineResponse200117`

NewInlineResponse200117 instantiates a new InlineResponse200117 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200117WithDefaults

`func NewInlineResponse200117WithDefaults() *InlineResponse200117`

NewInlineResponse200117WithDefaults instantiates a new InlineResponse200117 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSupportedMetrics

`func (o *InlineResponse200117) GetSupportedMetrics() []string`

GetSupportedMetrics returns the SupportedMetrics field if non-nil, zero value otherwise.

### GetSupportedMetricsOk

`func (o *InlineResponse200117) GetSupportedMetricsOk() (*[]string, bool)`

GetSupportedMetricsOk returns a tuple with the SupportedMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedMetrics

`func (o *InlineResponse200117) SetSupportedMetrics(v []string)`

SetSupportedMetrics sets SupportedMetrics field to given value.

### HasSupportedMetrics

`func (o *InlineResponse200117) HasSupportedMetrics() bool`

HasSupportedMetrics returns a boolean if a field has been set.

### GetCounts

`func (o *InlineResponse200117) GetCounts() InlineResponse200117Counts`

GetCounts returns the Counts field if non-nil, zero value otherwise.

### GetCountsOk

`func (o *InlineResponse200117) GetCountsOk() (*InlineResponse200117Counts, bool)`

GetCountsOk returns a tuple with the Counts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounts

`func (o *InlineResponse200117) SetCounts(v InlineResponse200117Counts)`

SetCounts sets Counts field to given value.

### HasCounts

`func (o *InlineResponse200117) HasCounts() bool`

HasCounts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


