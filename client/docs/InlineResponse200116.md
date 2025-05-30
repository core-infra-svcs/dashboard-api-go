# InlineResponse200116

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SupportedMetrics** | Pointer to **[]string** | List of metrics that are supported for alerts, based on available sensor devices in the network | [optional] 
**Counts** | Pointer to [**InlineResponse200116Counts**](InlineResponse200116Counts.md) |  | [optional] 

## Methods

### NewInlineResponse200116

`func NewInlineResponse200116() *InlineResponse200116`

NewInlineResponse200116 instantiates a new InlineResponse200116 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200116WithDefaults

`func NewInlineResponse200116WithDefaults() *InlineResponse200116`

NewInlineResponse200116WithDefaults instantiates a new InlineResponse200116 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSupportedMetrics

`func (o *InlineResponse200116) GetSupportedMetrics() []string`

GetSupportedMetrics returns the SupportedMetrics field if non-nil, zero value otherwise.

### GetSupportedMetricsOk

`func (o *InlineResponse200116) GetSupportedMetricsOk() (*[]string, bool)`

GetSupportedMetricsOk returns a tuple with the SupportedMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedMetrics

`func (o *InlineResponse200116) SetSupportedMetrics(v []string)`

SetSupportedMetrics sets SupportedMetrics field to given value.

### HasSupportedMetrics

`func (o *InlineResponse200116) HasSupportedMetrics() bool`

HasSupportedMetrics returns a boolean if a field has been set.

### GetCounts

`func (o *InlineResponse200116) GetCounts() InlineResponse200116Counts`

GetCounts returns the Counts field if non-nil, zero value otherwise.

### GetCountsOk

`func (o *InlineResponse200116) GetCountsOk() (*InlineResponse200116Counts, bool)`

GetCountsOk returns a tuple with the Counts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounts

`func (o *InlineResponse200116) SetCounts(v InlineResponse200116Counts)`

SetCounts sets Counts field to given value.

### HasCounts

`func (o *InlineResponse200116) HasCounts() bool`

HasCounts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


